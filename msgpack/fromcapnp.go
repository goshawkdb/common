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
		cto.Abort = make([]*ClientUpdate, abort.Len())
		for idx := range cto.Abort {
			cto.Abort[idx] = (&ClientUpdate{}).FromCapnp(abort.At(idx))
		}
	case msgs.CLIENTTXNOUTCOME_ERROR:
		cto.Error = msg.Error()
	default:
		panic("Unexpected Txn Outcome")
	}
	return cto
}

func (cu *ClientUpdate) FromCapnp(msg msgs.ClientUpdate) *ClientUpdate {
	cu.Version = msg.Version()
	actions := msg.Actions()
	cu.Actions = make([]*ClientAction, actions.Len())
	for idx := range cu.Actions {
		cu.Actions[idx] = (&ClientAction{}).FromCapnp(actions.At(idx))
	}
	return cu
}

func (a *ClientAction) FromCapnp(msg msgs.ClientAction) *ClientAction {
	a.VarId = msg.VarId()
	a.Read = nil
	a.Write = nil
	a.ReadWrite = nil
	a.Create = nil
	a.Delete = false
	switch msg.Which() {
	case msgs.CLIENTACTION_READ:
		a.Read = &ClientActionRead{Version: msg.Read().Version()}

	case msgs.CLIENTACTION_WRITE:
		write := msg.Write()
		refs := write.References()
		a.Write = &ClientActionWrite{
			Value:      write.Value(),
			References: make([]*ClientVarId, refs.Len()),
		}
		for idx := range a.Write.References {
			a.Write.References[idx] = (&ClientVarId{}).FromCapnp(refs.At(idx))
		}

	case msgs.CLIENTACTION_READWRITE:
		rw := msg.Readwrite()
		refs := rw.References()
		a.ReadWrite = &ClientActionReadWrite{
			Version:    rw.Version(),
			Value:      rw.Value(),
			References: make([]*ClientVarId, refs.Len()),
		}
		for idx := range a.ReadWrite.References {
			a.ReadWrite.References[idx] = (&ClientVarId{}).FromCapnp(refs.At(idx))
		}

	case msgs.CLIENTACTION_CREATE:
		create := msg.Create()
		refs := create.References()
		a.Create = &ClientActionCreate{
			Value:      create.Value(),
			References: make([]*ClientVarId, refs.Len()),
		}
		for idx := range a.Create.References {
			a.Create.References[idx] = (&ClientVarId{}).FromCapnp(refs.At(idx))
		}

	case msgs.CLIENTACTION_DELETE:
		a.Delete = true

	default:
		panic("Unexpected action type")
	}
	return a
}

func (cvi *ClientVarId) FromCapnp(msg msgs.ClientVarIdPos) *ClientVarId {
	cvi.VarId = msg.VarId()
	cvi.Capability = &Capability{}
	cvi.Capability.FromCapnp(msg.Capability())
	return cvi
}
