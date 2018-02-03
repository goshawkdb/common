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
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
		return
	}
	err = en.WriteBool(z.Read)
	if err != nil {
		return
	}
	// write "Write"
	err = en.Append(0xa5, 0x57, 0x72, 0x69, 0x74, 0x65)
	if err != nil {
		return
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
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
		case "Value":
			err = z.Value.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "Meta":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "AddSub":
					z.Meta.AddSub, err = dc.ReadBool()
					if err != nil {
						return
					}
				case "DelSub":
					z.Meta.DelSub, err = dc.ReadBytes(z.Meta.DelSub)
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
		return
	}
	err = en.WriteBytes(z.VarId)
	if err != nil {
		return
	}
	// write "Value"
	err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return
	}
	err = z.Value.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "Meta"
	// map header, size 2
	// write "AddSub"
	err = en.Append(0xa4, 0x4d, 0x65, 0x74, 0x61, 0x82, 0xa6, 0x41, 0x64, 0x64, 0x53, 0x75, 0x62)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Meta.AddSub)
	if err != nil {
		return
	}
	// write "DelSub"
	err = en.Append(0xa6, 0x44, 0x65, 0x6c, 0x53, 0x75, 0x62)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Meta.DelSub)
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
	// string "Value"
	o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o, err = z.Value.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "Meta"
	// map header, size 2
	// string "AddSub"
	o = append(o, 0xa4, 0x4d, 0x65, 0x74, 0x61, 0x82, 0xa6, 0x41, 0x64, 0x64, 0x53, 0x75, 0x62)
	o = msgp.AppendBool(o, z.Meta.AddSub)
	// string "DelSub"
	o = append(o, 0xa6, 0x44, 0x65, 0x6c, 0x53, 0x75, 0x62)
	o = msgp.AppendBytes(o, z.Meta.DelSub)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientAction) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
		case "Value":
			bts, err = z.Value.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "Meta":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "AddSub":
					z.Meta.AddSub, bts, err = msgp.ReadBoolBytes(bts)
					if err != nil {
						return
					}
				case "DelSub":
					z.Meta.DelSub, bts, err = msgp.ReadBytesBytes(bts, z.Meta.DelSub)
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
	s = 1 + 6 + msgp.BytesPrefixSize + len(z.VarId) + 6 + z.Value.Msgsize() + 5 + 1 + 7 + msgp.BoolSize + 7 + msgp.BytesPrefixSize + len(z.Meta.DelSub)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientActionMeta) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "AddSub":
			z.AddSub, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "DelSub":
			z.DelSub, err = dc.ReadBytes(z.DelSub)
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
func (z *ClientActionMeta) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "AddSub"
	err = en.Append(0x82, 0xa6, 0x41, 0x64, 0x64, 0x53, 0x75, 0x62)
	if err != nil {
		return
	}
	err = en.WriteBool(z.AddSub)
	if err != nil {
		return
	}
	// write "DelSub"
	err = en.Append(0xa6, 0x44, 0x65, 0x6c, 0x53, 0x75, 0x62)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.DelSub)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientActionMeta) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "AddSub"
	o = append(o, 0x82, 0xa6, 0x41, 0x64, 0x64, 0x53, 0x75, 0x62)
	o = msgp.AppendBool(o, z.AddSub)
	// string "DelSub"
	o = append(o, 0xa6, 0x44, 0x65, 0x6c, 0x53, 0x75, 0x62)
	o = msgp.AppendBytes(o, z.DelSub)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientActionMeta) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "AddSub":
			z.AddSub, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "DelSub":
			z.DelSub, bts, err = msgp.ReadBytesBytes(bts, z.DelSub)
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
func (z *ClientActionMeta) Msgsize() (s int) {
	s = 1 + 7 + msgp.BoolSize + 7 + msgp.BytesPrefixSize + len(z.DelSub)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientActionValue) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Missing":
			z.Missing, err = dc.ReadBool()
			if err != nil {
				return
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
					z.Create = new(ClientActionValueCreate)
				}
				err = z.Create.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Existing":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Existing = nil
			} else {
				if z.Existing == nil {
					z.Existing = new(ClientActionValueExisting)
				}
				var zb0002 uint32
				zb0002, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zb0002 > 0 {
					zb0002--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Read":
						z.Existing.Read, err = dc.ReadBool()
						if err != nil {
							return
						}
					case "Modify":
						if dc.IsNil() {
							err = dc.ReadNil()
							if err != nil {
								return
							}
							z.Existing.Modify = nil
						} else {
							if z.Existing.Modify == nil {
								z.Existing.Modify = new(ClientActionValueExistingModify)
							}
							err = z.Existing.Modify.DecodeMsg(dc)
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
func (z *ClientActionValue) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Missing"
	err = en.Append(0x83, 0xa7, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Missing)
	if err != nil {
		return
	}
	// write "Create"
	err = en.Append(0xa6, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65)
	if err != nil {
		return
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
	// write "Existing"
	err = en.Append(0xa8, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67)
	if err != nil {
		return
	}
	if z.Existing == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		// map header, size 2
		// write "Read"
		err = en.Append(0x82, 0xa4, 0x52, 0x65, 0x61, 0x64)
		if err != nil {
			return
		}
		err = en.WriteBool(z.Existing.Read)
		if err != nil {
			return
		}
		// write "Modify"
		err = en.Append(0xa6, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79)
		if err != nil {
			return
		}
		if z.Existing.Modify == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Existing.Modify.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientActionValue) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Missing"
	o = append(o, 0x83, 0xa7, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67)
	o = msgp.AppendBool(o, z.Missing)
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
	// string "Existing"
	o = append(o, 0xa8, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67)
	if z.Existing == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 2
		// string "Read"
		o = append(o, 0x82, 0xa4, 0x52, 0x65, 0x61, 0x64)
		o = msgp.AppendBool(o, z.Existing.Read)
		// string "Modify"
		o = append(o, 0xa6, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79)
		if z.Existing.Modify == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Existing.Modify.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientActionValue) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Missing":
			z.Missing, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
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
					z.Create = new(ClientActionValueCreate)
				}
				bts, err = z.Create.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Existing":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Existing = nil
			} else {
				if z.Existing == nil {
					z.Existing = new(ClientActionValueExisting)
				}
				var zb0002 uint32
				zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0002 > 0 {
					zb0002--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Read":
						z.Existing.Read, bts, err = msgp.ReadBoolBytes(bts)
						if err != nil {
							return
						}
					case "Modify":
						if msgp.IsNil(bts) {
							bts, err = msgp.ReadNilBytes(bts)
							if err != nil {
								return
							}
							z.Existing.Modify = nil
						} else {
							if z.Existing.Modify == nil {
								z.Existing.Modify = new(ClientActionValueExistingModify)
							}
							bts, err = z.Existing.Modify.UnmarshalMsg(bts)
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
func (z *ClientActionValue) Msgsize() (s int) {
	s = 1 + 8 + msgp.BoolSize + 7
	if z.Create == nil {
		s += msgp.NilSize
	} else {
		s += z.Create.Msgsize()
	}
	s += 9
	if z.Existing == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 5 + msgp.BoolSize + 7
		if z.Existing.Modify == nil {
			s += msgp.NilSize
		} else {
			s += z.Existing.Modify.Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientActionValueCreate) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.References) >= int(zb0002) {
				z.References = (z.References)[:zb0002]
			} else {
				z.References = make([]*ClientVarId, zb0002)
			}
			for za0001 := range z.References {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.References[za0001] = nil
				} else {
					if z.References[za0001] == nil {
						z.References[za0001] = new(ClientVarId)
					}
					var zb0003 uint32
					zb0003, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zb0003 > 0 {
						zb0003--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "VarId":
							z.References[za0001].VarId, err = dc.ReadBytes(z.References[za0001].VarId)
							if err != nil {
								return
							}
						case "Capability":
							if dc.IsNil() {
								err = dc.ReadNil()
								if err != nil {
									return
								}
								z.References[za0001].Capability = nil
							} else {
								if z.References[za0001].Capability == nil {
									z.References[za0001].Capability = new(Capability)
								}
								var zb0004 uint32
								zb0004, err = dc.ReadMapHeader()
								if err != nil {
									return
								}
								for zb0004 > 0 {
									zb0004--
									field, err = dc.ReadMapKeyPtr()
									if err != nil {
										return
									}
									switch msgp.UnsafeString(field) {
									case "Read":
										z.References[za0001].Capability.Read, err = dc.ReadBool()
										if err != nil {
											return
										}
									case "Write":
										z.References[za0001].Capability.Write, err = dc.ReadBool()
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
func (z *ClientActionValueCreate) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Value"
	err = en.Append(0x82, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Value)
	if err != nil {
		return
	}
	// write "References"
	err = en.Append(0xaa, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.References)))
	if err != nil {
		return
	}
	for za0001 := range z.References {
		if z.References[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "VarId"
			err = en.Append(0x82, 0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
			if err != nil {
				return
			}
			err = en.WriteBytes(z.References[za0001].VarId)
			if err != nil {
				return
			}
			// write "Capability"
			err = en.Append(0xaa, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79)
			if err != nil {
				return
			}
			if z.References[za0001].Capability == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				// map header, size 2
				// write "Read"
				err = en.Append(0x82, 0xa4, 0x52, 0x65, 0x61, 0x64)
				if err != nil {
					return
				}
				err = en.WriteBool(z.References[za0001].Capability.Read)
				if err != nil {
					return
				}
				// write "Write"
				err = en.Append(0xa5, 0x57, 0x72, 0x69, 0x74, 0x65)
				if err != nil {
					return
				}
				err = en.WriteBool(z.References[za0001].Capability.Write)
				if err != nil {
					return
				}
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientActionValueCreate) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Value"
	o = append(o, 0x82, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendBytes(o, z.Value)
	// string "References"
	o = append(o, 0xaa, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.References)))
	for za0001 := range z.References {
		if z.References[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "VarId"
			o = append(o, 0x82, 0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
			o = msgp.AppendBytes(o, z.References[za0001].VarId)
			// string "Capability"
			o = append(o, 0xaa, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79)
			if z.References[za0001].Capability == nil {
				o = msgp.AppendNil(o)
			} else {
				// map header, size 2
				// string "Read"
				o = append(o, 0x82, 0xa4, 0x52, 0x65, 0x61, 0x64)
				o = msgp.AppendBool(o, z.References[za0001].Capability.Read)
				// string "Write"
				o = append(o, 0xa5, 0x57, 0x72, 0x69, 0x74, 0x65)
				o = msgp.AppendBool(o, z.References[za0001].Capability.Write)
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientActionValueCreate) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.References) >= int(zb0002) {
				z.References = (z.References)[:zb0002]
			} else {
				z.References = make([]*ClientVarId, zb0002)
			}
			for za0001 := range z.References {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.References[za0001] = nil
				} else {
					if z.References[za0001] == nil {
						z.References[za0001] = new(ClientVarId)
					}
					var zb0003 uint32
					zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zb0003 > 0 {
						zb0003--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "VarId":
							z.References[za0001].VarId, bts, err = msgp.ReadBytesBytes(bts, z.References[za0001].VarId)
							if err != nil {
								return
							}
						case "Capability":
							if msgp.IsNil(bts) {
								bts, err = msgp.ReadNilBytes(bts)
								if err != nil {
									return
								}
								z.References[za0001].Capability = nil
							} else {
								if z.References[za0001].Capability == nil {
									z.References[za0001].Capability = new(Capability)
								}
								var zb0004 uint32
								zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
								if err != nil {
									return
								}
								for zb0004 > 0 {
									zb0004--
									field, bts, err = msgp.ReadMapKeyZC(bts)
									if err != nil {
										return
									}
									switch msgp.UnsafeString(field) {
									case "Read":
										z.References[za0001].Capability.Read, bts, err = msgp.ReadBoolBytes(bts)
										if err != nil {
											return
										}
									case "Write":
										z.References[za0001].Capability.Write, bts, err = msgp.ReadBoolBytes(bts)
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
func (z *ClientActionValueCreate) Msgsize() (s int) {
	s = 1 + 6 + msgp.BytesPrefixSize + len(z.Value) + 11 + msgp.ArrayHeaderSize
	for za0001 := range z.References {
		if z.References[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 6 + msgp.BytesPrefixSize + len(z.References[za0001].VarId) + 11
			if z.References[za0001].Capability == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 5 + msgp.BoolSize + 6 + msgp.BoolSize
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientActionValueExisting) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
		case "Modify":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Modify = nil
			} else {
				if z.Modify == nil {
					z.Modify = new(ClientActionValueExistingModify)
				}
				err = z.Modify.DecodeMsg(dc)
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
func (z *ClientActionValueExisting) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Read"
	err = en.Append(0x82, 0xa4, 0x52, 0x65, 0x61, 0x64)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Read)
	if err != nil {
		return
	}
	// write "Modify"
	err = en.Append(0xa6, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79)
	if err != nil {
		return
	}
	if z.Modify == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Modify.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientActionValueExisting) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Read"
	o = append(o, 0x82, 0xa4, 0x52, 0x65, 0x61, 0x64)
	o = msgp.AppendBool(o, z.Read)
	// string "Modify"
	o = append(o, 0xa6, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79)
	if z.Modify == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Modify.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientActionValueExisting) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
		case "Modify":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Modify = nil
			} else {
				if z.Modify == nil {
					z.Modify = new(ClientActionValueExistingModify)
				}
				bts, err = z.Modify.UnmarshalMsg(bts)
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
func (z *ClientActionValueExisting) Msgsize() (s int) {
	s = 1 + 5 + msgp.BoolSize + 7
	if z.Modify == nil {
		s += msgp.NilSize
	} else {
		s += z.Modify.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientActionValueExistingModify) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.References) >= int(zb0002) {
				z.References = (z.References)[:zb0002]
			} else {
				z.References = make([]*ClientVarId, zb0002)
			}
			for za0001 := range z.References {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.References[za0001] = nil
				} else {
					if z.References[za0001] == nil {
						z.References[za0001] = new(ClientVarId)
					}
					err = z.References[za0001].DecodeMsg(dc)
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
func (z *ClientActionValueExistingModify) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Value"
	err = en.Append(0x82, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Value)
	if err != nil {
		return
	}
	// write "References"
	err = en.Append(0xaa, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.References)))
	if err != nil {
		return
	}
	for za0001 := range z.References {
		if z.References[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.References[za0001].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientActionValueExistingModify) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Value"
	o = append(o, 0x82, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendBytes(o, z.Value)
	// string "References"
	o = append(o, 0xaa, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.References)))
	for za0001 := range z.References {
		if z.References[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.References[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientActionValueExistingModify) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.References) >= int(zb0002) {
				z.References = (z.References)[:zb0002]
			} else {
				z.References = make([]*ClientVarId, zb0002)
			}
			for za0001 := range z.References {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.References[za0001] = nil
				} else {
					if z.References[za0001] == nil {
						z.References[za0001] = new(ClientVarId)
					}
					bts, err = z.References[za0001].UnmarshalMsg(bts)
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
func (z *ClientActionValueExistingModify) Msgsize() (s int) {
	s = 1 + 6 + msgp.BytesPrefixSize + len(z.Value) + 11 + msgp.ArrayHeaderSize
	for za0001 := range z.References {
		if z.References[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.References[za0001].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
		return
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
		return
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
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
		case "Counter":
			z.Counter, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "Actions":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Actions) >= int(zb0002) {
				z.Actions = (z.Actions)[:zb0002]
			} else {
				z.Actions = make([]*ClientAction, zb0002)
			}
			for za0001 := range z.Actions {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Actions[za0001] = nil
				} else {
					if z.Actions[za0001] == nil {
						z.Actions[za0001] = new(ClientAction)
					}
					var zb0003 uint32
					zb0003, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zb0003 > 0 {
						zb0003--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "VarId":
							z.Actions[za0001].VarId, err = dc.ReadBytes(z.Actions[za0001].VarId)
							if err != nil {
								return
							}
						case "Value":
							err = z.Actions[za0001].Value.DecodeMsg(dc)
							if err != nil {
								return
							}
						case "Meta":
							var zb0004 uint32
							zb0004, err = dc.ReadMapHeader()
							if err != nil {
								return
							}
							for zb0004 > 0 {
								zb0004--
								field, err = dc.ReadMapKeyPtr()
								if err != nil {
									return
								}
								switch msgp.UnsafeString(field) {
								case "AddSub":
									z.Actions[za0001].Meta.AddSub, err = dc.ReadBool()
									if err != nil {
										return
									}
								case "DelSub":
									z.Actions[za0001].Meta.DelSub, err = dc.ReadBytes(z.Actions[za0001].Meta.DelSub)
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
func (z *ClientTxn) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Id"
	err = en.Append(0x83, 0xa2, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Id)
	if err != nil {
		return
	}
	// write "Counter"
	err = en.Append(0xa7, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Counter)
	if err != nil {
		return
	}
	// write "Actions"
	err = en.Append(0xa7, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Actions)))
	if err != nil {
		return
	}
	for za0001 := range z.Actions {
		if z.Actions[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 3
			// write "VarId"
			err = en.Append(0x83, 0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
			if err != nil {
				return
			}
			err = en.WriteBytes(z.Actions[za0001].VarId)
			if err != nil {
				return
			}
			// write "Value"
			err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return
			}
			err = z.Actions[za0001].Value.EncodeMsg(en)
			if err != nil {
				return
			}
			// write "Meta"
			// map header, size 2
			// write "AddSub"
			err = en.Append(0xa4, 0x4d, 0x65, 0x74, 0x61, 0x82, 0xa6, 0x41, 0x64, 0x64, 0x53, 0x75, 0x62)
			if err != nil {
				return
			}
			err = en.WriteBool(z.Actions[za0001].Meta.AddSub)
			if err != nil {
				return
			}
			// write "DelSub"
			err = en.Append(0xa6, 0x44, 0x65, 0x6c, 0x53, 0x75, 0x62)
			if err != nil {
				return
			}
			err = en.WriteBytes(z.Actions[za0001].Meta.DelSub)
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
	// string "Counter"
	o = append(o, 0xa7, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72)
	o = msgp.AppendUint64(o, z.Counter)
	// string "Actions"
	o = append(o, 0xa7, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Actions)))
	for za0001 := range z.Actions {
		if z.Actions[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 3
			// string "VarId"
			o = append(o, 0x83, 0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
			o = msgp.AppendBytes(o, z.Actions[za0001].VarId)
			// string "Value"
			o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
			o, err = z.Actions[za0001].Value.MarshalMsg(o)
			if err != nil {
				return
			}
			// string "Meta"
			// map header, size 2
			// string "AddSub"
			o = append(o, 0xa4, 0x4d, 0x65, 0x74, 0x61, 0x82, 0xa6, 0x41, 0x64, 0x64, 0x53, 0x75, 0x62)
			o = msgp.AppendBool(o, z.Actions[za0001].Meta.AddSub)
			// string "DelSub"
			o = append(o, 0xa6, 0x44, 0x65, 0x6c, 0x53, 0x75, 0x62)
			o = msgp.AppendBytes(o, z.Actions[za0001].Meta.DelSub)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientTxn) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
		case "Counter":
			z.Counter, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "Actions":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Actions) >= int(zb0002) {
				z.Actions = (z.Actions)[:zb0002]
			} else {
				z.Actions = make([]*ClientAction, zb0002)
			}
			for za0001 := range z.Actions {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Actions[za0001] = nil
				} else {
					if z.Actions[za0001] == nil {
						z.Actions[za0001] = new(ClientAction)
					}
					var zb0003 uint32
					zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zb0003 > 0 {
						zb0003--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "VarId":
							z.Actions[za0001].VarId, bts, err = msgp.ReadBytesBytes(bts, z.Actions[za0001].VarId)
							if err != nil {
								return
							}
						case "Value":
							bts, err = z.Actions[za0001].Value.UnmarshalMsg(bts)
							if err != nil {
								return
							}
						case "Meta":
							var zb0004 uint32
							zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
							if err != nil {
								return
							}
							for zb0004 > 0 {
								zb0004--
								field, bts, err = msgp.ReadMapKeyZC(bts)
								if err != nil {
									return
								}
								switch msgp.UnsafeString(field) {
								case "AddSub":
									z.Actions[za0001].Meta.AddSub, bts, err = msgp.ReadBoolBytes(bts)
									if err != nil {
										return
									}
								case "DelSub":
									z.Actions[za0001].Meta.DelSub, bts, err = msgp.ReadBytesBytes(bts, z.Actions[za0001].Meta.DelSub)
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
func (z *ClientTxn) Msgsize() (s int) {
	s = 1 + 3 + msgp.BytesPrefixSize + len(z.Id) + 8 + msgp.Uint64Size + 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Actions {
		if z.Actions[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 6 + msgp.BytesPrefixSize + len(z.Actions[za0001].VarId) + 6 + z.Actions[za0001].Value.Msgsize() + 5 + 1 + 7 + msgp.BoolSize + 7 + msgp.BytesPrefixSize + len(z.Actions[za0001].Meta.DelSub)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientTxnOutcome) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
		case "Counter":
			z.Counter, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "Commit":
			z.Commit, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "Abort":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Abort) >= int(zb0002) {
				z.Abort = (z.Abort)[:zb0002]
			} else {
				z.Abort = make([]*ClientAction, zb0002)
			}
			for za0001 := range z.Abort {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Abort[za0001] = nil
				} else {
					if z.Abort[za0001] == nil {
						z.Abort[za0001] = new(ClientAction)
					}
					var zb0003 uint32
					zb0003, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zb0003 > 0 {
						zb0003--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "VarId":
							z.Abort[za0001].VarId, err = dc.ReadBytes(z.Abort[za0001].VarId)
							if err != nil {
								return
							}
						case "Value":
							err = z.Abort[za0001].Value.DecodeMsg(dc)
							if err != nil {
								return
							}
						case "Meta":
							var zb0004 uint32
							zb0004, err = dc.ReadMapHeader()
							if err != nil {
								return
							}
							for zb0004 > 0 {
								zb0004--
								field, err = dc.ReadMapKeyPtr()
								if err != nil {
									return
								}
								switch msgp.UnsafeString(field) {
								case "AddSub":
									z.Abort[za0001].Meta.AddSub, err = dc.ReadBool()
									if err != nil {
										return
									}
								case "DelSub":
									z.Abort[za0001].Meta.DelSub, err = dc.ReadBytes(z.Abort[za0001].Meta.DelSub)
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
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
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
	// map header, size 6
	// write "Id"
	err = en.Append(0x86, 0xa2, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Id)
	if err != nil {
		return
	}
	// write "FinalId"
	err = en.Append(0xa7, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.FinalId)
	if err != nil {
		return
	}
	// write "Counter"
	err = en.Append(0xa7, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Counter)
	if err != nil {
		return
	}
	// write "Commit"
	err = en.Append(0xa6, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Commit)
	if err != nil {
		return
	}
	// write "Abort"
	err = en.Append(0xa5, 0x41, 0x62, 0x6f, 0x72, 0x74)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Abort)))
	if err != nil {
		return
	}
	for za0001 := range z.Abort {
		if z.Abort[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 3
			// write "VarId"
			err = en.Append(0x83, 0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
			if err != nil {
				return
			}
			err = en.WriteBytes(z.Abort[za0001].VarId)
			if err != nil {
				return
			}
			// write "Value"
			err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return
			}
			err = z.Abort[za0001].Value.EncodeMsg(en)
			if err != nil {
				return
			}
			// write "Meta"
			// map header, size 2
			// write "AddSub"
			err = en.Append(0xa4, 0x4d, 0x65, 0x74, 0x61, 0x82, 0xa6, 0x41, 0x64, 0x64, 0x53, 0x75, 0x62)
			if err != nil {
				return
			}
			err = en.WriteBool(z.Abort[za0001].Meta.AddSub)
			if err != nil {
				return
			}
			// write "DelSub"
			err = en.Append(0xa6, 0x44, 0x65, 0x6c, 0x53, 0x75, 0x62)
			if err != nil {
				return
			}
			err = en.WriteBytes(z.Abort[za0001].Meta.DelSub)
			if err != nil {
				return
			}
		}
	}
	// write "Error"
	err = en.Append(0xa5, 0x45, 0x72, 0x72, 0x6f, 0x72)
	if err != nil {
		return
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
	// map header, size 6
	// string "Id"
	o = append(o, 0x86, 0xa2, 0x49, 0x64)
	o = msgp.AppendBytes(o, z.Id)
	// string "FinalId"
	o = append(o, 0xa7, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x64)
	o = msgp.AppendBytes(o, z.FinalId)
	// string "Counter"
	o = append(o, 0xa7, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72)
	o = msgp.AppendUint64(o, z.Counter)
	// string "Commit"
	o = append(o, 0xa6, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74)
	o = msgp.AppendBool(o, z.Commit)
	// string "Abort"
	o = append(o, 0xa5, 0x41, 0x62, 0x6f, 0x72, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Abort)))
	for za0001 := range z.Abort {
		if z.Abort[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 3
			// string "VarId"
			o = append(o, 0x83, 0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
			o = msgp.AppendBytes(o, z.Abort[za0001].VarId)
			// string "Value"
			o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
			o, err = z.Abort[za0001].Value.MarshalMsg(o)
			if err != nil {
				return
			}
			// string "Meta"
			// map header, size 2
			// string "AddSub"
			o = append(o, 0xa4, 0x4d, 0x65, 0x74, 0x61, 0x82, 0xa6, 0x41, 0x64, 0x64, 0x53, 0x75, 0x62)
			o = msgp.AppendBool(o, z.Abort[za0001].Meta.AddSub)
			// string "DelSub"
			o = append(o, 0xa6, 0x44, 0x65, 0x6c, 0x53, 0x75, 0x62)
			o = msgp.AppendBytes(o, z.Abort[za0001].Meta.DelSub)
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
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
		case "Counter":
			z.Counter, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "Commit":
			z.Commit, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "Abort":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Abort) >= int(zb0002) {
				z.Abort = (z.Abort)[:zb0002]
			} else {
				z.Abort = make([]*ClientAction, zb0002)
			}
			for za0001 := range z.Abort {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Abort[za0001] = nil
				} else {
					if z.Abort[za0001] == nil {
						z.Abort[za0001] = new(ClientAction)
					}
					var zb0003 uint32
					zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zb0003 > 0 {
						zb0003--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "VarId":
							z.Abort[za0001].VarId, bts, err = msgp.ReadBytesBytes(bts, z.Abort[za0001].VarId)
							if err != nil {
								return
							}
						case "Value":
							bts, err = z.Abort[za0001].Value.UnmarshalMsg(bts)
							if err != nil {
								return
							}
						case "Meta":
							var zb0004 uint32
							zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
							if err != nil {
								return
							}
							for zb0004 > 0 {
								zb0004--
								field, bts, err = msgp.ReadMapKeyZC(bts)
								if err != nil {
									return
								}
								switch msgp.UnsafeString(field) {
								case "AddSub":
									z.Abort[za0001].Meta.AddSub, bts, err = msgp.ReadBoolBytes(bts)
									if err != nil {
										return
									}
								case "DelSub":
									z.Abort[za0001].Meta.DelSub, bts, err = msgp.ReadBytesBytes(bts, z.Abort[za0001].Meta.DelSub)
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
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
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
	s = 1 + 3 + msgp.BytesPrefixSize + len(z.Id) + 8 + msgp.BytesPrefixSize + len(z.FinalId) + 8 + msgp.Uint64Size + 7 + msgp.BoolSize + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Abort {
		if z.Abort[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 6 + msgp.BytesPrefixSize + len(z.Abort[za0001].VarId) + 6 + z.Abort[za0001].Value.Msgsize() + 5 + 1 + 7 + msgp.BoolSize + 7 + msgp.BytesPrefixSize + len(z.Abort[za0001].Meta.DelSub)
		}
	}
	s += 6 + msgp.StringPrefixSize + len(z.Error)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientVarId) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
				var zb0002 uint32
				zb0002, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zb0002 > 0 {
					zb0002--
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
		return
	}
	err = en.WriteBytes(z.VarId)
	if err != nil {
		return
	}
	// write "Capability"
	err = en.Append(0xaa, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79)
	if err != nil {
		return
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
			return
		}
		err = en.WriteBool(z.Capability.Read)
		if err != nil {
			return
		}
		// write "Write"
		err = en.Append(0xa5, 0x57, 0x72, 0x69, 0x74, 0x65)
		if err != nil {
			return
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
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
				var zb0002 uint32
				zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0002 > 0 {
					zb0002--
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
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
		return
	}
	err = en.WriteString(z.Product)
	if err != nil {
		return
	}
	// write "Version"
	err = en.Append(0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
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
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Roots) >= int(zb0002) {
				z.Roots = (z.Roots)[:zb0002]
			} else {
				z.Roots = make([]*Root, zb0002)
			}
			for za0001 := range z.Roots {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Roots[za0001] = nil
				} else {
					if z.Roots[za0001] == nil {
						z.Roots[za0001] = new(Root)
					}
					err = z.Roots[za0001].DecodeMsg(dc)
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
		return
	}
	err = en.WriteBytes(z.Namespace)
	if err != nil {
		return
	}
	// write "Roots"
	err = en.Append(0xa5, 0x52, 0x6f, 0x6f, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Roots)))
	if err != nil {
		return
	}
	for za0001 := range z.Roots {
		if z.Roots[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Roots[za0001].EncodeMsg(en)
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
	for za0001 := range z.Roots {
		if z.Roots[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Roots[za0001].MarshalMsg(o)
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
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Roots) >= int(zb0002) {
				z.Roots = (z.Roots)[:zb0002]
			} else {
				z.Roots = make([]*Root, zb0002)
			}
			for za0001 := range z.Roots {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Roots[za0001] = nil
				} else {
					if z.Roots[za0001] == nil {
						z.Roots[za0001] = new(Root)
					}
					bts, err = z.Roots[za0001].UnmarshalMsg(bts)
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
	for za0001 := range z.Roots {
		if z.Roots[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Roots[za0001].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Root) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
				var zb0002 uint32
				zb0002, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zb0002 > 0 {
					zb0002--
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
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "VarId"
	err = en.Append(0xa5, 0x56, 0x61, 0x72, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.VarId)
	if err != nil {
		return
	}
	// write "Capability"
	err = en.Append(0xaa, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79)
	if err != nil {
		return
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
			return
		}
		err = en.WriteBool(z.Capability.Read)
		if err != nil {
			return
		}
		// write "Write"
		err = en.Append(0xa5, 0x57, 0x72, 0x69, 0x74, 0x65)
		if err != nil {
			return
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
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
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
				var zb0002 uint32
				zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0002 > 0 {
					zb0002--
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
