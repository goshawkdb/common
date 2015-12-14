package capnp

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	"encoding/json"
	C "github.com/glycerine/go-capnproto"
	"io"
)

type Hello C.Struct

func NewHello(s *C.Segment) Hello      { return Hello(s.NewStruct(8, 3)) }
func NewRootHello(s *C.Segment) Hello  { return Hello(s.NewRootStruct(8, 3)) }
func AutoNewHello(s *C.Segment) Hello  { return Hello(s.NewStructAR(8, 3)) }
func ReadRootHello(s *C.Segment) Hello { return Hello(s.Root(0).ToStruct()) }
func (s Hello) Product() string        { return C.Struct(s).GetObject(0).ToText() }
func (s Hello) SetProduct(v string)    { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s Hello) Version() string        { return C.Struct(s).GetObject(1).ToText() }
func (s Hello) SetVersion(v string)    { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s Hello) PublicKey() []byte      { return C.Struct(s).GetObject(2).ToData() }
func (s Hello) SetPublicKey(v []byte)  { C.Struct(s).SetObject(2, s.Segment.NewData(v)) }
func (s Hello) IsClient() bool         { return C.Struct(s).Get1(0) }
func (s Hello) SetIsClient(v bool)     { C.Struct(s).Set1(0, v) }
func (s Hello) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"product\":")
	if err != nil {
		return err
	}
	{
		s := s.Product()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"version\":")
	if err != nil {
		return err
	}
	{
		s := s.Version()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"publicKey\":")
	if err != nil {
		return err
	}
	{
		s := s.PublicKey()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"isClient\":")
	if err != nil {
		return err
	}
	{
		s := s.IsClient()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Hello) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Hello) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("product = ")
	if err != nil {
		return err
	}
	{
		s := s.Product()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("version = ")
	if err != nil {
		return err
	}
	{
		s := s.Version()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("publicKey = ")
	if err != nil {
		return err
	}
	{
		s := s.PublicKey()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("isClient = ")
	if err != nil {
		return err
	}
	{
		s := s.IsClient()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Hello) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Hello_List C.PointerList

func NewHelloList(s *C.Segment, sz int) Hello_List { return Hello_List(s.NewCompositeList(8, 3, sz)) }
func (s Hello_List) Len() int                      { return C.PointerList(s).Len() }
func (s Hello_List) At(i int) Hello                { return Hello(C.PointerList(s).At(i).ToStruct()) }
func (s Hello_List) ToArray() []Hello {
	n := s.Len()
	a := make([]Hello, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Hello_List) Set(i int, item Hello) { C.PointerList(s).Set(i, C.Object(item)) }

type HelloFromClient C.Struct

func NewHelloFromClient(s *C.Segment) HelloFromClient { return HelloFromClient(s.NewStruct(0, 2)) }
func NewRootHelloFromClient(s *C.Segment) HelloFromClient {
	return HelloFromClient(s.NewRootStruct(0, 2))
}
func AutoNewHelloFromClient(s *C.Segment) HelloFromClient { return HelloFromClient(s.NewStructAR(0, 2)) }
func ReadRootHelloFromClient(s *C.Segment) HelloFromClient {
	return HelloFromClient(s.Root(0).ToStruct())
}
func (s HelloFromClient) Username() string     { return C.Struct(s).GetObject(0).ToText() }
func (s HelloFromClient) SetUsername(v string) { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s HelloFromClient) Password() []byte     { return C.Struct(s).GetObject(1).ToData() }
func (s HelloFromClient) SetPassword(v []byte) { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s HelloFromClient) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"username\":")
	if err != nil {
		return err
	}
	{
		s := s.Username()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"password\":")
	if err != nil {
		return err
	}
	{
		s := s.Password()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s HelloFromClient) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s HelloFromClient) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("username = ")
	if err != nil {
		return err
	}
	{
		s := s.Username()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("password = ")
	if err != nil {
		return err
	}
	{
		s := s.Password()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s HelloFromClient) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type HelloFromClient_List C.PointerList

func NewHelloFromClientList(s *C.Segment, sz int) HelloFromClient_List {
	return HelloFromClient_List(s.NewCompositeList(0, 2, sz))
}
func (s HelloFromClient_List) Len() int { return C.PointerList(s).Len() }
func (s HelloFromClient_List) At(i int) HelloFromClient {
	return HelloFromClient(C.PointerList(s).At(i).ToStruct())
}
func (s HelloFromClient_List) ToArray() []HelloFromClient {
	n := s.Len()
	a := make([]HelloFromClient, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s HelloFromClient_List) Set(i int, item HelloFromClient) {
	C.PointerList(s).Set(i, C.Object(item))
}

type HelloFromServer C.Struct

func NewHelloFromServer(s *C.Segment) HelloFromServer { return HelloFromServer(s.NewStruct(8, 5)) }
func NewRootHelloFromServer(s *C.Segment) HelloFromServer {
	return HelloFromServer(s.NewRootStruct(8, 5))
}
func AutoNewHelloFromServer(s *C.Segment) HelloFromServer { return HelloFromServer(s.NewStructAR(8, 5)) }
func ReadRootHelloFromServer(s *C.Segment) HelloFromServer {
	return HelloFromServer(s.Root(0).ToStruct())
}
func (s HelloFromServer) LocalHost() string         { return C.Struct(s).GetObject(0).ToText() }
func (s HelloFromServer) SetLocalHost(v string)     { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s HelloFromServer) TieBreak() uint32          { return C.Struct(s).Get32(0) }
func (s HelloFromServer) SetTieBreak(v uint32)      { C.Struct(s).Set32(0, v) }
func (s HelloFromServer) Namespace() []byte         { return C.Struct(s).GetObject(1).ToData() }
func (s HelloFromServer) SetNamespace(v []byte)     { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s HelloFromServer) TopologyDBVersion() []byte { return C.Struct(s).GetObject(2).ToData() }
func (s HelloFromServer) SetTopologyDBVersion(v []byte) {
	C.Struct(s).SetObject(2, s.Segment.NewData(v))
}
func (s HelloFromServer) Topology() Topology     { return Topology(C.Struct(s).GetObject(3).ToStruct()) }
func (s HelloFromServer) SetTopology(v Topology) { C.Struct(s).SetObject(3, C.Object(v)) }
func (s HelloFromServer) Root() VarIdPos         { return VarIdPos(C.Struct(s).GetObject(4).ToStruct()) }
func (s HelloFromServer) SetRoot(v VarIdPos)     { C.Struct(s).SetObject(4, C.Object(v)) }
func (s HelloFromServer) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"localHost\":")
	if err != nil {
		return err
	}
	{
		s := s.LocalHost()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"tieBreak\":")
	if err != nil {
		return err
	}
	{
		s := s.TieBreak()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"namespace\":")
	if err != nil {
		return err
	}
	{
		s := s.Namespace()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"topologyDBVersion\":")
	if err != nil {
		return err
	}
	{
		s := s.TopologyDBVersion()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"topology\":")
	if err != nil {
		return err
	}
	{
		s := s.Topology()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"root\":")
	if err != nil {
		return err
	}
	{
		s := s.Root()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s HelloFromServer) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s HelloFromServer) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("localHost = ")
	if err != nil {
		return err
	}
	{
		s := s.LocalHost()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("tieBreak = ")
	if err != nil {
		return err
	}
	{
		s := s.TieBreak()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("namespace = ")
	if err != nil {
		return err
	}
	{
		s := s.Namespace()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("topologyDBVersion = ")
	if err != nil {
		return err
	}
	{
		s := s.TopologyDBVersion()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("topology = ")
	if err != nil {
		return err
	}
	{
		s := s.Topology()
		err = s.WriteCapLit(b)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("root = ")
	if err != nil {
		return err
	}
	{
		s := s.Root()
		err = s.WriteCapLit(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s HelloFromServer) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type HelloFromServer_List C.PointerList

func NewHelloFromServerList(s *C.Segment, sz int) HelloFromServer_List {
	return HelloFromServer_List(s.NewCompositeList(8, 5, sz))
}
func (s HelloFromServer_List) Len() int { return C.PointerList(s).Len() }
func (s HelloFromServer_List) At(i int) HelloFromServer {
	return HelloFromServer(C.PointerList(s).At(i).ToStruct())
}
func (s HelloFromServer_List) ToArray() []HelloFromServer {
	n := s.Len()
	a := make([]HelloFromServer, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s HelloFromServer_List) Set(i int, item HelloFromServer) {
	C.PointerList(s).Set(i, C.Object(item))
}

type Message C.Struct
type Message_Which uint16

const (
	MESSAGE_HEARTBEAT           Message_Which = 0
	MESSAGE_CLIENTTXNSUBMISSION Message_Which = 1
	MESSAGE_CLIENTTXNOUTCOME    Message_Which = 2
	MESSAGE_TXNSUBMISSION       Message_Which = 3
	MESSAGE_SUBMISSIONOUTCOME   Message_Which = 4
	MESSAGE_SUBMISSIONCOMPLETE  Message_Which = 5
	MESSAGE_SUBMISSIONABORT     Message_Which = 6
	MESSAGE_ONEATXNVOTES        Message_Which = 7
	MESSAGE_ONEBTXNVOTES        Message_Which = 8
	MESSAGE_TWOATXNVOTES        Message_Which = 9
	MESSAGE_TWOBTXNVOTES        Message_Which = 10
	MESSAGE_TXNLOCALLYCOMPLETE  Message_Which = 11
	MESSAGE_TXNGLOBALLYCOMPLETE Message_Which = 12
)

func NewMessage(s *C.Segment) Message      { return Message(s.NewStruct(8, 1)) }
func NewRootMessage(s *C.Segment) Message  { return Message(s.NewRootStruct(8, 1)) }
func AutoNewMessage(s *C.Segment) Message  { return Message(s.NewStructAR(8, 1)) }
func ReadRootMessage(s *C.Segment) Message { return Message(s.Root(0).ToStruct()) }
func (s Message) Which() Message_Which     { return Message_Which(C.Struct(s).Get16(0)) }
func (s Message) SetHeartbeat()            { C.Struct(s).Set16(0, 0) }
func (s Message) ClientTxnSubmission() ClientTxn {
	return ClientTxn(C.Struct(s).GetObject(0).ToStruct())
}
func (s Message) SetClientTxnSubmission(v ClientTxn) {
	C.Struct(s).Set16(0, 1)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) ClientTxnOutcome() ClientTxnOutcome {
	return ClientTxnOutcome(C.Struct(s).GetObject(0).ToStruct())
}
func (s Message) SetClientTxnOutcome(v ClientTxnOutcome) {
	C.Struct(s).Set16(0, 2)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) TxnSubmission() Txn { return Txn(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetTxnSubmission(v Txn) {
	C.Struct(s).Set16(0, 3)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) SubmissionOutcome() Outcome { return Outcome(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetSubmissionOutcome(v Outcome) {
	C.Struct(s).Set16(0, 4)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) SubmissionComplete() TxnSubmissionComplete {
	return TxnSubmissionComplete(C.Struct(s).GetObject(0).ToStruct())
}
func (s Message) SetSubmissionComplete(v TxnSubmissionComplete) {
	C.Struct(s).Set16(0, 5)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) SubmissionAbort() TxnSubmissionAbort {
	return TxnSubmissionAbort(C.Struct(s).GetObject(0).ToStruct())
}
func (s Message) SetSubmissionAbort(v TxnSubmissionAbort) {
	C.Struct(s).Set16(0, 6)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) OneATxnVotes() OneATxnVotes { return OneATxnVotes(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetOneATxnVotes(v OneATxnVotes) {
	C.Struct(s).Set16(0, 7)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) OneBTxnVotes() OneBTxnVotes { return OneBTxnVotes(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetOneBTxnVotes(v OneBTxnVotes) {
	C.Struct(s).Set16(0, 8)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) TwoATxnVotes() TwoATxnVotes { return TwoATxnVotes(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetTwoATxnVotes(v TwoATxnVotes) {
	C.Struct(s).Set16(0, 9)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) TwoBTxnVotes() TwoBTxnVotes { return TwoBTxnVotes(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetTwoBTxnVotes(v TwoBTxnVotes) {
	C.Struct(s).Set16(0, 10)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) TxnLocallyComplete() TxnLocallyComplete {
	return TxnLocallyComplete(C.Struct(s).GetObject(0).ToStruct())
}
func (s Message) SetTxnLocallyComplete(v TxnLocallyComplete) {
	C.Struct(s).Set16(0, 11)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) TxnGloballyComplete() TxnGloballyComplete {
	return TxnGloballyComplete(C.Struct(s).GetObject(0).ToStruct())
}
func (s Message) SetTxnGloballyComplete(v TxnGloballyComplete) {
	C.Struct(s).Set16(0, 12)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	if s.Which() == MESSAGE_HEARTBEAT {
		_, err = b.WriteString("\"heartbeat\":")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == MESSAGE_CLIENTTXNSUBMISSION {
		_, err = b.WriteString("\"clientTxnSubmission\":")
		if err != nil {
			return err
		}
		{
			s := s.ClientTxnSubmission()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_CLIENTTXNOUTCOME {
		_, err = b.WriteString("\"clientTxnOutcome\":")
		if err != nil {
			return err
		}
		{
			s := s.ClientTxnOutcome()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_TXNSUBMISSION {
		_, err = b.WriteString("\"txnSubmission\":")
		if err != nil {
			return err
		}
		{
			s := s.TxnSubmission()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_SUBMISSIONOUTCOME {
		_, err = b.WriteString("\"submissionOutcome\":")
		if err != nil {
			return err
		}
		{
			s := s.SubmissionOutcome()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_SUBMISSIONCOMPLETE {
		_, err = b.WriteString("\"submissionComplete\":")
		if err != nil {
			return err
		}
		{
			s := s.SubmissionComplete()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_SUBMISSIONABORT {
		_, err = b.WriteString("\"submissionAbort\":")
		if err != nil {
			return err
		}
		{
			s := s.SubmissionAbort()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_ONEATXNVOTES {
		_, err = b.WriteString("\"oneATxnVotes\":")
		if err != nil {
			return err
		}
		{
			s := s.OneATxnVotes()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_ONEBTXNVOTES {
		_, err = b.WriteString("\"oneBTxnVotes\":")
		if err != nil {
			return err
		}
		{
			s := s.OneBTxnVotes()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_TWOATXNVOTES {
		_, err = b.WriteString("\"twoATxnVotes\":")
		if err != nil {
			return err
		}
		{
			s := s.TwoATxnVotes()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_TWOBTXNVOTES {
		_, err = b.WriteString("\"twoBTxnVotes\":")
		if err != nil {
			return err
		}
		{
			s := s.TwoBTxnVotes()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_TXNLOCALLYCOMPLETE {
		_, err = b.WriteString("\"txnLocallyComplete\":")
		if err != nil {
			return err
		}
		{
			s := s.TxnLocallyComplete()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_TXNGLOBALLYCOMPLETE {
		_, err = b.WriteString("\"txnGloballyComplete\":")
		if err != nil {
			return err
		}
		{
			s := s.TxnGloballyComplete()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Message) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Message) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	if s.Which() == MESSAGE_HEARTBEAT {
		_, err = b.WriteString("heartbeat = ")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == MESSAGE_CLIENTTXNSUBMISSION {
		_, err = b.WriteString("clientTxnSubmission = ")
		if err != nil {
			return err
		}
		{
			s := s.ClientTxnSubmission()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_CLIENTTXNOUTCOME {
		_, err = b.WriteString("clientTxnOutcome = ")
		if err != nil {
			return err
		}
		{
			s := s.ClientTxnOutcome()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_TXNSUBMISSION {
		_, err = b.WriteString("txnSubmission = ")
		if err != nil {
			return err
		}
		{
			s := s.TxnSubmission()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_SUBMISSIONOUTCOME {
		_, err = b.WriteString("submissionOutcome = ")
		if err != nil {
			return err
		}
		{
			s := s.SubmissionOutcome()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_SUBMISSIONCOMPLETE {
		_, err = b.WriteString("submissionComplete = ")
		if err != nil {
			return err
		}
		{
			s := s.SubmissionComplete()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_SUBMISSIONABORT {
		_, err = b.WriteString("submissionAbort = ")
		if err != nil {
			return err
		}
		{
			s := s.SubmissionAbort()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_ONEATXNVOTES {
		_, err = b.WriteString("oneATxnVotes = ")
		if err != nil {
			return err
		}
		{
			s := s.OneATxnVotes()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_ONEBTXNVOTES {
		_, err = b.WriteString("oneBTxnVotes = ")
		if err != nil {
			return err
		}
		{
			s := s.OneBTxnVotes()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_TWOATXNVOTES {
		_, err = b.WriteString("twoATxnVotes = ")
		if err != nil {
			return err
		}
		{
			s := s.TwoATxnVotes()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_TWOBTXNVOTES {
		_, err = b.WriteString("twoBTxnVotes = ")
		if err != nil {
			return err
		}
		{
			s := s.TwoBTxnVotes()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_TXNLOCALLYCOMPLETE {
		_, err = b.WriteString("txnLocallyComplete = ")
		if err != nil {
			return err
		}
		{
			s := s.TxnLocallyComplete()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == MESSAGE_TXNGLOBALLYCOMPLETE {
		_, err = b.WriteString("txnGloballyComplete = ")
		if err != nil {
			return err
		}
		{
			s := s.TxnGloballyComplete()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Message) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Message_List C.PointerList

func NewMessageList(s *C.Segment, sz int) Message_List {
	return Message_List(s.NewCompositeList(8, 1, sz))
}
func (s Message_List) Len() int         { return C.PointerList(s).Len() }
func (s Message_List) At(i int) Message { return Message(C.PointerList(s).At(i).ToStruct()) }
func (s Message_List) ToArray() []Message {
	n := s.Len()
	a := make([]Message, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Message_List) Set(i int, item Message) { C.PointerList(s).Set(i, C.Object(item)) }
