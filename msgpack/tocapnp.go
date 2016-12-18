package msgpack

import (
	capn "github.com/glycerine/go-capnproto"
	msgs "goshawkdb.io/common/capnp"
)

func (ctxn *ClientTxn) ToCapnp(seg *capn.Segment) msgs.ClientTxn {
	msg := msgs.NewClientTxn(seg)
	msg.SetId(ctxn.Id)
	msg.SetRetry(ctxn.Retry)
	actionsMsg := msgs.NewClientActionList(seg, len(ctxn.Actions))
	for idx, action := range ctxn.Actions {
		action.SetCapnp(seg, actionsMsg.At(idx))
	}
	msg.SetActions(actionsMsg)
	return msg
}

func (a *ClientAction) SetCapnp(seg *capn.Segment, msg msgs.ClientAction) {
	msg.SetVarId(a.VarId)
	switch {
	case a.Read != nil:
		read := a.Read
		msg.SetRead()
		readMsg := msg.Read()
		readMsg.SetVersion(read.Version)

	case a.Write != nil:
		write := a.Write
		msg.SetWrite()
		writeMsg := msg.Write()
		writeMsg.SetValue(write.Value)
		refsMsg := msgs.NewClientVarIdPosList(seg, len(write.References))
		for idx, ref := range write.References {
			ref.SetCapnp(seg, refsMsg.At(idx))
		}
		writeMsg.SetReferences(refsMsg)

	case a.ReadWrite != nil:
		rw := a.ReadWrite
		msg.SetWrite()
		rwMsg := msg.Readwrite()
		rwMsg.SetVersion(rw.Version)
		rwMsg.SetValue(rw.Value)
		refsMsg := msgs.NewClientVarIdPosList(seg, len(rw.References))
		for idx, ref := range rw.References {
			ref.SetCapnp(seg, refsMsg.At(idx))
		}
		rwMsg.SetReferences(refsMsg)

	case a.Create != nil:
		create := a.Create
		msg.SetCreate()
		createMsg := msg.Create()
		createMsg.SetValue(create.Value)
		refsMsg := msgs.NewClientVarIdPosList(seg, len(create.References))
		for idx, ref := range create.References {
			ref.SetCapnp(seg, refsMsg.At(idx))
		}
		createMsg.SetReferences(refsMsg)
	}
}

func (cvi *ClientVarId) SetCapnp(seg *capn.Segment, msg msgs.ClientVarIdPos) {
	msg.SetVarId(cvi.VarId)
	msg.SetCapability(cvi.Capability.ToCapnp(seg))
}

func (c *Capability) ToCapnp(seg *capn.Segment) msgs.Capability {
	msg := msgs.NewCapability(seg)
	switch {
	case c.Read && c.Write:
		msg.SetReadWrite()
	case c.Read:
		msg.SetRead()
	case c.Write:
		msg.SetWrite()
	default:
		msg.SetNone()
	}
	return msg
}
