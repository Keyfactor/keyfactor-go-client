package api

type CertificateStoreTypeGeneric struct {
	Name                string                                `json:"Name"`
	ShortName           string                                `json:"ShortName"`
	Capability          string                                `json:"Capability"`
	LocalStore          bool                                  `json:"LocalStore"`
	SupportedOperations *StoreTypeSupportedOperations         `json:"SupportedOperations"`
	Properties          *[]StoreTypePropertyDefinitionGeneric `json:"Properties"`
	EntryParameters     *[]EntryParameterGeneric              `json:"EntryParameters"`
	PasswordOptions     *StoreTypePasswordOptions             `json:"PasswordOptions"`
	//StorePathType       string                                `json:"StorePathType"` # This is not returned in the API and computed after POST
	StorePathValue    string `json:"StorePathValue"`
	PrivateKeyAllowed string `json:"PrivateKeyAllowed"`
	//JobProperties       *[]string                             `json:"JobProperties"` # This is not returned in the API and computed after POST
	ServerRequired     bool   `json:"ServerRequired"`
	PowerShell         bool   `json:"PowerShell"`
	BlueprintAllowed   bool   `json:"BlueprintAllowed"`
	CustomAliasAllowed string `json:"CustomAliasAllowed"`
}

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

type StoreTypePropertyDefinitionGeneric struct {
	Name         string      `json:"Name"`
	DisplayName  string      `json:"DisplayName"`
	Type         string      `json:"Type"`
	DependsOn    string      `json:"DependsOn"`
	DefaultValue interface{} `json:"DefaultValue"`
	Required     bool        `json:"Required"`
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

type EntryParameterGeneric struct {
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
	EntrySupported bool   `json:"EntrySupported"`
	StoreRequired  bool   `json:"StoreRequired"`
	Style          string `json:"Style"`
}

type StoreTypeSupportedOperations struct {
	Add        bool `json:"Add"`
	Create     bool `json:"Create"`
	Discovery  bool `json:"Discovery"`
	Enrollment bool `json:"Enrollment"`
	Remove     bool `json:"Remove"`
}
