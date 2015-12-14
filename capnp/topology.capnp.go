package capnp

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	"encoding/json"
	C "github.com/glycerine/go-capnproto"
	"io"
)

type Topology C.Struct

func NewTopology(s *C.Segment) Topology       { return Topology(s.NewStruct(8, 4)) }
func NewRootTopology(s *C.Segment) Topology   { return Topology(s.NewRootStruct(8, 4)) }
func AutoNewTopology(s *C.Segment) Topology   { return Topology(s.NewStructAR(8, 4)) }
func ReadRootTopology(s *C.Segment) Topology  { return Topology(s.Root(0).ToStruct()) }
func (s Topology) ClusterId() string          { return C.Struct(s).GetObject(0).ToText() }
func (s Topology) SetClusterId(v string)      { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s Topology) Version() uint32            { return C.Struct(s).Get32(0) }
func (s Topology) SetVersion(v uint32)        { C.Struct(s).Set32(0, v) }
func (s Topology) Hosts() C.TextList          { return C.TextList(C.Struct(s).GetObject(1)) }
func (s Topology) SetHosts(v C.TextList)      { C.Struct(s).SetObject(1, C.Object(v)) }
func (s Topology) F() uint8                   { return C.Struct(s).Get8(4) }
func (s Topology) SetF(v uint8)               { C.Struct(s).Set8(4, v) }
func (s Topology) MaxRMCount() uint8          { return C.Struct(s).Get8(5) }
func (s Topology) SetMaxRMCount(v uint8)      { C.Struct(s).Set8(5, v) }
func (s Topology) AsyncFlush() bool           { return C.Struct(s).Get1(48) }
func (s Topology) SetAsyncFlush(v bool)       { C.Struct(s).Set1(48, v) }
func (s Topology) Rms() C.UInt32List          { return C.UInt32List(C.Struct(s).GetObject(2)) }
func (s Topology) SetRms(v C.UInt32List)      { C.Struct(s).SetObject(2, C.Object(v)) }
func (s Topology) Accounts() Account_List     { return Account_List(C.Struct(s).GetObject(3)) }
func (s Topology) SetAccounts(v Account_List) { C.Struct(s).SetObject(3, C.Object(v)) }
func (s Topology) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"clusterId\":")
	if err != nil {
		return err
	}
	{
		s := s.ClusterId()
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
	_, err = b.WriteString("\"hosts\":")
	if err != nil {
		return err
	}
	{
		s := s.Hosts()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"f\":")
	if err != nil {
		return err
	}
	{
		s := s.F()
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
	_, err = b.WriteString("\"maxRMCount\":")
	if err != nil {
		return err
	}
	{
		s := s.MaxRMCount()
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
	_, err = b.WriteString("\"asyncFlush\":")
	if err != nil {
		return err
	}
	{
		s := s.AsyncFlush()
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
	_, err = b.WriteString("\"rms\":")
	if err != nil {
		return err
	}
	{
		s := s.Rms()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"accounts\":")
	if err != nil {
		return err
	}
	{
		s := s.Accounts()
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
func (s Topology) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Topology) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("clusterId = ")
	if err != nil {
		return err
	}
	{
		s := s.ClusterId()
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
	_, err = b.WriteString("hosts = ")
	if err != nil {
		return err
	}
	{
		s := s.Hosts()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("f = ")
	if err != nil {
		return err
	}
	{
		s := s.F()
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
	_, err = b.WriteString("maxRMCount = ")
	if err != nil {
		return err
	}
	{
		s := s.MaxRMCount()
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
	_, err = b.WriteString("asyncFlush = ")
	if err != nil {
		return err
	}
	{
		s := s.AsyncFlush()
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
	_, err = b.WriteString("rms = ")
	if err != nil {
		return err
	}
	{
		s := s.Rms()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("accounts = ")
	if err != nil {
		return err
	}
	{
		s := s.Accounts()
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
func (s Topology) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Topology_List C.PointerList

func NewTopologyList(s *C.Segment, sz int) Topology_List {
	return Topology_List(s.NewCompositeList(8, 4, sz))
}
func (s Topology_List) Len() int          { return C.PointerList(s).Len() }
func (s Topology_List) At(i int) Topology { return Topology(C.PointerList(s).At(i).ToStruct()) }
func (s Topology_List) ToArray() []Topology {
	n := s.Len()
	a := make([]Topology, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Topology_List) Set(i int, item Topology) { C.PointerList(s).Set(i, C.Object(item)) }

type Account C.Struct

func NewAccount(s *C.Segment) Account      { return Account(s.NewStruct(0, 2)) }
func NewRootAccount(s *C.Segment) Account  { return Account(s.NewRootStruct(0, 2)) }
func AutoNewAccount(s *C.Segment) Account  { return Account(s.NewStructAR(0, 2)) }
func ReadRootAccount(s *C.Segment) Account { return Account(s.Root(0).ToStruct()) }
func (s Account) Username() string         { return C.Struct(s).GetObject(0).ToText() }
func (s Account) SetUsername(v string)     { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s Account) Password() string         { return C.Struct(s).GetObject(1).ToText() }
func (s Account) SetPassword(v string)     { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s Account) WriteJSON(w io.Writer) error {
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
func (s Account) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Account) WriteCapLit(w io.Writer) error {
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
func (s Account) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Account_List C.PointerList

func NewAccountList(s *C.Segment, sz int) Account_List {
	return Account_List(s.NewCompositeList(0, 2, sz))
}
func (s Account_List) Len() int         { return C.PointerList(s).Len() }
func (s Account_List) At(i int) Account { return Account(C.PointerList(s).At(i).ToStruct()) }
func (s Account_List) ToArray() []Account {
	n := s.Len()
	a := make([]Account, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Account_List) Set(i int, item Account) { C.PointerList(s).Set(i, C.Object(item)) }
