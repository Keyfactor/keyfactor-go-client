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
	"encoding/json"
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

// TestCreateStore_DoesNotLogPlaintextServerPasswordInNestedProperties is a
// regression test for a confirmed leak that is distinct from the
// top-level-field leaks covered above: CreateStoreFctArgs (and
// UpdateStoreFctArgs) carry certificate store connection properties as
// PropertiesString, a JSON string PRE-SERIALIZED from the Properties map by
// CreateStore/UpdateStore themselves (see store.go) before the outer struct
// is marshaled again by sendRequest. That produces a JSON *string* value at
// the top level of the outer request object whose own contents are
// themselves JSON containing secrets such as ServerUsername/ServerPassword -
// exactly the shape real callers hit today (confirmed usage:
// terraform-provider-keyfactor's certificate store resource populates
// ServerPassword/ServerUsername in this Properties map on every create and
// update call). redactSensitiveValue's original string case treated any
// string leaf as opaque, so this nested JSON-as-a-string secret sailed
// through unredacted into the [TRACE] request-body log even though the
// top-level "Password" field redacted fine.
//
// This drives the real CreateStore function (not just the redaction helper
// in isolation) against a minimal mock server, so it reproduces exactly what
// a caller creating a certificate store observes in their logs.
func TestCreateStore_DoesNotLogPlaintextServerPasswordInNestedProperties(t *testing.T) {
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"Message":"synthetic failure for test"}`))
	}))
	defer srv.Close()

	c := newTestClient(srv)

	buf := captureGlobalLogOutput(t)

	// Return value is intentionally ignored: the server above always errors,
	// so CreateStore is expected to return a non-nil error. What matters for
	// this test is only what got written to the log before that error was
	// returned.
	_, _ = c.CreateStore(&CreateStoreFctArgs{
		ClientMachine: "test-client-machine",
		StorePath:     "/opt/test/store",
		AgentId:       "11111111-1111-1111-1111-111111111111",
		Properties: map[string]interface{}{
			"ServerUsername": "admin",
			"ServerPassword": canaryRecoveryPassword,
		},
	})

	logged := buf.String()
	if strings.Contains(logged, canaryRecoveryPassword) {
		t.Fatalf(
			"CreateStore logged the plaintext ServerPassword nested inside the pre-serialized Properties JSON string; captured log output:\n%s",
			logged,
		)
	}
}

// TestRedactSensitiveValue_UnwrapsNestedJSONEncodedAsString is a narrower
// unit test against redactSensitiveValue directly (rather than a full
// end-to-end CreateStore drive), pinning down the exact contract: a string
// leaf whose own contents decode to a JSON object or array gets redacted the
// same way a "real" nested object would, and the result is re-marshaled back
// to a string so the log output keeps the "valid JSON encoded as a string"
// shape instead of silently becoming a nested object.
func TestRedactSensitiveValue_UnwrapsNestedJSONEncodedAsString(t *testing.T) {
	outer := map[string]interface{}{
		"ClientMachine": "test-client-machine",
		"Properties":    `{"ServerUsername":"admin","ServerPassword":"` + canaryRecoveryPassword + `"}`,
	}

	redacted := redactSensitiveValue(outer, maxNestedJSONStringDepth)

	redactedBytes, err := json.Marshal(redacted)
	if err != nil {
		t.Fatalf("failed to marshal redacted value: %v", err)
	}
	if strings.Contains(string(redactedBytes), canaryRecoveryPassword) {
		t.Fatalf("redactSensitiveValue did not redact a secret nested inside a JSON-encoded-as-string leaf; got: %s", redactedBytes)
	}

	// The Properties value should still be a JSON string (not have been
	// promoted to a nested object), and it should still be valid JSON once
	// unwrapped, with ServerUsername left intact and only ServerPassword
	// redacted.
	redactedMap, ok := redacted.(map[string]interface{})
	if !ok {
		t.Fatalf("expected top-level redacted value to remain a map, got %T", redacted)
	}
	propertiesVal, ok := redactedMap["Properties"].(string)
	if !ok {
		t.Fatalf("expected Properties to remain a JSON-encoded string, got %T: %v", redactedMap["Properties"], redactedMap["Properties"])
	}
	var reparsed map[string]interface{}
	if err := json.Unmarshal([]byte(propertiesVal), &reparsed); err != nil {
		t.Fatalf("Properties string is no longer valid JSON after redaction: %v", err)
	}
	if reparsed["ServerUsername"] != "admin" {
		t.Fatalf("expected non-sensitive ServerUsername to survive redaction unchanged, got: %v", reparsed["ServerUsername"])
	}
	if reparsed["ServerPassword"] != redactedLogValue {
		t.Fatalf("expected ServerPassword to be redacted to %q, got: %v", redactedLogValue, reparsed["ServerPassword"])
	}
}

// TestRedactSensitiveValue_DepthGuardStopsOnDeeplyNestedJSONStrings ensures
// the recursion depth guard actually bounds how many levels of
// string-encoded-JSON get unwrapped, so adversarial or accidental
// JSON-of-string-of-JSON-of-string-of-... nesting can't cause unbounded
// recursion. It doesn't assert a specific leak/no-leak outcome beyond the
// guard depth (that's an explicit, documented tradeoff) - only that
// redaction terminates and produces valid JSON.
func TestRedactSensitiveValue_DepthGuardStopsOnDeeplyNestedJSONStrings(t *testing.T) {
	// Build a value nested well beyond maxNestedJSONStringDepth: each layer
	// is a JSON object whose single field's value is itself a JSON-encoded
	// string of the next layer down, innermost carrying the canary secret.
	value := canaryRecoveryPassword
	for i := 0; i < maxNestedJSONStringDepth+5; i++ {
		layer, err := json.Marshal(map[string]interface{}{"Secret": value})
		if err != nil {
			t.Fatalf("failed to build nested fixture at layer %d: %v", i, err)
		}
		value = string(layer)
	}

	var decoded interface{}
	if err := json.Unmarshal([]byte(value), &decoded); err != nil {
		t.Fatalf("failed to decode fixture: %v", err)
	}

	redacted := redactSensitiveValue(decoded, maxNestedJSONStringDepth)

	// This must not hang or panic (the real assertion here is that the test
	// completes at all); as a secondary check, confirm the result still
	// marshals to valid JSON.
	if _, err := json.Marshal(redacted); err != nil {
		t.Fatalf("redacted deeply-nested value failed to re-marshal: %v", err)
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
