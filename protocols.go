package common

import (
	"fmt"
	capn "github.com/glycerine/go-capnproto"
	"github.com/go-kit/kit/log"
	"io"
	"math/rand"
	"net"
	"time"
)

type Dialer interface {
	Dial() error
	Socket() net.Conn
	Close()
	RemoteHost() string
}

// TCP dialer

type TCPDialer struct {
	logger     log.Logger
	remoteHost string
	socket     net.Conn
}

func NewTCPDialer(socket net.Conn, remoteHost string, logger log.Logger) *TCPDialer {
	dialer := &TCPDialer{
		logger:     logger,
		remoteHost: remoteHost,
		socket:     socket,
	}
	return dialer
}

func (td *TCPDialer) Dial() error {
	if td.socket != nil {
		return nil
	}
	if len(td.remoteHost) == 0 {
		panic("Attempting to dial to empty remote host")
	}
	td.logger.Log("msg", "Attempting connection.", "remoteHost", td.remoteHost)
	tcpAddr, err := net.ResolveTCPAddr("tcp", td.remoteHost)
	if err != nil {
		return err
	}
	socket, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		if socket != nil {
			socket.Close()
		}
		return err
	}
	if err = ConfigureSocket(socket); err != nil {
		socket.Close()
		return err
	}
	td.socket = socket
	return nil
}

func (td *TCPDialer) Socket() net.Conn {
	return td.socket
}

func (td *TCPDialer) Close() {
	if td.socket != nil {
		td.socket.Close()
		td.socket = nil
	}
}

func (td *TCPDialer) RemoteHost() string {
	return td.remoteHost
}

func (td *TCPDialer) String() string {
	return fmt.Sprintf("TCPDialer to %s", td.remoteHost)
}

// TLS Capnp Handshaker Base

type TLSCapnpHandshakerBase struct {
	Dialer
	remoteHost     string
	rng            *rand.Rand
	beater         *TLSCapnpBeater
	buf            []byte
	bufWriteOffset int64
}

