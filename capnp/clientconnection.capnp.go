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

func NewHello(s *C.Segment) Hello      { return Hello(s.NewStruct(8, 2)) }
func NewRootHello(s *C.Segment) Hello  { return Hello(s.NewRootStruct(8, 2)) }
func AutoNewHello(s *C.Segment) Hello  { return Hello(s.NewStructAR(8, 2)) }
func ReadRootHello(s *C.Segment) Hello { return Hello(s.Root(0).ToStruct()) }
func (s Hello) Product() string        { return C.Struct(s).GetObject(0).ToText() }
func (s Hello) ProductBytes() []byte   { return C.Struct(s).GetObject(0).ToDataTrimLastByte() }
func (s Hello) SetProduct(v string)    { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s Hello) Version() string        { return C.Struct(s).GetObject(1).ToText() }
func (s Hello) VersionBytes() []byte   { return C.Struct(s).GetObject(1).ToDataTrimLastByte() }
func (s Hello) SetVersion(v string)    { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
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

func NewHelloList(s *C.Segment, sz int) Hello_List { return Hello_List(s.NewCompositeList(8, 2, sz)) }
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

type HelloClientFromServer C.Struct

func NewHelloClientFromServer(s *C.Segment) HelloClientFromServer {
	return HelloClientFromServer(s.NewStruct(0, 2))
}
func NewRootHelloClientFromServer(s *C.Segment) HelloClientFromServer {
	return HelloClientFromServer(s.NewRootStruct(0, 2))
}
func AutoNewHelloClientFromServer(s *C.Segment) HelloClientFromServer {
	return HelloClientFromServer(s.NewStructAR(0, 2))
}
func ReadRootHelloClientFromServer(s *C.Segment) HelloClientFromServer {
	return HelloClientFromServer(s.Root(0).ToStruct())
}
func (s HelloClientFromServer) Namespace() []byte     { return C.Struct(s).GetObject(0).ToData() }
func (s HelloClientFromServer) SetNamespace(v []byte) { C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s HelloClientFromServer) Roots() Root_List      { return Root_List(C.Struct(s).GetObject(1)) }
func (s HelloClientFromServer) SetRoots(v Root_List)  { C.Struct(s).SetObject(1, C.Object(v)) }
func (s HelloClientFromServer) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
	_, err = b.WriteString("\"roots\":")
	if err != nil {
		return err
	}
	{
		s := s.Roots()
		{
			err = b.WriteByte('[')
			if err != nil {
				return err
			}
			for i, s := range s.ToArray() {
				if i != 0 {
					_, err = b.WriteString(", ")
				}
				if err != nil {
					return err
				}
				err = s.WriteJSON(b)
				if err != nil {
					return err
				}
			}
			err = b.WriteByte(']')
		}
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
func (s HelloClientFromServer) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s HelloClientFromServer) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
	_, err = b.WriteString("roots = ")
	if err != nil {
		return err
	}
	{
		s := s.Roots()
		{
			err = b.WriteByte('[')
			if err != nil {
				return err
			}
			for i, s := range s.ToArray() {
				if i != 0 {
					_, err = b.WriteString(", ")
				}
				if err != nil {
					return err
				}
				err = s.WriteCapLit(b)
				if err != nil {
					return err
				}
			}
			err = b.WriteByte(']')
		}
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
func (s HelloClientFromServer) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type HelloClientFromServer_List C.PointerList

func NewHelloClientFromServerList(s *C.Segment, sz int) HelloClientFromServer_List {
	return HelloClientFromServer_List(s.NewCompositeList(0, 2, sz))
}
func (s HelloClientFromServer_List) Len() int { return C.PointerList(s).Len() }
func (s HelloClientFromServer_List) At(i int) HelloClientFromServer {
	return HelloClientFromServer(C.PointerList(s).At(i).ToStruct())
}
func (s HelloClientFromServer_List) ToArray() []HelloClientFromServer {
	n := s.Len()
	a := make([]HelloClientFromServer, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s HelloClientFromServer_List) Set(i int, item HelloClientFromServer) {
	C.PointerList(s).Set(i, C.Object(item))
}

type Root C.Struct

func NewRoot(s *C.Segment) Root               { return Root(s.NewStruct(0, 3)) }
func NewRootRoot(s *C.Segment) Root           { return Root(s.NewRootStruct(0, 3)) }
func AutoNewRoot(s *C.Segment) Root           { return Root(s.NewStructAR(0, 3)) }
func ReadRootRoot(s *C.Segment) Root          { return Root(s.Root(0).ToStruct()) }
func (s Root) Name() string                   { return C.Struct(s).GetObject(0).ToText() }
func (s Root) NameBytes() []byte              { return C.Struct(s).GetObject(0).ToDataTrimLastByte() }
func (s Root) SetName(v string)               { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s Root) VarId() []byte                  { return C.Struct(s).GetObject(1).ToData() }
func (s Root) SetVarId(v []byte)              { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s Root) Capabilities() Capabilities     { return Capabilities(C.Struct(s).GetObject(2).ToStruct()) }
func (s Root) SetCapabilities(v Capabilities) { C.Struct(s).SetObject(2, C.Object(v)) }
func (s Root) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"name\":")
	if err != nil {
		return err
	}
	{
		s := s.Name()
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
	_, err = b.WriteString("\"varId\":")
	if err != nil {
		return err
	}
	{
		s := s.VarId()
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
	_, err = b.WriteString("\"capabilities\":")
	if err != nil {
		return err
	}
	{
		s := s.Capabilities()
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
func (s Root) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Root) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("name = ")
	if err != nil {
		return err
	}
	{
		s := s.Name()
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
	_, err = b.WriteString("varId = ")
	if err != nil {
		return err
	}
	{
		s := s.VarId()
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
	_, err = b.WriteString("capabilities = ")
	if err != nil {
		return err
	}
	{
		s := s.Capabilities()
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
func (s Root) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Root_List C.PointerList

func NewRootList(s *C.Segment, sz int) Root_List { return Root_List(s.NewCompositeList(0, 3, sz)) }
func (s Root_List) Len() int                     { return C.PointerList(s).Len() }
func (s Root_List) At(i int) Root                { return Root(C.PointerList(s).At(i).ToStruct()) }
func (s Root_List) ToArray() []Root {
	n := s.Len()
	a := make([]Root, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Root_List) Set(i int, item Root) { C.PointerList(s).Set(i, C.Object(item)) }

type ClientMessage C.Struct
type ClientMessage_Which uint16

const (
	CLIENTMESSAGE_HEARTBEAT           ClientMessage_Which = 0
	CLIENTMESSAGE_CLIENTTXNSUBMISSION ClientMessage_Which = 1
	CLIENTMESSAGE_CLIENTTXNOUTCOME    ClientMessage_Which = 2
)

func NewClientMessage(s *C.Segment) ClientMessage      { return ClientMessage(s.NewStruct(8, 1)) }
func NewRootClientMessage(s *C.Segment) ClientMessage  { return ClientMessage(s.NewRootStruct(8, 1)) }
func AutoNewClientMessage(s *C.Segment) ClientMessage  { return ClientMessage(s.NewStructAR(8, 1)) }
func ReadRootClientMessage(s *C.Segment) ClientMessage { return ClientMessage(s.Root(0).ToStruct()) }
func (s ClientMessage) Which() ClientMessage_Which     { return ClientMessage_Which(C.Struct(s).Get16(0)) }
func (s ClientMessage) SetHeartbeat()                  { C.Struct(s).Set16(0, 0) }
func (s ClientMessage) ClientTxnSubmission() ClientTxn {
	return ClientTxn(C.Struct(s).GetObject(0).ToStruct())
}
func (s ClientMessage) SetClientTxnSubmission(v ClientTxn) {
	C.Struct(s).Set16(0, 1)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s ClientMessage) ClientTxnOutcome() ClientTxnOutcome {
	return ClientTxnOutcome(C.Struct(s).GetObject(0).ToStruct())
}
func (s ClientMessage) SetClientTxnOutcome(v ClientTxnOutcome) {
	C.Struct(s).Set16(0, 2)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s ClientMessage) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	if s.Which() == CLIENTMESSAGE_HEARTBEAT {
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
	if s.Which() == CLIENTMESSAGE_CLIENTTXNSUBMISSION {
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
	if s.Which() == CLIENTMESSAGE_CLIENTTXNOUTCOME {
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
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s ClientMessage) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s ClientMessage) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	if s.Which() == CLIENTMESSAGE_HEARTBEAT {
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
	if s.Which() == CLIENTMESSAGE_CLIENTTXNSUBMISSION {
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
	if s.Which() == CLIENTMESSAGE_CLIENTTXNOUTCOME {
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
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s ClientMessage) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type ClientMessage_List C.PointerList

func NewClientMessageList(s *C.Segment, sz int) ClientMessage_List {
	return ClientMessage_List(s.NewCompositeList(8, 1, sz))
}
func (s ClientMessage_List) Len() int { return C.PointerList(s).Len() }
func (s ClientMessage_List) At(i int) ClientMessage {
	return ClientMessage(C.PointerList(s).At(i).ToStruct())
}
func (s ClientMessage_List) ToArray() []ClientMessage {
	n := s.Len()
	a := make([]ClientMessage, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s ClientMessage_List) Set(i int, item ClientMessage) { C.PointerList(s).Set(i, C.Object(item)) }
