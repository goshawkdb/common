source $stdenv/setup

mkdir -p $TMP/src
cd $TMP/src
unpackFile $src
timestamp="$(stat --format=%y ${archiveTimeStampSrc})"
cd $TMP
rm -rf src

mkdir -p $TMP/home/rpmbuild/{RPMS,SRPMS,BUILD,SOURCES,SPECS,tmp}
mkdir $TMP/home/tmp

cp $tar/* $TMP/home/rpmbuild/SOURCES/
cp $spec $TMP/home/rpmbuild/SPECS/goshawkdb-server.spec
cd $TMP/home/rpmbuild
export HOME=$TMP/home

cat <<EOF >~/.rpmmacros
%_topdir   %(echo $HOME)/rpmbuild
%_tmppath  %{_topdir}/tmp
EOF

for t in $(find $TMP/home -depth | sort); do
  touch -ad "${timestamp}" $t
  touch -md "${timestamp}" $t
done

rpmbuild -bb SPECS/goshawkdb-server.spec

mkdir -p $out
cp $TMP/home/rpmbuild/RPMS/x86_64/goshawkdb-server-${goshawkdbVersion}-1.x86_64.rpm $out
