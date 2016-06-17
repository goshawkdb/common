package capnp

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	"encoding/json"
	C "github.com/glycerine/go-capnproto"
	"io"
)

type Capabilities C.Struct
type CapabilitiesReferences Capabilities
type CapabilitiesReferencesRead Capabilities
type CapabilitiesReferencesWrite Capabilities
type CapabilitiesReferencesRead_Which uint16

const (
	CAPABILITIESREFERENCESREAD_ALL  CapabilitiesReferencesRead_Which = 0
	CAPABILITIESREFERENCESREAD_ONLY CapabilitiesReferencesRead_Which = 1
)

type CapabilitiesReferencesWrite_Which uint16

const (
	CAPABILITIESREFERENCESWRITE_ALL  CapabilitiesReferencesWrite_Which = 0
	CAPABILITIESREFERENCESWRITE_ONLY CapabilitiesReferencesWrite_Which = 1
)

func NewCapabilities(s *C.Segment) Capabilities           { return Capabilities(s.NewStruct(8, 2)) }
func NewRootCapabilities(s *C.Segment) Capabilities       { return Capabilities(s.NewRootStruct(8, 2)) }
func AutoNewCapabilities(s *C.Segment) Capabilities       { return Capabilities(s.NewStructAR(8, 2)) }
func ReadRootCapabilities(s *C.Segment) Capabilities      { return Capabilities(s.Root(0).ToStruct()) }
func (s Capabilities) Value() ValueCapability             { return ValueCapability(C.Struct(s).Get16(0)) }
func (s Capabilities) SetValue(v ValueCapability)         { C.Struct(s).Set16(0, uint16(v)) }
func (s Capabilities) References() CapabilitiesReferences { return CapabilitiesReferences(s) }
func (s CapabilitiesReferences) Read() CapabilitiesReferencesRead {
	return CapabilitiesReferencesRead(s)
}
func (s CapabilitiesReferencesRead) Which() CapabilitiesReferencesRead_Which {
	return CapabilitiesReferencesRead_Which(C.Struct(s).Get16(2))
}
func (s CapabilitiesReferencesRead) SetAll()            { C.Struct(s).Set16(2, 0) }
func (s CapabilitiesReferencesRead) Only() C.UInt32List { return C.UInt32List(C.Struct(s).GetObject(0)) }
func (s CapabilitiesReferencesRead) SetOnly(v C.UInt32List) {
	C.Struct(s).Set16(2, 1)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s CapabilitiesReferences) Write() CapabilitiesReferencesWrite {
	return CapabilitiesReferencesWrite(s)
}
func (s CapabilitiesReferencesWrite) Which() CapabilitiesReferencesWrite_Which {
	return CapabilitiesReferencesWrite_Which(C.Struct(s).Get16(4))
}
func (s CapabilitiesReferencesWrite) SetAll() { C.Struct(s).Set16(4, 0) }
func (s CapabilitiesReferencesWrite) Only() C.UInt32List {
	return C.UInt32List(C.Struct(s).GetObject(1))
}
func (s CapabilitiesReferencesWrite) SetOnly(v C.UInt32List) {
	C.Struct(s).Set16(4, 1)
	C.Struct(s).SetObject(1, C.Object(v))
}
func (s Capabilities) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
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
		err = s.WriteJSON(b)
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
			err = b.WriteByte('{')
			if err != nil {
				return err
			}
			if s.Which() == CAPABILITIESREFERENCESREAD_ALL {
				_, err = b.WriteString("\"all\":")
				if err != nil {
					return err
				}
				_ = s
				_, err = b.WriteString("null")
				if err != nil {
					return err
				}
			}
			if s.Which() == CAPABILITIESREFERENCESREAD_ONLY {
				_, err = b.WriteString("\"only\":")
				if err != nil {
					return err
				}
				{
					s := s.Only()
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
			if s.Which() == CAPABILITIESREFERENCESWRITE_ALL {
				_, err = b.WriteString("\"all\":")
				if err != nil {
					return err
				}
				_ = s
				_, err = b.WriteString("null")
				if err != nil {
					return err
				}
			}
			if s.Which() == CAPABILITIESREFERENCESWRITE_ONLY {
				_, err = b.WriteString("\"only\":")
				if err != nil {
					return err
				}
				{
					s := s.Only()
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
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Capabilities) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Capabilities) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
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
		err = s.WriteCapLit(b)
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
			err = b.WriteByte('(')
			if err != nil {
				return err
			}
			if s.Which() == CAPABILITIESREFERENCESREAD_ALL {
				_, err = b.WriteString("all = ")
				if err != nil {
					return err
				}
				_ = s
				_, err = b.WriteString("null")
				if err != nil {
					return err
				}
			}
			if s.Which() == CAPABILITIESREFERENCESREAD_ONLY {
				_, err = b.WriteString("only = ")
				if err != nil {
					return err
				}
				{
					s := s.Only()
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
			if s.Which() == CAPABILITIESREFERENCESWRITE_ALL {
				_, err = b.WriteString("all = ")
				if err != nil {
					return err
				}
				_ = s
				_, err = b.WriteString("null")
				if err != nil {
					return err
				}
			}
			if s.Which() == CAPABILITIESREFERENCESWRITE_ONLY {
				_, err = b.WriteString("only = ")
				if err != nil {
					return err
				}
				{
					s := s.Only()
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
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Capabilities) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Capabilities_List C.PointerList

func NewCapabilitiesList(s *C.Segment, sz int) Capabilities_List {
	return Capabilities_List(s.NewCompositeList(8, 2, sz))
}
func (s Capabilities_List) Len() int { return C.PointerList(s).Len() }
func (s Capabilities_List) At(i int) Capabilities {
	return Capabilities(C.PointerList(s).At(i).ToStruct())
}
func (s Capabilities_List) ToArray() []Capabilities {
	n := s.Len()
	a := make([]Capabilities, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Capabilities_List) Set(i int, item Capabilities) { C.PointerList(s).Set(i, C.Object(item)) }

type ValueCapability uint16

const (
	VALUECAPABILITY_NONE      ValueCapability = 0
	VALUECAPABILITY_READ      ValueCapability = 1
	VALUECAPABILITY_WRITE     ValueCapability = 2
	VALUECAPABILITY_READWRITE ValueCapability = 3
)

func (c ValueCapability) String() string {
	switch c {
	case VALUECAPABILITY_NONE:
		return "none"
	case VALUECAPABILITY_READ:
		return "read"
	case VALUECAPABILITY_WRITE:
		return "write"
	case VALUECAPABILITY_READWRITE:
		return "readWrite"
	default:
		return ""
	}
}

func ValueCapabilityFromString(c string) ValueCapability {
	switch c {
	case "none":
		return VALUECAPABILITY_NONE
	case "read":
		return VALUECAPABILITY_READ
	case "write":
		return VALUECAPABILITY_WRITE
	case "readWrite":
		return VALUECAPABILITY_READWRITE
	default:
		return 0
	}
}

type ValueCapability_List C.PointerList

func NewValueCapabilityList(s *C.Segment, sz int) ValueCapability_List {
	return ValueCapability_List(s.NewUInt16List(sz))
}
func (s ValueCapability_List) Len() int                 { return C.UInt16List(s).Len() }
func (s ValueCapability_List) At(i int) ValueCapability { return ValueCapability(C.UInt16List(s).At(i)) }
func (s ValueCapability_List) ToArray() []ValueCapability {
	n := s.Len()
	a := make([]ValueCapability, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s ValueCapability_List) Set(i int, item ValueCapability) { C.UInt16List(s).Set(i, uint16(item)) }
func (s ValueCapability) WriteJSON(w io.Writer) error {
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
func (s ValueCapability) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s ValueCapability) WriteCapLit(w io.Writer) error {
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
func (s ValueCapability) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}
