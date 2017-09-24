package msgpack

import (
	"errors"
	"fmt"
	capn "github.com/glycerine/go-capnproto"
	"goshawkdb.io/common"
	msgs "goshawkdb.io/common/capnp"
)

func (ctxn *ClientTxn) ToCapnp(seg *capn.Segment) (*msgs.ClientTxn, error) {
	msg := msgs.NewClientTxn(seg)
	if len(ctxn.Id) != common.KeyLen {
		return nil, errors.New("Illegal ClientTxn: Invalid Txn Id (incorrect length).")
	}
	msg.SetId(ctxn.Id)
	msg.SetRetry(ctxn.Retry)
	if len(ctxn.Actions) == 0 {
		return nil, errors.New("Illegal ClientTxn: No actions.")
	}
	actionsMsg := msgs.NewClientActionList(seg, len(ctxn.Actions))
	for idx, action := range ctxn.Actions {
		if err := action.SetCapnp(seg, actionsMsg.At(idx)); err != nil {
			return nil, fmt.Errorf("Illegal ClientTxn: %v", err)
		}
	}
	msg.SetActions(actionsMsg)
	return &msg, nil
}

func (a *ClientAction) SetCapnp(seg *capn.Segment, msg msgs.ClientAction) error {
	if len(a.VarId) != common.KeyLen {
		return errors.New("Invalid Var Id in ClientAction (incorrect length).")
	}
	msg.SetVarId(a.VarId)
	isModified := a.Modified != nil
	if isModified {
		modified := a.Modified
		msg.SetModified()
		modifiedMsg := msg.Modified()
		modifiedMsg.SetValue(modified.Value)
		refsMsg := msgs.NewClientVarIdPosList(seg, len(modified.References))
		for idx, ref := range modified.References {
			if err := ref.SetCapnp(seg, refsMsg.At(idx)); err != nil {
				return err
			}
		}
		modifiedMsg.SetReferences(refsMsg)
	} else {
		msg.SetUnmodified()
	}
	needsModified := true
	switch a.ActionType {
	case CreateActionType:
		msg.SetActionType(msgs.CLIENTACTIONTYPE_CREATE)
	case ReadOnlyActionType:
		msg.SetActionType(msgs.CLIENTACTIONTYPE_READONLY)
		needsModified = false
	case WriteOnlyActionType:
		msg.SetActionType(msgs.CLIENTACTIONTYPE_WRITEONLY)
	case ReadWriteActionType:
		msg.SetActionType(msgs.CLIENTACTIONTYPE_READWRITE)
	case DeleteActionType:
		msg.SetActionType(msgs.CLIENTACTIONTYPE_DELETE)
		needsModified = false
	default:
		return fmt.Errorf("Invalid ClientActionType: %v", a.ActionType)
	}
	if needsModified != isModified {
		return fmt.Errorf("Modification required: %v; Modification provided: %v", needsModified, isModified)
	}
	return nil
}

func (cvi *ClientVarId) SetCapnp(seg *capn.Segment, msg msgs.ClientVarIdPos) error {
	if len(cvi.VarId) != common.KeyLen {
		return errors.New("Invalid VarId in Action References (incorrect length).")
	}
	msg.SetVarId(cvi.VarId)
	if cvi.Capability == nil {
		return errors.New("Invalid Capability in Action References (no capability supplied).")
	}
	msg.SetCapability(cvi.Capability.ToCapnp(seg))
	return nil
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
