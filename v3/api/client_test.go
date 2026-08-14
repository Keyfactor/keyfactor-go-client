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
	"net/http"
	"net/http/httptest"
	"net/url"
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

// TestNewKeyfactorClient_PlumbsClientTimeout is a regression test proving that
// a Server.ClientTimeout value survives NewKeyfactorClient's rebuild of the
// CommandAuthConfig. Before this fix, baseConfig never set HttpClientTimeout,
// so the rebuilt auth config (and everything derived from it, including
// BuildTransport's ResponseHeaderTimeout) silently fell back to
// DefaultClientTimeout (60s) regardless of what the caller configured,
// producing "net/http: timeout awaiting response headers" on long-running
// calls such as PFX enrollment.
func TestNewKeyfactorClient_PlumbsClientTimeout(t *testing.T) {
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
