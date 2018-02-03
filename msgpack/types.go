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
	Id []byte
	// The Counter is essentially a version indicator: it is used to
	// indicate the version a client is at, and so the client must
	// maintain, per connection, a counter value. The initial value of
	// the Counter is 0. When subscriptions are in use, an update from
	// a subscription can cause the counter to increase. The client
	// must always send a ClientTxn with Counter set to the value of
	// counter when the transaction started on the client, even if
	// updates arrived from the server during the transaction. This
	// value is used to detect when a transaction which is sent to the
	// server concurrently with the server sending an update due to a
	// subscription.
	Counter uint64
	Actions []*ClientAction
}

type ClientTxnOutcome struct {
	// The TxnId with which the client submitted this txn. Outcomes are
	// also used for subscription updates, in which case the Id will be
	// the Id of the transaction that created the subscription. The
	// client will never be sent duplicate updates for an object even
	// if the client has multiple subscriptions for the same object
	// (which is legal). But in this case it is undetermined which of
	// the subscriptions will receive the update; only that one of them
	// will.
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
	// If the outcome is for a subscription then the FinalId will be
	// empty.
	FinalId []byte
	// The updated counter. This value should be used for the next
	// transaction that is started after this outcome is received and
	// processed.
	Counter uint64
	// True iff the transaction committed. Commit, Abort and Error are
	// all mutually exclusive.
	Commit bool
	// If the txn failed to commit, and did not error, then
	// Abort.length will not be undefined. Abort will contain the
	// updates which the client should apply to its cache and then the
	// client should rerun the transaction. Note in some cases (see
	// notes on subscriptions below), the list may have 0 length.
	Abort []*ClientAction
	// If some error occured then this is provided.
	Error string
}

// All ClientActions must specify the VarId and contain an action:
// either a create, xor a read, xor a write, xor a read and write. An
// AddSub requires a read xor create, and a DelSub requires a read.
// When the client sends this type to the server (as part of a Txn
// Submission), it may not use Missing. When the server sends this
// type to the client (as part of a ClientTxnOutcome) only Existing
// with Write, and Missing will be used.
type ClientAction struct {
	// The Object Id. For newly created objects, the client should use
	// a uint64 in Big Endian as the first 8 bytes, followed by the
	// namespace as the latter 12 bytes. The client should maintain
	// independent counters for the created Object Id, and Txn Id.
	VarId []byte
	Value ClientActionValue
	Meta  ClientActionMeta
}

type ClientActionValue struct {
	// Missing is used by the server to indicate that it knows the
	// client's cached copy of this object is out of date, but the
	// server isn't able to provide an updated value at this point. So
	// the client must simply forget about this object, and if it needs
	// to access this object in the future, then it must reload it.
	Missing bool
	// Create is only used by the client when creating objects. It is
	// never used by the server for updates.
	Create *ClientActionValueCreate
	// The server will always use only Missing or Existing for updates.
	Existing *ClientActionValueExisting
}

type ClientActionValueCreate struct {
	Value      []byte
	References []*ClientVarId
}

type ClientActionValueExisting struct {
	// Set true iff the transaction read the current value. Never used
	// in updates.
	Read bool
	// Set to undefined if there was no modification.
	Modify *ClientActionValueExistingModify
}

type ClientActionValueExistingModify struct {
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

type ClientActionMeta struct {
	// To subscribe to an object, set AddSub to true. The action must
	// also be either a read or a create. Once the transaction commits,
	// the server will asynchronously send updates to the client
	// whenever a subscribed object is modified. The updates are not
	// limited just to subscribed objects: if there is another
	// transaction which modifies both a subscribed object and another
	// object that the server knows the client has cached, then the
	// server may choose to send an update to the client that contains
	// updates to both objects.
	// It is legal to have multiple subscriptions to the same
	// object. But see the notes above in ClientTxnOutcome.
	AddSub bool
	// To unsubscribe from an object, set DelSub to the Transaction Id
	// of the transaction in which the subscription was added. Note
	// that this will not cancel any subscription other than the
	// indicated object and transaction Id. I.e. you can use this to
	// remove individual objects from a subscription if you wish. If
	// you have multiple subscriptions to the same object then you must
	// cancel all the subscriptions in different transactions one by
	// one (it is not legal to have a transaction mention the same
	// object multiple times, so you can only cancel one subscription
	// per object in a transaction).
	DelSub []byte
}

// General notes on subscriptions.
//
// Basically, you want to maintain a mapping from object Id to a list
// of transactions Ids (and most likely callbacks) where each
// transaction Id is the Id of the transaction that created the
// subscription. Then, when you get an update from the server for a
// subscription (you can detect it's for a subscription because the
// txnId will be before the current next txnId), just iterate through
// the update actions, update the client cache, and figure out which
// callbacks you need to trigger by using this mapping. The mapping
// will also help when you want to delete subscriptions.
//
// When you submit a transaction to the server, once the server has
// received the transaction, it guarantees that it won't send any
// updates to the client until the outcome of the transaction is
// known. This makes the life of the client much easier as it
// eliminates many race conditions. Of course, there's still the
// possibility that the server sends subscription updates to the
// client at the exact same time that the client submits a transaction
// to the server. The client should process the updates it receives as
// normal (it's welcome to buffer all such updates until after the
// outcome of the transaction is known if it so wishes). The server
// will detect whether those updates invalidate the contents of the
// transaction (this is the purpose of the Counter), and if so, it
// will abort the transaction, asking the client to rerun the
// transaction, but the list of updates will be the empty list.
