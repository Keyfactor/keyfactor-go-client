package api_test

import (
	"github.com/Keyfactor/keyfactor-go-client/api"
	"io"
	"log"
	"net/http"
	"testing"
)

const (
	ActionCreate = "create"
	ActionDelete = "delete"
	ActionGet    = "get"
	ActionList   = "list"
	ActionUpdate = "update"
)

type storeTypeTestArgs struct {
	storeType *api.CertificateStoreType
	action    string
	id        int
	name      string
}
type storeTypeTestFields struct {
	hostname        string
	httpClient      *http.Client
	basicAuthString string
}
type storeTypeTest struct {
	name    string
	fields  storeTypeTestFields
	args    storeTypeTestArgs
	want    *api.CertificateStoreType
	wantErr bool
}

var testStoreType = &api.CertificateStoreType{
	Name:       "SampleStoreType13",
	ShortName:  "SAMPTYPE13",
	Capability: "SAMPTYPE13",
	SupportedOperations: &api.StoreTypeSupportedOperations{
		Add:        false,
		Create:     false,
		Discovery:  false,
		Enrollment: false,
		Remove:     false,
	},
	Properties:      &[]api.StoreTypePropertyDefinition{},
	EntryParameters: &[]api.EntryParameter{},
	PasswordOptions: &api.StoreTypePasswordOptions{
		EntrySupported: false,
		StoreRequired:  false,
		Style:          "Default",
	},
	PrivateKeyAllowed:  "forbidden",
	JobProperties:      &[]string{},
	ServerRequired:     false,
	PowerShell:         false,
	BlueprintAllowed:   false,
	CustomAliasAllowed: "Forbidden",
}

func runStoreTypeTests(t *testing.T, tests []storeTypeTest, c *api.Client) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.args.action {
			case ActionCreate:
				got, err := c.CreateStoreType(tt.args.storeType)
				if (err != nil) != tt.wantErr {
					t.Errorf("Client.CreateStoreType() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr {
					if got.Name != tt.want.Name {
						t.Errorf("Client.CreateStoreType() = %v, want %v", got, tt.want)
					}
				}
			case ActionDelete:
				s, gErr := c.GetCertificateStoreType(tt.args.storeType.ShortName)
				if gErr != nil {
					t.Errorf("Unable to cleanup created store %s. error = %v", tt.args.storeType.Name, gErr)
					return
				}
				got, err := c.DeleteCertificateStoreType(s.StoreType)
				if (err != nil) != tt.wantErr {
					t.Errorf("Client.DeleteStoreType() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr {
					if got.ID <= 0 {
						t.Errorf("Client.DeleteStoreType() = %v, want %v", got, tt.want)
					}
				}
			case ActionGet:
				// Test Get by Name first
				got, err := c.GetCertificateStoreType(tt.args.name)
				if (err != nil) != tt.wantErr {
					t.Errorf("Client.GetCertificateStoreType() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr {
					if got.Name != tt.want.Name {
						t.Errorf("Client.GetCertificateStoreType() = %v, want %v", got, tt.want)
					}
					// Test Get by ID
					got, err = c.GetCertificateStoreType(got.StoreType)
					if (err != nil) != tt.wantErr {
						t.Errorf("Client.GetCertificateStoreType() error = %v, wantErr %v", err, tt.wantErr)
						return
					}
					if !tt.wantErr {
						if got.Name != tt.want.Name {
							t.Errorf("Client.GetCertificateStoreType() = %v, want %v", got, tt.want)
						}
					}
				}
			case ActionList:
				got, err := c.ListCertificateStoreTypes()
				if (err != nil) != tt.wantErr {
					t.Errorf("Client.ListCertificateStoreTypes() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr {
					if len(*got) <= 0 {
						t.Errorf("Client.ListCertificateStoreTypes() = %v, want > %v", got, 1)
					}
				}
			case ActionUpdate:
				s, gErr := c.GetCertificateStoreType(tt.args.storeType.ShortName)
				if gErr != nil {
					t.Errorf("Unable to cleanup created store %s. error = %v", tt.args.storeType.Name, gErr)
					return
				}
				tt.args.storeType.StoreType = s.StoreType
				tt.args.storeType.ImportType = s.ImportType
				got, err := c.UpdateStoreType(tt.args.storeType)
				if (err != nil) != tt.wantErr {
					t.Errorf("Client.UpdateCertificateStoreType() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr {
					if got.Name == testStoreType.Name {
						t.Errorf("Client.UpdateCertificateStoreType() = %v, want %v", got.Name, tt.want.Name)
					} else if got.Name != tt.want.Name {
						t.Errorf("Client.UpdateCertificateStoreType() = %v, want %v", got, tt.want)
					}
					if got.PowerShell == testStoreType.PowerShell {
						t.Errorf("Client.UpdateCertificateStoreType() PowerShell = %v, want %v", got.PowerShell, !testStoreType.PowerShell)
					} else if got.PowerShell != tt.want.PowerShell {
						t.Errorf("Client.UpdateCertificateStoreType() PowerShell = %v, want %v", got.PowerShell, tt.want.PowerShell)
					}
				}
			}

		})
	}
}

// TODO
func TestClient_CreateStoreType(t *testing.T) {
	log.SetOutput(io.Discard)
	c, kfcErr := api.NewKeyfactorClient(&api.AuthConfig{})
	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}
	var tests = []storeTypeTest{
		{
			name:   "CreateStoreType",
			fields: storeTypeTestFields{},
			args: storeTypeTestArgs{
				name:      testStoreType.Name,
				storeType: testStoreType,
				action:    ActionCreate,
			},
			want:    testStoreType,
			wantErr: false,
		},
		{
			name:   "DeleteStoreType",
			fields: storeTypeTestFields{},
			args: storeTypeTestArgs{
				name:      testStoreType.Name,
				storeType: testStoreType,
				action:    ActionDelete,
			},
			want: &api.CertificateStoreType{
				Name: testStoreType.Name,
			},
			wantErr: false,
		},
	}
	runStoreTypeTests(t, tests, c)
}

