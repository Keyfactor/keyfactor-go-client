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

// maxNestedJSONStringDepth bounds how many levels of "JSON encoded as a
// string value" redactSensitiveValue will attempt to unmarshal and recurse
// into. Several request structs (e.g. CreateStoreFctArgs/UpdateStoreFctArgs's
// PropertiesString field) are populated by pre-serializing a
// map[string]interface{} to a JSON string before the outer struct itself is
// marshaled by sendRequest, producing a JSON string leaf whose *contents* are
// themselves JSON containing secrets (e.g. ServerUsername/ServerPassword).
// Without unwrapping these string-encoded-JSON leaves, redaction never sees
// the nested keys and secrets sail through unredacted into TRACE logs. The
// depth guard exists purely so adversarial or accidentally-deep
// string-of-JSON-of-string-of-JSON... nesting can't recurse unboundedly; a
// legitimate payload should never come close to this limit.
const maxNestedJSONStringDepth = 5

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

	redacted, err := json.Marshal(redactSensitiveValue(decoded, maxNestedJSONStringDepth))
	if err != nil {
		return jsonBytes
	}
	return redacted
}

// redactSensitiveValue recursively walks a decoded JSON value (as produced by
// encoding/json's default interface{} unmarshaling: map[string]interface{},
// []interface{}, a string, or another scalar), replacing the value of any
// map key that matches sensitiveLogFieldPattern with redactedLogValue.
//
// String leaves get one extra check: some request structs pre-serialize a
// map to a JSON string before the outer struct is marshaled again (e.g.
// CreateStoreFctArgs/UpdateStoreFctArgs's PropertiesString field), so a
// string leaf's own contents may themselves be JSON carrying secrets
// (ServerUsername/ServerPassword, etc.) that would otherwise sail through
// unredacted. If a string leaf successfully unmarshals as a
// map[string]interface{} or []interface{}, it is redacted the same way and
// re-marshaled back to a string, preserving its "valid JSON encoded as a
// string" shape in the log output. If it doesn't decode to one of those two
// container types (including "it isn't valid JSON at all"), it is left
// alone: that's the common case of an ordinary string value, not a nested
// payload to unwrap.
//
// remainingDepth bounds how many further levels of string-encoded-JSON will
// be unwrapped, so adversarial or accidentally deep nesting
// (JSON-of-string-of-JSON-of-string-of-...) can't recurse unboundedly. Once
// it reaches zero, string leaves are left as-is without attempting to parse
// them further; map/slice recursion is unaffected by this guard since it can
// only nest as deeply as the decoded value's own structure allows.
func redactSensitiveValue(v interface{}, remainingDepth int) interface{} {
	switch t := v.(type) {
	case map[string]interface{}:
		out := make(map[string]interface{}, len(t))
		for k, val := range t {
			if sensitiveLogFieldPattern.MatchString(k) {
				out[k] = redactedLogValue
			} else {
				out[k] = redactSensitiveValue(val, remainingDepth)
			}
		}
		return out
	case []interface{}:
		out := make([]interface{}, len(t))
		for i, val := range t {
			out[i] = redactSensitiveValue(val, remainingDepth)
		}
		return out
	case string:
		if remainingDepth <= 0 {
			return t
		}
		var nested interface{}
		if err := json.Unmarshal([]byte(t), &nested); err != nil {
			return t
		}
		switch nested.(type) {
		case map[string]interface{}, []interface{}:
			redactedNested := redactSensitiveValue(nested, remainingDepth-1)
			reMarshaled, err := json.Marshal(redactedNested)
			if err != nil {
				return t
			}
			return string(reMarshaled)
		default:
			// Decoded to a bare scalar (a JSON string/number/bool/null that
			// happens to be valid JSON on its own, e.g. the string "42" or
			// `"true"`) rather than a container - nothing to redact inside
			// it, and re-marshaling would just be a no-op wrapped in
			// pointless work. Leave the original string leaf unchanged.
			return t
		}
	default:
		return v
	}
}
