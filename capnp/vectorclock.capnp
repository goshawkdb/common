using Go = import "go.capnp";

$Go.package("capnp");
$Go.import("goshawkdb.io/common/capnp");

@0xf202ad7409b022cf;

struct VectorClock {
  varUuids @0: List(Data);
  values   @1: List(UInt64);
}
