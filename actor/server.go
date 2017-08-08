package actor

import (
	"fmt"
	"github.com/go-kit/kit/log"
)

type Server interface {
	Init(*Actor) (bool, error)
	HandleBeat() (bool, error)
	HandleMsg(Msg) (bool, error)
	HandleShutdown(error) bool
}

type ShutdownableActor interface {
	ShutdownSync()
}

type EnqueueFuncActor interface {
	EnqueueFuncAsync(func() (bool, error)) bool
}

func NewBasicServerInner(logger log.Logger) *BasicServerInner {
	return &BasicServerInner{Logger: logger}
}

func NewBasicServerOuter(mailbox *Mailbox) *BasicServerOuter {
	return &BasicServerOuter{mailbox: mailbox}
}

type BasicServerOuter struct {
	mailbox *Mailbox
}

type BasicServerInner struct {
	Logger log.Logger
}

type MsgShutdown struct{}

func (shutdown MsgShutdown) Error() string {
	return "Actor Shutdown"
}

type Shutdownable interface {
	ShutdownSync()
}

func (outer *BasicServerOuter) ShutdownSync() {
	if outer.mailbox.EnqueueMsg(MsgShutdown{}) {
		outer.mailbox.cellTail.Wait()
	}
}

type MsgExec interface {
	Exec() (bool, error)
}

type MsgExecFunc func() (bool, error)

func (mef MsgExecFunc) Exec() (bool, error) {
	return mef()
}

func (outer *BasicServerOuter) EnqueueFuncAsync(fun func() (bool, error)) bool {
	return outer.mailbox.EnqueueMsg(MsgExecFunc(fun))
}

func (inner *BasicServerInner) Init(*Actor) (bool, error) {
	return false, nil
}

func (inner *BasicServerInner) HandleBeat() (bool, error) {
	return false, nil
}

func (inner *BasicServerInner) HandleMsg(msg Msg) (terminate bool, err error) {
	switch msgT := msg.(type) {
	case MsgExec:
		terminate, err = msgT.Exec()
	case MsgShutdown:
		terminate = true
	default:
		panic(fmt.Sprintf("Received unexpected message: %#v", msg))
	}
	return
}

func (inner *BasicServerInner) HandleShutdown(err error) bool {
	if err == nil {
		inner.Logger.Log("msg", "Actor shutting down.")
	} else {
		inner.Logger.Log("msg", "Actor shutting down.", "error", err)
	}
	return true
}
