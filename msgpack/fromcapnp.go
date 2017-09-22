package msgpack

import (
	"fmt"
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
	ctxn.Retry = msg.Retry()
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
	if msg.Which() == msgs.CLIENTACTION_MODIFIED {
		modified := msg.Modified()
		msgRefs := modified.References()
		refs := make([]*ClientVarId, msgRefs.Len())
		for idx := range refs {
			refs[idx] = (&ClientVarId{}).FromCapnp(msgRefs.At(idx))
		}
		a.Modified = &ClientActionModify{
			Value:      modified.Value(),
			References: refs,
		}
	} else {
		a.Modified = nil
	}
	switch msg.ActionType() {
	case msgs.CLIENTACTIONTYPE_CREATE:
		a.ActionType = CreateActionType
	case msgs.CLIENTACTIONTYPE_READONLY:
		a.ActionType = ReadOnlyActionType
	case msgs.CLIENTACTIONTYPE_WRITEONLY:
		a.ActionType = WriteOnlyActionType
	case msgs.CLIENTACTIONTYPE_READWRITE:
		a.ActionType = ReadWriteActionType
	case msgs.CLIENTACTIONTYPE_DELETE:
		a.ActionType = DeleteActionType
	default:
		panic(fmt.Sprintf("Unexpected action type: %v", msg.ActionType()))
	}
	return a
}

func (cvi *ClientVarId) FromCapnp(msg msgs.ClientVarIdPos) *ClientVarId {
	cvi.VarId = msg.VarId()
	cvi.Capability = &Capability{}
	cvi.Capability.FromCapnp(msg.Capability())
	return cvi
}
