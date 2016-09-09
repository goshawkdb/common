package capnp

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	C "github.com/glycerine/go-capnproto"
	"io"
)

type Capability C.Struct
type Capability_Which uint16

const (
	CAPABILITY_NONE      Capability_Which = 0
	CAPABILITY_READ      Capability_Which = 1
	CAPABILITY_WRITE     Capability_Which = 2
	CAPABILITY_READWRITE Capability_Which = 3
)

func NewCapability(s *C.Segment) Capability      { return Capability(s.NewStruct(8, 0)) }
func NewRootCapability(s *C.Segment) Capability  { return Capability(s.NewRootStruct(8, 0)) }
func AutoNewCapability(s *C.Segment) Capability  { return Capability(s.NewStructAR(8, 0)) }
func ReadRootCapability(s *C.Segment) Capability { return Capability(s.Root(0).ToStruct()) }
func (s Capability) Which() Capability_Which     { return Capability_Which(C.Struct(s).Get16(0)) }
func (s Capability) SetNone()                    { C.Struct(s).Set16(0, 0) }
func (s Capability) SetRead()                    { C.Struct(s).Set16(0, 1) }
func (s Capability) SetWrite()                   { C.Struct(s).Set16(0, 2) }
func (s Capability) SetReadWrite()               { C.Struct(s).Set16(0, 3) }
func (s Capability) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	if s.Which() == CAPABILITY_NONE {
		_, err = b.WriteString("\"none\":")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == CAPABILITY_READ {
		_, err = b.WriteString("\"read\":")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == CAPABILITY_WRITE {
		_, err = b.WriteString("\"write\":")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == CAPABILITY_READWRITE {
		_, err = b.WriteString("\"readWrite\":")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
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
func (s Capability) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Capability) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	if s.Which() == CAPABILITY_NONE {
		_, err = b.WriteString("none = ")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == CAPABILITY_READ {
		_, err = b.WriteString("read = ")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == CAPABILITY_WRITE {
		_, err = b.WriteString("write = ")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == CAPABILITY_READWRITE {
		_, err = b.WriteString("readWrite = ")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
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
func (s Capability) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Capability_List C.PointerList

func NewCapabilityList(s *C.Segment, sz int) Capability_List {
	return Capability_List(s.NewCompositeList(8, 0, sz))
}
func (s Capability_List) Len() int            { return C.PointerList(s).Len() }
func (s Capability_List) At(i int) Capability { return Capability(C.PointerList(s).At(i).ToStruct()) }
func (s Capability_List) ToArray() []Capability {
	n := s.Len()
	a := make([]Capability, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Capability_List) Set(i int, item Capability) { C.PointerList(s).Set(i, C.Object(item)) }
