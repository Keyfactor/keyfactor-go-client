package api_test

import (
	"github.com/Keyfactor/keyfactor-go-client/api"
	"io"
	"log"
	"testing"
)

func TestClient_GetStoreContainer(t *testing.T) {
	log.SetOutput(io.Discard)
	c, kfcErr := api.NewKeyfactorClient(&api.AuthConfig{})
	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}
	containers, lErr := c.GetStoreContainers()
	if lErr != nil {
		t.Errorf("unable to list containers. %s", lErr)
		return
	}
	if len(*containers) == 0 {
		t.Errorf("no containers found in client.")
		return
	}
	containerID := (*containers)[0].Id
	containerName := (*containers)[0].Name
	type fields struct{}
	type args struct {
		id interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.CertStoreContainer
		wantErr bool
	}{
		{
			name:   "Get store container by ID",
			fields: fields{},
			args: args{
				id: *containerID,
			},
			want:    &api.CertStoreContainer{},
			wantErr: false,
		},
		{
			name:   "Get store container by invalid ID",
			fields: fields{},
			args: args{
				id: "-1",
			},
			want:    &api.CertStoreContainer{},
			wantErr: true,
		},
		{
			name:   "Get store container by name",
			fields: fields{},
			args: args{
				id: containerName,
			},
			want:    &api.CertStoreContainer{},
			wantErr: false,
		},
		{
			name:   "Get store container invalid by name",
			fields: fields{},
			args: args{
				id: "invalid-container-name",
			},
			want:    &api.CertStoreContainer{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GetStoreContainer(tt.args.id) // GET Store container by ID
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStoreContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_GetStoreContainers(t *testing.T) {
	log.SetOutput(io.Discard)
	log.SetOutput(io.Discard)
	c, kfcErr := api.NewKeyfactorClient(&api.AuthConfig{})
	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}

	type fields struct{}
	tests := []struct {
		name    string
		fields  fields
		want    *[]api.CertStoreContainer
		wantErr bool
	}{
		{
			name:    "List store containers",
			fields:  fields{},
			want:    &[]api.CertStoreContainer{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
