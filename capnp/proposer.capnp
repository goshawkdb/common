using Go = import "go.capnp";

$Go.package("capnp");
$Go.import("goshawkdb.io/common/capnp");

@0xa2ff51a4491e88a6;

struct ProposerState {
  acceptors @0: List(UInt32);
}
