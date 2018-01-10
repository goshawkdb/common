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
type ClientActionModified ClientAction
type ClientAction_Which uint16

const (
	CLIENTACTION_UNMODIFIED ClientAction_Which = 0
	CLIENTACTION_MODIFIED   ClientAction_Which = 1
)

func NewClientAction(s *C.Segment) ClientAction       { return ClientAction(s.NewStruct(8, 3)) }
func NewRootClientAction(s *C.Segment) ClientAction   { return ClientAction(s.NewRootStruct(8, 3)) }
func AutoNewClientAction(s *C.Segment) ClientAction   { return ClientAction(s.NewStructAR(8, 3)) }
func ReadRootClientAction(s *C.Segment) ClientAction  { return ClientAction(s.Root(0).ToStruct()) }
func (s ClientAction) Which() ClientAction_Which      { return ClientAction_Which(C.Struct(s).Get16(0)) }
func (s ClientAction) VarId() []byte                  { return C.Struct(s).GetObject(0).ToData() }
func (s ClientAction) SetVarId(v []byte)              { C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s ClientAction) SetUnmodified()                 { C.Struct(s).Set16(0, 0) }
func (s ClientAction) Modified() ClientActionModified { return ClientActionModified(s) }
func (s ClientAction) SetModified()                   { C.Struct(s).Set16(0, 1) }
func (s ClientActionModified) Value() []byte          { return C.Struct(s).GetObject(1).ToData() }
func (s ClientActionModified) SetValue(v []byte)      { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s ClientActionModified) References() ClientVarIdPos_List {
	return ClientVarIdPos_List(C.Struct(s).GetObject(2))
}
func (s ClientActionModified) SetReferences(v ClientVarIdPos_List) {
	C.Struct(s).SetObject(2, C.Object(v))
}
func (s ClientAction) ActionType() ClientActionType     { return ClientActionType(C.Struct(s).Get16(2)) }
func (s ClientAction) SetActionType(v ClientActionType) { C.Struct(s).Set16(2, uint16(v)) }
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
	if s.Which() == CLIENTACTION_UNMODIFIED {
		_, err = b.WriteString("\"unmodified\":")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == CLIENTACTION_MODIFIED {
		_, err = b.WriteString("\"modified\":")
		if err != nil {
			return err
		}
		{
			s := s.Modified()
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
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"actionType\":")
	if err != nil {
		return err
	}
	{
		s := s.ActionType()
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
	if s.Which() == CLIENTACTION_UNMODIFIED {
		_, err = b.WriteString("unmodified = ")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == CLIENTACTION_MODIFIED {
		_, err = b.WriteString("modified = ")
		if err != nil {
			return err
		}
		{
			s := s.Modified()
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
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("actionType = ")
	if err != nil {
		return err
	}
	{
		s := s.ActionType()
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
func (s ClientAction) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type ClientAction_List C.PointerList

func NewClientActionList(s *C.Segment, sz int) ClientAction_List {
	return ClientAction_List(s.NewCompositeList(8, 3, sz))
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

type ClientActionType uint16

const (
	CLIENTACTIONTYPE_CREATE          ClientActionType = 0
	CLIENTACTIONTYPE_READONLY        ClientActionType = 1
	CLIENTACTIONTYPE_WRITEONLY       ClientActionType = 2
	CLIENTACTIONTYPE_READWRITE       ClientActionType = 3
	CLIENTACTIONTYPE_DELETE          ClientActionType = 4
	CLIENTACTIONTYPE_ROLL            ClientActionType = 5
	CLIENTACTIONTYPE_ADDSUBSCRIPTION ClientActionType = 6
	CLIENTACTIONTYPE_DELSUBSCRIPTION ClientActionType = 7
)

func (c ClientActionType) String() string {
	switch c {
	case CLIENTACTIONTYPE_CREATE:
		return "create"
	case CLIENTACTIONTYPE_READONLY:
		return "readOnly"
	case CLIENTACTIONTYPE_WRITEONLY:
		return "writeOnly"
	case CLIENTACTIONTYPE_READWRITE:
		return "readWrite"
	case CLIENTACTIONTYPE_DELETE:
		return "delete"
	case CLIENTACTIONTYPE_ROLL:
		return "roll"
	case CLIENTACTIONTYPE_ADDSUBSCRIPTION:
		return "addSubscription"
	case CLIENTACTIONTYPE_DELSUBSCRIPTION:
		return "delSubscription"
	default:
		return ""
	}
}

func ClientActionTypeFromString(c string) ClientActionType {
	switch c {
	case "create":
		return CLIENTACTIONTYPE_CREATE
	case "readOnly":
		return CLIENTACTIONTYPE_READONLY
	case "writeOnly":
		return CLIENTACTIONTYPE_WRITEONLY
	case "readWrite":
		return CLIENTACTIONTYPE_READWRITE
	case "delete":
		return CLIENTACTIONTYPE_DELETE
	case "roll":
		return CLIENTACTIONTYPE_ROLL
	case "addSubscription":
		return CLIENTACTIONTYPE_ADDSUBSCRIPTION
	case "delSubscription":
		return CLIENTACTIONTYPE_DELSUBSCRIPTION
	default:
		return 0
	}
}

type ClientActionType_List C.PointerList

func NewClientActionTypeList(s *C.Segment, sz int) ClientActionType_List {
	return ClientActionType_List(s.NewUInt16List(sz))
}
func (s ClientActionType_List) Len() int { return C.UInt16List(s).Len() }
func (s ClientActionType_List) At(i int) ClientActionType {
	return ClientActionType(C.UInt16List(s).At(i))
}
func (s ClientActionType_List) ToArray() []ClientActionType {
	n := s.Len()
	a := make([]ClientActionType, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s ClientActionType_List) Set(i int, item ClientActionType) { C.UInt16List(s).Set(i, uint16(item)) }
func (s ClientActionType) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	buf, err = json.Marshal(s.String())
	if err != nil {
		return err
	}
	_, err = b.Write(buf)
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s ClientActionType) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s ClientActionType) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	_, err = b.WriteString(s.String())
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s ClientActionType) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

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
func (s ClientTxnOutcome) Abort() ClientAction_List {
	return ClientAction_List(C.Struct(s).GetObject(2))
}
func (s ClientTxnOutcome) SetAbort(v ClientAction_List) {
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
