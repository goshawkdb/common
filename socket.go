package common

import (
	"net"
	"time"
)

func ConfigureSocket(socket *net.TCPConn) error {
	if err := socket.SetKeepAlive(true); err != nil {
		return err
	}
	if err := socket.SetKeepAlivePeriod(time.Second); err != nil {
		return err
	}
	if err := socket.SetDeadline(time.Time{}); err != nil {
		return err
	}
	if err := socket.SetReadBuffer(ConnectionBufferSize); err != nil {
		return err
	}
	if err := socket.SetWriteBuffer(ConnectionBufferSize); err != nil {
		return err
	}
	return nil
}
