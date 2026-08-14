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
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Keyfactor/keyfactor-auth-client-go/auth_providers"
)

// newFakeCommandServer stands in for a Keyfactor Command instance for
// CommandAuthConfigBasic.Authenticate(), which performs a real GET against
// {host}/{apiPath}/Status/Endpoints as part of authentication. It always
// returns 200 with a valid JSON string array, regardless of credentials.
func newFakeCommandServer(t *testing.T) *httptest.Server {
	t.Helper()
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`["endpoint1"]`))
	}))
	t.Cleanup(server.Close)
	return server
}

// isolateKeyfactorEnv unsets ambient KEYFACTOR_* environment variables that
// CommandAuthConfig.ValidateAuthConfig() falls back to whenever the
// corresponding struct field is left at its zero value, restoring their
// original values (present-and-unset, or present-with-value) once the test
// completes. This makes tests that build a Server/CommandAuthConfig with an
// intentionally-zero field (e.g. ClientTimeout: 0 to exercise the "use the
// default" path, or SkipTLSVerify relying on a literal true) hermetic:
// without this, a developer or CI job with KEYFACTOR_CLIENT_TIMEOUT or
// KEYFACTOR_SKIP_VERIFY exported in their shell would get spurious failures
// or, worse, a silently-clobbered SkipVerify that rejects the test's
// self-signed httptest TLS cert.
//
// Note: t.Setenv(key, "") is NOT equivalent to unsetting - os.LookupEnv still
// reports the variable as present with an empty value, which is enough to
// take the "environment variable is set" branch in ValidateAuthConfig (e.g.
// strconv.Atoi("") fails silently and leaves HttpClientTimeout at 0 rather
// than falling through to DefaultClientTimeout). The variable must be
// actually removed from the environment.
func isolateKeyfactorEnv(t *testing.T, keys ...string) {
	t.Helper()
	for _, key := range keys {
		key := key
		originalValue, wasSet := os.LookupEnv(key)
		if err := os.Unsetenv(key); err != nil {
			t.Fatalf("failed to unset %s: %v", key, err)
		}
		t.Cleanup(func() {
			if wasSet {
				_ = os.Setenv(key, originalValue)
			} else {
				_ = os.Unsetenv(key)
			}
		})
	}
}

// TestNewKeyfactorClient_PlumbsClientTimeout is a regression test proving that
// a Server.ClientTimeout value survives NewKeyfactorClient's rebuild of the
// CommandAuthConfig. Before this fix, baseConfig never set HttpClientTimeout,
// so the rebuilt auth config (and everything derived from it, including
// BuildTransport's ResponseHeaderTimeout) silently fell back to
// DefaultClientTimeout (60s) regardless of what the caller configured,
// producing "net/http: timeout awaiting response headers" on long-running
// calls such as PFX enrollment.
func TestNewKeyfactorClient_PlumbsClientTimeout(t *testing.T) {
	isolateKeyfactorEnv(
		t,
		auth_providers.EnvKeyfactorClientTimeout,
		auth_providers.EnvKeyfactorSkipVerify,
		auth_providers.EnvKeyfactorPort,
		auth_providers.EnvKeyfactorCACert,
	)
	server := newFakeCommandServer(t)
	u, uErr := url.Parse(server.URL)
	if uErr != nil {
		t.Fatalf("failed to parse test server URL: %v", uErr)
	}

	cfg := &auth_providers.Server{
		Host:          u.Host,
		Username:      "user",
		Password:      "pass",
		APIPath:       "api",
		SkipTLSVerify: true,
		ClientTimeout: 300,
	}

	ctx := context.Background()
	client, err := NewKeyfactorClient(cfg, &ctx)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	basicCfg, ok := client.AuthClient.(*auth_providers.CommandAuthConfigBasic)
	if !ok {
		t.Fatalf("expected AuthClient to be *auth_providers.CommandAuthConfigBasic, got %T", client.AuthClient)
	}

	if basicCfg.HttpClientTimeout != 300 {
		t.Fatalf("expected HttpClientTimeout to be 300, got %d", basicCfg.HttpClientTimeout)
	}

	transport, tErr := basicCfg.CommandAuthConfig.BuildTransport()
	if tErr != nil {
		t.Fatalf("expected no error building transport, got %v", tErr)
	}

	expected := 300 * time.Second
	if transport.ResponseHeaderTimeout != expected {
		t.Fatalf("expected ResponseHeaderTimeout to be %v, got %v", expected, transport.ResponseHeaderTimeout)
	}
}

