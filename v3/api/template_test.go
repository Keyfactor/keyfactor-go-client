package api

import (
	"encoding/json"
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
