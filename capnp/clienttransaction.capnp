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
  retry   @1: Bool;
  actions @2: List(ClientAction);
}

struct ClientAction {
  varId      @0: Data;
  union {
    read :group {
      version    @1: Data;
    }
    write :group {
      value      @2: Data;
      references @3: List(ClientVarIdPos);
    }
    readwrite :group {
      version    @4: Data;
      value      @5: Data;
      references @6: List(ClientVarIdPos);
    }
    create :group {
      value      @7: Data;
      references @8: List(ClientVarIdPos);
    }
    delete       @9: Void;
    roll :group {
      version    @10: Data;
      value      @11: Data;
      references @12: List(ClientVarIdPos);
    }
  }
}

struct ClientTxnOutcome {
  id      @0: Data;
  finalId @1: Data;
  union {
    commit @2: Void;
    abort  @3: List(ClientUpdate);
    error  @4: Text;
  }
}

struct ClientUpdate {
  version @0: Data;
  actions @1: List(ClientAction);
}

struct ClientVarIdPos {
  varId        @0: Data;
  capabilities @1: Cap.Capabilities;
}
