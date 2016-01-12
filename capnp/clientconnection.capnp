using Go = import "go.capnp";

$Go.package("capnp");
$Go.import("goshawkdb.io/common/capnp");

@0xb807d103e02b2b62;

using CTxn = import "clienttransaction.capnp";

struct Hello {
 product   @0: Text;
 version   @1: Text;
 isClient  @2: Bool;
}

struct HelloClientFromServer {
 namespace @0: Data;
 rootId    @1: Data;
}

struct ClientMessage {
  union {
    heartbeat           @0:  Void;
    clientTxnSubmission @1:  CTxn.ClientTxn;
    clientTxnOutcome    @2:  CTxn.ClientTxnOutcome;
  }
}
