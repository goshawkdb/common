using Go = import "go.capnp";
$Go.package("capnp");
$Go.import("goshawkdb.io/common/capnp");

using Java = import "java.capnp";
$Java.package("io.goshawkdb.client.capnp");
$Java.outerClassname("TransactionCap");

@0xefb08c285e27a97b;

using Cap = import "capabilities.capnp";

struct ClientTxn {
  id      @0: Data;
  counter @1: UInt64;
  actions @2: List(ClientAction);
}

struct ClientAction {
  varId  @0: Data;
  value :union {
    missing @1: Void;
    create :group {
      value      @2: Data;
      references @3: List(ClientVarIdPos);
    }
    existing :group {
      read @4: Bool;
      modify :union {
        not  @5: Void;
        roll @6: Void;
        write :group {
          value      @7: Data;
          references @8: List(ClientVarIdPos);
        }
      }
    }
  }
  meta :group {
    addSub @9: Bool;  # requires create or read
    delSub @10: Data; # requires read
  }
}

struct ClientTxnOutcome {
  id      @0: Data;
  finalId @1: Data;
  counter @2: UInt64;
  union {
    commit @3: Void;
    abort  @4: List(ClientAction);
    error  @5: Text;
  }
}

struct ClientVarIdPos {
  varId      @0: Data;
  capability @1: Cap.Capability;
}
