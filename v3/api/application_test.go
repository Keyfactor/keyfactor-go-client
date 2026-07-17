package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// TestListApplications_Pagination verifies that ListApplications fetches all pages.
// Before the fix, only the first 50 results were returned; a newly created app on
// page 2+ would never be found by name, causing a consistent integration test failure.
func TestListApplications_Pagination(t *testing.T) {
	const totalApps = 150

	allApps := make([]ApplicationListItem, totalApps)
	for i := range allApps {
		allApps[i] = ApplicationListItem{Id: i + 1, Name: "app-" + strconv.Itoa(i+1)}
	}

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
		if start >= len(allApps) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("[]"))
			return
		}
		if end > len(allApps) {
			end = len(allApps)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(allApps[start:end])
	}))
	defer srv.Close()

	c := newTestClient(srv)

	apps, err := c.ListApplications()
	if err != nil {
		t.Fatalf("ListApplications() error: %v", err)
	}
	if len(apps) != totalApps {
		t.Errorf("ListApplications() returned %d apps, want %d (pagination broken)", len(apps), totalApps)
	}
	if apps[len(apps)-1].Id != totalApps {
		t.Errorf("last app ID = %d, want %d (page 2 results missing)", apps[len(apps)-1].Id, totalApps)
	}
}

// TestListApplications_SinglePage verifies a sub-pageSize result terminates the loop.
func TestListApplications_SinglePage(t *testing.T) {
	calls := 0
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]ApplicationListItem{{Id: 1, Name: "only-app"}})
	}))
	defer srv.Close()

	c := newTestClient(srv)

	apps, err := c.ListApplications()
	if err != nil {
		t.Fatalf("ListApplications() error: %v", err)
	}
	if len(apps) != 1 {
		t.Errorf("got %d apps, want 1", len(apps))
	}
	if calls != 1 {
		t.Errorf("server called %d times, want 1", calls)
	}
}
