package common

import (
	"bytes"
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

func (a *TxnId) Equal(b *TxnId) bool {
	return a == b || (a != nil && b != nil && bytes.Equal(a[:], b[:]))
}

func (a *TxnId) LessThan(b *TxnId) bool {
	switch {
	case b == nil:
		return false
	case a == nil:
		return true
	default:
		return bytes.Compare(a[:], b[:]) < 0
	}
}

func (a *VarUUId) Equal(b *VarUUId) bool {
	return a == b || (a != nil && b != nil && bytes.Equal(a[:], b[:]))
}

func (a *VarUUId) LessThan(b *VarUUId) bool {
	switch {
	case b == nil:
		return false
	case a == nil:
		return true
	default:
		return bytes.Compare(a[:], b[:]) < 0
	}
}

type VarUUIds []*VarUUId

func (vUUIds VarUUIds) Sort()              { sort.Sort(vUUIds) }
func (vUUIds VarUUIds) Len() int           { return len(vUUIds) }
func (vUUIds VarUUIds) Less(i, j int) bool { return vUUIds[i].LessThan(vUUIds[j]) }
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

func (rmIds RMIds) NonEmptyLen() int {
	count := 0
	for _, rmId := range rmIds {
		if rmId != RMIdEmpty {
			count++
		}
	}
	return count
}