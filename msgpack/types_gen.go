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
		case "Modified":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Modified = nil
			} else {
				if z.Modified == nil {
					z.Modified = new(ClientActionModify)
				}
				err = z.Modified.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "ActionType":
			{
				var zcmr uint8
				zcmr, err = dc.ReadUint8()
				if err != nil {
					return
				}
				z.ActionType = ClientActionType(zcmr)
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
	// map header, size 3
	// write "VarId"
	err = en.Append(0x83, 0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.VarId)
	if err != nil {
		return
	}
	// write "Modified"
	err = en.Append(0xa8, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64)
	if err != nil {
		return err
	}
	if z.Modified == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Modified.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "ActionType"
	err = en.Append(0xaa, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteUint8(uint8(z.ActionType))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientAction) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "VarId"
	o = append(o, 0x83, 0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
	o = msgp.AppendBytes(o, z.VarId)
	// string "Modified"
	o = append(o, 0xa8, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64)
	if z.Modified == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Modified.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "ActionType"
	o = append(o, 0xaa, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendUint8(o, uint8(z.ActionType))
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
		case "Modified":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Modified = nil
			} else {
				if z.Modified == nil {
					z.Modified = new(ClientActionModify)
				}
				bts, err = z.Modified.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "ActionType":
			{
				var zwht uint8
				zwht, bts, err = msgp.ReadUint8Bytes(bts)
				if err != nil {
					return
				}
				z.ActionType = ClientActionType(zwht)
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
	s = 1 + 6 + msgp.BytesPrefixSize + len(z.VarId) + 9
	if z.Modified == nil {
		s += msgp.NilSize
	} else {
		s += z.Modified.Msgsize()
	}
	s += 11 + msgp.Uint8Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientActionModify) DecodeMsg(dc *msgp.Reader) (err error) {
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
func (z *ClientActionModify) EncodeMsg(en *msgp.Writer) (err error) {
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
func (z *ClientActionModify) MarshalMsg(b []byte) (o []byte, err error) {
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
func (z *ClientActionModify) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
func (z *ClientActionModify) Msgsize() (s int) {
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
func (z *ClientActionType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zpks uint8
		zpks, err = dc.ReadUint8()
		if err != nil {
			return
		}
		(*z) = ClientActionType(zpks)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z ClientActionType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint8(uint8(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ClientActionType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint8(o, uint8(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientActionType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zjfb uint8
		zjfb, bts, err = msgp.ReadUint8Bytes(bts)
		if err != nil {
			return
		}
		(*z) = ClientActionType(zjfb)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ClientActionType) Msgsize() (s int) {
	s = msgp.Uint8Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zcxo uint32
	zcxo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcxo > 0 {
		zcxo--
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
	var zeff uint32
	zeff, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zeff > 0 {
		zeff--
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
	var zxpk uint32
	zxpk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxpk > 0 {
		zxpk--
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
			var zdnj uint32
			zdnj, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Actions) >= int(zdnj) {
				z.Actions = (z.Actions)[:zdnj]
			} else {
				z.Actions = make([]*ClientAction, zdnj)
			}
			for zrsw := range z.Actions {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Actions[zrsw] = nil
				} else {
					if z.Actions[zrsw] == nil {
						z.Actions[zrsw] = new(ClientAction)
					}
					err = z.Actions[zrsw].DecodeMsg(dc)
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
	for zrsw := range z.Actions {
		if z.Actions[zrsw] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Actions[zrsw].EncodeMsg(en)
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
	for zrsw := range z.Actions {
		if z.Actions[zrsw] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Actions[zrsw].MarshalMsg(o)
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
	var zobc uint32
	zobc, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zobc > 0 {
		zobc--
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
			var zsnv uint32
			zsnv, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Actions) >= int(zsnv) {
				z.Actions = (z.Actions)[:zsnv]
			} else {
				z.Actions = make([]*ClientAction, zsnv)
			}
			for zrsw := range z.Actions {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Actions[zrsw] = nil
				} else {
					if z.Actions[zrsw] == nil {
						z.Actions[zrsw] = new(ClientAction)
					}
					bts, err = z.Actions[zrsw].UnmarshalMsg(bts)
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
	for zrsw := range z.Actions {
		if z.Actions[zrsw] == nil {
			s += msgp.NilSize
		} else {
			s += z.Actions[zrsw].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientTxnOutcome) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
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
			var zpez uint32
			zpez, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Abort) >= int(zpez) {
				z.Abort = (z.Abort)[:zpez]
			} else {
				z.Abort = make([]*ClientAction, zpez)
			}
			for zkgt := range z.Abort {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Abort[zkgt] = nil
				} else {
					if z.Abort[zkgt] == nil {
						z.Abort[zkgt] = new(ClientAction)
					}
					err = z.Abort[zkgt].DecodeMsg(dc)
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
	for zkgt := range z.Abort {
		if z.Abort[zkgt] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Abort[zkgt].EncodeMsg(en)
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
	for zkgt := range z.Abort {
		if z.Abort[zkgt] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Abort[zkgt].MarshalMsg(o)
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
			var zqyh uint32
			zqyh, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Abort) >= int(zqyh) {
				z.Abort = (z.Abort)[:zqyh]
			} else {
				z.Abort = make([]*ClientAction, zqyh)
			}
			for zkgt := range z.Abort {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Abort[zkgt] = nil
				} else {
					if z.Abort[zkgt] == nil {
						z.Abort[zkgt] = new(ClientAction)
					}
					bts, err = z.Abort[zkgt].UnmarshalMsg(bts)
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
	for zkgt := range z.Abort {
		if z.Abort[zkgt] == nil {
			s += msgp.NilSize
		} else {
			s += z.Abort[zkgt].Msgsize()
		}
	}
	s += 6 + msgp.StringPrefixSize + len(z.Error)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientVarId) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zyzr uint32
	zyzr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zyzr > 0 {
		zyzr--
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
				var zywj uint32
				zywj, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zywj > 0 {
					zywj--
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
	var zjpj uint32
	zjpj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjpj > 0 {
		zjpj--
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
	var zrfe uint32
	zrfe, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrfe > 0 {
		zrfe--
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
	var zgmo uint32
	zgmo, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zgmo > 0 {
		zgmo--
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
	var zeth uint32
	zeth, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zeth > 0 {
		zeth--
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
			var zsbz uint32
			zsbz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Roots) >= int(zsbz) {
				z.Roots = (z.Roots)[:zsbz]
			} else {
				z.Roots = make([]*Root, zsbz)
			}
			for ztaf := range z.Roots {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Roots[ztaf] = nil
				} else {
					if z.Roots[ztaf] == nil {
						z.Roots[ztaf] = new(Root)
					}
					err = z.Roots[ztaf].DecodeMsg(dc)
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
	for ztaf := range z.Roots {
		if z.Roots[ztaf] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Roots[ztaf].EncodeMsg(en)
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
	for ztaf := range z.Roots {
		if z.Roots[ztaf] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Roots[ztaf].MarshalMsg(o)
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
	var zrjx uint32
	zrjx, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrjx > 0 {
		zrjx--
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
			var zawn uint32
			zawn, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Roots) >= int(zawn) {
				z.Roots = (z.Roots)[:zawn]
			} else {
				z.Roots = make([]*Root, zawn)
			}
			for ztaf := range z.Roots {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Roots[ztaf] = nil
				} else {
					if z.Roots[ztaf] == nil {
						z.Roots[ztaf] = new(Root)
					}
					bts, err = z.Roots[ztaf].UnmarshalMsg(bts)
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
	for ztaf := range z.Roots {
		if z.Roots[ztaf] == nil {
			s += msgp.NilSize
		} else {
			s += z.Roots[ztaf].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Root) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zwel uint32
	zwel, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zwel > 0 {
		zwel--
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
				var zrbe uint32
				zrbe, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zrbe > 0 {
					zrbe--
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
	var zmfd uint32
	zmfd, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zmfd > 0 {
		zmfd--
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
				var zzdc uint32
				zzdc, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zzdc > 0 {
					zzdc--
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