func NewTLSCapnpHandshakerBase(dialer Dialer) *TLSCapnpHandshakerBase {
	return &TLSCapnpHandshakerBase{
		Dialer: dialer,
		rng:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (tchb *TLSCapnpHandshakerBase) InternalShutdown() {
	if tchb.beater != nil {
		tchb.beater.Stop()
		tchb.beater = nil
	}
	// do not nuke out the dialer!
	tchb.Dialer.Close()
	tchb.buf = nil
	tchb.bufWriteOffset = 0
}

func (tchb *TLSCapnpHandshakerBase) Send(msg []byte) error {
	l := len(msg)
	for l > 0 {
		switch w, err := tchb.Socket().Write(msg); {
		case err != nil:
			return err
		case w == l:
			return nil
		default:
			msg = msg[w:]
			l -= w
		}
	}
	return nil
}

func (tchb *TLSCapnpHandshakerBase) SendMessage(msg []byte) error {
	return tchb.beater.SendMessage(msg)
}

// this exists for handshake so that we don't accidentally read bits of the TLS handshake and break that!
func (tchb *TLSCapnpHandshakerBase) ReadExactlyOne() (*capn.Segment, error) {
	if err := tchb.Socket().SetReadDeadline(time.Now().Add(HeartbeatInterval << 1)); err != nil {
		return nil, err
	}
	return capn.ReadFromStream(tchb.Socket(), nil)
}

func (tchb *TLSCapnpHandshakerBase) ReadOne() (*capn.Segment, error) {
	if tchb.bufWriteOffset > 0 {
		// still have data in the buffer, so first try decoding that
		if seg, err := tchb.attemptCapnpDecode(); seg != nil || err != nil {
			return seg, err
		}
	}
	for {
		if int64(len(tchb.buf)) == tchb.bufWriteOffset { // run out of buf space
			tchb.increaseBuffer()
		}
		if err := tchb.Socket().SetReadDeadline(time.Now().Add(HeartbeatInterval << 1)); err != nil {
			return nil, err
		}
		if readCount, err := tchb.Socket().Read(tchb.buf[tchb.bufWriteOffset:]); err != nil {
			return nil, err
		} else if readCount > 0 { // we read something; try to parse
			tchb.bufWriteOffset += int64(readCount)
			if seg, err := tchb.attemptCapnpDecode(); seg != nil || err != nil {
				return seg, err
			}
		}
	}
}

func (tchb *TLSCapnpHandshakerBase) increaseBuffer() {
	size := int64(ConnectionBufferSize)
	req := tchb.bufWriteOffset
	for size <= req {
		size *= 2
	}
	buf := make([]byte, size)
	copy(buf[:req], tchb.buf[:req])
	tchb.buf = buf
}

func (tchb *TLSCapnpHandshakerBase) attemptCapnpDecode() (*capn.Segment, error) {
	seg, used, err := capn.ReadFromMemoryZeroCopy(tchb.buf[:tchb.bufWriteOffset])
	if used > 0 {
		tchb.buf = tchb.buf[used:]
		tchb.bufWriteOffset -= used
	}
	if err == io.EOF {
		// ignore it - it just means we don't have enough data from the socket yet
		return nil, nil
	} else {
		return seg, err
	}
}

func (tchb *TLSCapnpHandshakerBase) CreateBeater(conn ConnectionActor, beatBytes []byte) {
	if tchb.beater == nil {
		tchb.beater = NewTLSCapnpBeater(tchb, conn, beatBytes)
		tchb.beater.Start()
	}
}

// TLSCapnpBeater

type TLSCapnpBeater struct {
	*TLSCapnpHandshakerBase
	conn         ConnectionActor
	beatBytes    []byte
	terminate    chan struct{}
	terminated   chan struct{}
	ticker       *time.Ticker
	mustSendBeat bool
}

func NewTLSCapnpBeater(tchb *TLSCapnpHandshakerBase, conn ConnectionActor, beatBytes []byte) *TLSCapnpBeater {
	return &TLSCapnpBeater{
		TLSCapnpHandshakerBase: tchb,
		beatBytes:              beatBytes,
		conn:                   conn,
		terminate:              make(chan struct{}),
		terminated:             make(chan struct{}),
		ticker:                 time.NewTicker(HeartbeatInterval >> 1),
		mustSendBeat:           true,
	}
}

func (b *TLSCapnpBeater) Start() {
	if b != nil {
		go b.tick()
	}
}

func (b *TLSCapnpBeater) Stop() {
	if b != nil {
		b.TLSCapnpHandshakerBase = nil
		select {
		case <-b.terminated:
		default:
			close(b.terminate)
			<-b.terminated
		}
	}
}

func (b *TLSCapnpBeater) tick() {
	defer func() {
		b.ticker.Stop()
		b.ticker = nil
		close(b.terminated)
	}()
	for {
		select {
		case <-b.terminate:
			return
		case <-b.ticker.C:
			if !b.conn.EnqueueError(b.beat) {
				return
			}
		}
	}
}

func (b *TLSCapnpBeater) beat() error {
	if b != nil && b.TLSCapnpHandshakerBase != nil {
		if b.mustSendBeat {
			return b.SendMessage(b.beatBytes)
		} else {
			b.mustSendBeat = true
		}
		// Useful for testing recovery from network brownouts
		/*
			if b.rng.Intn(15) == 0 && b.dialer != nil {
				return fmt.Errorf("Random death. Restarting connection.")
			}
		*/
	}
	return nil
}

func (b *TLSCapnpBeater) SendMessage(msg []byte) error {
	if b != nil {
		b.mustSendBeat = false
		return b.Send(msg)
	}
	return nil
}

// Reader

type SocketReader struct {
	conn ConnectionActor
	SocketMsgHandler
	terminate  chan struct{}
	terminated chan struct{}
}

type SocketMsgHandler interface {
	ReadAndHandleOneMsg() error
}

func NewSocketReader(conn ConnectionActor, smh SocketMsgHandler) *SocketReader {
	return &SocketReader{
		conn:             conn,
		SocketMsgHandler: smh,
		terminate:        make(chan struct{}),
		terminated:       make(chan struct{}),
	}
}

func (sr *SocketReader) Start() {
	if sr != nil {
		go sr.read()
	}
}

func (sr *SocketReader) Stop() {
	if sr != nil {
		select {
		case <-sr.terminated:
		default:
			close(sr.terminate)
			<-sr.terminated
		}
	}
}

func (sr *SocketReader) read() {
	defer close(sr.terminated)
	for {
		select {
		case <-sr.terminate:
			return
		default:
			if err := sr.ReadAndHandleOneMsg(); err != nil {
				sr.conn.EnqueueError(func() error { return err }) // connectionReadError{error: err})
				return
			}
		}
	}
}

type ConnectionActor interface {
	EnqueueError(func() error) bool
}
