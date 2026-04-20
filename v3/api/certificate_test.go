package api

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.mozilla.org/pkcs7"
)

// makeTestCA generates a minimal self-signed CA certificate for testing.
func makeTestCA(t *testing.T) (*rsa.PrivateKey, *x509.Certificate) {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate CA key: %v", err)
	}
	tmpl := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "Test Root CA"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(10 * 365 * 24 * time.Hour),
		IsCA:                  true,
		BasicConstraintsValid: true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatalf("create CA cert: %v", err)
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatalf("parse CA cert: %v", err)
	}
	return key, cert
}

// makeTestLeaf generates a minimal end-entity certificate signed by the given CA.
func makeTestLeaf(t *testing.T, caKey *rsa.PrivateKey, caCert *x509.Certificate) *x509.Certificate {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate leaf key: %v", err)
	}
	tmpl := &x509.Certificate{
		SerialNumber:          big.NewInt(2),
		Subject:               pkix.Name{CommonName: "test-leaf.example.com"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		IsCA:                  false,
		BasicConstraintsValid: true,
		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, caCert, &key.PublicKey, caKey)
	if err != nil {
		t.Fatalf("create leaf cert: %v", err)
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatalf("parse leaf cert: %v", err)
	}
	return cert
}

// buildP7B constructs a degenerate PKCS7 SignedData (P7B) containing the given
// certificates in the provided order, and returns the base64-encoded result.
func buildP7B(t *testing.T, certs ...*x509.Certificate) string {
	t.Helper()
	sd, err := pkcs7.NewSignedData([]byte{})
	if err != nil {
		t.Fatalf("NewSignedData: %v", err)
	}
	for _, c := range certs {
		sd.AddCertificate(c)
	}
	der, err := sd.Finish()
	if err != nil {
		t.Fatalf("Finish: %v", err)
	}
	return base64.StdEncoding.EncodeToString(der)
}

// TestConvertBase64P7BtoCertificates_RootFirstOrder demonstrates the bug:
// when the P7B contains the root CA as the first certificate (as DigiCert chains
// appear to be returned by Keyfactor), ConvertBase64P7BtoCertificates returns the
// certs in that same order, so certs[0] is the ROOT — not the leaf.
//
// The existing DownloadCertificate code blindly returns certs[0] as the leaf,
// which is wrong. This test reproduces the customer-reported issue where
// common_name, thumbprint etc. are populated with root CA data after a refresh.
func TestConvertBase64P7BtoCertificates_RootFirstOrder(t *testing.T) {
	caKey, caCert := makeTestCA(t)
	leafCert := makeTestLeaf(t, caKey, caCert)

	// Build P7B with root FIRST, leaf SECOND — mirrors what DigiCert chains look
	// like after Keyfactor's ChainOrder:"EndEntityFirst" is ignored by the P7B parser.
	p7bBase64 := buildP7B(t, caCert, leafCert)

	certs, err := ConvertBase64P7BtoCertificates(p7bBase64)
	if err != nil {
		t.Fatalf("ConvertBase64P7BtoCertificates: %v", err)
	}
	if len(certs) != 2 {
		t.Fatalf("expected 2 certs, got %d", len(certs))
	}

	t.Logf("certs[0]: CN=%q IsCA=%v", certs[0].Subject.CommonName, certs[0].IsCA)
	t.Logf("certs[1]: CN=%q IsCA=%v", certs[1].Subject.CommonName, certs[1].IsCA)

	// Document the bug: certs[0] is the root CA, not the leaf.
	// The current DownloadCertificate code returns certs[0] as "leaf", which is wrong.
	if certs[0].Subject.CommonName == "Test Root CA" {
		t.Logf("BUG REPRODUCED: certs[0] is the root CA (%q), not the leaf", certs[0].Subject.CommonName)
	}
	if certs[0].Subject.CommonName == leafCert.Subject.CommonName {
		t.Logf("Order preserved leaf-first — P7B ordering is end-entity-first in this run")
	}

	// The leaf cert must exist somewhere in the list.
	leafFound := false
	for _, c := range certs {
		if c.Subject.CommonName == "test-leaf.example.com" {
			leafFound = true
		}
	}
	if !leafFound {
		t.Error("leaf cert not found in parsed P7B at all")
	}
}

