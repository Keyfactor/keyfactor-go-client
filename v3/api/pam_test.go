// Copyright 2025 Keyfactor
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
	"strings"
	"testing"
)

func TestListPAMProviders(t *testing.T) {
	tests := []struct {
		name           string
		query          *GetPAMProviderQuery
		mockResponse   interface{}
		mockStatusCode int
		wantErr        bool
		wantCount      int
	}{
		{
			name:  "successful list without query",
			query: nil,
			mockResponse: []ProviderResponseLegacy{
				mockProviderResponseLegacy,
			},
			mockStatusCode: http.StatusOK,
			wantErr:        false,
			wantCount:      1,
		},
		{
			name: "successful list with query",
			query: &GetPAMProviderQuery{
				QueryString:  "Name -eq 'Test'",
				ReturnLimit:  10,
				PageReturned: 1,
			},
			mockResponse: []ProviderResponseLegacy{
				mockProviderResponseLegacy,
			},
			mockStatusCode: http.StatusOK,
			wantErr:        false,
			wantCount:      1,
		},
		{
			name:           "empty list",
			query:          nil,
			mockResponse:   []ProviderResponseLegacy{},
			mockStatusCode: http.StatusOK,
			wantErr:        false,
			wantCount:      0,
		},
		{
			name:           "server error",
			query:          nil,
			mockResponse:   map[string]string{"error": "internal server error"},
			mockStatusCode: http.StatusInternalServerError,
			wantErr:        true,
			wantCount:      0,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				server := httptest.NewTLSServer(
					http.HandlerFunc(
						func(w http.ResponseWriter, r *http.Request) {
							// Verify request path starts with expected endpoint
							if !strings.HasPrefix(r.URL.Path, "/KeyfactorAPI/PamProviders") {
								t.Errorf("Expected path to start with /KeyfactorAPI/PamProviders, got %s", r.URL.Path)
							}
							if r.Method != "GET" {
								t.Errorf("Expected GET method, got %s", r.Method)
							}

							w.WriteHeader(tt.mockStatusCode)
							_ = json.NewEncoder(w).Encode(tt.mockResponse)
						},
					),
				)
				defer server.Close()

				client := newTestClient(server)

				got, err := client.ListPAMProviders(tt.query)
				if (err != nil) != tt.wantErr {
					t.Errorf("ListPAMProviders() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr && got != nil {
					if len(*got) != tt.wantCount {
						t.Errorf("ListPAMProviders() got %d items, want %d", len(*got), tt.wantCount)
					}
				}
			},
		)
	}
}

func TestGetPAMProvider(t *testing.T) {
	tests := []struct {
		name           string
		providerId     int
		mockResponse   interface{}
		mockStatusCode int
		wantErr        bool
	}{
		{
			name:           "successful get",
			providerId:     1,
			mockResponse:   mockProviderResponseLegacy,
			mockStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name:           "not found",
			providerId:     999,
			mockResponse:   map[string]string{"error": "not found"},
			mockStatusCode: http.StatusNotFound,
			wantErr:        true,
		},
		{
			name:           "server error",
			providerId:     1,
			mockResponse:   map[string]string{"error": "internal server error"},
			mockStatusCode: http.StatusInternalServerError,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				server := httptest.NewTLSServer(
					http.HandlerFunc(
						func(w http.ResponseWriter, r *http.Request) {
							expectedPath := "/KeyfactorAPI/PamProviders/1"
							if tt.providerId == 999 {
								expectedPath = "/KeyfactorAPI/PamProviders/999"
							}
							if r.URL.Path != expectedPath {
								t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
							}
							if r.Method != "GET" {
								t.Errorf("Expected GET method, got %s", r.Method)
							}

							w.WriteHeader(tt.mockStatusCode)
							_ = json.NewEncoder(w).Encode(tt.mockResponse)
						},
					),
				)
				defer server.Close()

				client := newTestClient(server)

				got, err := client.GetPAMProvider(tt.providerId)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetPAMProvider() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr && got != nil {
					if got.Id != mockProviderResponseLegacy.Id {
						t.Errorf("GetPAMProvider() Id = %v, want %v", got.Id, mockProviderResponseLegacy.Id)
					}
				}
			},
		)
	}
}

