package actor

import (
	"fmt"
	cc "github.com/msackman/chancell"
	"runtime/debug"
)

func Spawn(server Server) (*Mailbox, error) {
	mailbox := &Mailbox{}
	actor := &Actor{
		Mailbox: mailbox,
		Server:  server,
	}

	var head *cc.ChanCellHead
	head, mailbox.cellTail = cc.NewChanCellTail(
		func(n int, cell *cc.ChanCell) {
			queryChan := make(chan Msg, n)
			cell.Open = func() { actor.queryChan = queryChan }
			cell.Close = func() { close(queryChan) }
			mailbox.enqueueQueryInner = func(msg Msg, curCell *cc.ChanCell, cont cc.CurCellConsumer) (bool, cc.CurCellConsumer) {
				if curCell == cell {
					select {
					case queryChan <- msg:
						return true, nil
					default:
						return false, nil
					}
				} else {
					return false, cont
				}
			}
		})
	mailbox.Terminated = mailbox.cellTail.Terminated

	initResult := make(chan error)
	go actor.loop(head, initResult)
	if initErr := <-initResult; initErr == nil {
		return mailbox, nil
	} else {
		return nil, initErr
	}
}

type Actor struct {
	*Mailbox
	queryChan <-chan Msg
	Server
}

func (actor *Actor) loop(head *cc.ChanCellHead, initResult chan error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("\n\nCaught panic in actor!\n Panic:\n  %s\n Stack:\n  %s\n\n",
				p, debug.Stack())
			panic(p)
		}
	}()
	var (
		queryChan <-chan Msg
		queryCell *cc.ChanCell
	)
	chanFun := func(cell *cc.ChanCell) { queryChan, queryCell = actor.queryChan, cell }
	head.WithCell(chanFun)
	terminate, err := actor.Init(actor)
	initResult <- err
	for !terminate {
		if !terminate && err == nil {
			terminate, err = actor.HandleBeat()
		}
		if !terminate && err == nil {
			if msg, ok := <-queryChan; ok {
				terminate, err = actor.HandleMsg(msg)
			} else {
				head.Next(queryCell, chanFun)
			}
		}
		if terminate || err != nil {
			terminate = actor.HandleShutdown(err)
		}
	}
	actor.cellTail.Terminate()
}
