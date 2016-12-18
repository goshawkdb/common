package msgpack

//go:generate msgp

type Hello struct {
	Product string
	Version string
}

type HelloClientFromServer struct {
	Namespace []byte
	Roots     []*Root
}

type Root struct {
	Name       string
	VarId      []byte
	Capability *Capability
}

type Capability struct {
	Read  bool
	Write bool
}

type ClientMessage struct {
	ClientTxnSubmission *ClientTxn
	ClientTxnOutcome    *ClientTxnOutcome
}

type ClientTxn struct {
	Id      []byte
	Retry   bool
	Actions []*ClientAction
}

type ClientTxnOutcome struct {
	Id      []byte
	FinalId []byte
	Commit  bool
	Abort   []*ClientUpdate
	Error   string
}

type ClientUpdate struct {
	Version []byte
	Actions []*ClientAction
}

type ClientAction struct {
	VarId     []byte
	Read      *ClientActionRead
	Write     *ClientActionWrite
	ReadWrite *ClientActionReadWrite
	Create    *ClientActionCreate
}

type ClientActionRead struct {
	Version []byte
}
type ClientActionWrite struct {
	Value      []byte
	References []*ClientVarId
}
type ClientActionReadWrite struct {
	Version    []byte
	Value      []byte
	References []*ClientVarId
}
type ClientActionCreate struct {
	Value      []byte
	References []*ClientVarId
}

type ClientVarId struct {
	VarId      []byte
	Capability *Capability
}
