package common

import (
	"bytes"
	"fmt"
	capn "github.com/glycerine/go-capnproto"
)

func SegToBytes(seg *capn.Segment) []byte {
	if seg == nil {
		panic("SegToBytes called with nil segment!")
	}
	buf := new(bytes.Buffer)
	if _, err := seg.WriteTo(buf); err != nil {
		panic(fmt.Sprintf("Error when writing segment to bytes: %v", err))
	}
	return buf.Bytes()
}
