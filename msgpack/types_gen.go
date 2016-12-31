package msgpack

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Capability) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxvk uint32
	zxvk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxvk > 0 {
		zxvk--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Read":
			z.Read, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "Write":
			z.Write, err = dc.ReadBool()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Capability) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Read"
	err = en.Append(0x82, 0xa4, 0x52, 0x65, 0x61, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Read)
	if err != nil {
		return
	}
	// write "Write"
	err = en.Append(0xa5, 0x57, 0x72, 0x69, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Write)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Capability) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Read"
	o = append(o, 0x82, 0xa4, 0x52, 0x65, 0x61, 0x64)
	o = msgp.AppendBool(o, z.Read)
	// string "Write"
	o = append(o, 0xa5, 0x57, 0x72, 0x69, 0x74, 0x65)
	o = msgp.AppendBool(o, z.Write)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Capability) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Read":
			z.Read, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "Write":
			z.Write, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Capability) Msgsize() (s int) {
	s = 1 + 5 + msgp.BoolSize + 6 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientAction) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbai uint32
	zbai, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbai > 0 {
		zbai--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "VarId":
			z.VarId, err = dc.ReadBytes(z.VarId)
			if err != nil {
				return
			}
		case "Read":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Read = nil
			} else {
				if z.Read == nil {
					z.Read = new(ClientActionRead)
				}
				var zcmr uint32
				zcmr, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zcmr > 0 {
					zcmr--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Version":
						z.Read.Version, err = dc.ReadBytes(z.Read.Version)
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		case "Write":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Write = nil
			} else {
				if z.Write == nil {
					z.Write = new(ClientActionWrite)
				}
				err = z.Write.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "ReadWrite":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.ReadWrite = nil
			} else {
				if z.ReadWrite == nil {
					z.ReadWrite = new(ClientActionReadWrite)
				}
				err = z.ReadWrite.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Create":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Create = nil
			} else {
				if z.Create == nil {
					z.Create = new(ClientActionCreate)
				}
				err = z.Create.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Delete":
			z.Delete, err = dc.ReadBool()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ClientAction) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "VarId"
	err = en.Append(0x86, 0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.VarId)
	if err != nil {
		return
	}
	// write "Read"
	err = en.Append(0xa4, 0x52, 0x65, 0x61, 0x64)
	if err != nil {
		return err
	}
	if z.Read == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		// map header, size 1
		// write "Version"
		err = en.Append(0x81, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
		if err != nil {
			return err
		}
		err = en.WriteBytes(z.Read.Version)
		if err != nil {
			return
		}
	}
	// write "Write"
	err = en.Append(0xa5, 0x57, 0x72, 0x69, 0x74, 0x65)
	if err != nil {
		return err
	}
	if z.Write == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Write.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "ReadWrite"
	err = en.Append(0xa9, 0x52, 0x65, 0x61, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65)
	if err != nil {
		return err
	}
	if z.ReadWrite == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.ReadWrite.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Create"
	err = en.Append(0xa6, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65)
	if err != nil {
		return err
	}
	if z.Create == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Create.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Delete"
	err = en.Append(0xa6, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Delete)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientAction) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "VarId"
	o = append(o, 0x86, 0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
	o = msgp.AppendBytes(o, z.VarId)
	// string "Read"
	o = append(o, 0xa4, 0x52, 0x65, 0x61, 0x64)
	if z.Read == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 1
		// string "Version"
		o = append(o, 0x81, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
		o = msgp.AppendBytes(o, z.Read.Version)
	}
	// string "Write"
	o = append(o, 0xa5, 0x57, 0x72, 0x69, 0x74, 0x65)
	if z.Write == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Write.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "ReadWrite"
	o = append(o, 0xa9, 0x52, 0x65, 0x61, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65)
	if z.ReadWrite == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.ReadWrite.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Create"
	o = append(o, 0xa6, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65)
	if z.Create == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Create.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Delete"
	o = append(o, 0xa6, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65)
	o = msgp.AppendBool(o, z.Delete)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientAction) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zajw uint32
	zajw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zajw > 0 {
		zajw--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "VarId":
			z.VarId, bts, err = msgp.ReadBytesBytes(bts, z.VarId)
			if err != nil {
				return
			}
		case "Read":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Read = nil
			} else {
				if z.Read == nil {
					z.Read = new(ClientActionRead)
				}
				var zwht uint32
				zwht, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zwht > 0 {
					zwht--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Version":
						z.Read.Version, bts, err = msgp.ReadBytesBytes(bts, z.Read.Version)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "Write":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Write = nil
			} else {
				if z.Write == nil {
					z.Write = new(ClientActionWrite)
				}
				bts, err = z.Write.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "ReadWrite":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.ReadWrite = nil
			} else {
				if z.ReadWrite == nil {
					z.ReadWrite = new(ClientActionReadWrite)
				}
				bts, err = z.ReadWrite.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Create":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Create = nil
			} else {
				if z.Create == nil {
					z.Create = new(ClientActionCreate)
				}
				bts, err = z.Create.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Delete":
			z.Delete, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ClientAction) Msgsize() (s int) {
	s = 1 + 6 + msgp.BytesPrefixSize + len(z.VarId) + 5
	if z.Read == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 8 + msgp.BytesPrefixSize + len(z.Read.Version)
	}
	s += 6
	if z.Write == nil {
		s += msgp.NilSize
	} else {
		s += z.Write.Msgsize()
	}
	s += 10
	if z.ReadWrite == nil {
		s += msgp.NilSize
	} else {
		s += z.ReadWrite.Msgsize()
	}
	s += 7
	if z.Create == nil {
		s += msgp.NilSize
	} else {
		s += z.Create.Msgsize()
	}
	s += 7 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientActionCreate) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zcua uint32
	zcua, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcua > 0 {
		zcua--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Value":
			z.Value, err = dc.ReadBytes(z.Value)
			if err != nil {
				return
			}
		case "References":
			var zxhx uint32
			zxhx, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.References) >= int(zxhx) {
				z.References = (z.References)[:zxhx]
			} else {
				z.References = make([]*ClientVarId, zxhx)
			}
			for zhct := range z.References {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.References[zhct] = nil
				} else {
					if z.References[zhct] == nil {
						z.References[zhct] = new(ClientVarId)
					}
					err = z.References[zhct].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ClientActionCreate) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Value"
	err = en.Append(0x82, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Value)
	if err != nil {
		return
	}
	// write "References"
	err = en.Append(0xaa, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.References)))
	if err != nil {
		return
	}
	for zhct := range z.References {
		if z.References[zhct] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.References[zhct].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientActionCreate) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Value"
	o = append(o, 0x82, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendBytes(o, z.Value)
	// string "References"
	o = append(o, 0xaa, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.References)))
	for zhct := range z.References {
		if z.References[zhct] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.References[zhct].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientActionCreate) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zlqf uint32
	zlqf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zlqf > 0 {
		zlqf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Value":
			z.Value, bts, err = msgp.ReadBytesBytes(bts, z.Value)
			if err != nil {
				return
			}
		case "References":
			var zdaf uint32
			zdaf, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.References) >= int(zdaf) {
				z.References = (z.References)[:zdaf]
			} else {
				z.References = make([]*ClientVarId, zdaf)
			}
			for zhct := range z.References {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.References[zhct] = nil
				} else {
					if z.References[zhct] == nil {
						z.References[zhct] = new(ClientVarId)
					}
					bts, err = z.References[zhct].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ClientActionCreate) Msgsize() (s int) {
	s = 1 + 6 + msgp.BytesPrefixSize + len(z.Value) + 11 + msgp.ArrayHeaderSize
	for zhct := range z.References {
		if z.References[zhct] == nil {
			s += msgp.NilSize
		} else {
			s += z.References[zhct].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientActionRead) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zpks uint32
	zpks, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zpks > 0 {
		zpks--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Version":
			z.Version, err = dc.ReadBytes(z.Version)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ClientActionRead) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Version"
	err = en.Append(0x81, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Version)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientActionRead) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Version"
	o = append(o, 0x81, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendBytes(o, z.Version)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientActionRead) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjfb uint32
	zjfb, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjfb > 0 {
		zjfb--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Version":
			z.Version, bts, err = msgp.ReadBytesBytes(bts, z.Version)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ClientActionRead) Msgsize() (s int) {
	s = 1 + 8 + msgp.BytesPrefixSize + len(z.Version)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientActionReadWrite) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zeff uint32
	zeff, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zeff > 0 {
		zeff--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Version":
			z.Version, err = dc.ReadBytes(z.Version)
			if err != nil {
				return
			}
		case "Value":
			z.Value, err = dc.ReadBytes(z.Value)
			if err != nil {
				return
			}
		case "References":
			var zrsw uint32
			zrsw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.References) >= int(zrsw) {
				z.References = (z.References)[:zrsw]
			} else {
				z.References = make([]*ClientVarId, zrsw)
			}
			for zcxo := range z.References {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.References[zcxo] = nil
				} else {
					if z.References[zcxo] == nil {
						z.References[zcxo] = new(ClientVarId)
					}
					err = z.References[zcxo].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ClientActionReadWrite) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Version"
	err = en.Append(0x83, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Version)
	if err != nil {
		return
	}
	// write "Value"
	err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Value)
	if err != nil {
		return
	}
	// write "References"
	err = en.Append(0xaa, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.References)))
	if err != nil {
		return
	}
	for zcxo := range z.References {
		if z.References[zcxo] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.References[zcxo].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientActionReadWrite) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Version"
	o = append(o, 0x83, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendBytes(o, z.Version)
	// string "Value"
	o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendBytes(o, z.Value)
	// string "References"
	o = append(o, 0xaa, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.References)))
	for zcxo := range z.References {
		if z.References[zcxo] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.References[zcxo].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientActionReadWrite) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zxpk uint32
	zxpk, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zxpk > 0 {
		zxpk--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Version":
			z.Version, bts, err = msgp.ReadBytesBytes(bts, z.Version)
			if err != nil {
				return
			}
		case "Value":
			z.Value, bts, err = msgp.ReadBytesBytes(bts, z.Value)
			if err != nil {
				return
			}
		case "References":
			var zdnj uint32
			zdnj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.References) >= int(zdnj) {
				z.References = (z.References)[:zdnj]
			} else {
				z.References = make([]*ClientVarId, zdnj)
			}
			for zcxo := range z.References {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.References[zcxo] = nil
				} else {
					if z.References[zcxo] == nil {
						z.References[zcxo] = new(ClientVarId)
					}
					bts, err = z.References[zcxo].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ClientActionReadWrite) Msgsize() (s int) {
	s = 1 + 8 + msgp.BytesPrefixSize + len(z.Version) + 6 + msgp.BytesPrefixSize + len(z.Value) + 11 + msgp.ArrayHeaderSize
	for zcxo := range z.References {
		if z.References[zcxo] == nil {
			s += msgp.NilSize
		} else {
			s += z.References[zcxo].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientActionWrite) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zsnv uint32
	zsnv, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zsnv > 0 {
		zsnv--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Value":
			z.Value, err = dc.ReadBytes(z.Value)
			if err != nil {
				return
			}
		case "References":
			var zkgt uint32
			zkgt, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.References) >= int(zkgt) {
				z.References = (z.References)[:zkgt]
			} else {
				z.References = make([]*ClientVarId, zkgt)
			}
			for zobc := range z.References {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.References[zobc] = nil
				} else {
					if z.References[zobc] == nil {
						z.References[zobc] = new(ClientVarId)
					}
					var zema uint32
					zema, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zema > 0 {
						zema--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "VarId":
							z.References[zobc].VarId, err = dc.ReadBytes(z.References[zobc].VarId)
							if err != nil {
								return
							}
						case "Capability":
							if dc.IsNil() {
								err = dc.ReadNil()
								if err != nil {
									return
								}
								z.References[zobc].Capability = nil
							} else {
								if z.References[zobc].Capability == nil {
									z.References[zobc].Capability = new(Capability)
								}
								var zpez uint32
								zpez, err = dc.ReadMapHeader()
								if err != nil {
									return
								}
								for zpez > 0 {
									zpez--
									field, err = dc.ReadMapKeyPtr()
									if err != nil {
										return
									}
									switch msgp.UnsafeString(field) {
									case "Read":
										z.References[zobc].Capability.Read, err = dc.ReadBool()
										if err != nil {
											return
										}
									case "Write":
										z.References[zobc].Capability.Write, err = dc.ReadBool()
										if err != nil {
											return
										}
									default:
										err = dc.Skip()
										if err != nil {
											return
										}
									}
								}
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ClientActionWrite) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Value"
	err = en.Append(0x82, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Value)
	if err != nil {
		return
	}
	// write "References"
	err = en.Append(0xaa, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.References)))
	if err != nil {
		return
	}
	for zobc := range z.References {
		if z.References[zobc] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "VarId"
			err = en.Append(0x82, 0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
			if err != nil {
				return err
			}
			err = en.WriteBytes(z.References[zobc].VarId)
			if err != nil {
				return
			}
			// write "Capability"
			err = en.Append(0xaa, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79)
			if err != nil {
				return err
			}
			if z.References[zobc].Capability == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				// map header, size 2
				// write "Read"
				err = en.Append(0x82, 0xa4, 0x52, 0x65, 0x61, 0x64)
				if err != nil {
					return err
				}
				err = en.WriteBool(z.References[zobc].Capability.Read)
				if err != nil {
					return
				}
				// write "Write"
				err = en.Append(0xa5, 0x57, 0x72, 0x69, 0x74, 0x65)
				if err != nil {
					return err
				}
				err = en.WriteBool(z.References[zobc].Capability.Write)
				if err != nil {
					return
				}
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientActionWrite) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Value"
	o = append(o, 0x82, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendBytes(o, z.Value)
	// string "References"
	o = append(o, 0xaa, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.References)))
	for zobc := range z.References {
		if z.References[zobc] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "VarId"
			o = append(o, 0x82, 0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
			o = msgp.AppendBytes(o, z.References[zobc].VarId)
			// string "Capability"
			o = append(o, 0xaa, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79)
			if z.References[zobc].Capability == nil {
				o = msgp.AppendNil(o)
			} else {
				// map header, size 2
				// string "Read"
				o = append(o, 0x82, 0xa4, 0x52, 0x65, 0x61, 0x64)
				o = msgp.AppendBool(o, z.References[zobc].Capability.Read)
				// string "Write"
				o = append(o, 0xa5, 0x57, 0x72, 0x69, 0x74, 0x65)
				o = msgp.AppendBool(o, z.References[zobc].Capability.Write)
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientActionWrite) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zqke uint32
	zqke, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zqke > 0 {
		zqke--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Value":
			z.Value, bts, err = msgp.ReadBytesBytes(bts, z.Value)
			if err != nil {
				return
			}
		case "References":
			var zqyh uint32
			zqyh, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.References) >= int(zqyh) {
				z.References = (z.References)[:zqyh]
			} else {
				z.References = make([]*ClientVarId, zqyh)
			}
			for zobc := range z.References {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.References[zobc] = nil
				} else {
					if z.References[zobc] == nil {
						z.References[zobc] = new(ClientVarId)
					}
					var zyzr uint32
					zyzr, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zyzr > 0 {
						zyzr--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "VarId":
							z.References[zobc].VarId, bts, err = msgp.ReadBytesBytes(bts, z.References[zobc].VarId)
							if err != nil {
								return
							}
						case "Capability":
							if msgp.IsNil(bts) {
								bts, err = msgp.ReadNilBytes(bts)
								if err != nil {
									return
								}
								z.References[zobc].Capability = nil
							} else {
								if z.References[zobc].Capability == nil {
									z.References[zobc].Capability = new(Capability)
								}
								var zywj uint32
								zywj, bts, err = msgp.ReadMapHeaderBytes(bts)
								if err != nil {
									return
								}
								for zywj > 0 {
									zywj--
									field, bts, err = msgp.ReadMapKeyZC(bts)
									if err != nil {
										return
									}
									switch msgp.UnsafeString(field) {
									case "Read":
										z.References[zobc].Capability.Read, bts, err = msgp.ReadBoolBytes(bts)
										if err != nil {
											return
										}
									case "Write":
										z.References[zobc].Capability.Write, bts, err = msgp.ReadBoolBytes(bts)
										if err != nil {
											return
										}
									default:
										bts, err = msgp.Skip(bts)
										if err != nil {
											return
										}
									}
								}
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ClientActionWrite) Msgsize() (s int) {
	s = 1 + 6 + msgp.BytesPrefixSize + len(z.Value) + 11 + msgp.ArrayHeaderSize
	for zobc := range z.References {
		if z.References[zobc] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 6 + msgp.BytesPrefixSize + len(z.References[zobc].VarId) + 11
			if z.References[zobc].Capability == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 5 + msgp.BoolSize + 6 + msgp.BoolSize
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zjpj uint32
	zjpj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zjpj > 0 {
		zjpj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ClientTxnSubmission":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.ClientTxnSubmission = nil
			} else {
				if z.ClientTxnSubmission == nil {
					z.ClientTxnSubmission = new(ClientTxn)
				}
				err = z.ClientTxnSubmission.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "ClientTxnOutcome":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.ClientTxnOutcome = nil
			} else {
				if z.ClientTxnOutcome == nil {
					z.ClientTxnOutcome = new(ClientTxnOutcome)
				}
				err = z.ClientTxnOutcome.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ClientMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "ClientTxnSubmission"
	err = en.Append(0x82, 0xb3, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x78, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	if z.ClientTxnSubmission == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.ClientTxnSubmission.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "ClientTxnOutcome"
	err = en.Append(0xb0, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x78, 0x6e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.ClientTxnOutcome == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.ClientTxnOutcome.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "ClientTxnSubmission"
	o = append(o, 0x82, 0xb3, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x78, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e)
	if z.ClientTxnSubmission == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.ClientTxnSubmission.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "ClientTxnOutcome"
	o = append(o, 0xb0, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x78, 0x6e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65)
	if z.ClientTxnOutcome == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.ClientTxnOutcome.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zzpf uint32
	zzpf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zzpf > 0 {
		zzpf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ClientTxnSubmission":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.ClientTxnSubmission = nil
			} else {
				if z.ClientTxnSubmission == nil {
					z.ClientTxnSubmission = new(ClientTxn)
				}
				bts, err = z.ClientTxnSubmission.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "ClientTxnOutcome":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.ClientTxnOutcome = nil
			} else {
				if z.ClientTxnOutcome == nil {
					z.ClientTxnOutcome = new(ClientTxnOutcome)
				}
				bts, err = z.ClientTxnOutcome.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ClientMessage) Msgsize() (s int) {
	s = 1 + 20
	if z.ClientTxnSubmission == nil {
		s += msgp.NilSize
	} else {
		s += z.ClientTxnSubmission.Msgsize()
	}
	s += 17
	if z.ClientTxnOutcome == nil {
		s += msgp.NilSize
	} else {
		s += z.ClientTxnOutcome.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientTxn) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zgmo uint32
	zgmo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zgmo > 0 {
		zgmo--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, err = dc.ReadBytes(z.Id)
			if err != nil {
				return
			}
		case "Retry":
			z.Retry, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "Actions":
			var ztaf uint32
			ztaf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Actions) >= int(ztaf) {
				z.Actions = (z.Actions)[:ztaf]
			} else {
				z.Actions = make([]*ClientAction, ztaf)
			}
			for zrfe := range z.Actions {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Actions[zrfe] = nil
				} else {
					if z.Actions[zrfe] == nil {
						z.Actions[zrfe] = new(ClientAction)
					}
					err = z.Actions[zrfe].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ClientTxn) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Id"
	err = en.Append(0x83, 0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Id)
	if err != nil {
		return
	}
	// write "Retry"
	err = en.Append(0xa5, 0x52, 0x65, 0x74, 0x72, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Retry)
	if err != nil {
		return
	}
	// write "Actions"
	err = en.Append(0xa7, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Actions)))
	if err != nil {
		return
	}
	for zrfe := range z.Actions {
		if z.Actions[zrfe] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Actions[zrfe].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientTxn) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Id"
	o = append(o, 0x83, 0xa2, 0x49, 0x64)
	o = msgp.AppendBytes(o, z.Id)
	// string "Retry"
	o = append(o, 0xa5, 0x52, 0x65, 0x74, 0x72, 0x79)
	o = msgp.AppendBool(o, z.Retry)
	// string "Actions"
	o = append(o, 0xa7, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Actions)))
	for zrfe := range z.Actions {
		if z.Actions[zrfe] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Actions[zrfe].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientTxn) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zeth uint32
	zeth, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zeth > 0 {
		zeth--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, bts, err = msgp.ReadBytesBytes(bts, z.Id)
			if err != nil {
				return
			}
		case "Retry":
			z.Retry, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "Actions":
			var zsbz uint32
			zsbz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Actions) >= int(zsbz) {
				z.Actions = (z.Actions)[:zsbz]
			} else {
				z.Actions = make([]*ClientAction, zsbz)
			}
			for zrfe := range z.Actions {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Actions[zrfe] = nil
				} else {
					if z.Actions[zrfe] == nil {
						z.Actions[zrfe] = new(ClientAction)
					}
					bts, err = z.Actions[zrfe].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ClientTxn) Msgsize() (s int) {
	s = 1 + 3 + msgp.BytesPrefixSize + len(z.Id) + 6 + msgp.BoolSize + 8 + msgp.ArrayHeaderSize
	for zrfe := range z.Actions {
		if z.Actions[zrfe] == nil {
			s += msgp.NilSize
		} else {
			s += z.Actions[zrfe].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientTxnOutcome) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zawn uint32
	zawn, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zawn > 0 {
		zawn--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, err = dc.ReadBytes(z.Id)
			if err != nil {
				return
			}
		case "FinalId":
			z.FinalId, err = dc.ReadBytes(z.FinalId)
			if err != nil {
				return
			}
		case "Commit":
			z.Commit, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "Abort":
			var zwel uint32
			zwel, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Abort) >= int(zwel) {
				z.Abort = (z.Abort)[:zwel]
			} else {
				z.Abort = make([]*ClientUpdate, zwel)
			}
			for zrjx := range z.Abort {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Abort[zrjx] = nil
				} else {
					if z.Abort[zrjx] == nil {
						z.Abort[zrjx] = new(ClientUpdate)
					}
					err = z.Abort[zrjx].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "Error":
			z.Error, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ClientTxnOutcome) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "Id"
	err = en.Append(0x85, 0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Id)
	if err != nil {
		return
	}
	// write "FinalId"
	err = en.Append(0xa7, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.FinalId)
	if err != nil {
		return
	}
	// write "Commit"
	err = en.Append(0xa6, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Commit)
	if err != nil {
		return
	}
	// write "Abort"
	err = en.Append(0xa5, 0x41, 0x62, 0x6f, 0x72, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Abort)))
	if err != nil {
		return
	}
	for zrjx := range z.Abort {
		if z.Abort[zrjx] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Abort[zrjx].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "Error"
	err = en.Append(0xa5, 0x45, 0x72, 0x72, 0x6f, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Error)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientTxnOutcome) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Id"
	o = append(o, 0x85, 0xa2, 0x49, 0x64)
	o = msgp.AppendBytes(o, z.Id)
	// string "FinalId"
	o = append(o, 0xa7, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x64)
	o = msgp.AppendBytes(o, z.FinalId)
	// string "Commit"
	o = append(o, 0xa6, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74)
	o = msgp.AppendBool(o, z.Commit)
	// string "Abort"
	o = append(o, 0xa5, 0x41, 0x62, 0x6f, 0x72, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Abort)))
	for zrjx := range z.Abort {
		if z.Abort[zrjx] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Abort[zrjx].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "Error"
	o = append(o, 0xa5, 0x45, 0x72, 0x72, 0x6f, 0x72)
	o = msgp.AppendString(o, z.Error)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientTxnOutcome) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zrbe uint32
	zrbe, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrbe > 0 {
		zrbe--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, bts, err = msgp.ReadBytesBytes(bts, z.Id)
			if err != nil {
				return
			}
		case "FinalId":
			z.FinalId, bts, err = msgp.ReadBytesBytes(bts, z.FinalId)
			if err != nil {
				return
			}
		case "Commit":
			z.Commit, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "Abort":
			var zmfd uint32
			zmfd, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Abort) >= int(zmfd) {
				z.Abort = (z.Abort)[:zmfd]
			} else {
				z.Abort = make([]*ClientUpdate, zmfd)
			}
			for zrjx := range z.Abort {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Abort[zrjx] = nil
				} else {
					if z.Abort[zrjx] == nil {
						z.Abort[zrjx] = new(ClientUpdate)
					}
					bts, err = z.Abort[zrjx].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "Error":
			z.Error, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ClientTxnOutcome) Msgsize() (s int) {
	s = 1 + 3 + msgp.BytesPrefixSize + len(z.Id) + 8 + msgp.BytesPrefixSize + len(z.FinalId) + 7 + msgp.BoolSize + 6 + msgp.ArrayHeaderSize
	for zrjx := range z.Abort {
		if z.Abort[zrjx] == nil {
			s += msgp.NilSize
		} else {
			s += z.Abort[zrjx].Msgsize()
		}
	}
	s += 6 + msgp.StringPrefixSize + len(z.Error)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientUpdate) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zelx uint32
	zelx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zelx > 0 {
		zelx--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Version":
			z.Version, err = dc.ReadBytes(z.Version)
			if err != nil {
				return
			}
		case "Actions":
			var zbal uint32
			zbal, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Actions) >= int(zbal) {
				z.Actions = (z.Actions)[:zbal]
			} else {
				z.Actions = make([]*ClientAction, zbal)
			}
			for zzdc := range z.Actions {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Actions[zzdc] = nil
				} else {
					if z.Actions[zzdc] == nil {
						z.Actions[zzdc] = new(ClientAction)
					}
					err = z.Actions[zzdc].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ClientUpdate) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Version"
	err = en.Append(0x82, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Version)
	if err != nil {
		return
	}
	// write "Actions"
	err = en.Append(0xa7, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Actions)))
	if err != nil {
		return
	}
	for zzdc := range z.Actions {
		if z.Actions[zzdc] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Actions[zzdc].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientUpdate) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Version"
	o = append(o, 0x82, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendBytes(o, z.Version)
	// string "Actions"
	o = append(o, 0xa7, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Actions)))
	for zzdc := range z.Actions {
		if z.Actions[zzdc] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Actions[zzdc].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientUpdate) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjqz uint32
	zjqz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjqz > 0 {
		zjqz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Version":
			z.Version, bts, err = msgp.ReadBytesBytes(bts, z.Version)
			if err != nil {
				return
			}
		case "Actions":
			var zkct uint32
			zkct, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Actions) >= int(zkct) {
				z.Actions = (z.Actions)[:zkct]
			} else {
				z.Actions = make([]*ClientAction, zkct)
			}
			for zzdc := range z.Actions {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Actions[zzdc] = nil
				} else {
					if z.Actions[zzdc] == nil {
						z.Actions[zzdc] = new(ClientAction)
					}
					bts, err = z.Actions[zzdc].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ClientUpdate) Msgsize() (s int) {
	s = 1 + 8 + msgp.BytesPrefixSize + len(z.Version) + 8 + msgp.ArrayHeaderSize
	for zzdc := range z.Actions {
		if z.Actions[zzdc] == nil {
			s += msgp.NilSize
		} else {
			s += z.Actions[zzdc].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientVarId) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var ztmt uint32
	ztmt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for ztmt > 0 {
		ztmt--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "VarId":
			z.VarId, err = dc.ReadBytes(z.VarId)
			if err != nil {
				return
			}
		case "Capability":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Capability = nil
			} else {
				if z.Capability == nil {
					z.Capability = new(Capability)
				}
				var ztco uint32
				ztco, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for ztco > 0 {
					ztco--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Read":
						z.Capability.Read, err = dc.ReadBool()
						if err != nil {
							return
						}
					case "Write":
						z.Capability.Write, err = dc.ReadBool()
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ClientVarId) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "VarId"
	err = en.Append(0x82, 0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.VarId)
	if err != nil {
		return
	}
	// write "Capability"
	err = en.Append(0xaa, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79)
	if err != nil {
		return err
	}
	if z.Capability == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		// map header, size 2
		// write "Read"
		err = en.Append(0x82, 0xa4, 0x52, 0x65, 0x61, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.Capability.Read)
		if err != nil {
			return
		}
		// write "Write"
		err = en.Append(0xa5, 0x57, 0x72, 0x69, 0x74, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.Capability.Write)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientVarId) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "VarId"
	o = append(o, 0x82, 0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
	o = msgp.AppendBytes(o, z.VarId)
	// string "Capability"
	o = append(o, 0xaa, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79)
	if z.Capability == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 2
		// string "Read"
		o = append(o, 0x82, 0xa4, 0x52, 0x65, 0x61, 0x64)
		o = msgp.AppendBool(o, z.Capability.Read)
		// string "Write"
		o = append(o, 0xa5, 0x57, 0x72, 0x69, 0x74, 0x65)
		o = msgp.AppendBool(o, z.Capability.Write)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientVarId) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zana uint32
	zana, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zana > 0 {
		zana--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "VarId":
			z.VarId, bts, err = msgp.ReadBytesBytes(bts, z.VarId)
			if err != nil {
				return
			}
		case "Capability":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Capability = nil
			} else {
				if z.Capability == nil {
					z.Capability = new(Capability)
				}
				var ztyy uint32
				ztyy, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for ztyy > 0 {
					ztyy--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Read":
						z.Capability.Read, bts, err = msgp.ReadBoolBytes(bts)
						if err != nil {
							return
						}
					case "Write":
						z.Capability.Write, bts, err = msgp.ReadBoolBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ClientVarId) Msgsize() (s int) {
	s = 1 + 6 + msgp.BytesPrefixSize + len(z.VarId) + 11
	if z.Capability == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 5 + msgp.BoolSize + 6 + msgp.BoolSize
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Hello) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zinl uint32
	zinl, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zinl > 0 {
		zinl--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Product":
			z.Product, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Version":
			z.Version, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Hello) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Product"
	err = en.Append(0x82, 0xa7, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Product)
	if err != nil {
		return
	}
	// write "Version"
	err = en.Append(0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Version)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Hello) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Product"
	o = append(o, 0x82, 0xa7, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74)
	o = msgp.AppendString(o, z.Product)
	// string "Version"
	o = append(o, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Version)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Hello) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zare uint32
	zare, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zare > 0 {
		zare--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Product":
			z.Product, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Version":
			z.Version, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Hello) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.Product) + 8 + msgp.StringPrefixSize + len(z.Version)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *HelloClientFromServer) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zixj uint32
	zixj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zixj > 0 {
		zixj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Namespace":
			z.Namespace, err = dc.ReadBytes(z.Namespace)
			if err != nil {
				return
			}
		case "Roots":
			var zrsc uint32
			zrsc, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Roots) >= int(zrsc) {
				z.Roots = (z.Roots)[:zrsc]
			} else {
				z.Roots = make([]*Root, zrsc)
			}
			for zljy := range z.Roots {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Roots[zljy] = nil
				} else {
					if z.Roots[zljy] == nil {
						z.Roots[zljy] = new(Root)
					}
					err = z.Roots[zljy].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *HelloClientFromServer) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Namespace"
	err = en.Append(0x82, 0xa9, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Namespace)
	if err != nil {
		return
	}
	// write "Roots"
	err = en.Append(0xa5, 0x52, 0x6f, 0x6f, 0x74, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Roots)))
	if err != nil {
		return
	}
	for zljy := range z.Roots {
		if z.Roots[zljy] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Roots[zljy].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *HelloClientFromServer) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Namespace"
	o = append(o, 0x82, 0xa9, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65)
	o = msgp.AppendBytes(o, z.Namespace)
	// string "Roots"
	o = append(o, 0xa5, 0x52, 0x6f, 0x6f, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Roots)))
	for zljy := range z.Roots {
		if z.Roots[zljy] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Roots[zljy].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *HelloClientFromServer) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zctn uint32
	zctn, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zctn > 0 {
		zctn--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Namespace":
			z.Namespace, bts, err = msgp.ReadBytesBytes(bts, z.Namespace)
			if err != nil {
				return
			}
		case "Roots":
			var zswy uint32
			zswy, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Roots) >= int(zswy) {
				z.Roots = (z.Roots)[:zswy]
			} else {
				z.Roots = make([]*Root, zswy)
			}
			for zljy := range z.Roots {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Roots[zljy] = nil
				} else {
					if z.Roots[zljy] == nil {
						z.Roots[zljy] = new(Root)
					}
					bts, err = z.Roots[zljy].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *HelloClientFromServer) Msgsize() (s int) {
	s = 1 + 10 + msgp.BytesPrefixSize + len(z.Namespace) + 6 + msgp.ArrayHeaderSize
	for zljy := range z.Roots {
		if z.Roots[zljy] == nil {
			s += msgp.NilSize
		} else {
			s += z.Roots[zljy].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Root) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var znsg uint32
	znsg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for znsg > 0 {
		znsg--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "VarId":
			z.VarId, err = dc.ReadBytes(z.VarId)
			if err != nil {
				return
			}
		case "Capability":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Capability = nil
			} else {
				if z.Capability == nil {
					z.Capability = new(Capability)
				}
				var zrus uint32
				zrus, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zrus > 0 {
					zrus--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Read":
						z.Capability.Read, err = dc.ReadBool()
						if err != nil {
							return
						}
					case "Write":
						z.Capability.Write, err = dc.ReadBool()
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Root) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Name"
	err = en.Append(0x83, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "VarId"
	err = en.Append(0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.VarId)
	if err != nil {
		return
	}
	// write "Capability"
	err = en.Append(0xaa, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79)
	if err != nil {
		return err
	}
	if z.Capability == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		// map header, size 2
		// write "Read"
		err = en.Append(0x82, 0xa4, 0x52, 0x65, 0x61, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.Capability.Read)
		if err != nil {
			return
		}
		// write "Write"
		err = en.Append(0xa5, 0x57, 0x72, 0x69, 0x74, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.Capability.Write)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Root) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Name"
	o = append(o, 0x83, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "VarId"
	o = append(o, 0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
	o = msgp.AppendBytes(o, z.VarId)
	// string "Capability"
	o = append(o, 0xaa, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79)
	if z.Capability == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 2
		// string "Read"
		o = append(o, 0x82, 0xa4, 0x52, 0x65, 0x61, 0x64)
		o = msgp.AppendBool(o, z.Capability.Read)
		// string "Write"
		o = append(o, 0xa5, 0x57, 0x72, 0x69, 0x74, 0x65)
		o = msgp.AppendBool(o, z.Capability.Write)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Root) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsvm uint32
	zsvm, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsvm > 0 {
		zsvm--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "VarId":
			z.VarId, bts, err = msgp.ReadBytesBytes(bts, z.VarId)
			if err != nil {
				return
			}
		case "Capability":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Capability = nil
			} else {
				if z.Capability == nil {
					z.Capability = new(Capability)
				}
				var zaoz uint32
				zaoz, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zaoz > 0 {
					zaoz--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Read":
						z.Capability.Read, bts, err = msgp.ReadBoolBytes(bts)
						if err != nil {
							return
						}
					case "Write":
						z.Capability.Write, bts, err = msgp.ReadBoolBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Root) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 6 + msgp.BytesPrefixSize + len(z.VarId) + 11
	if z.Capability == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 5 + msgp.BoolSize + 6 + msgp.BoolSize
	}
	return
}
