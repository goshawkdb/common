source $stdenv/setup

mkdir -p $TMP/src
cd $TMP/src
unpackFile $src
timestamp="$(stat --format=%y ${archiveTimeStampSrc})"
cd $TMP

mkdir $TMP/goshawkdb-server-${goshawkdbVersion}
cp $TMP/src/$license $TMP/goshawkdb-server-${goshawkdbVersion}/LICENSE
chmod 444 $TMP/goshawkdb-server-${goshawkdbVersion}/LICENSE

cp -r $server/bin $TMP/goshawkdb-server-${goshawkdbVersion}/
for t in $(find $TMP/goshawkdb-server-${goshawkdbVersion} -depth | sort); do
  touch -ad "${timestamp}" $t
  touch -md "${timestamp}" $t
done

mkdir -p $out
fakeroot -- tar cJf $out/goshawkdb-server-0.1.tar.xz goshawkdb-server-${goshawkdbVersion}

