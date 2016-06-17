using Go = import "go.capnp";
$Go.package("capnp");
$Go.import("goshawkdb.io/common/capnp");

using Java = import "java.capnp";
$Java.package("io.goshawkdb.client.capnp");
$Java.outerClassname("TransactionCap");

@0xd9bafc2f09656f8c;

struct Capabilities {
  value @0: ValueCapability;
  references :group {
    read :union {
      all  @1: Void;
      only @2: List(UInt32);
    }
    write :union {
      all  @3: Void;
      only @4: List(UInt32);
    }
  }
}

enum ValueCapability {
  none      @0;
  read      @1;
  write     @2;
  readWrite @3;
}