// TestNewKeyfactorClient_DefaultClientTimeout confirms the zero-value
// (unset) case still falls back to the library default rather than 0s,
// preserving pre-fix behavior for callers who don't set ClientTimeout.
func TestNewKeyfactorClient_DefaultClientTimeout(t *testing.T) {
	isolateKeyfactorEnv(
		t,
		auth_providers.EnvKeyfactorClientTimeout,
		auth_providers.EnvKeyfactorSkipVerify,
		auth_providers.EnvKeyfactorPort,
		auth_providers.EnvKeyfactorCACert,
	)
	server := newFakeCommandServer(t)
	u, uErr := url.Parse(server.URL)
	if uErr != nil {
		t.Fatalf("failed to parse test server URL: %v", uErr)
	}

	cfg := &auth_providers.Server{
		Host:          u.Host,
		Username:      "user",
		Password:      "pass",
		APIPath:       "api",
		SkipTLSVerify: true,
	}

	ctx := context.Background()
	client, err := NewKeyfactorClient(cfg, &ctx)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	basicCfg, ok := client.AuthClient.(*auth_providers.CommandAuthConfigBasic)
	if !ok {
		t.Fatalf("expected AuthClient to be *auth_providers.CommandAuthConfigBasic, got %T", client.AuthClient)
	}

	if basicCfg.HttpClientTimeout != auth_providers.DefaultClientTimeout {
		t.Fatalf("expected HttpClientTimeout to fall back to default %d, got %d", auth_providers.DefaultClientTimeout, basicCfg.HttpClientTimeout)
	}
}

// perCallTransportAuthConfig is a minimal AuthConfig test double that mimics
// the real behavior of keyfactor-auth-client-go's CommandConfigOauth and
// CommandAuthConfigBasic GetHttpClient() implementations: every call builds a
// brand new *http.Transport (and therefore a brand new, empty connection
// pool) rather than reusing one. It exists to prove that Client caches the
// *http.Client it gets back rather than calling GetHttpClient() (and paying
// for a fresh transport/connection pool) on every request.
type perCallTransportAuthConfig struct {
	server         *httptest.Server
	getClientCalls int32
}

func (a *perCallTransportAuthConfig) GetServerConfig() *auth_providers.Server {
	return &auth_providers.Server{
		Host:          a.server.URL,
		APIPath:       "KeyfactorAPI",
		SkipTLSVerify: true,
	}
}

func (a *perCallTransportAuthConfig) GetHttpClient() (*http.Client, error) {
	atomic.AddInt32(&a.getClientCalls, 1)
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}, nil
}

func (a *perCallTransportAuthConfig) Authenticate() error { return nil }

func (a *perCallTransportAuthConfig) GetCommandVersion() string { return "25.1.0.0" }

// TestClient_ReusesHttpClientAcrossRequests is a regression test for a
// resource leak: sendRequest used to call c.AuthClient.GetHttpClient() on
// every single request. Since the real AuthConfig implementations build a
// brand new http.Transport (and connection pool) per call, and that
// transport's IdleConnTimeout is derived from the configured
// HttpClientTimeout, every API call opened its own never-reused connection
// whose socket lingered until IdleConnTimeout fired - amplified by the fix
// that plumbs a caller-configured ClientTimeout (which can be arbitrarily
// large, e.g. 1800s) all the way through instead of the fixed 60s default.
//
// This test drives Client.sendRequest directly across multiple requests and
// asserts both that AuthConfig.GetHttpClient() is invoked at most once
// (proving the *http.Client is cached) and that the underlying TCP
// connection is reused rather than growing linearly with the request count.
func TestClient_ReusesHttpClientAcrossRequests(t *testing.T) {
	var newConns int32
	srv := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`[]`))
			},
		),
	)
	srv.Config.ConnState = func(_ net.Conn, state http.ConnState) {
		if state == http.StateNew {
			atomic.AddInt32(&newConns, 1)
		}
	}
	srv.StartTLS()
	t.Cleanup(srv.Close)

	auth := &perCallTransportAuthConfig{server: srv}
	client := NewKeyfactorClientWithAuth(auth, nil)

	const requestCount = 10
	for i := 0; i < requestCount; i++ {
		resp, err := client.sendRequest(
			&request{
				Method:   http.MethodGet,
				Endpoint: "CertificateStoreContainers",
				Headers:  &apiHeaders{},
			},
		)
		if err != nil {
			t.Fatalf("request %d failed: %v", i, err)
		}
		// Fully drain and close the body so the underlying transport is free
		// to return the connection to its idle pool for reuse.
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}

	if calls := atomic.LoadInt32(&auth.getClientCalls); calls != 1 {
		t.Fatalf(
			"expected AuthClient.GetHttpClient to be called exactly once across %d requests (client should be cached), got %d calls",
			requestCount,
			calls,
		)
	}

	if conns := atomic.LoadInt32(&newConns); conns > 2 {
		t.Fatalf(
			"expected the TCP connection to be reused across %d sequential requests (at most ~1-2 new connections), observed %d new connections",
			requestCount,
			conns,
		)
	}
}
