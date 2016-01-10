# nix-build 0.1
# Should get 6f76q9ahm991in9715csjs6d3nnbl156-go1.5-goshawkdb-server-dist-bin

with import <nixpkgs> {}; with go15Packages;

let
  skiplist = buildFromGitHub {
    rev = "57733164b18444c51f63e9a80f1693961dde8036";
    owner = "msackman";
    repo = "skiplist";
    sha256 = "1al58q82p2c8cbngrijiwrpbk2vk07k7xzk8r8093hl188jphsbp";
  };

  chancell = buildFromGitHub {
    rev = "f422164a269c10a3ec7495720dc97100d598fb98";
    owner = "msackman";
    repo = "chancell";
    sha256 = "0n9ng03akn8kczfs8ivz7jhig6qzjkblpg29y2fdpl7mrwqbx17n";
  };

  gomdb = buildFromGitHub {
    rev = "e5477472276299169a7eda58c5f0d0d615c758c8";
    owner = "msackman";
    repo = "gomdb";
    sha256 = "04wsdq17qm0mb7lahlqi420p51c5d2gpk79mklmzqhm9gczci5wf";
    propagatedBuildInputs = [ lmdb chancell ];
  };

  rbtree = buildFromGitHub {
    rev = "cd7940bb26b149ce2faf398e7c63fff01aa7b394";
    owner = "glycerine";
    repo = "rbtree";
    sha256 = "1xnl4m9yn998jj1xlws10d9sgq7jvira8i5w4vls1mgphd4hx0zg";
  };

  capnp = buildFromGitHub {
    rev = "db36ab24140ac1737830f5a0d8c87ea26f367fbb";
    owner = "glycerine";
    repo = "go-capnproto";
    sha256 = "1hqc04nak6v5lr305hcpfmvphc5v24p1ds08crf9yabzmqv1fiaj";
    propagatedBuildInputs = [ rbtree ];
  };

  goshawkdb-common = buildGoPackage rec {
    name = "goshawkdb-common";
    goPackagePath = "goshawkdb.io/common";
    rev = "df8631857532";
    src = fetchhg {
      inherit rev;
      url = "https://src.${goPackagePath}";
      sha256 = "0d55ixvaa7xl8jp14xqz60yfja1cw09iz88bpdgcf78mf5vd34cz";
    };
    propagatedBuildInputs = [ capnp ];
  };

  goshawkdb-server = buildGoPackage rec {
    name = "goshawkdb-server";
    goPackagePath = "goshawkdb.io/server";
    rev = "debcf84b047b";
    src = fetchhg {
      inherit rev;
      url = "https://src.${goPackagePath}";
      sha256 = "0f7b65dlpjhjb028mszmxj1wx4nnnz37hf9cv5h2sbxwwwzas0ki";
    };
    propagatedBuildInputs = [ goshawkdb-common capnp skiplist chancell gomdb crypto ];
  };

  goshawkdb-server-dist = buildGoPackage rec {
    name = "goshawkdb-server-dist";
    goPackagePath = "goshawkdb.io/server";
    rev = "debcf84b047b";
    src = fetchhg {
      inherit rev;
      url = "https://src.${goPackagePath}";
      sha256 = "0f7b65dlpjhjb028mszmxj1wx4nnnz37hf9cv5h2sbxwwwzas0ki";
    };
    propagatedBuildInputs = [ goshawkdb-common capnp skiplist chancell gomdb crypto ];
    buildInpots = [ patchelf ];
    postInstall = ''
      find $bin/bin -type f | while read t; do
        patchelf --set-interpreter /lib64/ld-linux-x86-64.so.2 $t
      done
    '';
  };

in
  goshawkdb-server-dist.bin
