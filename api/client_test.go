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
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestClient_sendRequest(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		request *request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
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
				got, err := c.sendRequest(tt.args.request)
				if (err != nil) != tt.wantErr {
					t.Errorf("sendRequest() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("sendRequest() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestNewKeyfactorClient(t *testing.T) {
	type args struct {
		auth *AuthConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := NewKeyfactorClient(tt.args.auth)
				if (err != nil) != tt.wantErr {
					t.Errorf("NewKeyfactorClient() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewKeyfactorClient() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_buildBasicAuthString(t *testing.T) {
	type args struct {
		auth *AuthConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := buildBasicAuthString(tt.args.auth); got != tt.want {
					t.Errorf("buildBasicAuthString() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_getTimestamp(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := getTimestamp(); got != tt.want {
					t.Errorf("getTimestamp() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_loginToKeyfactor(t *testing.T) {
	type args struct {
		auth *AuthConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := loginToKeyfactor(tt.args.auth)
				fmt.Print(got)
				if (err != nil) != tt.wantErr {
					t.Errorf("loginToKeyfactor() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("loginToKeyfactor() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
