package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
)

func TestClient_GetStoreContainer(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	kfClient, connErr := initClient()
	if connErr != nil {
		t.Errorf("Error connecting to Keyfactor: %s", connErr)
		return
	}
	containerID := os.Getenv("TEST_KEYFACTOR_STORE_CONTAINER_ID")
	containerName := os.Getenv("TEST_KEYFACTOR_STORE_CONTAINER_NAME")

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
		want    *CertStoreContainer
		wantErr bool
	}{
		{
			name: "Get store container by ID",
			fields: fields{
				hostname:        kfClient.hostname,
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			},
			args: args{
				id: containerID,
			},
			want:    &CertStoreContainer{},
			wantErr: false,
		},
		{
			name: "Get store container by invalid ID",
			fields: fields{
				hostname:        kfClient.hostname,
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			},
			args: args{
				id: "-1",
			},
			want:    &CertStoreContainer{},
			wantErr: true,
		},
		{
			name: "Get store container by name",
			fields: fields{
				hostname:        kfClient.hostname,
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			},
			args: args{
				id: containerName,
			},
			want:    &CertStoreContainer{},
			wantErr: false,
		},
		{
			name: "Get store container invalid by name",
			fields: fields{
				hostname:        kfClient.hostname,
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			},
			args: args{
				id: "invalid-container-name",
			},
			want:    &CertStoreContainer{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				hostname:        tt.fields.hostname,
				httpClient:      tt.fields.httpClient,
				basicAuthString: tt.fields.basicAuthString,
			}
			var idInt int
			_, cErr := strconv.Atoi(tt.args.id.(string))
			var got *CertStoreContainer
			var err error
			if cErr == nil {
				idInt, _ = strconv.Atoi(tt.args.id.(string))
				got, err = c.GetStoreContainer(idInt) // GET Store container by ID
			} else {
				got, err = c.GetStoreContainer(tt.args.id) // GET Store container by name
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("GetStoreContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if cErr == nil {
					// Check that ID is correct
					if *got.Id != idInt {
						t.Errorf("GetStoreContainer() got = %v, want %v", got.Id, idInt)
					}
				} else {
					// Check that name is correct
					if got.Name != tt.args.id {
						t.Errorf("GetStoreContainer() got = %v, want %v", got.Name, tt.args.id)
					}
				}
			}
		})
	}
}

func TestClient_GetStoreContainers(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	kfClient, _ := initClient()

	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *[]CertStoreContainer
		wantErr bool
	}{
		{
			name: "List store containers",
			fields: fields{
				hostname:        kfClient.hostname,
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			},
			want:    &[]CertStoreContainer{},
			wantErr: false,
		},
		{
			name: "List store containers failure",
			fields: fields{
				hostname:        "",
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			},
			want:    &[]CertStoreContainer{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				hostname:        tt.fields.hostname,
				httpClient:      tt.fields.httpClient,
				basicAuthString: tt.fields.basicAuthString,
			}
			got, err := c.GetStoreContainers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStoreContainers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Check that we got at least one store container
				if len(*got) == 0 {
					t.Errorf("GetStoreContainers() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
