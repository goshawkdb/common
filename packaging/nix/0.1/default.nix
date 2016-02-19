# nix-build 0.1

with import <nixpkgs> {}; with go15Packages;

let
  goshawkdbVersion = "0.1";
  self = rec {
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

    lmdb0 = lmdb.overrideDerivation (oldAttrs: {
      postInstall = ''
        mv $out/lib/liblmdb.so $out/lib/liblmdb.so.0.0.0
        ln $out/lib/liblmdb.so.0.0.0 $out/lib/liblmdb.so.0
      '';
    });

    gomdb = buildFromGitHub {
      rev = "e5477472276299169a7eda58c5f0d0d615c758c8";
      owner = "msackman";
      repo = "gomdb";
      sha256 = "04wsdq17qm0mb7lahlqi420p51c5d2gpk79mklmzqhm9gczci5wf";
      propagatedBuildInputs = [ lmdb0 chancell ];
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
      rev = "goshawkdb_${goshawkdbVersion}";
      src = fetchhg {
        inherit rev;
        url = "https://src.${goPackagePath}";
        sha256 = "04fy03rga2nd399grr8ws7fmk3g6gnjqwndclrr1dc87p91yh0jy";
      };
      propagatedBuildInputs = [ capnp ];
    };

    goshawkdb-server = buildGoPackage rec {
      name = "goshawkdb-server";
      goPackagePath = "goshawkdb.io/server";
      src = fetchurl {
        url = "https://src.goshawkdb.io/server/archive/goshawkdb_${goshawkdbVersion}.tar.gz";
        sha256 = "0sq7p5m8aqm1mqdm5qid4lh1hdrn26yi0pcz624crwcyp5nh891k";
      } // {
        archiveTimeStampSrc = "server-goshawkdb_${goshawkdbVersion}/.hg_archival.txt";
        license = "server-goshawkdb_${goshawkdbVersion}/LICENSE";
      };
      buildInputs = [ goshawkdb-common capnp skiplist chancell gomdb crypto ];
    };

    goshawkdb-server-dist = stdenv.mkDerivation {
      name = "goshawkdb-server-dist";
      buildInputs = [ patchelf binutils ];
      src = goshawkdb-server.bin;
      builder = ./builder-patchelf.sh;
    };

    goshawkdb-server-deb = stdenv.mkDerivation rec {
      name = "goshawkdb-server-deb";
      server = goshawkdb-server-dist;
      debian = ./debian;
      builder = ./builder-deb.sh;
      buildInputs = [ dpkg fakeroot ];
      inherit (goshawkdb-server) src;
      inherit (goshawkdb-server.src) archiveTimeStampSrc license;
    };

    goshawkdb-server-tar = stdenv.mkDerivation rec {
      name = "goshawkdb-server-tar";
      server = goshawkdb-server-dist;
      debian = ./debian;
      builder = ./builder-tar.sh;
      buildInputs = [ fakeroot ];
      inherit (goshawkdb-server) src;
      inherit (goshawkdb-server.src) archiveTimeStampSrc license;
      inherit goshawkdbVersion;
    };

    goshawkdb-server-rpm = stdenv.mkDerivation rec {
      name = "goshawkdb-server-rpm";
      server = goshawkdb-server-dist;
      tar = goshawkdb-server-tar;
      spec = ./rpm/goshawkdb-server.spec;
      builder = ./builder-rpm.sh;
      buildInputs = [ rpm file libfaketime ];
      inherit (goshawkdb-server) src;
      inherit (goshawkdb-server.src) archiveTimeStampSrc;
      inherit goshawkdbVersion;
    };
  };
in
  self
