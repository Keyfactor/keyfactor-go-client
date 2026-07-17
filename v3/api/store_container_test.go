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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// TestGetStoreContainers_Pagination verifies that GetStoreContainers fetches
// all pages. Before the fix, only the server's first page (default page size)
// was returned; a container sorted beyond that first page would never be
// found by ID via the list-endpoint fallback in the provider's
// lookupContainerNameByID, causing the container name to silently read back
// as null even though the container assignment itself was intact.
func TestGetStoreContainers_Pagination(t *testing.T) {
	const totalContainers = 150

	allContainers := make([]CertStoreContainer, totalContainers)
	for i := range allContainers {
		id := i + 1
		allContainers[i] = CertStoreContainer{
			Id:   &id,
			Name: "Container-" + strconv.Itoa(id),
		}
	}
	// Place the known-problematic late-sorted container at index 118 (position 119, past a 100-item first page).
	lateId := 999
	allContainers[118] = CertStoreContainer{Id: &lateId, Name: "zzz-late-sorted-container"}

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
		if start >= len(allContainers) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("[]"))
			return
		}
		if end > len(allContainers) {
			end = len(allContainers)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(allContainers[start:end])
	}))
	defer srv.Close()

	c := newTestClient(srv)

	containers, err := c.GetStoreContainers()
	if err != nil {
		t.Fatalf("GetStoreContainers() error: %v", err)
	}
	if containers == nil || len(*containers) != totalContainers {
		got := 0
		if containers != nil {
			got = len(*containers)
		}
		t.Errorf("GetStoreContainers() returned %d containers, want %d (pagination broken)", got, totalContainers)
	}

	found := false
	for _, container := range *containers {
		if container.Name == "zzz-late-sorted-container" {
			found = true
			if container.Id == nil || *container.Id != lateId {
				t.Errorf("target container Id = %v, want %d", container.Id, lateId)
			}
			break
		}
	}
	if !found {
		t.Errorf("target container %q (position 119) not found — page 2+ results missing", "zzz-late-sorted-container")
	}
}

// TestGetStoreContainers_MaxPagesGuard verifies that GetStoreContainers
// aborts with an error when the server always returns a full page
// (simulating a server that ignores pagination and would otherwise cause an
// infinite loop / unbounded memory growth).
func TestGetStoreContainers_MaxPagesGuard(t *testing.T) {
	fullPage := make([]CertStoreContainer, 100)
	for i := range fullPage {
		id := i + 1
		fullPage[i] = CertStoreContainer{Id: &id, Name: "Container-" + strconv.Itoa(id)}
	}

	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Always return a full page regardless of PageReturned — simulates a
		// server that ignores paging parameters.
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fullPage)
	}))
	defer srv.Close()

	// Lower the safety bound so the test terminates quickly.
	orig := getStoreContainersMaxPages
	getStoreContainersMaxPages = 3
	defer func() { getStoreContainersMaxPages = orig }()

	c := newTestClient(srv)

	_, err := c.GetStoreContainers()
	if err == nil {
		t.Fatal("GetStoreContainers() expected an error when max pages exceeded, got nil")
	}
}

// TestGetStoreContainers_SinglePage verifies that a sub-pageSize result
// terminates the pagination loop in a single call.
func TestGetStoreContainers_SinglePage(t *testing.T) {
	calls := 0
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		id := 1
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]CertStoreContainer{{Id: &id, Name: "OnlyContainer"}})
	}))
	defer srv.Close()

	c := newTestClient(srv)

	containers, err := c.GetStoreContainers()
	if err != nil {
		t.Fatalf("GetStoreContainers() error: %v", err)
	}
	if containers == nil || len(*containers) != 1 {
		t.Errorf("got %v containers, want 1", containers)
	}
	if calls != 1 {
		t.Errorf("server called %d times, want 1", calls)
	}
}
