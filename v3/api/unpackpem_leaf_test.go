package api

// Regression tests for UnpackPEM leaf selection. UnpackPEM previously assumed
// certificates[0] was the end-entity leaf, which returns the ROOT when Keyfactor
// Command sends a non-leaf-first PEM bundle (e.g. externally-rooted chains
// returned root-first). It now selects the leaf by chain topology via
// findLeafCert, matching DownloadCertificate.

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"testing"
	"time"
)

func pemCert(c *x509.Certificate) string {
	return string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: c.Raw}))
}

// makeTestIntermediateCA creates an intermediate CA signed by the given parent.
func makeTestIntermediateCA(t *testing.T, parentKey *rsa.PrivateKey, parentCert *x509.Certificate) (*rsa.PrivateKey, *x509.Certificate) {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate intermediate key: %v", err)
	}
	tmpl := &x509.Certificate{
		SerialNumber:          big.NewInt(10),
		Subject:               pkix.Name{CommonName: "Test Intermediate CA"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(5 * 365 * 24 * time.Hour),
		IsCA:                  true,
		BasicConstraintsValid: true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, parentCert, &key.PublicKey, parentKey)
	if err != nil {
		t.Fatalf("create intermediate cert: %v", err)
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatalf("parse intermediate cert: %v", err)
	}
	return key, cert
}

// assertUnpackLeaf runs UnpackPEM and asserts the selected leaf is the expected
// end-entity (non-CA) cert and the CA chain has the expected length.
func assertUnpackLeaf(t *testing.T, bundle, wantCN string, wantChainLen int) {
	t.Helper()
	_, certificate, caCerts, err := UnpackPEM(bundle, "")
	if err != nil {
		t.Fatalf("UnpackPEM: %v", err)
	}
	block, _ := pem.Decode([]byte(certificate))
	if block == nil {
		t.Fatalf("returned leaf is not a PEM block: %q", certificate)
	}
	c, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Fatalf("parse returned leaf: %v", err)
	}
	if c.IsCA {
		t.Errorf("UnpackPEM returned a CA as the leaf: CN=%q IsCA=true", c.Subject.CommonName)
	}
	if c.Subject.CommonName != wantCN {
		t.Errorf("leaf CN = %q, want %q", c.Subject.CommonName, wantCN)
	}
	if len(caCerts) != wantChainLen {
		t.Errorf("caCertificates length = %d, want %d", len(caCerts), wantChainLen)
	}
}

// TestUnpackPEM_LeafSelection verifies the end-entity leaf is selected for every
// bundle ordering, including the root-first orderings that fooled the old
// positional certificates[0] logic.
func TestUnpackPEM_LeafSelection(t *testing.T) {
	rootKey, root := makeTestCA(t)
	intKey, intermediate := makeTestIntermediateCA(t, rootKey, root)
	leaf := makeTestLeaf(t, intKey, intermediate)

	rootPEM := pemCert(root)
	intPEM := pemCert(intermediate)
	leafPEM := pemCert(leaf)

	cases := []struct {
		name     string
		bundle   string
		chainLen int
	}{
		{"root-first-3cert", rootPEM + intPEM + leafPEM, 2}, // the bug trigger
		{"leaf-first-3cert", leafPEM + intPEM + rootPEM, 2},
		{"shuffled-3cert", rootPEM + leafPEM + intPEM, 2},
		{"root-first-2cert", rootPEM + leafPEM, 1}, // the bug trigger
		{"leaf-first-2cert", leafPEM + rootPEM, 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assertUnpackLeaf(t, tc.bundle, "test-leaf.example.com", tc.chainLen)
		})
	}
}

// TestUnpackPEM_WithPrivateKey_RootFirst verifies that, with a private key block
// present in a root-first bundle, both the key is extracted and the leaf (not the
// root) is selected.
func TestUnpackPEM_WithPrivateKey_RootFirst(t *testing.T) {
	rootKey, root := makeTestCA(t)
	leaf := makeTestLeaf(t, rootKey, root)

	k, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}
	keyDER, err := x509.MarshalPKCS8PrivateKey(k)
	if err != nil {
		t.Fatalf("marshal key: %v", err)
	}
	keyPEM := string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyDER}))

	// key + ROOT-first certs
	bundle := keyPEM + pemCert(root) + pemCert(leaf)

	privateKey, certificate, caCerts, err := UnpackPEM(bundle, "")
	if err != nil {
		t.Fatalf("UnpackPEM: %v", err)
	}
	if privateKey == "" {
		t.Error("expected private key to be extracted")
	}
	block, _ := pem.Decode([]byte(certificate))
	if block == nil {
		t.Fatalf("returned leaf is not a PEM block")
	}
	c, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Fatalf("parse returned leaf: %v", err)
	}
	if c.IsCA || c.Subject.CommonName != "test-leaf.example.com" {
		t.Errorf("leaf = CN=%q IsCA=%v, want test-leaf.example.com (non-CA)", c.Subject.CommonName, c.IsCA)
	}
	if len(caCerts) != 1 {
		t.Errorf("caCertificates length = %d, want 1", len(caCerts))
	}
}

// TestUnpackPEM_SingleCert returns the only cert as the leaf, no chain.
func TestUnpackPEM_SingleCert(t *testing.T) {
	rootKey, root := makeTestCA(t)
	leaf := makeTestLeaf(t, rootKey, root)
	_, certificate, caCerts, err := UnpackPEM(pemCert(leaf), "")
	if err != nil {
		t.Fatalf("UnpackPEM: %v", err)
	}
	if certificate == "" {
		t.Fatal("expected a certificate")
	}
	if len(caCerts) != 0 {
		t.Errorf("caCertificates length = %d, want 0", len(caCerts))
	}
}
