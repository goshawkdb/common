package common

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	capn "github.com/glycerine/go-capnproto"
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

type ValueCaps uint8

const (
	None      ValueCaps = iota
	Read                = iota
	Write               = iota
	ReadWrite           = iota
)

func (vc ValueCaps) String() string {
	switch vc {
	case None:
		return "None"
	case Read:
		return "Read"
	case Write:
		return "Write"
	default:
		return "ReadWrite"
	}
}

type ReferenceCapabilities struct {
	All  bool
	Only []uint32
}

func (a ReferenceCapabilities) Equal(b ReferenceCapabilities) bool {
	if !(a.All == b.All &&
		len(a.Only) == len(b.Only)) {
		return false
	}
	for idx, index := range a.Only {
		if index != b.Only[idx] {
			return false
		}
	}
	return true
}

func (rc ReferenceCapabilities) String() string {
	if rc.All {
		return "All"
	}
	return fmt.Sprintf("OnlyIndices: %v", rc.Only)
}

type Capabilities struct {
	Value      ValueCaps
	References struct {
		Read  ReferenceCapabilities
		Write ReferenceCapabilities
	}
}

func (c *Capabilities) String() string {
	return fmt.Sprintf("Capabilities(Value: %v; References: Read: %v; Write: %v)",
		c.Value, c.References.Read, c.References.Write)
}

func (a *Capabilities) Equal(b *Capabilities) bool {
	return a.Value == b.Value &&
		a.References.Read.Equal(b.References.Read) &&
		a.References.Write.Equal(b.References.Write)
}
