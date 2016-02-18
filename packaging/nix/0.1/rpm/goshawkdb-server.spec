Summary: GoshawkDB distributed data store server
Name: goshawkdb-server
Version: 0.1
Release: 1
Requires: lmdb-libs >= 0.9.10
Group: Applications/Databases
License: AGPLv3
SOURCE0: %{buildroot}/%{name}-%{version}.tar.xz

%description
%{summary}

%prep
%setup -q

%build
# Empty section.

%install
rm -rf %{buildroot}
mkdir -p  %{buildroot}

# in builddir
mkdir -p %{buildroot}/usr
cp -a bin %{buildroot}/usr/
mkdir -p %{buildroot}/usr/share/doc/%{name}-%{version}
cp LICENSE %{buildroot}/usr/share/doc/%{name}-%{version}/

%clean
rm -rf %{buildroot}

%files
%defattr(-,root,root,-)
/usr/share/doc/%{name}-%{version}/LICENSE
/usr/bin/consistencychecker
/usr/bin/goshawkdb


%changelog
* Fri Dec 18 2015  Matthew Sackman <matthew@goshawkdb.io> 1.0-1
- Initial release

EOF
