package common

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	capn "github.com/glycerine/go-capnproto"
	msgs "goshawkdb.io/common/capnp"
	"sort"
)

const (
	RMIdEmpty = RMId(0)
	KeyLen    = 20
	ClientLen = KeyLen - 8
)

type RMId uint32

func (rmId RMId) String() string {
	if rmId == RMIdEmpty {
		return "RM:Empty"
	} else {
		return fmt.Sprintf("RM:%x", uint32(rmId))
	}
}

type KeyType [KeyLen]byte

type VarUUId KeyType

func MakeVarUUId(data []byte) *VarUUId {
	uuid := VarUUId([KeyLen]byte{})
	copy(uuid[:], data)
	return &uuid
}

func MakeVarUUIdFromStr(str string) *VarUUId {
	data, err := hex.DecodeString(str)
	if err != nil {
		return nil
	}
	return MakeVarUUId(data)
}

func (vUUId VarUUId) String() string {
	h := ([]byte)(hex.EncodeToString(vUUId[:]))
	return fmt.Sprintf("VarUUId:%s", h)
}

func (vUUId VarUUId) ConnectionCount() uint32 {
	return binary.BigEndian.Uint32(vUUId[8:12])
}

func (vUUId VarUUId) BootCount() uint32 {
	return binary.BigEndian.Uint32(vUUId[12:16])
}

func (vUUId VarUUId) RMId() RMId {
	return RMId(binary.BigEndian.Uint32(vUUId[16:20]))
}

type TxnId KeyType

func MakeTxnId(data []byte) *TxnId {
	id := TxnId([KeyLen]byte{})
	copy(id[:], data)
	return &id
}

func (txnId TxnId) String() string {
	t := txnId[:]
	client, conn, boot, rmId := hex.EncodeToString(t[0:8]), hex.EncodeToString(t[8:12]), hex.EncodeToString(t[12:16]), hex.EncodeToString(t[16:20])
	return fmt.Sprintf("TxnId:%s-%s-%s-%v", client, conn, boot, rmId)
}

func (txnId TxnId) ClientId() [ClientLen]byte {
	client := [ClientLen]byte{}
	copy(client[:], txnId[8:])
	return client
}

func (txnId TxnId) ConnectionCount() uint32 {
	return binary.BigEndian.Uint32(txnId[8:12])
}

func (txnId TxnId) BootCount() uint32 {
	return binary.BigEndian.Uint32(txnId[12:16])
}

func (txnId TxnId) RMId() RMId {
	return RMId(binary.BigEndian.Uint32(txnId[16:20]))
}

type Cmp int8

const (
	LT Cmp = iota - 1
	EQ
	GT
)

func (a *TxnId) Compare(b *TxnId) Cmp {
	switch {
	case a == b:
		return EQ
	case a == nil:
		return LT
	case b == nil:
		return GT
	default:
		switch cmp := bytes.Compare(a[:], b[:]); {
		case cmp < 0:
			return LT
		case cmp > 0:
			return GT
		default:
			return EQ
		}
	}
}

func (a *VarUUId) Compare(b *VarUUId) Cmp {
	switch {
	case a == b:
		return EQ
	case a == nil:
		return LT
	case b == nil:
		return GT
	default:
		switch cmp := bytes.Compare(a[:], b[:]); {
		case cmp < 0:
			return LT
		case cmp > 0:
			return GT
		default:
			return EQ
		}
	}
}

type VarUUIds []*VarUUId

func (vUUIds VarUUIds) Sort()              { sort.Sort(vUUIds) }
func (vUUIds VarUUIds) Len() int           { return len(vUUIds) }
func (vUUIds VarUUIds) Less(i, j int) bool { return vUUIds[i].Compare(vUUIds[j]) == LT }
func (vUUIds VarUUIds) Swap(i, j int)      { vUUIds[i], vUUIds[j] = vUUIds[j], vUUIds[i] }

type Positions capn.UInt8List

func (a *Positions) Equal(b *Positions) bool {
	if a == nil || b == nil {
		return a == b
	}
	aCap, bCap := (*capn.UInt8List)(a), (*capn.UInt8List)(b)
	if aCap.Len() != bCap.Len() {
		return false
	}
	for idx, l := 0, aCap.Len(); idx < l; idx++ {
		if aV, bV := aCap.At(idx), bCap.At(idx); aV != bV {
			return false
		}
	}
	return true
}