// TestConvertBase64P7BtoCertificates_LeafFirstOrder verifies that when the P7B
// has the leaf first, the current code works correctly (certs[0] = leaf).
// This shows the code is position-dependent, not validated.
func TestConvertBase64P7BtoCertificates_LeafFirstOrder(t *testing.T) {
	caKey, caCert := makeTestCA(t)
	leafCert := makeTestLeaf(t, caKey, caCert)

	// Build P7B with leaf FIRST, root SECOND.
	p7bBase64 := buildP7B(t, leafCert, caCert)

	certs, err := ConvertBase64P7BtoCertificates(p7bBase64)
	if err != nil {
		t.Fatalf("ConvertBase64P7BtoCertificates: %v", err)
	}
	if len(certs) != 2 {
		t.Fatalf("expected 2 certs, got %d", len(certs))
	}

	t.Logf("certs[0]: CN=%q IsCA=%v", certs[0].Subject.CommonName, certs[0].IsCA)
	t.Logf("certs[1]: CN=%q IsCA=%v", certs[1].Subject.CommonName, certs[1].IsCA)

	// When the leaf is first, the current code happens to return the right cert.
	// But this is accidental — it depends on P7B cert ordering.
	if certs[0].Subject.CommonName == leafCert.Subject.CommonName {
		t.Logf("leaf-first P7B: certs[0] is correctly the leaf (by luck of ordering)")
	}
}

// mockDownloadServer returns a TLS test server that responds to any POST with a
// JSON body {"Content": p7bBase64}.
func mockDownloadServer(t *testing.T, p7bBase64 string) *httptest.Server {
	t.Helper()
	body, err := json.Marshal(map[string]string{"Content": p7bBase64})
	if err != nil {
		t.Fatalf("marshal mock response: %v", err)
	}
	return httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(body)
	}))
}

// TestDownloadCertificate_RootFirstP7B is the primary regression test for the
// cert-chain ordering bug. When the server returns a P7B with the root CA first
// (as DigiCert chains do), DownloadCertificate must still return the end-entity
// cert as the leaf — not the root CA.
//
// This test FAILS against the buggy code and PASSES after the fix.
func TestDownloadCertificate_RootFirstP7B(t *testing.T) {
	caKey, caCert := makeTestCA(t)
	leafCert := makeTestLeaf(t, caKey, caCert)

	// Root-first ordering — the problematic case reported by customers using DigiCert.
	server := mockDownloadServer(t, buildP7B(t, caCert, leafCert))
	defer server.Close()

	client := newTestClient(server)
	leaf, chain, _, err := client.DownloadCertificate(1, "", "", "", 0, "P7B")
	if err != nil {
		t.Fatalf("DownloadCertificate: %v", err)
	}
	if leaf == nil {
		t.Fatal("leaf is nil")
	}
	if chain == nil || len(chain) != 2 {
		t.Fatalf("expected chain of length 2, got %v", chain)
	}

	// Core assertion: the returned leaf must be the end-entity, not the CA.
	if leaf.IsCA {
		t.Errorf("DownloadCertificate returned a CA cert as leaf: CN=%q IsCA=%v; want end-entity CN=%q",
			leaf.Subject.CommonName, leaf.IsCA, leafCert.Subject.CommonName)
	}
	if leaf.Subject.CommonName != leafCert.Subject.CommonName {
		t.Errorf("leaf CN = %q, want %q", leaf.Subject.CommonName, leafCert.Subject.CommonName)
	}
}

// TestDownloadCertificate_LeafFirstP7B verifies the common case where the server
// returns a leaf-first P7B (most internal CAs). The fix must not break this.
func TestDownloadCertificate_LeafFirstP7B(t *testing.T) {
	caKey, caCert := makeTestCA(t)
	leafCert := makeTestLeaf(t, caKey, caCert)

	// Leaf-first ordering — the "lucky" path that has always worked.
	server := mockDownloadServer(t, buildP7B(t, leafCert, caCert))
	defer server.Close()

	client := newTestClient(server)
	leaf, chain, _, err := client.DownloadCertificate(1, "", "", "", 0, "P7B")
	if err != nil {
		t.Fatalf("DownloadCertificate: %v", err)
	}
	if leaf == nil {
		t.Fatal("leaf is nil")
	}
	if chain == nil || len(chain) != 2 {
		t.Fatalf("expected chain of length 2, got %v", chain)
	}

	if leaf.IsCA {
		t.Errorf("DownloadCertificate returned a CA cert as leaf: CN=%q IsCA=%v; want end-entity CN=%q",
			leaf.Subject.CommonName, leaf.IsCA, leafCert.Subject.CommonName)
	}
	if leaf.Subject.CommonName != leafCert.Subject.CommonName {
		t.Errorf("leaf CN = %q, want %q", leaf.Subject.CommonName, leafCert.Subject.CommonName)
	}
}
