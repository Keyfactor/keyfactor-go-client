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
	"io"
	"log"
	"testing"
)

func TestClient_GetStoreContainer(t *testing.T) {
	log.SetOutput(io.Discard)
	c, kfcErr := NewKeyfactorClient(&AuthConfig{})
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
		want    *CertStoreContainer
		wantErr bool
	}{
		{
			name:   "Get store container by ID",
			fields: fields{},
			args: args{
				id: *containerID,
			},
			want:    &CertStoreContainer{},
			wantErr: false,
		},
		{
			name:   "Get store container by invalid ID",
			fields: fields{},
			args: args{
				id: "-1",
			},
			want:    &CertStoreContainer{},
			wantErr: true,
		},
		{
			name:   "Get store container by name",
			fields: fields{},
			args: args{
				id: containerName,
			},
			want:    &CertStoreContainer{},
			wantErr: false,
		},
		{
			name:   "Get store container invalid by name",
			fields: fields{},
			args: args{
				id: "invalid-container-name",
			},
			want:    &CertStoreContainer{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				_, err := c.GetStoreContainer(tt.args.id) // GET Store container by ID
				if (err != nil) != tt.wantErr {
					t.Errorf("GetStoreContainer() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			},
		)
	}
}

func TestClient_GetStoreContainers(t *testing.T) {
	log.SetOutput(io.Discard)
	log.SetOutput(io.Discard)
	c, kfcErr := NewKeyfactorClient(&AuthConfig{})
	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}

	type fields struct{}
	tests := []struct {
		name    string
		fields  fields
		want    *[]CertStoreContainer
		wantErr bool
	}{
		{
			name:    "List store containers",
			fields:  fields{},
			want:    &[]CertStoreContainer{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
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
			},
		)
	}
}
