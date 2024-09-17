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

func TestClient_GetTemplate(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		Id interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetTemplateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := &Client{
					hostname:        tt.fields.hostname,
					httpClient:      tt.fields.httpClient,
					basicAuthString: tt.fields.basicAuthString,
				}
				got, err := c.GetTemplate(tt.args.Id)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetTemplate() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetTemplate() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestClient_GetTemplates(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []GetTemplateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := &Client{
					hostname:        tt.fields.hostname,
					httpClient:      tt.fields.httpClient,
					basicAuthString: tt.fields.basicAuthString,
				}
				got, err := c.GetTemplates()
				if (err != nil) != tt.wantErr {
					t.Errorf("GetTemplates() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetTemplates() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestClient_UpdateTemplate(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		uta *UpdateTemplateArg
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UpdateTemplateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := &Client{
					hostname:        tt.fields.hostname,
					httpClient:      tt.fields.httpClient,
					basicAuthString: tt.fields.basicAuthString,
				}
				got, err := c.UpdateTemplate(tt.args.uta)
				if (err != nil) != tt.wantErr {
					t.Errorf("UpdateTemplate() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("UpdateTemplate() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
