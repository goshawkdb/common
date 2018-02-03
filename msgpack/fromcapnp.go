package msgpack

import (
	msgs "goshawkdb.io/common/capnp"
)

func (r *Root) FromCapnp(msg msgs.Root) *Root {
	r.Name = msg.Name()
	r.VarId = msg.VarId()
	r.Capability.FromCapnp(msg.Capability())
	return r
}

func (c *Capability) FromCapnp(msg msgs.Capability) *Capability {
	switch msg.Which() {
	case msgs.CAPABILITY_NONE:
		c.Read = false
		c.Write = false
	case msgs.CAPABILITY_READ:
		c.Read = true
		c.Write = false
	case msgs.CAPABILITY_WRITE:
		c.Read = false
		c.Write = true
	case msgs.CAPABILITY_READWRITE:
		c.Read = true
		c.Write = true
	default:
		panic("Unexpected capability")
	}
	return c
}

func (ctxn *ClientTxn) FromCapnp(msg msgs.ClientTxn) *ClientTxn {
	ctxn.Id = msg.Id()
	ctxn.Counter = msg.Counter()
	actions := msg.Actions()
	ctxn.Actions = make([]*ClientAction, actions.Len())
	for idx := range ctxn.Actions {
		ctxn.Actions[idx] = (&ClientAction{}).FromCapnp(actions.At(idx))
	}
	return ctxn
}

func (cto *ClientTxnOutcome) FromCapnp(msg msgs.ClientTxnOutcome) *ClientTxnOutcome {
	cto.Id = msg.Id()
	cto.FinalId = msg.FinalId()
	cto.Counter = msg.Counter()
	cto.Commit = false
	cto.Abort = nil
	cto.Error = ""
	switch msg.Which() {
	case msgs.CLIENTTXNOUTCOME_COMMIT:
		cto.Commit = true
	case msgs.CLIENTTXNOUTCOME_ABORT:
		abort := msg.Abort()
		cto.Abort = make([]*ClientAction, abort.Len())
		for idx := range cto.Abort {
			cto.Abort[idx] = (&ClientAction{}).FromCapnp(abort.At(idx))
		}
	case msgs.CLIENTTXNOUTCOME_ERROR:
		cto.Error = msg.Error()
	default:
		panic("Unexpected Txn Outcome")
	}
	return cto
}

func (a *ClientAction) FromCapnp(msg msgs.ClientAction) *ClientAction {
	a.VarId = msg.VarId()
	aValue := &a.Value
	msgValue := msg.Value()
	switch msgValue.Which() {
	case msgs.CLIENTACTIONVALUE_MISSING:
		aValue.Missing = true
	case msgs.CLIENTACTIONVALUE_CREATE:
		msgCreate := msgValue.Create()
		msgRefs := msgCreate.References()
		aRefs := make([]*ClientVarId, msgRefs.Len())
		for idx := range aRefs {
			aRefs[idx] = (&ClientVarId{}).FromCapnp(msgRefs.At(idx))
		}
		aValue.Create = &ClientActionValueCreate{
			Value:      msgCreate.Value(),
			References: aRefs,
		}
	case msgs.CLIENTACTIONVALUE_EXISTING:
		msgExisting := msgValue.Existing()
		aValue.Existing = &ClientActionValueExisting{
			Read: msgExisting.Read(),
		}
		msgModify := msgExisting.Modify()
		if msgModify.Which() == msgs.CLIENTACTIONVALUEEXISTINGMODIFY_WRITE {
			msgWrite := msgModify.Write()
			msgRefs := msgWrite.References()
			aRefs := make([]*ClientVarId, msgRefs.Len())
			for idx := range aRefs {
				aRefs[idx] = (&ClientVarId{}).FromCapnp(msgRefs.At(idx))
			}
			aValue.Existing.Modify = &ClientActionValueExistingModify{
				Value:      msgWrite.Value(),
				References: aRefs,
			}
		}
	}
	msgMeta := msg.Meta()
	aMeta := &a.Meta
	aMeta.AddSub = msgMeta.AddSub()
	if delSub := msgMeta.DelSub(); len(delSub) != 0 {
		aMeta.DelSub = delSub
	}
	return a
}

func (cvi *ClientVarId) FromCapnp(msg msgs.ClientVarIdPos) *ClientVarId {
	cvi.VarId = msg.VarId()
	cvi.Capability = &Capability{}
	cvi.Capability.FromCapnp(msg.Capability())
	return cvi
}
