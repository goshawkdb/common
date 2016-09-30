package capnp

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	"encoding/json"
	C "github.com/glycerine/go-capnproto"
	"io"
)

type ClientTxn C.Struct

func NewClientTxn(s *C.Segment) ClientTxn          { return ClientTxn(s.NewStruct(8, 2)) }
func NewRootClientTxn(s *C.Segment) ClientTxn      { return ClientTxn(s.NewRootStruct(8, 2)) }
func AutoNewClientTxn(s *C.Segment) ClientTxn      { return ClientTxn(s.NewStructAR(8, 2)) }
func ReadRootClientTxn(s *C.Segment) ClientTxn     { return ClientTxn(s.Root(0).ToStruct()) }
func (s ClientTxn) Id() []byte                     { return C.Struct(s).GetObject(0).ToData() }
func (s ClientTxn) SetId(v []byte)                 { C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s ClientTxn) Retry() bool                    { return C.Struct(s).Get1(0) }
func (s ClientTxn) SetRetry(v bool)                { C.Struct(s).Set1(0, v) }
func (s ClientTxn) Actions() ClientAction_List     { return ClientAction_List(C.Struct(s).GetObject(1)) }
func (s ClientTxn) SetActions(v ClientAction_List) { C.Struct(s).SetObject(1, C.Object(v)) }
func (s ClientTxn) WriteJSON(w io.Writer) error {
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
	_, err = b.WriteString("\"retry\":")
	if err != nil {
		return err
	}
	{
		s := s.Retry()
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
	_, err = b.WriteString("\"actions\":")
	if err != nil {
		return err
	}
	{
		s := s.Actions()
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
func (s ClientTxn) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s ClientTxn) WriteCapLit(w io.Writer) error {
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
	_, err = b.WriteString("retry = ")
	if err != nil {
		return err
	}
	{
		s := s.Retry()
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
	_, err = b.WriteString("actions = ")
	if err != nil {
		return err
	}
	{
		s := s.Actions()
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
func (s ClientTxn) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type ClientTxn_List C.PointerList

func NewClientTxnList(s *C.Segment, sz int) ClientTxn_List {
	return ClientTxn_List(s.NewCompositeList(8, 2, sz))
}
func (s ClientTxn_List) Len() int           { return C.PointerList(s).Len() }
func (s ClientTxn_List) At(i int) ClientTxn { return ClientTxn(C.PointerList(s).At(i).ToStruct()) }
func (s ClientTxn_List) ToArray() []ClientTxn {
	n := s.Len()
	a := make([]ClientTxn, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s ClientTxn_List) Set(i int, item ClientTxn) { C.PointerList(s).Set(i, C.Object(item)) }

type ClientAction C.Struct
type ClientActionRead ClientAction
type ClientActionWrite ClientAction
type ClientActionReadwrite ClientAction
type ClientActionCreate ClientAction
type ClientActionRoll ClientAction
type ClientAction_Which uint16

const (
	CLIENTACTION_READ      ClientAction_Which = 0
	CLIENTACTION_WRITE     ClientAction_Which = 1
	CLIENTACTION_READWRITE ClientAction_Which = 2
	CLIENTACTION_CREATE    ClientAction_Which = 3
	CLIENTACTION_DELETE    ClientAction_Which = 4
	CLIENTACTION_ROLL      ClientAction_Which = 5
)

func NewClientAction(s *C.Segment) ClientAction      { return ClientAction(s.NewStruct(8, 4)) }
func NewRootClientAction(s *C.Segment) ClientAction  { return ClientAction(s.NewRootStruct(8, 4)) }
func AutoNewClientAction(s *C.Segment) ClientAction  { return ClientAction(s.NewStructAR(8, 4)) }
func ReadRootClientAction(s *C.Segment) ClientAction { return ClientAction(s.Root(0).ToStruct()) }
func (s ClientAction) Which() ClientAction_Which     { return ClientAction_Which(C.Struct(s).Get16(0)) }
func (s ClientAction) VarId() []byte                 { return C.Struct(s).GetObject(0).ToData() }
func (s ClientAction) SetVarId(v []byte)             { C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s ClientAction) Read() ClientActionRead        { return ClientActionRead(s) }
func (s ClientAction) SetRead()                      { C.Struct(s).Set16(0, 0) }
func (s ClientActionRead) Version() []byte           { return C.Struct(s).GetObject(1).ToData() }
func (s ClientActionRead) SetVersion(v []byte)       { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s ClientAction) Write() ClientActionWrite      { return ClientActionWrite(s) }
func (s ClientAction) SetWrite()                     { C.Struct(s).Set16(0, 1) }
func (s ClientActionWrite) Value() []byte            { return C.Struct(s).GetObject(1).ToData() }
func (s ClientActionWrite) SetValue(v []byte)        { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s ClientActionWrite) References() ClientVarIdPos_List {
	return ClientVarIdPos_List(C.Struct(s).GetObject(2))
}
func (s ClientActionWrite) SetReferences(v ClientVarIdPos_List) { C.Struct(s).SetObject(2, C.Object(v)) }
func (s ClientAction) Readwrite() ClientActionReadwrite         { return ClientActionReadwrite(s) }
func (s ClientAction) SetReadwrite()                            { C.Struct(s).Set16(0, 2) }
func (s ClientActionReadwrite) Version() []byte                 { return C.Struct(s).GetObject(1).ToData() }
func (s ClientActionReadwrite) SetVersion(v []byte)             { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s ClientActionReadwrite) Value() []byte                   { return C.Struct(s).GetObject(2).ToData() }
func (s ClientActionReadwrite) SetValue(v []byte)               { C.Struct(s).SetObject(2, s.Segment.NewData(v)) }
func (s ClientActionReadwrite) References() ClientVarIdPos_List {
	return ClientVarIdPos_List(C.Struct(s).GetObject(3))
}
func (s ClientActionReadwrite) SetReferences(v ClientVarIdPos_List) {
	C.Struct(s).SetObject(3, C.Object(v))
}
func (s ClientAction) Create() ClientActionCreate { return ClientActionCreate(s) }
func (s ClientAction) SetCreate()                 { C.Struct(s).Set16(0, 3) }
func (s ClientActionCreate) Value() []byte        { return C.Struct(s).GetObject(1).ToData() }
func (s ClientActionCreate) SetValue(v []byte)    { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s ClientActionCreate) References() ClientVarIdPos_List {
	return ClientVarIdPos_List(C.Struct(s).GetObject(2))
}
func (s ClientActionCreate) SetReferences(v ClientVarIdPos_List) {
	C.Struct(s).SetObject(2, C.Object(v))
}
func (s ClientAction) SetDelete()              { C.Struct(s).Set16(0, 4) }
func (s ClientAction) Roll() ClientActionRoll  { return ClientActionRoll(s) }
func (s ClientAction) SetRoll()                { C.Struct(s).Set16(0, 5) }
func (s ClientActionRoll) Version() []byte     { return C.Struct(s).GetObject(1).ToData() }
func (s ClientActionRoll) SetVersion(v []byte) { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s ClientActionRoll) Value() []byte       { return C.Struct(s).GetObject(2).ToData() }
func (s ClientActionRoll) SetValue(v []byte)   { C.Struct(s).SetObject(2, s.Segment.NewData(v)) }
func (s ClientActionRoll) References() ClientVarIdPos_List {
	return ClientVarIdPos_List(C.Struct(s).GetObject(3))
}
func (s ClientActionRoll) SetReferences(v ClientVarIdPos_List) { C.Struct(s).SetObject(3, C.Object(v)) }
func (s ClientAction) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
	if s.Which() == CLIENTACTION_READ {
		_, err = b.WriteString("\"read\":")
		if err != nil {
			return err
		}
		{
			s := s.Read()
			err = b.WriteByte('{')
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
			err = b.WriteByte('}')
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == CLIENTACTION_WRITE {
		_, err = b.WriteString("\"write\":")
		if err != nil {
			return err
		}
		{
			s := s.Write()
			err = b.WriteByte('{')
			if err != nil {
				return err
			}
			_, err = b.WriteString("\"value\":")
			if err != nil {
				return err
			}
			{
				s := s.Value()
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
			_, err = b.WriteString("\"references\":")
			if err != nil {
				return err
			}
			{
				s := s.References()
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
		}
	}
	if s.Which() == CLIENTACTION_READWRITE {
		_, err = b.WriteString("\"readwrite\":")
		if err != nil {
			return err
		}
		{
			s := s.Readwrite()
			err = b.WriteByte('{')
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
			_, err = b.WriteString("\"value\":")
			if err != nil {
				return err
			}
			{
				s := s.Value()
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
			_, err = b.WriteString("\"references\":")
			if err != nil {
				return err
			}
			{
				s := s.References()
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
		}
	}
	if s.Which() == CLIENTACTION_CREATE {
		_, err = b.WriteString("\"create\":")
		if err != nil {
			return err
		}
		{
			s := s.Create()
			err = b.WriteByte('{')
			if err != nil {
				return err
			}
			_, err = b.WriteString("\"value\":")
			if err != nil {
				return err
			}
			{
				s := s.Value()
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
			_, err = b.WriteString("\"references\":")
			if err != nil {
				return err
			}
			{
				s := s.References()
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
		}
	}
	if s.Which() == CLIENTACTION_DELETE {
		_, err = b.WriteString("\"delete\":")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == CLIENTACTION_ROLL {
		_, err = b.WriteString("\"roll\":")
		if err != nil {
			return err
		}
		{
			s := s.Roll()
			err = b.WriteByte('{')
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
			_, err = b.WriteString("\"value\":")
			if err != nil {
				return err
			}
			{
				s := s.Value()
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
			_, err = b.WriteString("\"references\":")
			if err != nil {
				return err
			}
			{
				s := s.References()
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
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s ClientAction) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s ClientAction) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
	if s.Which() == CLIENTACTION_READ {
		_, err = b.WriteString("read = ")
		if err != nil {
			return err
		}
		{
			s := s.Read()
			err = b.WriteByte('(')
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
			err = b.WriteByte(')')
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == CLIENTACTION_WRITE {
		_, err = b.WriteString("write = ")
		if err != nil {
			return err
		}
		{
			s := s.Write()
			err = b.WriteByte('(')
			if err != nil {
				return err
			}
			_, err = b.WriteString("value = ")
			if err != nil {
				return err
			}
			{
				s := s.Value()
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
			_, err = b.WriteString("references = ")
			if err != nil {
				return err
			}
			{
				s := s.References()
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
		}
	}
	if s.Which() == CLIENTACTION_READWRITE {
		_, err = b.WriteString("readwrite = ")
		if err != nil {
			return err
		}
		{
			s := s.Readwrite()
			err = b.WriteByte('(')
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
			_, err = b.WriteString("value = ")
			if err != nil {
				return err
			}
			{
				s := s.Value()
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
			_, err = b.WriteString("references = ")
			if err != nil {
				return err
			}
			{
				s := s.References()
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
		}
	}
	if s.Which() == CLIENTACTION_CREATE {
		_, err = b.WriteString("create = ")
		if err != nil {
			return err
		}
		{
			s := s.Create()
			err = b.WriteByte('(')
			if err != nil {
				return err
			}
			_, err = b.WriteString("value = ")
			if err != nil {
				return err
			}
			{
				s := s.Value()
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
			_, err = b.WriteString("references = ")
			if err != nil {
				return err
			}
			{
				s := s.References()
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
		}
	}
	if s.Which() == CLIENTACTION_DELETE {
		_, err = b.WriteString("delete = ")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == CLIENTACTION_ROLL {
		_, err = b.WriteString("roll = ")
		if err != nil {
			return err
		}
		{
			s := s.Roll()
			err = b.WriteByte('(')
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
			_, err = b.WriteString("value = ")
			if err != nil {
				return err
			}
			{
				s := s.Value()
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
			_, err = b.WriteString("references = ")
			if err != nil {
				return err
			}
			{
				s := s.References()
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
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s ClientAction) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type ClientAction_List C.PointerList

func NewClientActionList(s *C.Segment, sz int) ClientAction_List {
	return ClientAction_List(s.NewCompositeList(8, 4, sz))
}
func (s ClientAction_List) Len() int { return C.PointerList(s).Len() }
func (s ClientAction_List) At(i int) ClientAction {
	return ClientAction(C.PointerList(s).At(i).ToStruct())
}
func (s ClientAction_List) ToArray() []ClientAction {
	n := s.Len()
	a := make([]ClientAction, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s ClientAction_List) Set(i int, item ClientAction) { C.PointerList(s).Set(i, C.Object(item)) }

type ClientTxnOutcome C.Struct
type ClientTxnOutcome_Which uint16

const (
	CLIENTTXNOUTCOME_COMMIT ClientTxnOutcome_Which = 0
	CLIENTTXNOUTCOME_ABORT  ClientTxnOutcome_Which = 1
	CLIENTTXNOUTCOME_ERROR  ClientTxnOutcome_Which = 2
)

func NewClientTxnOutcome(s *C.Segment) ClientTxnOutcome { return ClientTxnOutcome(s.NewStruct(8, 3)) }
func NewRootClientTxnOutcome(s *C.Segment) ClientTxnOutcome {
	return ClientTxnOutcome(s.NewRootStruct(8, 3))
}
func AutoNewClientTxnOutcome(s *C.Segment) ClientTxnOutcome {
	return ClientTxnOutcome(s.NewStructAR(8, 3))
}
func ReadRootClientTxnOutcome(s *C.Segment) ClientTxnOutcome {
	return ClientTxnOutcome(s.Root(0).ToStruct())
}
func (s ClientTxnOutcome) Which() ClientTxnOutcome_Which {
	return ClientTxnOutcome_Which(C.Struct(s).Get16(0))
}
func (s ClientTxnOutcome) Id() []byte          { return C.Struct(s).GetObject(0).ToData() }
func (s ClientTxnOutcome) SetId(v []byte)      { C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s ClientTxnOutcome) FinalId() []byte     { return C.Struct(s).GetObject(1).ToData() }
func (s ClientTxnOutcome) SetFinalId(v []byte) { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s ClientTxnOutcome) SetCommit()          { C.Struct(s).Set16(0, 0) }
func (s ClientTxnOutcome) Abort() ClientUpdate_List {
	return ClientUpdate_List(C.Struct(s).GetObject(2))
}
func (s ClientTxnOutcome) SetAbort(v ClientUpdate_List) {
	C.Struct(s).Set16(0, 1)
	C.Struct(s).SetObject(2, C.Object(v))
}
func (s ClientTxnOutcome) Error() string      { return C.Struct(s).GetObject(2).ToText() }
func (s ClientTxnOutcome) ErrorBytes() []byte { return C.Struct(s).GetObject(2).ToDataTrimLastByte() }
func (s ClientTxnOutcome) SetError(v string) {
	C.Struct(s).Set16(0, 2)
	C.Struct(s).SetObject(2, s.Segment.NewText(v))
}
func (s ClientTxnOutcome) WriteJSON(w io.Writer) error {
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
	_, err = b.WriteString("\"finalId\":")
	if err != nil {
		return err
	}
	{
		s := s.FinalId()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	if s.Which() == CLIENTTXNOUTCOME_COMMIT {
		_, err = b.WriteString("\"commit\":")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == CLIENTTXNOUTCOME_ABORT {
		_, err = b.WriteString("\"abort\":")
		if err != nil {
			return err
		}
		{
			s := s.Abort()
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
	}
	if s.Which() == CLIENTTXNOUTCOME_ERROR {
		_, err = b.WriteString("\"error\":")
		if err != nil {
			return err
		}
		{
			s := s.Error()
			buf, err = json.Marshal(s)
			if err != nil {
				return err
			}
			_, err = b.Write(buf)
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
func (s ClientTxnOutcome) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s ClientTxnOutcome) WriteCapLit(w io.Writer) error {
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
	_, err = b.WriteString("finalId = ")
	if err != nil {
		return err
	}
	{
		s := s.FinalId()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	if s.Which() == CLIENTTXNOUTCOME_COMMIT {
		_, err = b.WriteString("commit = ")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == CLIENTTXNOUTCOME_ABORT {
		_, err = b.WriteString("abort = ")
		if err != nil {
			return err
		}
		{
			s := s.Abort()
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
	}
	if s.Which() == CLIENTTXNOUTCOME_ERROR {
		_, err = b.WriteString("error = ")
		if err != nil {
			return err
		}
		{
			s := s.Error()
			buf, err = json.Marshal(s)
			if err != nil {
				return err
			}
			_, err = b.Write(buf)
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
func (s ClientTxnOutcome) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type ClientTxnOutcome_List C.PointerList

func NewClientTxnOutcomeList(s *C.Segment, sz int) ClientTxnOutcome_List {
	return ClientTxnOutcome_List(s.NewCompositeList(8, 3, sz))
}
func (s ClientTxnOutcome_List) Len() int { return C.PointerList(s).Len() }
func (s ClientTxnOutcome_List) At(i int) ClientTxnOutcome {
	return ClientTxnOutcome(C.PointerList(s).At(i).ToStruct())
}
func (s ClientTxnOutcome_List) ToArray() []ClientTxnOutcome {
	n := s.Len()
	a := make([]ClientTxnOutcome, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s ClientTxnOutcome_List) Set(i int, item ClientTxnOutcome) {
	C.PointerList(s).Set(i, C.Object(item))
}

type ClientUpdate C.Struct

func NewClientUpdate(s *C.Segment) ClientUpdate       { return ClientUpdate(s.NewStruct(0, 2)) }
func NewRootClientUpdate(s *C.Segment) ClientUpdate   { return ClientUpdate(s.NewRootStruct(0, 2)) }
func AutoNewClientUpdate(s *C.Segment) ClientUpdate   { return ClientUpdate(s.NewStructAR(0, 2)) }
func ReadRootClientUpdate(s *C.Segment) ClientUpdate  { return ClientUpdate(s.Root(0).ToStruct()) }
func (s ClientUpdate) Version() []byte                { return C.Struct(s).GetObject(0).ToData() }
func (s ClientUpdate) SetVersion(v []byte)            { C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s ClientUpdate) Actions() ClientAction_List     { return ClientAction_List(C.Struct(s).GetObject(1)) }
func (s ClientUpdate) SetActions(v ClientAction_List) { C.Struct(s).SetObject(1, C.Object(v)) }
func (s ClientUpdate) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
	_, err = b.WriteString("\"actions\":")
	if err != nil {
		return err
	}
	{
		s := s.Actions()
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
func (s ClientUpdate) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s ClientUpdate) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
	_, err = b.WriteString("actions = ")
	if err != nil {
		return err
	}
	{
		s := s.Actions()
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
func (s ClientUpdate) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type ClientUpdate_List C.PointerList

func NewClientUpdateList(s *C.Segment, sz int) ClientUpdate_List {
	return ClientUpdate_List(s.NewCompositeList(0, 2, sz))
}
func (s ClientUpdate_List) Len() int { return C.PointerList(s).Len() }
func (s ClientUpdate_List) At(i int) ClientUpdate {
	return ClientUpdate(C.PointerList(s).At(i).ToStruct())
}
func (s ClientUpdate_List) ToArray() []ClientUpdate {
	n := s.Len()
	a := make([]ClientUpdate, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s ClientUpdate_List) Set(i int, item ClientUpdate) { C.PointerList(s).Set(i, C.Object(item)) }

type ClientVarIdPos C.Struct

func NewClientVarIdPos(s *C.Segment) ClientVarIdPos      { return ClientVarIdPos(s.NewStruct(0, 2)) }
func NewRootClientVarIdPos(s *C.Segment) ClientVarIdPos  { return ClientVarIdPos(s.NewRootStruct(0, 2)) }
func AutoNewClientVarIdPos(s *C.Segment) ClientVarIdPos  { return ClientVarIdPos(s.NewStructAR(0, 2)) }
func ReadRootClientVarIdPos(s *C.Segment) ClientVarIdPos { return ClientVarIdPos(s.Root(0).ToStruct()) }
func (s ClientVarIdPos) VarId() []byte                   { return C.Struct(s).GetObject(0).ToData() }
func (s ClientVarIdPos) SetVarId(v []byte)               { C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s ClientVarIdPos) Capability() Capability {
	return Capability(C.Struct(s).GetObject(1).ToStruct())
}
func (s ClientVarIdPos) SetCapability(v Capability) { C.Struct(s).SetObject(1, C.Object(v)) }
func (s ClientVarIdPos) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
	_, err = b.WriteString("\"capability\":")
	if err != nil {
		return err
	}
	{
		s := s.Capability()
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
func (s ClientVarIdPos) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s ClientVarIdPos) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
	_, err = b.WriteString("capability = ")
	if err != nil {
		return err
	}
	{
		s := s.Capability()
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
func (s ClientVarIdPos) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type ClientVarIdPos_List C.PointerList

func NewClientVarIdPosList(s *C.Segment, sz int) ClientVarIdPos_List {
	return ClientVarIdPos_List(s.NewCompositeList(0, 2, sz))
}
func (s ClientVarIdPos_List) Len() int { return C.PointerList(s).Len() }
func (s ClientVarIdPos_List) At(i int) ClientVarIdPos {
	return ClientVarIdPos(C.PointerList(s).At(i).ToStruct())
}
func (s ClientVarIdPos_List) ToArray() []ClientVarIdPos {
	n := s.Len()
	a := make([]ClientVarIdPos, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s ClientVarIdPos_List) Set(i int, item ClientVarIdPos) { C.PointerList(s).Set(i, C.Object(item)) }
