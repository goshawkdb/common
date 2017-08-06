package actor

import (
	cc "github.com/msackman/chancell"
)

type Msg interface{}

type MsgSync interface {
	InitMsg(actor EnqueueMsgActor)
	Close() bool
	Wait() bool
}

type MsgSyncQuery struct {
	WaitChan       chan struct{}
	TerminatedChan <-chan struct{}
}

func (msq *MsgSyncQuery) InitMsg(actor EnqueueMsgActor) {
	msq.WaitChan = make(chan struct{})
	msq.TerminatedChan = actor.TerminatedChan()
}

func (msq *MsgSyncQuery) Close() bool {
	select {
	case <-msq.WaitChan:
		return false
	default:
		close(msq.WaitChan)
		return true
	}
}

func (msq *MsgSyncQuery) MustClose() {
	if !msq.Close() {
		panic("Cannot close MsgSyncQuery!")
	}
}

// Returns true iff Close() is called on MsgSyncQuery (i.e
// MsgSyncQuery.WaitChan is closed).  Blocks until either
// MsgSyncQuery.WaitChan is closed, or MsgSyncQuery.TerminatedChan
// is closed.
func (msq *MsgSyncQuery) Wait() bool {
	select {
	case <-msq.WaitChan:
		return true
	case <-msq.TerminatedChan:
		// if they both were ready to receive on, we could still end
		// up here, so we need to test at this point whether or not
		// finished can be received from.
		select {
		case <-msq.WaitChan:
			return true
		default:
			return false
		}
	}
}

type Mailbox struct {
	cellTail          *cc.ChanCellTail
	enqueueQueryInner func(Msg, *cc.ChanCell, cc.CurCellConsumer) (bool, cc.CurCellConsumer)
	Terminated        <-chan struct{}
}

type EnqueueMsgActor interface {
	EnqueueMsg(Msg) bool
	TerminatedChan() <-chan struct{}
}

type queryCapture struct {
	mailbox *Mailbox
	msg     Msg
}

func (qc *queryCapture) consumer(cell *cc.ChanCell) (bool, cc.CurCellConsumer) {
	return qc.mailbox.enqueueQueryInner(qc.msg, cell, qc.consumer)
}

func (mailbox *Mailbox) EnqueueMsg(msg Msg) bool {
	qc := &queryCapture{mailbox: mailbox, msg: msg}
	return mailbox.cellTail.WithCell(qc.consumer)
}

func (mailbox *Mailbox) TerminatedChan() <-chan struct{} {
	return mailbox.Terminated
}
