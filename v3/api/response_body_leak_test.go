// Copyright 2024 Keyfactor
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Keyfactor/keyfactor-auth-client-go/auth_providers"
	"go.mozilla.org/pkcs7"
)

// leakTestMaxConns mirrors the shape of the production transport built by
// keyfactor-auth-client-go (MaxConnsPerHost: 10, no Client.Timeout) but with a
// tiny cap so a leak of just a few response bodies is enough to wedge the
// client.
const leakTestMaxConns = 2

// leakTestAuthConfig is a minimal AuthConfig whose *http.Client uses a
// transport with a hard per-host connection cap and no client timeout, exactly
// like the production auth providers. A request that never releases its
// connection (an unclosed/undrained response body) permanently consumes one
// slot; once the cap is exhausted the next request blocks forever in
// net/http.(*Transport).getConn.
type leakTestAuthConfig struct {
	server *httptest.Server
	client *http.Client
}

func newLeakTestAuthConfig(t *testing.T, handler http.Handler) *leakTestAuthConfig {
	t.Helper()
	srv := httptest.NewTLSServer(handler)
	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // httptest self-signed cert
		MaxConnsPerHost:     leakTestMaxConns,
		MaxIdleConnsPerHost: leakTestMaxConns,
	}
	a := &leakTestAuthConfig{server: srv, client: &http.Client{Transport: tr}} // Timeout deliberately 0
	t.Cleanup(func() {
		// Unblock anything wedged by a leak before closing so the failing run
		// reports promptly instead of hanging in Server.Close.
		srv.CloseClientConnections()
		tr.CloseIdleConnections()
		srv.Close()
	})
	return a
}

func (a *leakTestAuthConfig) GetServerConfig() *auth_providers.Server {
	return &auth_providers.Server{Host: a.server.URL, APIPath: "KeyfactorAPI", SkipTLSVerify: true}
}
func (a *leakTestAuthConfig) GetHttpClient() (*http.Client, error) { return a.client, nil }
func (a *leakTestAuthConfig) Authenticate() error                  { return nil }
func (a *leakTestAuthConfig) GetCommandVersion() string            { return "25.1.0.0" }

// runWithDeadline runs fn and fails the test if it does not return within d.
// A wedged client never returns, so this is how a connection leak shows up.
func runWithDeadline(t *testing.T, d time.Duration, what string, fn func() error) {
	t.Helper()
	done := make(chan error, 1)
	go func() { done <- fn() }()
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("%s: %v", what, err)
		}
	case <-time.After(d):
		t.Fatalf("%s: did not complete within %s - client is wedged waiting for a connection "+
			"(response bodies are being leaked; MaxConnsPerHost=%d exhausted)", what, d, leakTestMaxConns)
	}
}

// selfSignedP7B returns a base64 PKCS#7 bundle containing one self-signed
// certificate, in the shape Certificates/Download returns it.
func selfSignedP7B(t *testing.T) string {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	tmpl := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "leak-test"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	p7, err := pkcs7.DegenerateCertificate(der)
	if err != nil {
		t.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(p7)
}

// TestDownloadCertificate_DoesNotLeakConnections calls DownloadCertificate far
// more times than the transport's per-host connection cap. Before the fix the
// response body was decoded but never closed, so each call pinned a
// connection and the (cap+1)th call blocked forever.
func TestDownloadCertificate_DoesNotLeakConnections(t *testing.T) {
	p7b := selfSignedP7B(t)
	payload, _ := json.Marshal(map[string]string{"Content": p7b})

	for _, mode := range []string{"content-length", "chunked"} {
		t.Run(mode, func(t *testing.T) {
			var served int32
			auth := newLeakTestAuthConfig(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				atomic.AddInt32(&served, 1)
				w.Header().Set("Content-Type", "application/json")
				if mode == "chunked" {
					// Flushing before the handler returns forces chunked
					// transfer encoding (no Content-Length), like
					// ASP.NET/IIS streaming responses.
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write(payload[:len(payload)/2])
					w.(http.Flusher).Flush()
					_, _ = w.Write(payload[len(payload)/2:])
					// Flush the last data chunk and hold the terminating zero-length
					// chunk back, so a json.Decoder that stops as soon as the value
					// is complete never observes io.EOF (deterministic, not racy).
					w.(http.Flusher).Flush()
					time.Sleep(100 * time.Millisecond)
					return
				}
				w.Header().Set("Content-Length", fmt.Sprint(len(payload)))
				_, _ = w.Write(payload)
			}))
			c := &Client{AuthClient: auth}

			const calls = leakTestMaxConns * 5
			runWithDeadline(t, 10*time.Second, "DownloadCertificate loop", func() error {
				for i := 0; i < calls; i++ {
					leaf, _, _, err := c.DownloadCertificate(i+1, "", "", "", 0, "PEM")
					if err != nil {
						return fmt.Errorf("call %d: %w", i, err)
					}
					if leaf == nil {
						return fmt.Errorf("call %d: nil leaf", i)
					}
				}
				return nil
			})
			if got := atomic.LoadInt32(&served); got != calls {
				t.Fatalf("server saw %d requests, want %d", got, calls)
			}
		})
	}
}

