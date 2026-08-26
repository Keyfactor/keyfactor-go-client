package api

// Provider represents a PAM provider (full model)
type Provider struct {
	Id                      int                          `json:"Id,omitempty"`
	Name                    string                       `json:"Name"`
	Area                    int                          `json:"Area,omitempty"`
	ProviderType            ProviderType                 `json:"ProviderType"`
	ProviderTypeParamValues *[]PamProviderTypeParamValue `json:"ProviderTypeParamValues,omitempty"`
	SecuredAreaId           *int                         `json:"SecuredAreaId,omitempty"`
	Remote                  bool                         `json:"Remote,omitempty"`
	IsInUse                 bool                         `json:"IsInUse,omitempty"`
	IsLocalDB               bool                         `json:"IsLocalDB,omitempty"`
}

// ProviderCreateRequestTypeParamValue represents a parameter value in a provider creation request
type ProviderCreateRequestTypeParamValue struct {
	Id                int                                     `json:"Id,omitempty"`
	Value             *string                                 `json:"Value,omitempty"`
	InstanceId        *int                                    `json:"InstanceId,omitempty"`
	InstanceGuid      *string                                 `json:"InstanceGuid,omitempty"` // UUID format
	ProviderTypeParam *ProviderCreateRequestProviderTypeParam `json:"ProviderTypeParam,omitempty"`
}

// ProviderCreateRequestProviderType represents a provider type reference in a provider creation request
type ProviderCreateRequestProviderType struct {
	Id string `json:"Id,omitempty"` // UUID format
}

// ProviderCreateRequest represents a request to create a PAM provider
type ProviderCreateRequest struct {
	Name                    string                       `json:"Name"`
	Remote                  bool                         `json:"Remote,omitempty"`
	Area                    int                          `json:"Area,omitempty"`
	ProviderType            ProviderType                 `json:"ProviderType"`
	ProviderTypeParamValues *[]PamProviderTypeParamValue `json:"ProviderTypeParamValues,omitempty"`
	SecuredAreaId           *int                         `json:"SecuredAreaId,omitempty"`
}

// ProviderUpdateRequestLegacy represents a request to update a PAM provider (legacy format)
type ProviderUpdateRequestLegacy struct {
	Id                      int                          `json:"Id"`
	Name                    string                       `json:"Name"`
	Remote                  bool                         `json:"Remote,omitempty"`
	Area                    int                          `json:"Area,omitempty"`
	ProviderType            ProviderType                 `json:"ProviderType"`
	ProviderTypeParamValues *[]PamProviderTypeParamValue `json:"ProviderTypeParamValues,omitempty"`
	SecuredAreaId           *int                         `json:"SecuredAreaId,omitempty"`
}

// ProviderResponseLegacy represents a PAM provider response (legacy format)
type ProviderResponseLegacy struct {
	Id                      int                                  `json:"Id,omitempty"`
	Name                    *string                              `json:"Name,omitempty"`
	Area                    int                                  `json:"Area,omitempty"`
	ProviderType            *ProviderType                        `json:"ProviderType,omitempty"`
	ProviderTypeParamValues *[]PamProviderTypeParamValueResponse `json:"ProviderTypeParamValues,omitempty"`
	SecuredAreaId           *int                                 `json:"SecuredAreaId,omitempty"`
	Remote                  bool                                 `json:"Remote,omitempty"`
}

// LocalPAMEntryCreateRequest represents a request to create a local PAM entry
type LocalPAMEntryCreateRequest struct {
	SecretName  string  `json:"SecretName"`
	Description *string `json:"Description,omitempty"`
	SecretValue string  `json:"SecretValue"`
}

// LocalPAMEntryUpdateRequest represents a request to update a local PAM entry
type LocalPAMEntryUpdateRequest struct {
	SecretName  string  `json:"SecretName"`
	Description *string `json:"Description,omitempty"`
	SecretValue *string `json:"SecretValue,omitempty"`
}

// LocalPAMEntryResponse represents a local PAM entry response
type LocalPAMEntryResponse struct {
	ProviderId  int     `json:"ProviderId,omitempty"`
	SecretName  *string `json:"SecretName,omitempty"`
	Description *string `json:"Description,omitempty"`
}

// KeyfactorSecret represents a Keyfactor secret
type KeyfactorSecret struct {
	Value                       interface{}                  `json:"Value,omitempty"`
	SecretTypeGuid              string                       `json:"SecretTypeGuid,omitempty"` // UUID format
	InstanceId                  *int                         `json:"InstanceId,omitempty"`
	InstanceGuid                *string                      `json:"InstanceGuid,omitempty"` // UUID format
	ProviderTypeParameterValues *[]PamProviderTypeParamValue `json:"ProviderTypeParameterValues,omitempty"`
	ProviderId                  *int                         `json:"ProviderId,omitempty"`
	IsManaged                   bool                         `json:"IsManaged,omitempty"`
	SecretType                  SecretType                   `json:"SecretType,omitempty"`
	RemoteProviderName          *string                      `json:"RemoteProviderName,omitempty"`
	HasValue                    bool                         `json:"HasValue,omitempty"`
}

// KeyfactorAPISecret represents a Keyfactor API secret
type KeyfactorAPISecret struct {
	SecretValue *string            `json:"SecretValue,omitempty"`
	Parameters  map[string]*string `json:"Parameters,omitempty"`
	Provider    *int               `json:"Provider,omitempty"`
}
