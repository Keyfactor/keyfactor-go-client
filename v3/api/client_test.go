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
	"strings"
	"sync"
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

// TestClient_ConcurrentRequestsNotCappedByMaxConnsPerHost is an end-to-end
// regression test closing the loop on a finding this package's own
// http.Client-caching fix (see TestClient_ReusesHttpClientAcrossRequests)
// caused but could not fix locally: caching a single *http.Client means the
// transport keyfactor-auth-client-go builds is now reused for the lifetime
// of the Client instance, so any nonzero MaxConnsPerHost on that transport
// stops being a harmless per-request default and becomes a permanent,
// unqueued-timeout ceiling on concurrent in-flight requests for the whole
// process - e.g. `terraform apply -parallelism=25` would silently serialize
// into batches of N with no bound on how long excess requests wait, since
// neither the cached client's Timeout (0, unset) nor its requests' contexts
// impose one.
//
// keyfactor-auth-client-go previously hardcoded MaxConnsPerHost: 10 on this
// transport. Its own fix (auth_core.go's newHTTPTransport, now pinned at 0 /
// unbounded to match net/http.DefaultTransport) was verified from that
// repo's side by inspecting the constructed *http.Transport's field value.
// This test verifies the fix end-to-end from this repo's perspective instead
// of trusting that inspection alone: it builds a real Client via
// NewKeyfactorClient (exactly as production code does), retrieves its cached
// *http.Client via getHttpClient(), and drives 25 concurrent requests through
// it against a real httptest server, asserting the server actually observes
// well more than 10 requests in flight at once rather than serializing into
// batches of 10.
func TestClient_ConcurrentRequestsNotCappedByMaxConnsPerHost(t *testing.T) {
	const (
		concurrentRequests = 25
		holdDuration       = 200 * time.Millisecond
	)

	var (
		mu          sync.Mutex
		inFlight    int
		maxInFlight int
	)

	srv := httptest.NewTLSServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				// The initial CommandAuthConfigBasic.Authenticate() call made
				// by NewKeyfactorClient below hits this same handler; letting
				// it fall through the same slow path is harmless since it
				// happens once, sequentially, before the concurrent phase
				// starts timing anything.
				mu.Lock()
				inFlight++
				if inFlight > maxInFlight {
					maxInFlight = inFlight
				}
				mu.Unlock()

				time.Sleep(holdDuration)

				mu.Lock()
				inFlight--
				mu.Unlock()

				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("x-keyfactor-product-version", "99.9.9")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`["endpoint1"]`))
			},
		),
	)
	t.Cleanup(srv.Close)

	u, uErr := url.Parse(srv.URL)
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
		t.Fatalf("NewKeyfactorClient failed: %v", err)
	}

	httpClient, hErr := client.getHttpClient()
	if hErr != nil {
		t.Fatalf("getHttpClient failed: %v", hErr)
	}

	// Reset the counters: the single sequential Authenticate() call above
	// already touched inFlight/maxInFlight and this resets the baseline so
	// the assertion below reflects only the concurrent phase.
	mu.Lock()
	inFlight = 0
	maxInFlight = 0
	mu.Unlock()

	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req, rErr := http.NewRequest(http.MethodGet, srv.URL+"/KeyfactorAPI/concurrent-probe", nil)
			if rErr != nil {
				t.Errorf("failed to build request: %v", rErr)
				return
			}
			resp, dErr := httpClient.Do(req)
			if dErr != nil {
				t.Errorf("request failed: %v", dErr)
				return
			}
			_ = resp.Body.Close()
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)

	mu.Lock()
	observedMax := maxInFlight
	mu.Unlock()

	t.Logf(
		"observed max in-flight requests: %d/%d, wall time: %s (old MaxConnsPerHost=10 cap measured ~10 in-flight/909ms for a comparable batch; unbounded measured ~315ms)",
		observedMax,
		concurrentRequests,
		elapsed,
	)

	// The old hardcoded MaxConnsPerHost: 10 would cap this at exactly 10
	// no matter how many requests are fired concurrently. Assert well above
	// that ceiling (comfortably below concurrentRequests to tolerate
	// scheduler jitter) to prove the requests are not being serialized into
	// batches of 10.
	const minAcceptableMaxInFlight = 15
	if observedMax <= 10 {
		t.Fatalf(
			"expected max concurrent in-flight requests to exceed the old MaxConnsPerHost=10 ceiling, got %d (elapsed %s) - concurrency ceiling regression",
			observedMax,
			elapsed,
		)
	}
	if observedMax < minAcceptableMaxInFlight {
		t.Fatalf(
			"expected max concurrent in-flight requests to be close to %d (unbounded), got only %d (elapsed %s)",
			concurrentRequests,
			observedMax,
			elapsed,
		)
	}

	// Wall time is a secondary signal: fully serialized into batches of 10
	// would take ceil(25/10)*holdDuration ~= 3*200ms = 600ms; unbounded
	// concurrency should complete in roughly one holdDuration plus overhead.
	maxAcceptableElapsed := holdDuration * 2
	if elapsed > maxAcceptableElapsed {
		t.Fatalf(
			"expected wall time close to a single %s hold duration for unbounded concurrency, got %s (elapsed too long, suggests serialization)",
			holdDuration,
			elapsed,
		)
	}
}

