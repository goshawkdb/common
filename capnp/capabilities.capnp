using Go = import "go.capnp";
$Go.package("capnp");
$Go.import("goshawkdb.io/common/capnp");

using Java = import "java.capnp";
$Java.package("io.goshawkdb.client.capnp");
$Java.outerClassname("TransactionCap");

@0xd9bafc2f09656f8c;

struct Capability {
  union {
    none      @0: Void;
    read      @1: Void;
    write     @2: Void;
    readWrite @3: Void;
  }
}
