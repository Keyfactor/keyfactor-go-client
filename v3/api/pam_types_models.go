// Copyright 2025 Keyfactor
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

// PamParameterDataType represents the data type of a PAM parameter
// 1 = string, 2 = secret
type PamParameterDataType int

const (
	PamParameterDataTypeString PamParameterDataType = 1
	PamParameterDataTypeSecret PamParameterDataType = 2
)

// SecretType represents the type of secret in the system
// 0-4 are valid values
type SecretType int

// ProviderTypeParameterResponse represents a parameter for a PAM provider type
type ProviderTypeParameterResponse struct {
	Id            int                  `json:"Id,omitempty"`
	Name          string               `json:"Name,omitempty"`
	DisplayName   *string              `json:"DisplayName,omitempty"`
	DataType      PamParameterDataType `json:"DataType,omitempty"`
	InstanceLevel bool                 `json:"InstanceLevel,omitempty"`
}

// ProviderTypeResponse represents a PAM provider type
type ProviderTypeResponse struct {
	Id         string                           `json:"Id,omitempty"` // UUID format
	Name       string                           `json:"Name,omitempty"`
	Parameters *[]ProviderTypeParameterResponse `json:"Parameters,omitempty"`
}

// ProviderTypeParameterCreateRequest represents a request to create a PAM provider type parameter
type ProviderTypeParameterCreateRequest struct {
	Name          string               `json:"Name"`
	DisplayName   *string              `json:"DisplayName,omitempty"`
	DataType      PamParameterDataType `json:"DataType,omitempty"`
	InstanceLevel bool                 `json:"InstanceLevel,omitempty"`
}

// ProviderTypeCreateRequest represents a request to create a PAM provider type
type ProviderTypeCreateRequest struct {
	Name       string                                `json:"Name"`
	Parameters *[]ProviderTypeParameterCreateRequest `json:"Parameters,omitempty"`
}

// ProviderCreateRequestProviderTypeParam represents a provider type parameter in a provider creation request
type ProviderCreateRequestProviderTypeParam struct {
	Id            int     `json:"Id,omitempty"`
	Name          string  `json:"Name,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty"`
	InstanceLevel bool    `json:"InstanceLevel,omitempty"`
}

// PamProviderTypeParam represents a provider type parameter (full model) for PAM operations
type PamProviderTypeParam struct {
	Id            int                  `json:"Id,omitempty"`
	Name          string               `json:"Name,omitempty"`
	DisplayName   *string              `json:"DisplayName,omitempty"`
	DataType      PamParameterDataType `json:"DataType,omitempty"`
	InstanceLevel bool                 `json:"InstanceLevel,omitempty"`
	ProviderType  *ProviderType        `json:"ProviderType,omitempty"`
}

// ProviderType represents a PAM provider type (full model)
type ProviderType struct {
	Id                 string                  `json:"Id,omitempty"` // UUID format
	Name               string                  `json:"Name,omitempty"`
	ProviderTypeParams *[]PamProviderTypeParam `json:"ProviderTypeParams,omitempty"`
}

// PamProviderTypeParamValue represents a parameter value for a PAM provider type
type PamProviderTypeParamValue struct {
	Id                int                   `json:"Id,omitempty"`
	Value             *string               `json:"Value,omitempty"`
	ParameterId       int                   `json:"ParameterId,omitempty"`
	InstanceId        *int                  `json:"InstanceId,omitempty"`
	InstanceGuid      *string               `json:"InstanceGuid,omitempty"` // UUID format
	Provider          *Provider             `json:"Provider,omitempty"`
	ProviderTypeParam *PamProviderTypeParam `json:"ProviderTypeParam,omitempty"`
}

// PamProviderTypeParamValueResponse represents a parameter value response for a PAM provider type
type PamProviderTypeParamValueResponse struct {
	Id                int                            `json:"Id,omitempty"`
	Value             *string                        `json:"Value,omitempty"`
	InstanceId        *int                           `json:"InstanceId,omitempty"`
	InstanceGuid      *string                        `json:"InstanceGuid,omitempty"` // UUID format
	ProviderTypeParam *ProviderTypeParameterResponse `json:"ProviderTypeParam,omitempty"`
}
