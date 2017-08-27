package common

import (
	"sync/atomic"
)

type ChannelWaitGroup struct {
	WaitChan chan struct{}
	count    uint64
}

func NewChannelWaitGroup() *ChannelWaitGroup {
	return &ChannelWaitGroup{
		WaitChan: make(chan struct{}, 0),
		count:    0,
	}
}

func (cmg *ChannelWaitGroup) Wait() {
	<-cmg.WaitChan
}

func (cmg *ChannelWaitGroup) Done() {
	select {
	case <-cmg.WaitChan:
		panic("ChannelWaitGroup.Done() called on already completed ChannelWaitGroup.")
	default:
		if atomic.AddUint64(&cmg.count, ^uint64(0)) == 0 {
			close(cmg.WaitChan)
		}
	}
}

func (cmg *ChannelWaitGroup) Add(delta uint64) {
	select {
	case <-cmg.WaitChan:
		panic("ChannelWaitGroup.Add() called on already completed ChannelWaitGroup.")
	default:
		atomic.AddUint64(&cmg.count, delta)
	}
}

func (cmg *ChannelWaitGroup) WaitUntilEither(other <-chan struct{}) {
	select {
	case <-cmg.WaitChan:
	case <-other:
	}
}
