package api

import (
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
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Hostname:        tt.fields.hostname,
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
		})
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
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewKeyfactorClient(tt.args.auth)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKeyfactorClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKeyfactorClient() got = %v, want %v", got, tt.want)
			}
		})
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
		t.Run(tt.name, func(t *testing.T) {
			if got := buildBasicAuthString(tt.args.auth); got != tt.want {
				t.Errorf("buildBasicAuthString() = %v, want %v", got, tt.want)
			}
		})
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
		t.Run(tt.name, func(t *testing.T) {
			if got := getTimestamp(); got != tt.want {
				t.Errorf("getTimestamp() = %v, want %v", got, tt.want)
			}
		})
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
		t.Run(tt.name, func(t *testing.T) {
			got, err := loginToKeyfactor(tt.args.auth)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginToKeyfactor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loginToKeyfactor() got = %v, want %v", got, tt.want)
			}
		})
	}
}