func TestClient_DeleteCertificateStoreType(t *testing.T) {
	log.SetOutput(io.Discard)
	c, kfcErr := api.NewKeyfactorClient(&api.AuthConfig{})
	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}

	tests := []storeTypeTest{
		{
			name:   "CreateStoreType",
			fields: storeTypeTestFields{},
			args: storeTypeTestArgs{
				name:      testStoreType.ShortName,
				storeType: testStoreType,
				action:    ActionCreate,
			},
			want:    testStoreType,
			wantErr: false,
		},
		{
			name:   "DeleteStoreType",
			fields: storeTypeTestFields{},
			args: storeTypeTestArgs{
				name:      testStoreType.ShortName,
				storeType: testStoreType,
				action:    ActionDelete,
			},
		},
	}
	if c != nil {
		runStoreTypeTests(t, tests, c)
	}

}

func TestClient_GetCertificateStoreType(t *testing.T) {
	log.SetOutput(io.Discard)
	c, kfcErr := api.NewKeyfactorClient(&api.AuthConfig{})
	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}
	var tests = []storeTypeTest{
		{
			name:   "CreateStoreType",
			fields: storeTypeTestFields{},
			args: storeTypeTestArgs{
				name:      testStoreType.ShortName,
				storeType: testStoreType,
				action:    ActionCreate,
			},
			want:    testStoreType,
			wantErr: false,
		},
		{
			name:   "GetStoreType",
			fields: storeTypeTestFields{},
			args: storeTypeTestArgs{
				name:      testStoreType.ShortName,
				storeType: testStoreType,
				action:    ActionGet,
			},
			want:    testStoreType,
			wantErr: false,
		},
		{
			name:   "DeleteStoreType",
			fields: storeTypeTestFields{},
			args: storeTypeTestArgs{
				name:      testStoreType.ShortName,
				storeType: testStoreType,
				action:    ActionDelete,
			},
		},
	}
	runStoreTypeTests(t, tests, c)
}

func TestClient_ListCertificateStoreTypes(t *testing.T) {
	log.SetOutput(io.Discard)
	c, kfcErr := api.NewKeyfactorClient(&api.AuthConfig{})
	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}
	tests := []storeTypeTest{
		{
			name:   "ListStoreTypes",
			fields: storeTypeTestFields{},
			args: storeTypeTestArgs{
				name:      testStoreType.ShortName,
				storeType: testStoreType,
				action:    ActionList,
			},
			want:    testStoreType,
			wantErr: false,
		},
	}
	runStoreTypeTests(t, tests, c)
}

func TestClient_UpdateStoreType(t *testing.T) {
	log.SetOutput(io.Discard)
	c, kfcErr := api.NewKeyfactorClient(&api.AuthConfig{})
	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}
	updatedStoreType := *testStoreType
	updatedStoreType.Name = "TestStoreType8"
	updatedStoreType.PowerShell = !testStoreType.PowerShell

	tests := []storeTypeTest{
		{
			name:   "CreateStoreType",
			fields: storeTypeTestFields{},
			args: storeTypeTestArgs{
				name:      testStoreType.ShortName,
				storeType: testStoreType,
				action:    ActionCreate,
			},
			want:    testStoreType,
			wantErr: false,
		},
		{
			name:   "UpdateStoreType",
			fields: storeTypeTestFields{},
			args: storeTypeTestArgs{
				name:      updatedStoreType.ShortName,
				storeType: &updatedStoreType,
				action:    ActionUpdate,
			},
			want:    &updatedStoreType,
			wantErr: false,
		},
		{
			name:   "DeleteStoreType",
			fields: storeTypeTestFields{},
			args: storeTypeTestArgs{
				name:      testStoreType.ShortName,
				storeType: testStoreType,
				action:    ActionDelete,
			},
			want:    testStoreType,
			wantErr: false,
		},
	}
	runStoreTypeTests(t, tests, c)
}
