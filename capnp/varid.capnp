using Go = import "go.capnp";

$Go.package("capnp");
$Go.import("goshawkdb.io/common/capnp");

@0xdc0a03440ff7a67f;

struct VarIdPos {
  id        @0: Data;
  positions @1: List(UInt8);
}
