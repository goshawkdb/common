package capnp

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	"encoding/json"
	C "github.com/glycerine/go-capnproto"
	"io"
)

type VarIdPos C.Struct

func NewVarIdPos(s *C.Segment) VarIdPos       { return VarIdPos(s.NewStruct(0, 2)) }
func NewRootVarIdPos(s *C.Segment) VarIdPos   { return VarIdPos(s.NewRootStruct(0, 2)) }
func AutoNewVarIdPos(s *C.Segment) VarIdPos   { return VarIdPos(s.NewStructAR(0, 2)) }
func ReadRootVarIdPos(s *C.Segment) VarIdPos  { return VarIdPos(s.Root(0).ToStruct()) }
func (s VarIdPos) Id() []byte                 { return C.Struct(s).GetObject(0).ToData() }
func (s VarIdPos) SetId(v []byte)             { C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s VarIdPos) Positions() C.UInt8List     { return C.UInt8List(C.Struct(s).GetObject(1)) }
func (s VarIdPos) SetPositions(v C.UInt8List) { C.Struct(s).SetObject(1, C.Object(v)) }
func (s VarIdPos) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"id\":")
	if err != nil {
		return err
	}
	{
		s := s.Id()
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
	_, err = b.WriteString("\"positions\":")
	if err != nil {
		return err
	}
	{
		s := s.Positions()
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
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s VarIdPos) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s VarIdPos) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("id = ")
	if err != nil {
		return err
	}
	{
		s := s.Id()
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
	_, err = b.WriteString("positions = ")
	if err != nil {
		return err
	}
	{
		s := s.Positions()
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
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s VarIdPos) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type VarIdPos_List C.PointerList

func NewVarIdPosList(s *C.Segment, sz int) VarIdPos_List {
	return VarIdPos_List(s.NewCompositeList(0, 2, sz))
}
func (s VarIdPos_List) Len() int          { return C.PointerList(s).Len() }
func (s VarIdPos_List) At(i int) VarIdPos { return VarIdPos(C.PointerList(s).At(i).ToStruct()) }
func (s VarIdPos_List) ToArray() []VarIdPos {
	n := s.Len()
	a := make([]VarIdPos, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s VarIdPos_List) Set(i int, item VarIdPos) { C.PointerList(s).Set(i, C.Object(item)) }
