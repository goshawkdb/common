source $stdenv/setup
mkdir -p $TMP/debian/DEBIAN
cp $debian/DEBIAN/control $TMP/debian/DEBIAN
chmod 755 $TMP/debian/DEBIAN
mkdir -p $TMP/debian/usr
cp -r $server/bin $TMP/debian/usr/
cd $TMP/debian
md5sum $(find usr -depth -type f | sort) > $TMP/debian/DEBIAN/md5sums
cd $TMP
chmod -R 755 $TMP/debian/usr/bin
for t in $(find $TMP/debian -depth); do
  touch -at 197001010000 $t
  touch -mt 197001010000 $t
done
mkdir -p $TMP/deb
fakeroot -- dpkg-deb --build debian deb
name=$(ls $TMP/deb)
ar x $TMP/deb/$name
for t in $(ar t $TMP/deb/$name); do
  touch -at 197001010000 $t
  touch -mt 197001010000 $t
done
mkdir -p $out
ar cq $out/$name $(ar t $TMP/deb/$name)
