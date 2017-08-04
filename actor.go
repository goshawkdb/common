package common

type MsgSync interface {
	Init() chan struct{}
	Close() bool
}

type MsgSyncQuery struct {
	resultChan chan struct{}
}

func (msq *MsgSyncQuery) Init() chan struct{} {
	msq.resultChan = make(chan struct{})
	return msq.resultChan
}

func (msq *MsgSyncQuery) Close() bool {
	select {
	case <-msq.resultChan:
		return false
	default:
		close(msq.resultChan)
		return true
	}
}
