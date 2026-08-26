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
	"testing"

	"github.com/Keyfactor/keyfactor-auth-client-go/auth_providers"
)

// mockAuthConfig implements AuthConfig interface for testing
type mockAuthConfig struct {
	serverConfig *auth_providers.Server
	httpClient   *http.Client
}

func (m *mockAuthConfig) GetServerConfig() *auth_providers.Server {
	return m.serverConfig
}

func (m *mockAuthConfig) GetHttpClient() (*http.Client, error) {
	return m.httpClient, nil
}

func (m *mockAuthConfig) Authenticate() error {
	return nil
}

func (m *mockAuthConfig) GetCommandVersion() string {
	return "25.1.0.0"
}

// newTestClient creates a test client with mock server
func newTestClient(server *httptest.Server) *Client {
	return &Client{
		AuthClient: &mockAuthConfig{
			serverConfig: &auth_providers.Server{
				Host:          server.URL,
				APIPath:       "/KeyfactorAPI",
				SkipTLSVerify: true,
			},
			httpClient: server.Client(),
		},
	}
}

// Mock response data
var (
	mockProviderTypeId   = "550e8400-e29b-41d4-a716-446655440000"
	mockProviderTypeName = "CyberArk"

	mockProviderTypeResponse = ProviderTypeResponse{
		Id:   mockProviderTypeId,
		Name: mockProviderTypeName,
		Parameters: &[]ProviderTypeParameterResponse{
			{
				Id:            1,
				Name:          "Username",
				DisplayName:   stringPtr("User Name"),
				DataType:      PamParameterDataTypeString,
				InstanceLevel: false,
			},
			{
				Id:            2,
				Name:          "Password",
				DisplayName:   stringPtr("Password"),
				DataType:      PamParameterDataTypeSecret,
				InstanceLevel: true,
			},
		},
	}

	mockProviderResponseLegacy = ProviderResponseLegacy{
		Id:   1,
		Name: stringPtr("Test Provider"),
		Area: 1,
		ProviderType: &ProviderType{
			Id:   mockProviderTypeId,
			Name: mockProviderTypeName,
		},
		SecuredAreaId: intPtr(1),
		Remote:        false,
	}

	mockLocalPAMEntry = LocalPAMEntryResponse{
		ProviderId:  1,
		SecretName:  stringPtr("test-secret"),
		Description: stringPtr("Test secret description"),
	}
)

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func TestListPAMProviderTypes(t *testing.T) {
	tests := []struct {
		name           string
		mockResponse   interface{}
		mockStatusCode int
		wantErr        bool
		wantCount      int
	}{
		{
			name: "successful list",
			mockResponse: []ProviderTypeResponse{
				mockProviderTypeResponse,
			},
			mockStatusCode: http.StatusOK,
			wantErr:        false,
			wantCount:      1,
		},
		{
			name:           "empty list",
			mockResponse:   []ProviderTypeResponse{},
			mockStatusCode: http.StatusOK,
			wantErr:        false,
			wantCount:      0,
		},
		{
			name:           "server error",
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
							// Verify request
							if r.URL.Path != "/KeyfactorAPI/PamProviders/Types" {
								t.Errorf("Expected path /KeyfactorAPI/PamProviders/Types, got %s", r.URL.Path)
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

				got, err := client.ListPAMProviderTypes()
				if (err != nil) != tt.wantErr {
					t.Errorf("ListPAMProviderTypes() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr && got != nil {
					if len(*got) != tt.wantCount {
						t.Errorf("ListPAMProviderTypes() got %d items, want %d", len(*got), tt.wantCount)
					}
				}
			},
		)
	}
}

func TestGetPAMProviderType(t *testing.T) {
	tests := []struct {
		name           string
		providerId     string
		mockResponse   interface{}
		mockStatusCode int
		wantErr        bool
	}{
		{
			name:           "successful get",
			providerId:     mockProviderTypeId,
			mockResponse:   mockProviderTypeResponse,
			mockStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name:           "not found",
			providerId:     "nonexistent-id",
			mockResponse:   map[string]string{"error": "not found"},
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
							expectedPath := "/KeyfactorAPI/PamProviders/Types/" + tt.providerId
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

				got, err := client.GetPAMProviderType(tt.providerId)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetPAMProviderType() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr && got != nil {
					if got.Id != mockProviderTypeId {
						t.Errorf("GetPAMProviderType() Id = %v, want %v", got.Id, mockProviderTypeId)
					}
				}
			},
		)
	}
}

