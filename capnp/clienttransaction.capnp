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
    unmodified   @1: Void;
    modified :group {
      value      @2: Data;
      references @3: List(ClientVarIdPos);
    }
  }
  actionType @4: ClientActionType;
}

enum ClientActionType {
  create          @0;
  readOnly        @1;
  writeOnly       @2;
  readWrite       @3;
  delete          @4;
  roll            @5;
  addSubscription @6;
  delSubscription @7;
}

struct ClientTxnOutcome {
  id      @0: Data;
  finalId @1: Data;
  union {
    commit @2: Void;
    abort  @3: List(ClientAction);
    error  @4: Text;
  }
}

struct ClientVarIdPos {
  varId      @0: Data;
  capability @1: Cap.Capability;
}