// TestSendRequest_DoesNotLeakConnections drives sendRequest directly for every
// status-code branch, including callers that never touch the returned body
// and error responses whose body the library decodes itself, and asserts the
// connection is released each time.
func TestSendRequest_DoesNotLeakConnections(t *testing.T) {
	bigBody := `{"Message":"` + strings.Repeat("x", 64*1024) + `"}`

	cases := []struct {
		name       string
		status     int
		body       string
		wantErr    bool
		chunked    bool
		readsBody  bool
		wantBodyEq string
	}{
		{name: "200 body never read by caller", status: 200, body: bigBody},
		{name: "200 chunked body never read by caller", status: 200, body: bigBody, chunked: true},
		{name: "200 body fully decoded by caller", status: 200, body: `{"ok":true}`, readsBody: true, wantBodyEq: `{"ok":true}`},
		{name: "204 no content", status: 204},
		{name: "404 json error body", status: 404, body: bigBody, wantErr: true},
		{name: "404 non-json body", status: 404, body: "<html>nope</html>", wantErr: true},
		{name: "401 body", status: 401, body: bigBody, wantErr: true},
		{name: "500 json error body", status: 500, body: bigBody, wantErr: true},
		{name: "500 chunked json error body", status: 500, body: bigBody, wantErr: true, chunked: true},
		{name: "400 non-json body", status: 400, body: "bad", wantErr: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			auth := newLeakTestAuthConfig(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if tc.chunked {
					w.WriteHeader(tc.status)
					half := len(tc.body) / 2
					_, _ = w.Write([]byte(tc.body[:half]))
					w.(http.Flusher).Flush()
					_, _ = w.Write([]byte(tc.body[half:]))
					w.(http.Flusher).Flush()
					time.Sleep(100 * time.Millisecond) // hold back the terminating chunk
					return
				}
				if tc.status != 204 {
					w.Header().Set("Content-Length", fmt.Sprint(len(tc.body)))
				}
				w.WriteHeader(tc.status)
				if tc.status != 204 {
					_, _ = w.Write([]byte(tc.body))
				}
			}))
			c := &Client{AuthClient: auth}

			runWithDeadline(t, 10*time.Second, "sendRequest loop", func() error {
				for i := 0; i < leakTestMaxConns*5; i++ {
					resp, err := c.sendRequest(&request{
						Method:   "GET",
						Endpoint: "Leak/Test",
						Headers:  &apiHeaders{},
					})
					if tc.wantErr {
						if err == nil {
							return fmt.Errorf("call %d: expected error for status %d", i, tc.status)
						}
						if strings.Contains(err.Error(), "context deadline") {
							return fmt.Errorf("call %d: unexpected transport error: %w", i, err)
						}
						continue
					}
					if err != nil {
						return fmt.Errorf("call %d: %w", i, err)
					}
					if tc.readsBody {
						var v map[string]any
						if dErr := json.NewDecoder(resp.Body).Decode(&v); dErr != nil {
							return fmt.Errorf("call %d: decode: %w", i, dErr)
						}
						if v["ok"] != true {
							return fmt.Errorf("call %d: wrong body %v", i, v)
						}
					}
					// Callers may or may not Close; both must be safe.
					if i%2 == 0 {
						_ = resp.Body.Close()
					}
				}
				return nil
			})
		})
	}
}
