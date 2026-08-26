package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// TestGetTemplates_Pagination verifies that GetTemplates fetches all pages.
// Before the fix, only the first 50 results were returned; a template sorted
// beyond position 50 (e.g. position 169 out of 278) would never be found by
// name, causing "Error template name not found" in keyfactor_template_role_binding.
func TestGetTemplates_Pagination(t *testing.T) {
	const totalTemplates = 278

	allTemplates := make([]GetTemplateResponse, totalTemplates)
	for i := range allTemplates {
		allTemplates[i] = GetTemplateResponse{
			Id:         i + 1,
			CommonName: "Template-" + strconv.Itoa(i+1),
		}
	}
	// Place the known-problematic late-sorted template at index 168 (position 169).
	allTemplates[168] = GetTemplateResponse{Id: 243, CommonName: "zzz-late-sorted-template"}

	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, _ := strconv.Atoi(r.URL.Query().Get("PageReturned"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("ReturnLimit"))
		if page < 1 {
			page = 1
		}
		if limit < 1 {
			limit = 50
		}
		start := (page - 1) * limit
		end := start + limit
		if start >= len(allTemplates) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("[]"))
			return
		}
		if end > len(allTemplates) {
			end = len(allTemplates)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(allTemplates[start:end])
	}))
	defer srv.Close()

	c := newTestClient(srv)

	templates, err := c.GetTemplates()
	if err != nil {
		t.Fatalf("GetTemplates() error: %v", err)
	}
	if len(templates) != totalTemplates {
		t.Errorf("GetTemplates() returned %d templates, want %d (pagination broken)", len(templates), totalTemplates)
	}

	// Verify the late-sorted target template is present.
	found := false
	for _, tmpl := range templates {
		if tmpl.CommonName == "zzz-late-sorted-template" {
			found = true
			if tmpl.Id != 243 {
				t.Errorf("target template ID = %d, want 243", tmpl.Id)
			}
			break
		}
	}
	if !found {
		t.Errorf("target template %q (position 169) not found — page 2+ results missing", "zzz-late-sorted-template")
	}
}

// TestGetTemplates_MaxPagesGuard verifies that GetTemplates aborts with an error
// when the server always returns a full page (simulating a server that ignores
// pagination and would otherwise cause an infinite loop / unbounded memory growth).
func TestGetTemplates_MaxPagesGuard(t *testing.T) {
	// Build a fixed full page of pageSize (100) items.
	fullPage := make([]GetTemplateResponse, 100)
	for i := range fullPage {
		fullPage[i] = GetTemplateResponse{Id: i + 1, CommonName: "Template-" + strconv.Itoa(i+1)}
	}

	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Always return a full page regardless of PageReturned — simulates a
		// server that ignores paging parameters.
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fullPage)
	}))
	defer srv.Close()

	// Lower the safety bound so the test terminates quickly.
	orig := getTemplatesMaxPages
	getTemplatesMaxPages = 3
	defer func() { getTemplatesMaxPages = orig }()

	c := newTestClient(srv)

	_, err := c.GetTemplates()
	if err == nil {
		t.Fatal("GetTemplates() expected an error when max pages exceeded, got nil")
	}
}

// TestGetTemplates_SinglePage verifies that a sub-pageSize result terminates
// the pagination loop in a single call.
func TestGetTemplates_SinglePage(t *testing.T) {
	calls := 0
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]GetTemplateResponse{{Id: 1, CommonName: "OnlyTemplate"}})
	}))
	defer srv.Close()

	c := newTestClient(srv)

	templates, err := c.GetTemplates()
	if err != nil {
		t.Fatalf("GetTemplates() error: %v", err)
	}
	if len(templates) != 1 {
		t.Errorf("got %d templates, want 1", len(templates))
	}
	if calls != 1 {
		t.Errorf("server called %d times, want 1", calls)
	}
}