func TestCreatePAMProvider(t *testing.T) {
	createRequest := &ProviderCreateRequest{
		Name:   "New PAM Provider",
		Remote: false,
		Area:   1,
		ProviderType: ProviderType{
			Id:   mockProviderTypeId,
			Name: &mockProviderTypeName,
		},
		SecuredAreaId: intPtr(1),
	}

	tests := []struct {
		name           string
		request        *ProviderCreateRequest
		mockResponse   interface{}
		mockStatusCode int
		wantErr        bool
	}{
		{
			name:           "successful create",
			request:        createRequest,
			mockResponse:   mockProviderResponseLegacy,
			mockStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name:           "bad request - missing name",
			request:        &ProviderCreateRequest{},
			mockResponse:   map[string]string{"error": "name is required"},
			mockStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name:           "server error",
			request:        createRequest,
			mockResponse:   map[string]string{"error": "internal server error"},
			mockStatusCode: http.StatusInternalServerError,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				server := httptest.NewTLSServer(
					http.HandlerFunc(
						func(w http.ResponseWriter, r *http.Request) {
							if r.URL.Path != "/KeyfactorAPI/PamProviders" {
								t.Errorf("Expected path /KeyfactorAPI/PamProviders, got %s", r.URL.Path)
							}
							if r.Method != "POST" {
								t.Errorf("Expected POST method, got %s", r.Method)
							}

							// Verify request body
							var receivedRequest ProviderCreateRequest
							if err := json.NewDecoder(r.Body).Decode(&receivedRequest); err != nil {
								t.Errorf("Failed to decode request body: %v", err)
							}
							if receivedRequest.Name != tt.request.Name {
								t.Errorf("Expected name %s, got %s", tt.request.Name, receivedRequest.Name)
							}

							w.WriteHeader(tt.mockStatusCode)
							_ = json.NewEncoder(w).Encode(tt.mockResponse)
						},
					),
				)
				defer server.Close()

				client := newTestClient(server)

				got, err := client.CreatePAMProvider(tt.request)
				if (err != nil) != tt.wantErr {
					t.Errorf("CreatePAMProvider() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr && got != nil {
					if got.Id != mockProviderResponseLegacy.Id {
						t.Errorf("CreatePAMProvider() Id = %v, want %v", got.Id, mockProviderResponseLegacy.Id)
					}
				}
			},
		)
	}
}

func TestUpdatePAMProvider(t *testing.T) {
	updateRequest := &ProviderUpdateRequestLegacy{
		Id:     1,
		Name:   "Updated PAM Provider",
		Remote: false,
		Area:   1,
		ProviderType: ProviderType{
			Id:   mockProviderTypeId,
			Name: &mockProviderTypeName,
		},
		SecuredAreaId: intPtr(1),
	}

	updatedResponse := ProviderResponseLegacy{
		Id:   1,
		Name: stringPtr("Updated PAM Provider"),
		Area: 1,
		ProviderType: &ProviderType{
			Id:   mockProviderTypeId,
			Name: &mockProviderTypeName,
		},
		SecuredAreaId: intPtr(1),
		Remote:        false,
	}

	tests := []struct {
		name           string
		request        *ProviderUpdateRequestLegacy
		mockResponse   interface{}
		mockStatusCode int
		wantErr        bool
	}{
		{
			name:           "successful update",
			request:        updateRequest,
			mockResponse:   updatedResponse,
			mockStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "not found",
			request: &ProviderUpdateRequestLegacy{
				Id:   999,
				Name: "Nonexistent Provider",
			},
			mockResponse:   map[string]string{"error": "not found"},
			mockStatusCode: http.StatusNotFound,
			wantErr:        true,
		},
		{
			name:           "bad request",
			request:        &ProviderUpdateRequestLegacy{Id: 1},
			mockResponse:   map[string]string{"error": "invalid request"},
			mockStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				server := httptest.NewTLSServer(
					http.HandlerFunc(
						func(w http.ResponseWriter, r *http.Request) {
							if r.URL.Path != "/KeyfactorAPI/PamProviders" {
								t.Errorf("Expected path /KeyfactorAPI/PamProviders, got %s", r.URL.Path)
							}
							if r.Method != "PUT" {
								t.Errorf("Expected PUT method, got %s", r.Method)
							}

							// Verify request body
							var receivedRequest ProviderUpdateRequestLegacy
							if err := json.NewDecoder(r.Body).Decode(&receivedRequest); err != nil {
								t.Errorf("Failed to decode request body: %v", err)
							}
							if receivedRequest.Id != tt.request.Id {
								t.Errorf("Expected Id %d, got %d", tt.request.Id, receivedRequest.Id)
							}

							w.WriteHeader(tt.mockStatusCode)
							_ = json.NewEncoder(w).Encode(tt.mockResponse)
						},
					),
				)
				defer server.Close()

				client := newTestClient(server)

				got, err := client.UpdatePAMProvider(tt.request)
				if (err != nil) != tt.wantErr {
					t.Errorf("UpdatePAMProvider() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr && got != nil {
					if *got.Name != "Updated PAM Provider" {
						t.Errorf("UpdatePAMProvider() Name = %v, want %v", *got.Name, "Updated PAM Provider")
					}
				}
			},
		)
	}
}

