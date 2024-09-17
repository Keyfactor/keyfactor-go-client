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
	"net/http"
	"reflect"
	"testing"
)

func TestClient_CreateSecurityIdentity(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		csia *CreateSecurityIdentityArg
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CreateSecurityIdentityResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := &Client{
					Hostname:        tt.fields.hostname,
					httpClient:      tt.fields.httpClient,
					basicAuthString: tt.fields.basicAuthString,
				}
				got, err := c.CreateSecurityIdentity(tt.args.csia)
				if (err != nil) != tt.wantErr {
					t.Errorf("CreateSecurityIdentity() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("CreateSecurityIdentity() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestClient_CreateSecurityRole(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		input *CreateSecurityRoleArg
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CreateSecurityRoleResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := &Client{
					Hostname:        tt.fields.hostname,
					httpClient:      tt.fields.httpClient,
					basicAuthString: tt.fields.basicAuthString,
				}
				got, err := c.CreateSecurityRole(tt.args.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("CreateSecurityRole() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("CreateSecurityRole() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestClient_DeleteSecurityIdentity(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := &Client{
					Hostname:        tt.fields.hostname,
					httpClient:      tt.fields.httpClient,
					basicAuthString: tt.fields.basicAuthString,
				}
				if err := c.DeleteSecurityIdentity(tt.args.id); (err != nil) != tt.wantErr {
					t.Errorf("DeleteSecurityIdentity() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}

func TestClient_DeleteSecurityRole(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := &Client{
					Hostname:        tt.fields.hostname,
					httpClient:      tt.fields.httpClient,
					basicAuthString: tt.fields.basicAuthString,
				}
				if err := c.DeleteSecurityRole(tt.args.id); (err != nil) != tt.wantErr {
					t.Errorf("DeleteSecurityRole() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}

func TestClient_GetSecurityIdentities(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []GetSecurityIdentityResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := &Client{
					Hostname:        tt.fields.hostname,
					httpClient:      tt.fields.httpClient,
					basicAuthString: tt.fields.basicAuthString,
				}
				got, err := c.GetSecurityIdentities()
				if (err != nil) != tt.wantErr {
					t.Errorf("GetSecurityIdentities() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetSecurityIdentities() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestClient_GetSecurityRole(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		id interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetSecurityRoleResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := &Client{
					Hostname:        tt.fields.hostname,
					httpClient:      tt.fields.httpClient,
					basicAuthString: tt.fields.basicAuthString,
				}
				got, err := c.GetSecurityRole(tt.args.id)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetSecurityRole() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetSecurityRole() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestClient_GetSecurityRoles(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []GetSecurityRolesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := &Client{
					Hostname:        tt.fields.hostname,
					httpClient:      tt.fields.httpClient,
					basicAuthString: tt.fields.basicAuthString,
				}
				got, err := c.GetSecurityRoles()
				if (err != nil) != tt.wantErr {
					t.Errorf("GetSecurityRoles() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetSecurityRoles() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestClient_UpdateSecurityRole(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		input *UpdateSecurityRoleArg
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UpdateSecurityRoleResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := &Client{
					Hostname:        tt.fields.hostname,
					httpClient:      tt.fields.httpClient,
					basicAuthString: tt.fields.basicAuthString,
				}
				got, err := c.UpdateSecurityRole(tt.args.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("UpdateSecurityRole() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("UpdateSecurityRole() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