func (p *Positions) String() string {
	return fmt.Sprint((*capn.UInt8List)(p).ToArray())
}

type RMIds []RMId

func (rmIds RMIds) Sort()              { sort.Sort(rmIds) }
func (rmIds RMIds) Len() int           { return len(rmIds) }
func (rmIds RMIds) Less(i, j int) bool { return rmIds[i] < rmIds[j] }
func (rmIds RMIds) Swap(i, j int)      { rmIds[i], rmIds[j] = rmIds[j], rmIds[i] }

func (a RMIds) Equal(b RMIds) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if len(a) != len(b) {
		return false
	}
	for idx, rA := range a {
		if b[idx] != rA {
			return false
		}
	}
	return true
}

func (rmIds RMIds) EmptyLen() int {
	count := 0
	for _, rmId := range rmIds {
		if rmId == RMIdEmpty {
			count++
		}
	}
	return count
}

func (rmIds RMIds) NonEmptyLen() int {
	count := 0
	for _, rmId := range rmIds {
		if rmId != RMIdEmpty {
			count++
		}
	}
	return count
}

func (rmIds RMIds) NonEmpty() RMIds {
	nel := rmIds.NonEmptyLen()
	if nel == len(rmIds) {
		return rmIds
	}
	nonEmpty := make([]RMId, 0, nel)
	for _, rmId := range rmIds {
		if rmId != RMIdEmpty {
			nonEmpty = append(nonEmpty, rmId)
		}
	}
	return nonEmpty
}

var MaxCapsCap *Capabilities

func init() {
	seg := capn.NewBuffer(nil)
	cap := msgs.NewCapabilities(seg)
	cap.SetValue(msgs.VALUECAPABILITY_READWRITE)
	ref := cap.References()
	ref.Read().SetAll()
	ref.Write().SetAll()
	MaxCapsCap = NewCapabilities(cap)
}

type Capabilities struct {
	msgs.Capabilities
}

func NewCapabilities(c msgs.Capabilities) *Capabilities {
	return &Capabilities{Capabilities: c}
}

func (a *Capabilities) Equal(b *Capabilities) bool {
	if a.Value() != b.Value() {
		return false
	}
	aRefs, bRefs := a.References(), b.References()
	if aRefs.Read().Which() != bRefs.Read().Which() ||
		aRefs.Write().Which() != bRefs.Write().Which() {
		return false
	}
	if aRefs.Read().Which() == msgs.CAPABILITIESREFERENCESREAD_ONLY {
		aOnly, bOnly := aRefs.Read().Only().ToArray(), bRefs.Read().Only().ToArray()
		if len(aOnly) != len(bOnly) {
			return false
		}
		for idx, aIndex := range aOnly {
			if aIndex != bOnly[idx] {
				return false
			}
		}
	}
	if aRefs.Write().Which() == msgs.CAPABILITIESREFERENCESWRITE_ONLY {
		aOnly, bOnly := aRefs.Write().Only().ToArray(), bRefs.Write().Only().ToArray()
		if len(aOnly) != len(bOnly) {
			return false
		}
		for idx, aIndex := range aOnly {
			if aIndex != bOnly[idx] {
				return false
			}
		}
	}
	return true
}

func (c *Capabilities) String() string {
	value := ""
	switch c.Value() {
	case msgs.VALUECAPABILITY_NONE:
		value = "None"
	case msgs.VALUECAPABILITY_READ:
		value = "Read"
	case msgs.VALUECAPABILITY_WRITE:
		value = "Write"
	case msgs.VALUECAPABILITY_READWRITE:
		value = "ReadWrite"
	}
	refsRead := ""
	if c.References().Read().Which() == msgs.CAPABILITIESREFERENCESREAD_ALL {
		refsRead = "All"
	} else {
		refsRead = fmt.Sprintf("%v", c.References().Read().Only().ToArray())
	}
	refsWrite := ""
	if c.References().Write().Which() == msgs.CAPABILITIESREFERENCESWRITE_ALL {
		refsWrite = "All"
	} else {
		refsWrite = fmt.Sprintf("%v", c.References().Write().Only().ToArray())
	}
	return fmt.Sprintf("{capability: value: %s; refsRead: %s; refsWrite: %s}", value, refsRead, refsWrite)
}

