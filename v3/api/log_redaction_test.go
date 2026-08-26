// Copyright 2026 Keyfactor
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
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// canaryRecoveryPassword is a value distinctive enough that its presence in
// captured log output can only be explained by the password itself leaking,
// not by an unrelated log line coincidentally matching.
const canaryRecoveryPassword = "Sup3rS3cr3t-Canary-Password-DoNotLeak"

// captureGlobalLogOutput redirects Go's global *log* package output (the
// same package this library's client.go/certificate.go call sites use, and
// the same package initLogger redirects to a TerraformLogger/tflog in
// production) to an in-memory buffer for the duration of the test, restoring
// whatever writer was previously configured afterward. This lets tests
// assert on the exact text that would have been logged without needing a
// live tflog/terraform-plugin-log sink.
func captureGlobalLogOutput(t *testing.T) *bytes.Buffer {
	t.Helper()
	var buf bytes.Buffer
	original := log.Writer()
	log.SetOutput(&buf)
	t.Cleanup(func() {
		log.SetOutput(original)
	})
	return &buf
}

// TestRecoverCertificate_DoesNotLogPlaintextPassword is a regression test for
// a confirmed leak: RecoverCertificate's args (including the private-key
// recovery Password) used to be dumped verbatim via
// `log.Println("[DEBUG] RecoverCertificate: Recovering certificate with
// args:", rca)` -- reachable on ordinary Read/Update/import private-key
// recovery paths at DEBUG level, a routine troubleshooting verbosity, not
// something requiring an unusually verbose log level to trigger.
//
// This drives the real RecoverCertificate function (not just the log
// statement in isolation) against a minimal mock server, so it reproduces
// exactly what a caller doing certificate recovery observes in their logs.
// The server intentionally returns a non-2xx status so RecoverCertificate
// exits with an error shortly after the log statement under test, without
// requiring a valid PFX/PKCS12 response body to be fabricated.
func TestRecoverCertificate_DoesNotLogPlaintextPassword(t *testing.T) {
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"Message":"synthetic failure for test"}`))
	}))
	defer srv.Close()

	c := newTestClient(srv)

	buf := captureGlobalLogOutput(t)

	// Return values are intentionally ignored: the server above always
	// errors, so RecoverCertificate is expected to return a non-nil error.
	// What matters for this test is only what got written to the log before
	// that error was returned.
	_, _, _, _, _ = c.RecoverCertificate(123, "", "", "", canaryRecoveryPassword, 0, "PFX")

	logged := buf.String()
	if strings.Contains(logged, canaryRecoveryPassword) {
		t.Fatalf(
			"RecoverCertificate logged the plaintext recovery password via Go's global log package; captured log output:\n%s",
			logged,
		)
	}
}

// TestEnrollPFXV2_DoesNotLogPlaintextPassword is a regression test for a
// confirmed leak: EnrollPFXV2's request struct (whose Payload carries
// ea.Password, the PFX private-key protection password) used to be dumped
// verbatim via `log.Println("[TRACE] Request: ", keyfactorAPIStruct)`.
//
// This drives the real EnrollPFXV2 function against a minimal mock server
// that always errors, so the function returns shortly after the log
// statement under test without requiring a valid enrollment response body.
func TestEnrollPFXV2_DoesNotLogPlaintextPassword(t *testing.T) {
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"Message":"synthetic failure for test"}`))
	}))
	defer srv.Close()

	c := newTestClient(srv)

	buf := captureGlobalLogOutput(t)

	_, _ = c.EnrollPFXV2(&EnrollPFXFctArgsV2{
		Template:      "test-template",
		CertFormat:    "PFX",
		SubjectString: "CN=test.example.com",
		Password:      canaryRecoveryPassword,
	})

	logged := buf.String()
	if strings.Contains(logged, canaryRecoveryPassword) {
		t.Fatalf(
			"EnrollPFXV2 logged the plaintext PFX password via Go's global log package; captured log output:\n%s",
			logged,
		)
	}
}
