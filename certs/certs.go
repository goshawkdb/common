package certs

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"net"
	"time"
)

type CertificatePrivateKeyPair struct {
	CertificatePEM string
	PrivateKeyPEM  string
	Certificate    []byte // DER
}

type NodeCertificatePrivateKeyPair struct {
	CertificateRoot *x509.Certificate
	Certificate     []byte // DER
	PrivateKey      *ecdsa.PrivateKey
}

func newKey() (*ecdsa.PrivateKey, []byte, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	privKeyDER, err := x509.MarshalECPrivateKey(privKey)
	if err != nil {
		return nil, nil, err
	}

	return privKey, privKeyDER, nil
}

func getCertificateFields() (time.Time, time.Time, []byte, []byte) {
	notBefore := time.Now()
	// just to be safe, minus one hour
	notBefore = notBefore.Add(-time.Hour)
	// certs last 200 years.
	notAfter := notBefore.AddDate(200, 0, 0)

	serialBytes := make([]byte, 8)
	rand.Read(serialBytes)

	subjectKeyId := make([]byte, 16)
	rand.Read(subjectKeyId)

	return notBefore, notAfter, serialBytes, subjectKeyId
}

func NewClusterCertificate() (*CertificatePrivateKeyPair, error) {
	privKey, privKeyDER, err := newKey()
	if err != nil {
		return nil, err
	}

	notBefore, notAfter, serialBytes, subjectKeyId := getCertificateFields()

	template := x509.Certificate{
		SerialNumber: new(big.Int).SetBytes(serialBytes[:]),
		Subject: pkix.Name{
			Organization: []string{"GoshawkDB"},
			CommonName:   "Cluster CA Root Certificate",
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		BasicConstraintsValid: true,
		SubjectKeyId:          subjectKeyId[:],
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{},
		KeyUsage:              x509.KeyUsageCertSign,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privKey.PublicKey, privKey)
	if err != nil {
		return nil, err
	}

	var certBuf, privBuf bytes.Buffer
	pem.Encode(&certBuf, &pem.Block{Bytes: derBytes, Type: "CERTIFICATE"})
	pem.Encode(&privBuf, &pem.Block{Bytes: privKeyDER, Type: "EC PRIVATE KEY"})

	return &CertificatePrivateKeyPair{
		Certificate:    derBytes,
		CertificatePEM: string(certBuf.Bytes()),
		PrivateKeyPEM:  string(privBuf.Bytes()),
	}, nil
}

func parseCertificate(certPEM []byte) (*x509.Certificate, []byte, error) {
	certBlock, rest := pem.Decode([]byte(certPEM))
	if certBlock == nil || certBlock.Type != "CERTIFICATE" {
		return nil, nil, errors.New("Didn't find expected Certificate in PEM data")
	}

	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}

	return cert, rest, nil
}

func parseCertificatePrivate(certificate []byte) (*x509.Certificate, *ecdsa.PrivateKey, error) {
	cert, rest, err := parseCertificate(certificate)
	if err != nil {
		return nil, nil, err
	}

	privBlock, _ := pem.Decode(rest)
	if privBlock == nil || privBlock.Type != "EC PRIVATE KEY" {
		return nil, nil, errors.New("Didn't find expected EC Private Key in PEM data")
	}

	privKey, err := x509.ParseECPrivateKey(privBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}

	return cert, privKey, nil
}

func NewClientCertificate(certificate []byte) (*CertificatePrivateKeyPair, error) {
	rootCert, rootKey, err := ExtractAndVerifyCertificate(certificate)
	if err != nil {
		return nil, err
	}

	privKey, privKeyDER, err := newKey()
	if err != nil {
		return nil, err
	}

	notBefore, notAfter, serialBytes, subjectKeyId := getCertificateFields()

	template := x509.Certificate{
		SerialNumber: new(big.Int).SetBytes(serialBytes[:]),
		Subject: pkix.Name{
			Organization: []string{"GoshawkDB"},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		BasicConstraintsValid: true,
		SubjectKeyId:          subjectKeyId[:],
		IsCA:                  false,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, rootCert, &privKey.PublicKey, rootKey)
	if err != nil {
		return nil, err
	}

	var certBuf, privBuf bytes.Buffer
	pem.Encode(&certBuf, &pem.Block{Bytes: certDER, Type: "CERTIFICATE"})
	pem.Encode(&privBuf, &pem.Block{Bytes: privKeyDER, Type: "EC PRIVATE KEY"})

	return &CertificatePrivateKeyPair{
		Certificate:    certDER,
		CertificatePEM: string(certBuf.Bytes()),
		PrivateKeyPEM:  string(privBuf.Bytes()),
	}, nil
}

func GenerateNodeCertificatePrivateKeyPair(certificate []byte, host string, ip net.IP, ou string) (*NodeCertificatePrivateKeyPair, error) {
	rootCert, rootKey, err := ExtractAndVerifyCertificate(certificate)
	if err != nil {
		return nil, err
	}

	privKey, _, err := newKey()
	if err != nil {
		return nil, err
	}

	notBefore, notAfter, serialBytes, subjectKeyId := getCertificateFields()

	template := x509.Certificate{
		SerialNumber: new(big.Int).SetBytes(serialBytes[:]),
		Subject: pkix.Name{
			Organization: []string{"GoshawkDB"},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		BasicConstraintsValid: true,
		SubjectKeyId:          subjectKeyId[:],
		IsCA:                  false,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature,
	}
	if len(host) > 0 {
		template.DNSNames = []string{host}
	}
	if ip != nil {
		template.IPAddresses = []net.IP{ip}
	}
	if len(ou) > 0 {
		template.Subject.OrganizationalUnit = []string{ou}
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, rootCert, &privKey.PublicKey, rootKey)
	if err != nil {
		return nil, err
	}

	return &NodeCertificatePrivateKeyPair{
		CertificateRoot: rootCert,
		Certificate:     certDER,
		PrivateKey:      privKey,
	}, nil
}

func ExtractAndVerifyCertificate(certificate []byte) (*x509.Certificate, *ecdsa.PrivateKey, error) {
	cert, privKey, err := parseCertificatePrivate(certificate)
	if err != nil {
		return nil, nil, err
	}
	keyPubKey := privKey.PublicKey
	certPubKey, ok := cert.PublicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, errors.New("Certificate Public Key is of the wrong type")
	}
	if !(certPubKey.Curve == keyPubKey.Curve &&
		certPubKey.X.Cmp(keyPubKey.X) == 0 &&
		certPubKey.Y.Cmp(keyPubKey.Y) == 0) {
		return nil, nil, errors.New("Certificate Public Key is unrelated to private key")
	}
	return cert, privKey, nil
}
