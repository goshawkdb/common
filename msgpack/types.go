package msgpack

//go:generate msgp

// Sent by the client to the server once the socket is established,
// and concurrently sent by the server to the client once the socket
// is established.
type Hello struct {
	Product string
	Version string
}

// Sent by the server to the client after the Hello's have been
// exchanged; assuming everything's with the Hello.
type HelloClientFromServer struct {
	// 12 bytes. These form the latter 12 bytes of the 20-byte TxnIds
	// and new VarIds.
	Namespace []byte
	// The roots this client account has been configured with.
	Roots []*Root
}

// Part of the HelloClientFromServer msg
type Root struct {
	Name       string
	VarId      []byte
	Capability *Capability
}

// Indication whether the client can read or write the object in
// question.
type Capability struct {
	Read  bool
	Write bool
}

// General container for messages between client and server once the
// connection is up (i.e. post the HelloClientFromServer msg). The
// client only ever sends ClientMessages with ClientTxnSubmission, and
// the client only ever receives ClientMessages with
// ClientTxnOutcomes.
type ClientMessage struct {
	ClientTxnSubmission *ClientTxn
	ClientTxnOutcome    *ClientTxnOutcome
}

type ClientTxn struct {
	// The TxnId. The first 8 bytes are a uint64 in Big Endian, the
	// remaining 12 bytes are the contents of the Namespace from
	// HelloClientFromServer. The number that is encoded into the first
	// 8 bytes should start from 0, and always increase. See also
	// FinalId below.
	Id      []byte
	Retry   bool // True iff this is a Retry txn
	Actions []*ClientAction
}

type ClientTxnOutcome struct {
	// The TxnId with which the client submitted this txn.
	Id []byte
	// The final TxnId of this txn. The latter 12 bytes are guaranteed
	// to be unchanged. The first 8 bytes should be read as a uint64 in
	// Big Endian. This number is guaranteed to be >= the corresponding
	// 8 bytes in the Id that the client submitted the txn with. The
	// client should update itself to ensure that the next txn it
	// submits has an Id > the FinalId for this number. Gaps are
	// acceptable, but pointless. The reason the FinalId can be
	// different from the Id is that the server may itself encounter
	// situations in which it must resubmit the txn which requires a
	// new txn Id.
	FinalId []byte
	// True iff the transaction committed. Commit, Abort and Error are
	// all mutually exclusive.
	Commit bool
	// If the txn failed to commit, and did not error, then
	// Abort.length will be > 0. Abort will contain the updates which
	// the client should apply to its cache and then the client should
	// rerun the transaction.
	Abort []*ClientUpdate
	// If some error occured then this is provided.
	Error string
}

// Essentially, a ClientUpdate is the write components of some other
// transaction that has collided with this client's own transaction.
type ClientUpdate struct {
	// The TxnId of the other transaction
	Version []byte
	// The actions which should be applied against the client
	// cache. This will only consist of Write and Delete actions.
	Actions []*ClientAction
}

// All ClientActions must specify the VarId. The remaining fields are
// mutually exclusive: exactly 1 of Read, Write, ReadWrite, Create and
// Delete should be specified. When the client sends this type to the
// server (as part of a Txn Submission), Delete is never true. When
// the server sends this type to the client (as part of a
// ClientUpdate), only Write and Delete will ever be used.
type ClientAction struct {
	// The Object Id. For newly created objects, the client should use
	// a uint64 in Big Endian as the first 8 bytes, followed by the
	// namespace as the latter 12 bytes. The client should maintain
	// independent counters for the created Object Id, and Txn Id.
	VarId     []byte
	Read      *ClientActionRead
	Write     *ClientActionWrite
	ReadWrite *ClientActionReadWrite
	Create    *ClientActionCreate
	// Used by the server to indicate that it knows the client's cached
	// copy of this object is out of date, but the server isn't able to
	// provide an updated value at this point. So the client must
	// simply forget about this object, and if it needs to access this
	// object in the future, then it must reload it (i.e. submit a read
	// only transaction with object version 0).
	Delete bool
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

// Pointers in GoshawkDB not only indicate the Id of the object being
// pointed to, but also contain a capability which constains the
// actions that can be performed against the destination object. See
// https://goshawkdb.io/documentation/capabilities.html for more
// details.
type ClientVarId struct {
	VarId      []byte
	Capability *Capability
}
