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
	Name                string                         `json:"Name,omitempty"`
	ShortName           string                         `json:"ShortName,omitempty"`
	Capability          string                         `json:"Capability,omitempty"`
	StoreType           int                            `json:"StoreType,omitempty"`
	ImportType          int                            `json:"ImportType,omitempty"`
	LocalStore          bool                           `json:"LocalStore,omitempty"`
	SupportedOperations *StoreTypeSupportedOperations  `json:"SupportedOperations,omitempty"`
	Properties          *[]StoreTypePropertyDefinition `json:"Properties,omitempty"`
	EntryParameters     *[]EntryParameter              `json:"EntryParameters,omitempty"`
	PasswordOptions     *StoreTypePasswordOptions      `json:"PasswordOptions,omitempty"`
	StorePathType       string                         `json:"StorePathType,omitempty"`
	StorePathValue      string                         `json:"StorePathValue,omitempty"`
	PrivateKeyAllowed   string                         `json:"PrivateKeyAllowed,omitempty"`
	JobProperties       *[]string                      `json:"JobProperties,omitempty"`
	ServerRequired      bool                           `json:"ServerRequired,omitempty"`
	PowerShell          bool                           `json:"PowerShell,omitempty"`
	BlueprintAllowed    bool                           `json:"BlueprintAllowed,omitempty"`
	CustomAliasAllowed  string                         `json:"CustomAliasAllowed,omitempty"`
	ServerRegistration  int                            `json:"ServerRegistration,omitempty"`
	InventoryEndpoint   string                         `json:"InventoryEndpoint,omitempty"`
	InventoryJobType    string                         `json:"InventoryJobType,omitempty"`
	ManagementJobType   string                         `json:"ManagementJobType,omitempty"`
	DiscoveryJobType    string                         `json:"DiscoveryJobType,omitempty"`
	EnrollmentJobType   string                         `json:"EnrollmentJobType,omitempty"`
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
	DependsOn    interface{} `json:"DependsOn"`
	DefaultValue interface{} `json:"DefaultValue"`
	Required     bool        `json:"Required"`
}

type StoreTypePropertyDefinition struct {
	StoreTypeID  int         `json:"StoreTypeId"`
	Name         string      `json:"Name"`
	DisplayName  string      `json:"DisplayName"`
	Type         string      `json:"Type"`
	DependsOn    interface{} `json:"DependsOn"`
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
