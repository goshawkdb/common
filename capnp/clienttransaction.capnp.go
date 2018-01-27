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
func (s ClientTxn) Counter() uint32                { return C.Struct(s).Get32(0) }
func (s ClientTxn) SetCounter(v uint32)            { C.Struct(s).Set32(0, v) }
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
	_, err = b.WriteString("\"counter\":")
	if err != nil {
		return err
	}
	{
		s := s.Counter()
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
	_, err = b.WriteString("counter = ")
	if err != nil {
		return err
	}
	{
		s := s.Counter()
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
type ClientActionValue ClientAction
type ClientActionValueCreate ClientAction
type ClientActionValueExisting ClientAction
type ClientActionValueExistingModify ClientAction
type ClientActionValueExistingModifyWrite ClientAction
type ClientActionMeta ClientAction
type ClientActionValue_Which uint16

const (
	CLIENTACTIONVALUE_MISSING  ClientActionValue_Which = 0
	CLIENTACTIONVALUE_CREATE   ClientActionValue_Which = 1
	CLIENTACTIONVALUE_EXISTING ClientActionValue_Which = 2
)

type ClientActionValueExistingModify_Which uint16

const (
	CLIENTACTIONVALUEEXISTINGMODIFY_NOT   ClientActionValueExistingModify_Which = 0
	CLIENTACTIONVALUEEXISTINGMODIFY_ROLL  ClientActionValueExistingModify_Which = 1
	CLIENTACTIONVALUEEXISTINGMODIFY_WRITE ClientActionValueExistingModify_Which = 2
)

func NewClientAction(s *C.Segment) ClientAction      { return ClientAction(s.NewStruct(8, 4)) }
func NewRootClientAction(s *C.Segment) ClientAction  { return ClientAction(s.NewRootStruct(8, 4)) }
func AutoNewClientAction(s *C.Segment) ClientAction  { return ClientAction(s.NewStructAR(8, 4)) }
func ReadRootClientAction(s *C.Segment) ClientAction { return ClientAction(s.Root(0).ToStruct()) }
func (s ClientAction) VarId() []byte                 { return C.Struct(s).GetObject(0).ToData() }
func (s ClientAction) SetVarId(v []byte)             { C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s ClientAction) Value() ClientActionValue      { return ClientActionValue(s) }
func (s ClientActionValue) Which() ClientActionValue_Which {
	return ClientActionValue_Which(C.Struct(s).Get16(0))
}
func (s ClientActionValue) SetMissing()                     { C.Struct(s).Set16(0, 0) }
func (s ClientActionValue) Create() ClientActionValueCreate { return ClientActionValueCreate(s) }
func (s ClientActionValue) SetCreate()                      { C.Struct(s).Set16(0, 1) }
func (s ClientActionValueCreate) Value() []byte             { return C.Struct(s).GetObject(1).ToData() }
func (s ClientActionValueCreate) SetValue(v []byte)         { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s ClientActionValueCreate) References() ClientVarIdPos_List {
	return ClientVarIdPos_List(C.Struct(s).GetObject(2))
}
func (s ClientActionValueCreate) SetReferences(v ClientVarIdPos_List) {
	C.Struct(s).SetObject(2, C.Object(v))
}
func (s ClientActionValue) Existing() ClientActionValueExisting { return ClientActionValueExisting(s) }
func (s ClientActionValue) SetExisting()                        { C.Struct(s).Set16(0, 2) }
func (s ClientActionValueExisting) Read() bool                  { return C.Struct(s).Get1(16) }
func (s ClientActionValueExisting) SetRead(v bool)              { C.Struct(s).Set1(16, v) }
func (s ClientActionValueExisting) Modify() ClientActionValueExistingModify {
	return ClientActionValueExistingModify(s)
}
func (s ClientActionValueExistingModify) Which() ClientActionValueExistingModify_Which {
	return ClientActionValueExistingModify_Which(C.Struct(s).Get16(4))
}
func (s ClientActionValueExistingModify) SetNot()  { C.Struct(s).Set16(4, 0) }
func (s ClientActionValueExistingModify) SetRoll() { C.Struct(s).Set16(4, 1) }
func (s ClientActionValueExistingModify) Write() ClientActionValueExistingModifyWrite {
	return ClientActionValueExistingModifyWrite(s)
}
func (s ClientActionValueExistingModify) SetWrite()          { C.Struct(s).Set16(4, 2) }
func (s ClientActionValueExistingModifyWrite) Value() []byte { return C.Struct(s).GetObject(1).ToData() }
func (s ClientActionValueExistingModifyWrite) SetValue(v []byte) {
	C.Struct(s).SetObject(1, s.Segment.NewData(v))
}
func (s ClientActionValueExistingModifyWrite) References() ClientVarIdPos_List {
	return ClientVarIdPos_List(C.Struct(s).GetObject(2))
}
func (s ClientActionValueExistingModifyWrite) SetReferences(v ClientVarIdPos_List) {
	C.Struct(s).SetObject(2, C.Object(v))
}
func (s ClientAction) Meta() ClientActionMeta { return ClientActionMeta(s) }
func (s ClientActionMeta) AddSub() bool       { return C.Struct(s).Get1(17) }
func (s ClientActionMeta) SetAddSub(v bool)   { C.Struct(s).Set1(17, v) }
func (s ClientActionMeta) DelSub() []byte     { return C.Struct(s).GetObject(3).ToData() }
func (s ClientActionMeta) SetDelSub(v []byte) { C.Struct(s).SetObject(3, s.Segment.NewData(v)) }
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
		err = b.WriteByte('{')
		if err != nil {
			return err
		}
		if s.Which() == CLIENTACTIONVALUE_MISSING {
			_, err = b.WriteString("\"missing\":")
			if err != nil {
				return err
			}
			_ = s
			_, err = b.WriteString("null")
			if err != nil {
				return err
			}
		}
		if s.Which() == CLIENTACTIONVALUE_CREATE {
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
		if s.Which() == CLIENTACTIONVALUE_EXISTING {
			_, err = b.WriteString("\"existing\":")
			if err != nil {
				return err
			}
			{
				s := s.Existing()
				err = b.WriteByte('{')
				if err != nil {
					return err
				}
				_, err = b.WriteString("\"read\":")
				if err != nil {
					return err
				}
				{
					s := s.Read()
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
				_, err = b.WriteString("\"modify\":")
				if err != nil {
					return err
				}
				{
					s := s.Modify()
					err = b.WriteByte('{')
					if err != nil {
						return err
					}
					if s.Which() == CLIENTACTIONVALUEEXISTINGMODIFY_NOT {
						_, err = b.WriteString("\"not\":")
						if err != nil {
							return err
						}
						_ = s
						_, err = b.WriteString("null")
						if err != nil {
							return err
						}
					}
					if s.Which() == CLIENTACTIONVALUEEXISTINGMODIFY_ROLL {
						_, err = b.WriteString("\"roll\":")
						if err != nil {
							return err
						}
						_ = s
						_, err = b.WriteString("null")
						if err != nil {
							return err
						}
					}
					if s.Which() == CLIENTACTIONVALUEEXISTINGMODIFY_WRITE {
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
					err = b.WriteByte('}')
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
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"meta\":")
	if err != nil {
		return err
	}
	{
		s := s.Meta()
		err = b.WriteByte('{')
		if err != nil {
			return err
		}
		_, err = b.WriteString("\"addSub\":")
		if err != nil {
			return err
		}
		{
			s := s.AddSub()
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
		_, err = b.WriteString("\"delSub\":")
		if err != nil {
			return err
		}
		{
			s := s.DelSub()
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
		err = b.WriteByte('(')
		if err != nil {
			return err
		}
		if s.Which() == CLIENTACTIONVALUE_MISSING {
			_, err = b.WriteString("missing = ")
			if err != nil {
				return err
			}
			_ = s
			_, err = b.WriteString("null")
			if err != nil {
				return err
			}
		}
		if s.Which() == CLIENTACTIONVALUE_CREATE {
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
		if s.Which() == CLIENTACTIONVALUE_EXISTING {
			_, err = b.WriteString("existing = ")
			if err != nil {
				return err
			}
			{
				s := s.Existing()
				err = b.WriteByte('(')
				if err != nil {
					return err
				}
				_, err = b.WriteString("read = ")
				if err != nil {
					return err
				}
				{
					s := s.Read()
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
				_, err = b.WriteString("modify = ")
				if err != nil {
					return err
				}
				{
					s := s.Modify()
					err = b.WriteByte('(')
					if err != nil {
						return err
					}
					if s.Which() == CLIENTACTIONVALUEEXISTINGMODIFY_NOT {
						_, err = b.WriteString("not = ")
						if err != nil {
							return err
						}
						_ = s
						_, err = b.WriteString("null")
						if err != nil {
							return err
						}
					}
					if s.Which() == CLIENTACTIONVALUEEXISTINGMODIFY_ROLL {
						_, err = b.WriteString("roll = ")
						if err != nil {
							return err
						}
						_ = s
						_, err = b.WriteString("null")
						if err != nil {
							return err
						}
					}
					if s.Which() == CLIENTACTIONVALUEEXISTINGMODIFY_WRITE {
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
					err = b.WriteByte(')')
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
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("meta = ")
	if err != nil {
		return err
	}
	{
		s := s.Meta()
		err = b.WriteByte('(')
		if err != nil {
			return err
		}
		_, err = b.WriteString("addSub = ")
		if err != nil {
			return err
		}
		{
			s := s.AddSub()
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
		_, err = b.WriteString("delSub = ")
		if err != nil {
			return err
		}
		{
			s := s.DelSub()
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
	return ClientTxnOutcome_Which(C.Struct(s).Get16(4))
}
func (s ClientTxnOutcome) Id() []byte          { return C.Struct(s).GetObject(0).ToData() }
func (s ClientTxnOutcome) SetId(v []byte)      { C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s ClientTxnOutcome) FinalId() []byte     { return C.Struct(s).GetObject(1).ToData() }
func (s ClientTxnOutcome) SetFinalId(v []byte) { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s ClientTxnOutcome) Commit() uint32      { return C.Struct(s).Get32(0) }
func (s ClientTxnOutcome) SetCommit(v uint32)  { C.Struct(s).Set16(4, 0); C.Struct(s).Set32(0, v) }
func (s ClientTxnOutcome) Abort() ClientAction_List {
	return ClientAction_List(C.Struct(s).GetObject(2))
}
func (s ClientTxnOutcome) SetAbort(v ClientAction_List) {
	C.Struct(s).Set16(4, 1)
	C.Struct(s).SetObject(2, C.Object(v))
}
func (s ClientTxnOutcome) Error() string      { return C.Struct(s).GetObject(2).ToText() }
func (s ClientTxnOutcome) ErrorBytes() []byte { return C.Struct(s).GetObject(2).ToDataTrimLastByte() }
func (s ClientTxnOutcome) SetError(v string) {
	C.Struct(s).Set16(4, 2)
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
		{
			s := s.Commit()
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
		{
			s := s.Commit()
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
