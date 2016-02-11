source $stdenv/setup
mkdir -p $out/bin
cp -r $src/bin $out/
for t in $(ls $out/bin); do
  patchelf --set-interpreter /lib64/ld-linux-x86-64.so.2 $out/bin/$t
  patchelf --set-rpath '' $out/bin/$t
  strip -o $out/bin/$t.stripped $out/bin/$t
  mv $out/bin/$t.stripped $out/bin/$t
done
