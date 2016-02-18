source $stdenv/setup

mkdir -p $TMP/src
cd $TMP/src
unpackFile $src
timestamp="$(stat --format=%y ${archiveTimeStampSrc})"
cd $TMP

mkdir -p $TMP/debian/DEBIAN
cp $debian/DEBIAN/control $TMP/debian/DEBIAN
chmod 755 $TMP/debian/DEBIAN
mkdir -p $TMP/debian/usr/share/doc/goshawkdb-server
cp $debian/copyright $TMP/debian/usr/share/doc/goshawkdb-server/copyright
chmod 644 $TMP/debian/usr/share/doc/goshawkdb-server/copyright
cat $TMP/src/$license | sed -e 's/^$/./; s/^/ /;' >> $TMP/debian/usr/share/doc/goshawkdb-server/copyright
cp -r $server/bin $TMP/debian/usr/
cd $TMP/debian
md5sum $(find usr -depth -type f | sort) > $TMP/debian/DEBIAN/md5sums
cd $TMP
chmod -R 755 $TMP/debian/usr/bin
for t in $(find $TMP/debian -depth); do
  touch -ad "${timestamp}" $t
  touch -md "${timestamp}" $t
done
mkdir -p $TMP/deb
fakeroot -- dpkg-deb --build debian deb
rm -rf $TMP/debian

rm -rf $TMP/src

name=$(ls $TMP/deb)
ar x $TMP/deb/$name
for t in $(ar t $TMP/deb/$name); do
  touch -ad "${timestamp}" $t
  touch -md "${timestamp}" $t
done
mkdir -p $out
ar cq $out/$name $(ar t $TMP/deb/$name)
