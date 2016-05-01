source $stdenv/setup
mkdir -p $out/bin
mkdir -p $TMP
cp -ar $src/bin $TMP/
chmod a+w $TMP/bin
for t in $(ls $TMP/bin); do
  chmod a+w $TMP/bin/$t
  patchelf --set-interpreter /lib64/ld-linux-x86-64.so.2 $TMP/bin/$t
  patchelf --set-rpath '' $TMP/bin/$t
  strip -o $TMP/bin/$t.stripped $TMP/bin/$t
  mv $TMP/bin/$t.stripped $out/bin/$t
done
