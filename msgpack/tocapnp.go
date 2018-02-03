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
	msg.SetCounter(ctxn.Counter)
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
	msgValue := msg.Value()
	aValue := &a.Value
	if aValue.Missing {
		msgValue.SetMissing()

	} else if aCreate := aValue.Create; aCreate != nil {
		msgValue.SetCreate()
		msgCreate := msgValue.Create()
		msgCreate.SetValue(aCreate.Value)
		refsMsg := msgs.NewClientVarIdPosList(seg, len(aCreate.References))
		for idx, ref := range aCreate.References {
			if err := ref.SetCapnp(seg, refsMsg.At(idx)); err != nil {
				return err
			}
		}
		msgCreate.SetReferences(refsMsg)

	} else if aExisting := aValue.Existing; aExisting != nil {
		msgValue.SetExisting()
		msgExisting := msgValue.Existing()
		msgExisting.SetRead(aExisting.Read)
		msgModify := msgExisting.Modify()
		if aModify := aExisting.Modify; aModify == nil {
			msgModify.SetNot()
		} else {
			msgModify.SetWrite()
			msgWrite := msgModify.Write()
			msgWrite.SetValue(aModify.Value)
			refsMsg := msgs.NewClientVarIdPosList(seg, len(aModify.References))
			for idx, ref := range aModify.References {
				if err := ref.SetCapnp(seg, refsMsg.At(idx)); err != nil {
					return err
				}
			}
			msgWrite.SetReferences(refsMsg)
		}

	} else {
		return errors.New("Invalid Action: no value action found.")
	}
	aMeta := &a.Meta
	msgMeta := msg.Meta()
	msgMeta.SetAddSub(aMeta.AddSub)
	if len(aMeta.DelSub) != 0 {
		msgMeta.SetDelSub(aMeta.DelSub)
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
