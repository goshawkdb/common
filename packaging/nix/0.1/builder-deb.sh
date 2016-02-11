source $stdenv/setup
mkdir -p $TMP/debian/DEBIAN
cp $debian/DEBIAN/control $TMP/debian/DEBIAN
chmod 755 $TMP/debian/DEBIAN
mkdir -p $TMP/debian/usr
cp -r $server/bin $TMP/debian/usr/
chmod -R 755 $TMP/debian/usr/bin
mkdir -p $out
fakeroot -- dpkg-deb --build debian $out
