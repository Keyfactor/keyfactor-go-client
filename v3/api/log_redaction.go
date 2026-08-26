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
	"encoding/json"
	"regexp"
)

// sensitiveLogFieldPattern matches JSON object keys that are known or likely
// to carry secret material: certificate/PFX passwords, PAM secret values,
// private key material, API tokens, etc. It is intentionally broad (a
// case-insensitive substring match, not an exact-name allowlist) because this
// package's request payloads are logged generically at TRACE/DEBUG level
// without per-call-site awareness of which fields are sensitive - erring
// toward redacting a borderline non-sensitive field (e.g. "PasswordOptions",
// a policy-description object, not a secret) in a debug log is a far cheaper
// mistake than missing a real secret.
var sensitiveLogFieldPattern = regexp.MustCompile(`(?i)(password|passphrase|secret|privatekey|private_key|pfx|apikey|api_key|accesstoken|access_token|clientsecret|client_secret|token)`)

const redactedLogValue = "[REDACTED]"

// redactSensitiveJSONForLogging returns a copy of jsonBytes (expected to be
// the JSON encoding of an API request/response payload) with the values of
// any object key matching sensitiveLogFieldPattern replaced by
// redactedLogValue, so that logging the payload for troubleshooting doesn't
// also leak certificate/PFX recovery passwords or other secrets into the
// application log (which historically was set to Go's global *log* package,
// bypassing any per-call caller-side masking applied to this library's own
// tflog calls).
//
// This function is used purely to sanitize what gets written to a log line;
// it never mutates or replaces the original bytes used to build the actual
// outgoing HTTP request body.
//
// If jsonBytes doesn't decode as JSON (e.g. it's the literal "null", or
// malformed), it is returned unchanged: there is no structured content to
// redact, and logging a bare scalar isn't a plausible secret-leak vector on
// its own.
func redactSensitiveJSONForLogging(jsonBytes []byte) []byte {
	var decoded interface{}
	if err := json.Unmarshal(jsonBytes, &decoded); err != nil {
		return jsonBytes
	}

	redacted, err := json.Marshal(redactSensitiveValue(decoded))
	if err != nil {
		return jsonBytes
	}
	return redacted
}

// redactSensitiveValue recursively walks a decoded JSON value (as produced by
// encoding/json's default interface{} unmarshaling: map[string]interface{},
// []interface{}, or a scalar), replacing the value of any map key that
// matches sensitiveLogFieldPattern with redactedLogValue.
func redactSensitiveValue(v interface{}) interface{} {
	switch t := v.(type) {
	case map[string]interface{}:
		out := make(map[string]interface{}, len(t))
		for k, val := range t {
			if sensitiveLogFieldPattern.MatchString(k) {
				out[k] = redactedLogValue
			} else {
				out[k] = redactSensitiveValue(val)
			}
		}
		return out
	case []interface{}:
		out := make([]interface{}, len(t))
		for i, val := range t {
			out[i] = redactSensitiveValue(val)
		}
		return out
	default:
		return v
	}
}
