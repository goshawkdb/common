using Go = import "go.capnp";

$Go.package("capnp");
$Go.import("goshawkdb.io/common/capnp");

@0xb30c67519ab66651;

using PTV = import "paxostxnvote.capnp";
using Outcome = import "outcome.capnp";
using CTxn = import "clienttransaction.capnp";
using Txn = import "transaction.capnp";
using TxnCompletion = import "txncompletion.capnp";
using Topology = import "topology.capnp";
using Var = import "var.capnp";

struct Hello {
 product   @0: Text;
 version   @1: Text;
 publicKey @2: Data;
 isClient  @3: Bool;
}

struct HelloFromClient {
 username @0: Text;
 password @1: Data;
}

struct HelloFromServer {
 localHost         @0: Text;
 tieBreak          @1: UInt32;
 namespace         @2: Data;
 topologyDBVersion @3: Data;
 topology          @4: Topology.Topology;
 root              @5: Var.VarIdPos;
}

struct Message {
  union {
    heartbeat           @0:  Void;
    clientTxnSubmission @1:  CTxn.ClientTxn;
    clientTxnOutcome    @2:  CTxn.ClientTxnOutcome;
    txnSubmission       @3:  Txn.Txn;
    submissionOutcome   @4:  Outcome.Outcome;
    submissionComplete  @5:  TxnCompletion.TxnSubmissionComplete;
    submissionAbort     @6:  TxnCompletion.TxnSubmissionAbort;
    oneATxnVotes        @7:  PTV.OneATxnVotes;
    oneBTxnVotes        @8:  PTV.OneBTxnVotes;
    twoATxnVotes        @9:  PTV.TwoATxnVotes;
    twoBTxnVotes        @10: PTV.TwoBTxnVotes;
    txnLocallyComplete  @11: TxnCompletion.TxnLocallyComplete;
    txnGloballyComplete @12: TxnCompletion.TxnGloballyComplete;
  }
}