// TestGetTemplateResponse_TemplatePolicy_Decode verifies that GetTemplateResponse
// decodes the "TemplatePolicy" object Command returns for templates linked to an
// enrollment pattern (PrimaryKeyAlgorithms/AlternativeKeyAlgorithms, wildcard/key-reuse
// flags, etc). Before the fix GetTemplateResponse had no TemplatePolicy field at all,
// so encoding/json silently dropped it and buildTemplateRoleBindingUpdateArg-style
// callers had nothing to copy forward.
func TestGetTemplateResponse_TemplatePolicy_Decode(t *testing.T) {
	// Trimmed down, real shape captured from a live Command 25.4.1 GET /Templates/{id}
	// response for a template linked to an enrollment pattern.
	body := `{
		"Id": 4,
		"CommonName": "Server_tlsServerAuth-1y",
		"UseAllowedRequesters": true,
		"AllowedRequesters": ["Administrator", "InstanceOwner"],
		"TemplatePolicy": {
			"TemplateId": 4,
			"AllowKeyReuse": true,
			"AllowWildcards": true,
			"RFCEnforcement": null,
			"CertificateOwnerRole": 0,
			"PrimaryKeyAlgorithms": [
				{"name": "RSA", "bit_lengths": [2048, 3072, 4096], "curves": []},
				{"name": "Ed25519", "bit_lengths": [255], "curves": []}
			],
			"AlternativeKeyAlgorithms": []
		}
	}`

	var resp GetTemplateResponse
	if err := json.Unmarshal([]byte(body), &resp); err != nil {
		t.Fatalf("json.Unmarshal() error: %v", err)
	}

	if resp.TemplatePolicy == nil {
		t.Fatal("GetTemplateResponse.TemplatePolicy = nil, want non-nil (field not decoded)")
	}
	if got, want := len(resp.TemplatePolicy.PrimaryKeyAlgorithms), 2; got != want {
		t.Fatalf("len(TemplatePolicy.PrimaryKeyAlgorithms) = %d, want %d", got, want)
	}
	if got, want := resp.TemplatePolicy.PrimaryKeyAlgorithms[0].Name, "RSA"; got != want {
		t.Errorf("PrimaryKeyAlgorithms[0].Name = %q, want %q", got, want)
	}
	if got, want := resp.TemplatePolicy.PrimaryKeyAlgorithms[0].BitLengths, []int{2048, 3072, 4096}; len(got) != len(want) {
		t.Errorf("PrimaryKeyAlgorithms[0].BitLengths = %v, want %v", got, want)
	}
	if resp.TemplatePolicy.AllowKeyReuse == nil || !*resp.TemplatePolicy.AllowKeyReuse {
		t.Errorf("TemplatePolicy.AllowKeyReuse = %v, want true", resp.TemplatePolicy.AllowKeyReuse)
	}
}

