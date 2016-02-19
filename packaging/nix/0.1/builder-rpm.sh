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
%_topdir   %{getenv:HOME}/rpmbuild
%_tmppath  %{_topdir}/tmp

# https://fedoraproject.org/wiki/Packaging:RPMMacros
%_sysconfdir        /etc
%_prefix            /usr
%_exec_prefix       %{_prefix}
%_bindir            %{_exec_prefix}/bin
%_libdir            %{_exec_prefix}/%{_lib}
%_libexecdir        %{_exec_prefix}/libexec
%_sbindir           %{_exec_prefix}/sbin
%_sharedstatedir    /var/lib
%_datarootdir       %{_prefix}/share
%_datadir           %{_datarootdir}
%_includedir        %{_prefix}/include
%_infodir           /usr/share/info
%_mandir            /usr/share/man
%_localstatedir     /var
%_initddir          %{_sysconfdir}/rc.d/init.d
%timestamp          ${timestamp}
EOF

faketime "${timestamp}" rpmbuild -bb SPECS/goshawkdb-server.spec

mkdir -p $out
cp $TMP/home/rpmbuild/RPMS/x86_64/goshawkdb-server-${goshawkdbVersion}-1.x86_64.rpm $out
