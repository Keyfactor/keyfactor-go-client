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

func TestClient_GetAllMetadataFields(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []MetadataField
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
				got, err := c.GetAllMetadataFields()
				if (err != nil) != tt.wantErr {
					t.Errorf("GetAllMetadataFields() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetAllMetadataFields() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestClient_UpdateMetadata(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		um *UpdateMetadataArgs
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
				if err := c.UpdateMetadata(tt.args.um); (err != nil) != tt.wantErr {
					t.Errorf("UpdateMetadata() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}
