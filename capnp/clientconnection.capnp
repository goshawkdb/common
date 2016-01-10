using Go = import "go.capnp";

$Go.package("capnp");
$Go.import("goshawkdb.io/common/capnp");

@0xb807d103e02b2b62;

using Topology = import "topology.capnp";
using VarId = import "varid.capnp";
using CTxn = import "clienttransaction.capnp";

struct Hello {
 product   @0: Text;
 version   @1: Text;
 isClient  @2: Bool;
}

struct HelloFromServer {
 localHost         @0: Text;
 tieBreak          @1: UInt32;
 namespace         @2: Data;
 topologyDBVersion @3: Data;
 topology          @4: Topology.Topology;
 root              @5: VarId.VarIdPos;
}

struct ClientMessage {
  union {
    heartbeat           @0:  Void;
    clientTxnSubmission @1:  CTxn.ClientTxn;
    clientTxnOutcome    @2:  CTxn.ClientTxnOutcome;
  }
}
