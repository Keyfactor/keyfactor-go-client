package api

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClient_GetCAList(t *testing.T) {
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []CA
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				hostname:        tt.fields.hostname,
				httpClient:      tt.fields.httpClient,
				basicAuthString: tt.fields.basicAuthString,
			}
			got, err := c.GetCAList()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCAList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCAList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