// TestSendRequest_ContextDeadlineExceeded_NoSilentRetry reproduces the first half
// of a confirmed HIGH-severity bug in sendRequest: on a client-side timeout
// ("context deadline exceeded"), the function used to transparently retry the
// request (up to MAX_CONTEXT_DEADLINE_RETRIES times) and, if a retry
// succeeded, return that success with no indication a timeout ever happened.
//
// That silently swallowed the fact that the *original* request may have
// already succeeded server-side (e.g. a certificate enrollment), and
// defeated callers -- like terraform-provider-keyfactor's orphaned-PFX
// recovery logic -- that specifically match on "context deadline exceeded"
// to trigger their own safe recovery search instead of blindly re-issuing a
// non-idempotent request.
//
// This test sets up a server that only delays its FIRST response beyond the
// client timeout and answers instantly after that -- i.e. exactly the shape
// that used to be masked into a silent "success after retry". It asserts
// sendRequest now returns the timeout error immediately, without contacting
// the server a second time.
func TestSendRequest_ContextDeadlineExceeded_NoSilentRetry(t *testing.T) {
	var requestCount int32

	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&requestCount, 1)
		if n == 1 {
			// First request: sleep well past the client timeout below, so the
			// client gives up on it -- but note the server continues to
			// process it and will "complete" it right after. A retry that
			// followed would succeed immediately, which is exactly the
			// scenario that used to be silently swallowed.
			time.Sleep(300 * time.Millisecond)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	c := &Client{
		AuthClient: &mockAuthConfig{
			serverConfig: newTestServerConfig(srv),
			httpClient:   testHTTPClientWithTimeout(srv, 50*time.Millisecond),
		},
	}

	resp, err := c.sendRequest(&request{Method: http.MethodGet, Endpoint: "test", Headers: &apiHeaders{}})

	if err == nil {
		t.Fatalf("sendRequest() returned nil error, want a context-deadline-exceeded error (retry must not silently succeed)")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("sendRequest() error = %q, want it to contain %q", err.Error(), "context deadline exceeded")
	}
	if resp != nil {
		t.Fatalf("sendRequest() returned a non-nil response alongside an error: %+v", resp)
	}
	if got := atomic.LoadInt32(&requestCount); got != 1 {
		t.Fatalf(
			"server was hit %d time(s), want exactly 1 -- sendRequest must not silently retry a timed-out request",
			got,
		)
	}
}

// TestSendRequest_ContextDeadlineExceeded_NoPanic reproduces the second half
// of the bug: when every attempt fails with a "context deadline exceeded"
// shaped error, sendRequest's response variable was never reassigned from
// its original nil value, and the surrounding switch had no `return` for
// this case -- so control fell through to `resp.StatusCode` on a nil
// *http.Response, panicking the caller (e.g. crashing `terraform apply`
// outright) instead of returning a normal, handleable error.
//
// This test uses a server that always delays past the client timeout, so
// every attempt sendRequest makes hits the same failure shape. It asserts a
// clean error is returned -- never a panic.
func TestSendRequest_ContextDeadlineExceeded_NoPanic(t *testing.T) {
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(300 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c := &Client{
		AuthClient: &mockAuthConfig{
			serverConfig: newTestServerConfig(srv),
			httpClient:   testHTTPClientWithTimeout(srv, 50*time.Millisecond),
		},
	}

	var (
		resp *http.Response
		err  error
	)

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("sendRequest() panicked: %v", r)
			}
		}()
		resp, err = c.sendRequest(&request{Method: http.MethodGet, Endpoint: "test", Headers: &apiHeaders{}})
	}()

	if err == nil {
		t.Fatalf("sendRequest() returned nil error, want a context-deadline-exceeded error")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("sendRequest() error = %q, want it to contain %q", err.Error(), "context deadline exceeded")
	}
	if resp != nil {
		t.Fatalf("sendRequest() returned a non-nil response alongside an error: %+v", resp)
	}
}

// newTestServerConfig builds a minimal *auth_providers.Server pointing at the
// given httptest server, matching the pattern used by newTestClient in
// pam_types_test.go.
func newTestServerConfig(server *httptest.Server) *auth_providers.Server {
	return &auth_providers.Server{
		Host:          server.URL,
		APIPath:       "/KeyfactorAPI",
		SkipTLSVerify: true,
	}
}

// testHTTPClientWithTimeout returns an *http.Client that trusts the given
// httptest TLS server's certificate (sendRequest always forces the https
// scheme, so a plain httptest.NewServer can't be used directly) but with a
// short overall Timeout so requests to a deliberately slow handler produce a
// real "context deadline exceeded" error, identical in shape to what a
// production client sees against a genuinely slow/unreachable Command
// server.
func testHTTPClientWithTimeout(server *httptest.Server, timeout time.Duration) *http.Client {
	base := server.Client()
	return &http.Client{
		Transport: base.Transport,
		Timeout:   timeout,
	}
}
