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
	actionCount := 0
	if a.Read != nil {
		actionCount++
	}
	if a.Write != nil {
		actionCount++
	}
	if a.ReadWrite != nil {
		actionCount++
	}
	if a.Create != nil {
		actionCount++
	}
	if actionCount != 1 {
		return errors.New("Invalid ClientAction: exactly one action type must be non-nil.")
	}
	switch {
	case a.Read != nil:
		read := a.Read
		msg.SetRead()
		readMsg := msg.Read()
		if len(read.Version) != common.KeyLen {
			return errors.New("Invalid read.Version in ClientAction (incorrect length).")
		}
		readMsg.SetVersion(read.Version)

	case a.Write != nil:
		write := a.Write
		msg.SetWrite()
		writeMsg := msg.Write()
		writeMsg.SetValue(write.Value)
		refsMsg := msgs.NewClientVarIdPosList(seg, len(write.References))
		for idx, ref := range write.References {
			if err := ref.SetCapnp(seg, refsMsg.At(idx)); err != nil {
				return err
			}
		}
		writeMsg.SetReferences(refsMsg)

	case a.ReadWrite != nil:
		rw := a.ReadWrite
		msg.SetReadwrite()
		rwMsg := msg.Readwrite()
		if len(rw.Version) != common.KeyLen {
			return errors.New("Invalid ReadWrite.Version in ClientAction (incorrect length).")
		}
		rwMsg.SetVersion(rw.Version)
		rwMsg.SetValue(rw.Value)
		refsMsg := msgs.NewClientVarIdPosList(seg, len(rw.References))
		for idx, ref := range rw.References {
			if err := ref.SetCapnp(seg, refsMsg.At(idx)); err != nil {
				return err
			}
		}
		rwMsg.SetReferences(refsMsg)

	case a.Create != nil:
		create := a.Create
		msg.SetCreate()
		createMsg := msg.Create()
		createMsg.SetValue(create.Value)
		refsMsg := msgs.NewClientVarIdPosList(seg, len(create.References))
		for idx, ref := range create.References {
			if err := ref.SetCapnp(seg, refsMsg.At(idx)); err != nil {
				return err
			}
		}
		createMsg.SetReferences(refsMsg)
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
