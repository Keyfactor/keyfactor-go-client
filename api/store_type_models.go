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

type CertificateStoreType struct {
	Name                string                         `json:"Name"`
	ShortName           string                         `json:"ShortName"`
	Capability          string                         `json:"Capability"`
	StoreType           int                            `json:"StoreType"`
	ImportType          int                            `json:"ImportType"`
	LocalStore          bool                           `json:"LocalStore"`
	SupportedOperations *StoreTypeSupportedOperations  `json:"SupportedOperations"`
	Properties          *[]StoreTypePropertyDefinition `json:"Properties"`
	EntryParameters     *[]EntryParameter              `json:"EntryParameters"`
	PasswordOptions     *StoreTypePasswordOptions      `json:"PasswordOptions"`
	StorePathType       string                         `json:"StorePathType"`
	StorePathValue      string                         `json:"StorePathValue"`
	PrivateKeyAllowed   string                         `json:"PrivateKeyAllowed"`
	JobProperties       *[]string                      `json:"JobProperties"`
	ServerRequired      bool                           `json:"ServerRequired"`
	PowerShell          bool                           `json:"PowerShell"`
	BlueprintAllowed    bool                           `json:"BlueprintAllowed"`
	CustomAliasAllowed  string                         `json:"CustomAliasAllowed"`
	ServerRegistration  int                            `json:"ServerRegistration"`
	InventoryEndpoint   string                         `json:"InventoryEndpoint"`
	InventoryJobType    string                         `json:"InventoryJobType"`
	ManagementJobType   string                         `json:"ManagementJobType"`
	DiscoveryJobType    string                         `json:"DiscoveryJobType"`
	EnrollmentJobType   string                         `json:"EnrollmentJobType"`
}

type CertStoreTypeResponseList []struct {
	CertStoreTypeResponse
}

type DeleteStoreType struct {
	ID int `json:"id"`
}

type StoreTypePropertyDefinition struct {
	StoreTypeID  int         `json:"StoreTypeId;omitempty"`
	Name         string      `json:"Name"`
	DisplayName  string      `json:"DisplayName"`
	Type         string      `json:"Type"`
	DependsOn    string      `json:"DependsOn"`
	DefaultValue interface{} `json:"DefaultValue"`
	Required     bool        `json:"Required"`
}

type EntryParameter struct {
	StoreTypeId  int    `json:"StoreTypeId"`
	Name         string `json:"Name"`
	DisplayName  string `json:"DisplayName"`
	Type         string `json:"Type"`
	RequiredWhen struct {
		HasPrivateKey  bool `json:"HasPrivateKey"`
		OnAdd          bool `json:"OnAdd"`
		OnRemove       bool `json:"OnRemove"`
		OnReenrollment bool `json:"OnReenrollment"`
	}
	DependsOn    string `json:"DependsOn"`
	DefaultValue string `json:"DefaultValue"`
	Options      string `json:"Options"`
}

type StoreTypePasswordOptions struct {
	EntrySupported bool   `json:"EntrySupported,omitempty"`
	StoreRequired  bool   `json:"StoreRequired,omitempty"`
	Style          string `json:"string,Style,omitempty"`
}

type StoreTypeSupportedOperations struct {
	Add        bool `json:"Add"`
	Create     bool `json:"Create"`
	Discovery  bool `json:"Discovery"`
	Enrollment bool `json:"Enrollment"`
	Remove     bool `json:"Remove"`
}