func (a *Capabilities) Union(b *Capabilities) *Capabilities {
	switch {
	case a == b:
		return a
	case a == MaxCapsCap || b == MaxCapsCap:
		return MaxCapsCap
	case a == nil:
		return b
	case b == nil:
		return a
	}

	aValue := a.Value()
	aRefsRead := a.References().Read()
	aRefsWrite := a.References().Write()

	bValue := b.Value()
	bRefsRead := b.References().Read()
	bRefsWrite := b.References().Write()

	valueRead := aValue == msgs.VALUECAPABILITY_READWRITE || aValue == msgs.VALUECAPABILITY_READ ||
		bValue == msgs.VALUECAPABILITY_READWRITE || bValue == msgs.VALUECAPABILITY_READ
	valueWrite := aValue == msgs.VALUECAPABILITY_READWRITE || aValue == msgs.VALUECAPABILITY_WRITE ||
		bValue == msgs.VALUECAPABILITY_READWRITE || bValue == msgs.VALUECAPABILITY_WRITE
	refsReadAll := aRefsRead.Which() == msgs.CAPABILITIESREFERENCESREAD_ALL || bRefsRead.Which() == msgs.CAPABILITIESREFERENCESREAD_ONLY
	refsWriteAll := aRefsWrite.Which() == msgs.CAPABILITIESREFERENCESWRITE_ALL || bRefsWrite.Which() == msgs.CAPABILITIESREFERENCESWRITE_ALL

	if valueRead && valueWrite && refsReadAll && refsWriteAll {
		return MaxCapsCap
	}

	seg := capn.NewBuffer(nil)
	cap := msgs.NewCapabilities(seg)
	switch {
	case valueRead && valueWrite:
		cap.SetValue(msgs.VALUECAPABILITY_READWRITE)
	case valueWrite:
		cap.SetValue(msgs.VALUECAPABILITY_WRITE)
	case valueRead:
		cap.SetValue(msgs.VALUECAPABILITY_WRITE)
	default:
		cap.SetValue(msgs.VALUECAPABILITY_NONE)
	}

	if refsReadAll {
		cap.References().Read().SetAll()
	} else {
		aOnly, bOnly := aRefsRead.Only().ToArray(), bRefsRead.Only().ToArray()
		cap.References().Read().SetOnly(mergeOnliesSeg(seg, aOnly, bOnly))
	}

	if refsWriteAll {
		cap.References().Write().SetAll()
	} else {
		aOnly, bOnly := aRefsWrite.Only().ToArray(), bRefsWrite.Only().ToArray()
		cap.References().Write().SetOnly(mergeOnliesSeg(seg, aOnly, bOnly))
	}

	return NewCapabilities(cap)
}

func mergeOnliesSeg(seg *capn.Segment, a, b []uint32) capn.UInt32List {
	only := mergeOnlies(a, b)

	cap := seg.NewUInt32List(len(only))
	for idx, index := range only {
		cap.Set(idx, index)
	}
	return cap
}

func mergeOnlies(a, b []uint32) []uint32 {
	only := make([]uint32, 0, len(a)+len(b))
	for len(a) > 0 && len(b) > 0 {
		aIndex, bIndex := a[0], b[0]
		switch {
		case aIndex < bIndex:
			only = append(only, aIndex)
			a = a[1:]
		case aIndex > bIndex:
			only = append(only, bIndex)
			b = b[1:]
		default:
			only = append(only, aIndex)
			a = a[1:]
			b = b[1:]
		}
	}
	if len(a) > 0 {
		only = append(only, a...)
	} else {
		only = append(only, b...)
	}

	return only
}

type SortUInt32 []uint32

func (nums SortUInt32) Sort()              { sort.Sort(nums) }
func (nums SortUInt32) Len() int           { return len(nums) }
func (nums SortUInt32) Less(i, j int) bool { return nums[i] < nums[j] }
func (nums SortUInt32) Swap(i, j int)      { nums[i], nums[j] = nums[j], nums[i] }