func TestDeletePAMProvider(t *testing.T) {
	tests := []struct {
		name           string
		providerId     int
		mockStatusCode int
		wantErr        bool
	}{
		{
			name:           "successful delete",
			providerId:     1,
			mockStatusCode: http.StatusNoContent,
			wantErr:        false,
		},
		{
			name:           "not found",
			providerId:     999,
			mockStatusCode: http.StatusNotFound,
			wantErr:        true,
		},
		{
			name:           "provider in use",
			providerId:     1,
			mockStatusCode: http.StatusConflict,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				server := httptest.NewTLSServer(
					http.HandlerFunc(
						func(w http.ResponseWriter, r *http.Request) {
							expectedPath := "/KeyfactorAPI/PamProviders/1"
							if tt.providerId == 999 {
								expectedPath = "/KeyfactorAPI/PamProviders/999"
							}
							if r.URL.Path != expectedPath {
								t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
							}
							if r.Method != "DELETE" {
								t.Errorf("Expected DELETE method, got %s", r.Method)
							}

							w.WriteHeader(tt.mockStatusCode)
						},
					),
				)
				defer server.Close()

				client := newTestClient(server)

				err := client.DeletePAMProvider(tt.providerId)
				if (err != nil) != tt.wantErr {
					t.Errorf("DeletePAMProvider() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}

func TestListLocalPAMEntries(t *testing.T) {
	tests := []struct {
		name           string
		providerId     int
		query          *GetPAMProviderQuery
		mockResponse   interface{}
		mockStatusCode int
		wantErr        bool
		wantCount      int
	}{
		{
			name:       "successful list without query",
			providerId: 1,
			query:      nil,
			mockResponse: []LocalPAMEntryResponse{
				mockLocalPAMEntry,
			},
			mockStatusCode: http.StatusOK,
			wantErr:        false,
			wantCount:      1,
		},
		{
			name:       "successful list with query",
			providerId: 1,
			query: &GetPAMProviderQuery{
				ReturnLimit: 10,
			},
			mockResponse: []LocalPAMEntryResponse{
				mockLocalPAMEntry,
			},
			mockStatusCode: http.StatusOK,
			wantErr:        false,
			wantCount:      1,
		},
		{
			name:           "empty list",
			providerId:     1,
			query:          nil,
			mockResponse:   []LocalPAMEntryResponse{},
			mockStatusCode: http.StatusOK,
			wantErr:        false,
			wantCount:      0,
		},
		{
			name:           "provider not found",
			providerId:     999,
			query:          nil,
			mockResponse:   map[string]string{"error": "provider not found"},
			mockStatusCode: http.StatusNotFound,
			wantErr:        true,
			wantCount:      0,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				server := httptest.NewTLSServer(
					http.HandlerFunc(
						func(w http.ResponseWriter, r *http.Request) {
							expectedPath := "/KeyfactorAPI/PamProviders/Local/1/Entries"
							if tt.providerId == 999 {
								expectedPath = "/KeyfactorAPI/PamProviders/Local/999/Entries"
							}
							if !strings.HasPrefix(r.URL.Path, expectedPath) {
								t.Errorf("Expected path to start with %s, got %s", expectedPath, r.URL.Path)
							}
							if r.Method != "GET" {
								t.Errorf("Expected GET method, got %s", r.Method)
							}

							w.WriteHeader(tt.mockStatusCode)
							_ = json.NewEncoder(w).Encode(tt.mockResponse)
						},
					),
				)
				defer server.Close()

				client := newTestClient(server)

				got, err := client.ListLocalPAMEntries(tt.providerId, tt.query)
				if (err != nil) != tt.wantErr {
					t.Errorf("ListLocalPAMEntries() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr && got != nil {
					if len(*got) != tt.wantCount {
						t.Errorf("ListLocalPAMEntries() got %d items, want %d", len(*got), tt.wantCount)
					}
				}
			},
		)
	}
}

func TestCreateLocalPAMEntry(t *testing.T) {
	createRequest := &LocalPAMEntryCreateRequest{
		SecretName:  "new-secret",
		Description: stringPtr("New secret description"),
		SecretValue: "super-secret-value",
	}

	createdResponse := LocalPAMEntryResponse{
		ProviderId:  1,
		SecretName:  stringPtr("new-secret"),
		Description: stringPtr("New secret description"),
	}

	tests := []struct {
		name           string
		providerId     int
		request        *LocalPAMEntryCreateRequest
		mockResponse   interface{}
		mockStatusCode int
		wantErr        bool
	}{
		{
			name:           "successful create",
			providerId:     1,
			request:        createRequest,
			mockResponse:   createdResponse,
			mockStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name:       "bad request - missing secret name",
			providerId: 1,
			request: &LocalPAMEntryCreateRequest{
				SecretValue: "value",
			},
			mockResponse:   map[string]string{"error": "secret name is required"},
			mockStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name:           "provider not found",
			providerId:     999,
			request:        createRequest,
			mockResponse:   map[string]string{"error": "provider not found"},
			mockStatusCode: http.StatusNotFound,
			wantErr:        true,
		},
		{
			name:           "duplicate secret name",
			providerId:     1,
			request:        createRequest,
			mockResponse:   map[string]string{"error": "secret already exists"},
			mockStatusCode: http.StatusConflict,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				server := httptest.NewTLSServer(
					http.HandlerFunc(
						func(w http.ResponseWriter, r *http.Request) {
							expectedPath := "/KeyfactorAPI/PamProviders/Local/1/Entries"
							if tt.providerId == 999 {
								expectedPath = "/KeyfactorAPI/PamProviders/Local/999/Entries"
							}
							if r.URL.Path != expectedPath {
								t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
							}
							if r.Method != "POST" {
								t.Errorf("Expected POST method, got %s", r.Method)
							}

							// Verify request body
							var receivedRequest LocalPAMEntryCreateRequest
							if err := json.NewDecoder(r.Body).Decode(&receivedRequest); err != nil {
								t.Errorf("Failed to decode request body: %v", err)
							}
							if receivedRequest.SecretName != tt.request.SecretName {
								t.Errorf(
									"Expected SecretName %s, got %s",
									tt.request.SecretName,
									receivedRequest.SecretName,
								)
							}

							w.WriteHeader(tt.mockStatusCode)
							_ = json.NewEncoder(w).Encode(tt.mockResponse)
						},
					),
				)
				defer server.Close()

				client := newTestClient(server)

				got, err := client.CreateLocalPAMEntry(tt.providerId, tt.request)
				if (err != nil) != tt.wantErr {
					t.Errorf("CreateLocalPAMEntry() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr && got != nil {
					if *got.SecretName != "new-secret" {
						t.Errorf("CreateLocalPAMEntry() SecretName = %v, want %v", *got.SecretName, "new-secret")
					}
				}
			},
		)
	}
}

func TestUpdateLocalPAMEntry(t *testing.T) {
	updateRequest := &LocalPAMEntryUpdateRequest{
		SecretName:  "updated-secret",
		Description: stringPtr("Updated description"),
		SecretValue: stringPtr("updated-value"),
	}

	updatedResponse := LocalPAMEntryResponse{
		ProviderId:  1,
		SecretName:  stringPtr("updated-secret"),
		Description: stringPtr("Updated description"),
	}

	tests := []struct {
		name           string
		providerId     int
		secretName     string
		request        *LocalPAMEntryUpdateRequest
		mockResponse   interface{}
		mockStatusCode int
		wantErr        bool
	}{
		{
			name:           "successful update",
			providerId:     1,
			secretName:     "test-secret",
			request:        updateRequest,
			mockResponse:   updatedResponse,
			mockStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name:       "secret not found",
			providerId: 1,
			secretName: "nonexistent",
			request: &LocalPAMEntryUpdateRequest{
				SecretName: "nonexistent",
			},
			mockResponse:   map[string]string{"error": "secret not found"},
			mockStatusCode: http.StatusNotFound,
			wantErr:        true,
		},
		{
			name:           "provider not found",
			providerId:     999,
			secretName:     "test-secret",
			request:        updateRequest,
			mockResponse:   map[string]string{"error": "provider not found"},
			mockStatusCode: http.StatusNotFound,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				server := httptest.NewTLSServer(
					http.HandlerFunc(
						func(w http.ResponseWriter, r *http.Request) {
							expectedPath := "/KeyfactorAPI/PamProviders/Local/1/Entries/test-secret"
							if tt.providerId == 999 {
								expectedPath = "/KeyfactorAPI/PamProviders/Local/999/Entries/test-secret"
							}
							if tt.secretName == "nonexistent" {
								expectedPath = "/KeyfactorAPI/PamProviders/Local/1/Entries/nonexistent"
							}
							if r.URL.Path != expectedPath {
								t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
							}
							if r.Method != "PUT" {
								t.Errorf("Expected PUT method, got %s", r.Method)
							}

							// Verify request body
							var receivedRequest LocalPAMEntryUpdateRequest
							if err := json.NewDecoder(r.Body).Decode(&receivedRequest); err != nil {
								t.Errorf("Failed to decode request body: %v", err)
							}
							if receivedRequest.SecretName != tt.request.SecretName {
								t.Errorf(
									"Expected SecretName %s, got %s",
									tt.request.SecretName,
									receivedRequest.SecretName,
								)
							}

							w.WriteHeader(tt.mockStatusCode)
							_ = json.NewEncoder(w).Encode(tt.mockResponse)
						},
					),
				)
				defer server.Close()

				client := newTestClient(server)

				got, err := client.UpdateLocalPAMEntry(tt.providerId, tt.secretName, tt.request)
				if (err != nil) != tt.wantErr {
					t.Errorf("UpdateLocalPAMEntry() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr && got != nil {
					if *got.SecretName != "updated-secret" {
						t.Errorf("UpdateLocalPAMEntry() SecretName = %v, want %v", *got.SecretName, "updated-secret")
					}
				}
			},
		)
	}
}

func TestDeleteLocalPAMEntry(t *testing.T) {
	tests := []struct {
		name           string
		providerId     int
		secretName     string
		mockStatusCode int
		wantErr        bool
	}{
		{
			name:           "successful delete",
			providerId:     1,
			secretName:     "test-secret",
			mockStatusCode: http.StatusNoContent,
			wantErr:        false,
		},
		{
			name:           "secret not found",
			providerId:     1,
			secretName:     "nonexistent",
			mockStatusCode: http.StatusNotFound,
			wantErr:        true,
		},
		{
			name:           "provider not found",
			providerId:     999,
			secretName:     "test-secret",
			mockStatusCode: http.StatusNotFound,
			wantErr:        true,
		},
		{
			name:           "secret in use",
			providerId:     1,
			secretName:     "in-use-secret",
			mockStatusCode: http.StatusConflict,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				server := httptest.NewTLSServer(
					http.HandlerFunc(
						func(w http.ResponseWriter, r *http.Request) {
							expectedPath := "/KeyfactorAPI/PamProviders/Local/1/Entries/test-secret"
							if tt.providerId == 999 {
								expectedPath = "/KeyfactorAPI/PamProviders/Local/999/Entries/test-secret"
							}
							if tt.secretName == "nonexistent" {
								expectedPath = "/KeyfactorAPI/PamProviders/Local/1/Entries/nonexistent"
							}
							if tt.secretName == "in-use-secret" {
								expectedPath = "/KeyfactorAPI/PamProviders/Local/1/Entries/in-use-secret"
							}
							if r.URL.Path != expectedPath {
								t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
							}
							if r.Method != "DELETE" {
								t.Errorf("Expected DELETE method, got %s", r.Method)
							}

							w.WriteHeader(tt.mockStatusCode)
						},
					),
				)
				defer server.Close()

				client := newTestClient(server)

				err := client.DeleteLocalPAMEntry(tt.providerId, tt.secretName)
				if (err != nil) != tt.wantErr {
					t.Errorf("DeleteLocalPAMEntry() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}
