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

func TestClient_AddCertificateToStores(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		config *AddCertificateToStore
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
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
				got, err := c.AddCertificateToStores(tt.args.config)
				if (err != nil) != tt.wantErr {
					t.Errorf("AddCertificateToStores() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("AddCertificateToStores() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestClient_CreateStore(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		ca *CreateStoreFctArgs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CreateStoreResponse
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
				got, err := c.CreateStore(tt.args.ca)
				if (err != nil) != tt.wantErr {
					t.Errorf("CreateStore() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("CreateStore() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestClient_DeleteCertificateStore(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		storeId string
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
				if err := c.DeleteCertificateStore(tt.args.storeId); (err != nil) != tt.wantErr {
					t.Errorf("DeleteCertificateStore() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}

func TestClient_GetCertStoreInventory(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		storeId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]CertStoreInventory
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
				got, err := c.GetCertStoreInventory(tt.args.storeId)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetCertStoreInventory() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetCertStoreInventory() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestClient_GetCertificateStoreByID(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		storeId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetCertificateStoreResponse
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
				got, err := c.GetCertificateStoreByID(tt.args.storeId)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetCertificateStoreByID() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetCertificateStoreByID() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestClient_ListCertificateStores(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *[]GetCertificateStoreResponse
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
				got, err := c.ListCertificateStores(nil)
				if (err != nil) != tt.wantErr {
					t.Errorf("ListCertificateStores() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ListCertificateStores() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestClient_RemoveCertificateFromStores(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		config *RemoveCertificateFromStore
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
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
				got, err := c.RemoveCertificateFromStores(tt.args.config)
				if (err != nil) != tt.wantErr {
					t.Errorf("RemoveCertificateFromStores() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("RemoveCertificateFromStores() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestClient_UpdateStore(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		ua *UpdateStoreFctArgs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UpdateStoreResponse
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
				got, err := c.UpdateStore(tt.args.ua)
				if (err != nil) != tt.wantErr {
					t.Errorf("UpdateStore() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("UpdateStore() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_buildPropertiesInterface(t *testing.T) {
	type args struct {
		properties map[string]string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := buildPropertiesInterface(tt.args.properties); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("buildPropertiesInterface() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_unmarshalPropertiesString(t *testing.T) {
	type args struct {
		properties string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := unmarshalPropertiesString(tt.args.properties); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("unmarshalPropertiesString() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_validateCreateStoreArgs(t *testing.T) {
	type args struct {
		ca *CreateStoreFctArgs
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if err := validateCreateStoreArgs(tt.args.ca); (err != nil) != tt.wantErr {
					t.Errorf("validateCreateStoreArgs() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}

func Test_validateUpdateStoreArgs(t *testing.T) {
	type args struct {
		ca *UpdateStoreFctArgs
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if err := validateUpdateStoreArgs(tt.args.ca); (err != nil) != tt.wantErr {
					t.Errorf("validateUpdateStoreArgs() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}