// TestUpdateTemplateArg_TemplatePolicy_Roundtrip verifies that an UpdateTemplateArg
// with TemplatePolicy set actually serializes that field onto the wire when passed to
// UpdateTemplate. This is the crux of the "'Policies' cannot be empty" bug fix
// (Keyfactor/terraform-provider-keyfactor#180): Command's PUT /Templates rejects the
// full-replace update outright for templates linked to an enrollment pattern unless
// TemplatePolicy.PrimaryKeyAlgorithms/AlternativeKeyAlgorithms round-trip the
// previously-fetched values. Before the fix, UpdateTemplateArg had no TemplatePolicy
// field, so this could never be sent — this test would fail to compile against the
// pre-fix struct definition.
func TestUpdateTemplateArg_TemplatePolicy_Roundtrip(t *testing.T) {
	var receivedBody []byte

	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		receivedBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(receivedBody)
	}))
	defer srv.Close()

	c := newTestClient(srv)

	allowKeyReuse := true
	allowWildcards := true
	useAllowedRequesters := true
	allowedRequesters := []string{"Administrator", "InstanceOwner"}
	arg := &UpdateTemplateArg{
		Id:                   4,
		UseAllowedRequesters: &useAllowedRequesters,
		AllowedRequesters:    &allowedRequesters,
		TemplatePolicy: &TemplatePolicy{
			TemplateId:     4,
			AllowKeyReuse:  &allowKeyReuse,
			AllowWildcards: &allowWildcards,
			PrimaryKeyAlgorithms: []TemplateKeyAlgorithm{
				{Name: "RSA", BitLengths: []int{2048, 3072, 4096}},
				{Name: "Ed25519", BitLengths: []int{255}},
			},
		},
	}

	if _, err := c.UpdateTemplate(arg); err != nil {
		t.Fatalf("UpdateTemplate() error: %v", err)
	}

	var onWire map[string]interface{}
	if err := json.Unmarshal(receivedBody, &onWire); err != nil {
		t.Fatalf("failed to decode request body sent to server: %v", err)
	}

	policy, ok := onWire["TemplatePolicy"].(map[string]interface{})
	if !ok {
		t.Fatalf("request body sent to server has no TemplatePolicy object; got keys: %v", onWire)
	}
	primaryAlgos, ok := policy["PrimaryKeyAlgorithms"].([]interface{})
	if !ok || len(primaryAlgos) != 2 {
		t.Fatalf("TemplatePolicy.PrimaryKeyAlgorithms on the wire = %v, want 2 entries", policy["PrimaryKeyAlgorithms"])
	}
}

// TestUpdateTemplateArg_KeyUsage_SerializesAsInt verifies that UpdateTemplateArg.KeyUsage
// serializes onto the wire as a JSON number, matching Command's TemplateUpdateRequest
// swagger schema ({"type":"integer","format":"int32"}) confirmed against a live v25.5
// instance. Before the fix, KeyUsage was typed *bool, which serialized as a JSON boolean
// and produced a live HTTP 400 from Command ("Unexpected character encountered while
// parsing value: t. Path 'KeyUsage'"). This also verifies the value returned by
// GetTemplateResponse.KeyUsage (an int) can be assigned directly to
// UpdateTemplateArg.KeyUsage without a type conversion, since both now agree on int.
func TestUpdateTemplateArg_KeyUsage_SerializesAsInt(t *testing.T) {
	var receivedBody []byte

	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		receivedBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(receivedBody)
	}))
	defer srv.Close()

	c := newTestClient(srv)

	// Simulate a real read-modify-write: KeyUsage comes straight off a
	// GetTemplateResponse (int) with no bool<->int conversion required.
	fetched := GetTemplateResponse{Id: 4, KeyUsage: 160} // digitalSignature|keyEncipherment
	keyUsage := fetched.KeyUsage

	arg := &UpdateTemplateArg{
		Id:       4,
		KeyUsage: &keyUsage,
	}

	if _, err := c.UpdateTemplate(arg); err != nil {
		t.Fatalf("UpdateTemplate() error: %v", err)
	}

	var onWire map[string]interface{}
	if err := json.Unmarshal(receivedBody, &onWire); err != nil {
		t.Fatalf("failed to decode request body sent to server: %v", err)
	}

	rawKeyUsage, ok := onWire["KeyUsage"]
	if !ok {
		t.Fatalf("request body sent to server has no KeyUsage field; got keys: %v", onWire)
	}
	switch v := rawKeyUsage.(type) {
	case float64:
		if v != 160 {
			t.Errorf("KeyUsage on the wire = %v, want 160", v)
		}
	case bool:
		t.Fatalf("KeyUsage on the wire is a JSON boolean (%v); Command's API expects an int32 bitmask and returns HTTP 400 for a boolean payload", v)
	default:
		t.Fatalf("KeyUsage on the wire has unexpected type %T (value %v), want a JSON number", v, v)
	}
}