func TestCreatePAMProviderType(t *testing.T) {
	createRequest := &ProviderTypeCreateRequest{
		Name: "New Provider Type",
		Parameters: &[]ProviderTypeParameterCreateRequest{
			{
				Name:          "ApiKey",
				DisplayName:   stringPtr("API Key"),
				DataType:      PamParameterDataTypeSecret,
				InstanceLevel: true,
			},
		},
	}

	tests := []struct {
		name           string
		request        *ProviderTypeCreateRequest
		mockResponse   interface{}
		mockStatusCode int
		wantErr        bool
	}{
		{
			name:           "successful create",
			request:        createRequest,
			mockResponse:   mockProviderTypeResponse,
			mockStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name:           "bad request",
			request:        createRequest,
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
							if r.URL.Path != "/KeyfactorAPI/PamProviders/Types" {
								t.Errorf("Expected path /KeyfactorAPI/PamProviders/Types, got %s", r.URL.Path)
							}
							if r.Method != "POST" {
								t.Errorf("Expected POST method, got %s", r.Method)
							}

							// Verify request body
							var receivedRequest ProviderTypeCreateRequest
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

				got, err := client.CreatePAMProviderType(tt.request)
				if (err != nil) != tt.wantErr {
					t.Errorf("CreatePAMProviderType() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr && got != nil {
					if got.Id != mockProviderTypeId {
						t.Errorf("CreatePAMProviderType() Id = %v, want %v", got.Id, mockProviderTypeId)
					}
				}
			},
		)
	}
}

func TestDeletePAMProviderType(t *testing.T) {
	tests := []struct {
		name           string
		providerId     string
		mockStatusCode int
		wantErr        bool
	}{
		{
			name:           "successful delete",
			providerId:     mockProviderTypeId,
			mockStatusCode: http.StatusNoContent,
			wantErr:        false,
		},
		{
			name:           "not found",
			providerId:     "nonexistent-id",
			mockStatusCode: http.StatusNotFound,
			wantErr:        true,
		},
		{
			name:           "in use",
			providerId:     mockProviderTypeId,
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
							expectedPath := "/KeyfactorAPI/PamProviders/Types/" + tt.providerId
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

				err := client.DeletePAMProviderType(tt.providerId)
				if (err != nil) != tt.wantErr {
					t.Errorf("DeletePAMProviderType() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}

func TestGetPAMProviderQuery_toQueryString(t *testing.T) {
	tests := []struct {
		name  string
		query *GetPAMProviderQuery
		want  string
	}{
		{
			name:  "nil query",
			query: nil,
			want:  "",
		},
		{
			name:  "empty query",
			query: &GetPAMProviderQuery{},
			want:  "",
		},
		{
			name: "full query",
			query: &GetPAMProviderQuery{
				QueryString:   "Name -eq 'Test'",
				PageReturned:  1,
				ReturnLimit:   10,
				SortField:     "Name",
				SortAscending: 0,
			},
			want: "QueryString=Name -eq 'Test'&PageReturned=1&ReturnLimit=10&SortField=Name&SortAscending=0",
		},
		{
			name: "partial query",
			query: &GetPAMProviderQuery{
				QueryString: "Name -eq 'Test'",
				ReturnLimit: 10,
			},
			want: "QueryString=Name -eq 'Test'&ReturnLimit=10",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.query.toQueryString()
				if got != tt.want {
					t.Errorf("GetPAMProviderQuery.toQueryString() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
