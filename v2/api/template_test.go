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
		t.Run(tt.name, func(t *testing.T) {
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
		})
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
		t.Run(tt.name, func(t *testing.T) {
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
		})
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
		t.Run(tt.name, func(t *testing.T) {
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
		})
	}
}
