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

// EnrollmentPatternCreateRequest represents the request structure for creating a new enrollment pattern
type EnrollmentPatternCreateRequest struct {
	Template               int                                     `json:"Template"`
	Name                   string                                  `json:"Name"`
	Description            *string                                 `json:"Description,omitempty"`
	TemplateDefault        bool                                    `json:"TemplateDefault,omitempty"`
	AssociatedRoles        []string                                `json:"AssociatedRoles,omitempty"`
	UseADPermissions       bool                                    `json:"UseADPermissions,omitempty"`
	CertificateAuthorities []int                                   `json:"CertificateAuthorities,omitempty"`
	AllowedEnrollmentTypes int                                     `json:"AllowedEnrollmentTypes,omitempty"`
	Regexes                []EnrollmentPatternRegexesRequest       `json:"Regexes,omitempty"`
	MetadataFields         []EnrollmentPatternMetadataFieldRequest `json:"MetadataFields,omitempty"`
	RestrictCAs            bool                                    `json:"RestrictCAs,omitempty"`
	Policies               EnrollmentPatternPolicyRequest          `json:"Policies"`
	Defaults               []EnrollmentPatternDefaultRequest       `json:"Defaults,omitempty"`
	EnrollmentFields       []EnrollmentPatternFieldRequest         `json:"EnrollmentFields,omitempty"`
}

// EnrollmentPatternRequest represents the request structure for updating an enrollment pattern
type EnrollmentPatternRequest struct {
	Name                   string                                  `json:"Name"`
	Description            string                                  `json:"Description,omitempty"`
	TemplateDefault        bool                                    `json:"TemplateDefault,omitempty"`
	AssociatedRoles        []string                                `json:"AssociatedRoles,omitempty"`
	UseADPermissions       bool                                    `json:"UseADPermissions,omitempty"`
	CertificateAuthorities []int                                   `json:"CertificateAuthorities,omitempty"`
	AllowedEnrollmentTypes int                                     `json:"AllowedEnrollmentTypes,omitempty"`
	Regexes                []EnrollmentPatternRegexesRequest       `json:"Regexes,omitempty"`
	MetadataFields         []EnrollmentPatternMetadataFieldRequest `json:"MetadataFields,omitempty"`
	RestrictCAs            bool                                    `json:"RestrictCAs,omitempty"`
	Policies               EnrollmentPatternPolicyRequest          `json:"Policies"`
	Defaults               []EnrollmentPatternDefaultRequest       `json:"Defaults,omitempty"`
	EnrollmentFields       []EnrollmentPatternFieldRequest         `json:"EnrollmentFields,omitempty"`
}

// EnrollmentPatternResponse represents the response structure for enrollment pattern operations
type EnrollmentPatternResponse struct {
	ID                     int                                       `json:"Id,omitempty"`
	Name                   string                                    `json:"Name,omitempty"`
	Description            string                                    `json:"Description,omitempty"`
	Template               *EnrollmentPatternTemplateResponse        `json:"Template,omitempty"`
	TemplateDefault        bool                                      `json:"TemplateDefault,omitempty"`
	UseADPermissions       bool                                      `json:"UseADPermissions,omitempty"`
	AssociatedRoles        []EnrollmentPatternAssociatedRoleResponse `json:"AssociatedRoles,omitempty"`
	CertificateAuthorities []EnrollmentPatternCAResponse             `json:"CertificateAuthorities,omitempty"`
	AllowedEnrollmentTypes *int                                      `json:"AllowedEnrollmentTypes,omitempty"`
	Regexes                []EnrollmentPatternRegexesResponse        `json:"Regexes,omitempty"`
	MetadataFields         []EnrollmentPatternMetadataFieldResponse  `json:"MetadataFields,omitempty"`
	RestrictCAs            bool                                      `json:"RestrictCAs,omitempty"`
	Policies               *EnrollmentPatternPolicyResponse          `json:"Policies,omitempty"`
	Defaults               []EnrollmentPatternDefaultResponse        `json:"Defaults,omitempty"`
	EnrollmentFields       []EnrollmentPatternFieldResponse          `json:"EnrollmentFields,omitempty"`
}

// EnrollmentPatternRegexesRequest represents regex validation rules for enrollment patterns
type EnrollmentPatternRegexesRequest struct {
	SubjectPart   string `json:"SubjectPart"`
	Regex         string `json:"Regex,omitempty"`
	Error         string `json:"Error,omitempty"`
	CaseSensitive bool   `json:"CaseSensitive,omitempty"`
}

// EnrollmentPatternRegexesResponse represents regex validation rules in responses
type EnrollmentPatternRegexesResponse struct {
	SubjectPart   string `json:"SubjectPart,omitempty"`
	Regex         string `json:"Regex,omitempty"`
	Error         string `json:"Error,omitempty"`
	CaseSensitive bool   `json:"CaseSensitive,omitempty"`
}

// EnrollmentPatternPolicyRequest represents policy settings for enrollment patterns
type EnrollmentPatternPolicyRequest struct {
	AllowKeyReuse                   *bool                    `json:"AllowKeyReuse,omitempty"`
	AllowWildcards                  *bool                    `json:"AllowWildcards,omitempty"`
	RFCEnforcement                  *bool                    `json:"RFCEnforcement,omitempty"`
	CertificateOwnerRole            *int                     `json:"CertificateOwnerRole,omitempty"`
	DefaultCertificateOwnerRoleId   *int                     `json:"DefaultCertificateOwnerRoleId,omitempty"`
	DefaultCertificateOwnerRoleName *string                  `json:"DefaultCertificateOwnerRoleName,omitempty"`
	DefaultCertificateOwnerOverride bool                     `json:"DefaultCertificateOwnerOverride,omitempty"`
	PrimaryKeyAlgorithms            []AlgorithmDataRequestV2 `json:"PrimaryKeyAlgorithms,omitempty"`
	AlternativeKeyAlgorithms        []AlgorithmDataRequestV2 `json:"AlternativeKeyAlgorithms,omitempty"`
}

// EnrollmentPatternPolicyResponse represents policy settings in responses
type EnrollmentPatternPolicyResponse struct {
	AllowKeyReuse                   bool                    `json:"AllowKeyReuse,omitempty"`
	AllowWildcards                  bool                    `json:"AllowWildcards,omitempty"`
	RFCEnforcement                  bool                    `json:"RFCEnforcement,omitempty"`
	CertificateOwnerRole            int                     `json:"CertificateOwnerRole,omitempty"`
	DefaultCertificateOwnerRoleId   int                     `json:"DefaultCertificateOwnerRoleId,omitempty"`
	DefaultCertificateOwnerRoleName string                  `json:"DefaultCertificateOwnerRoleName,omitempty"`
	DefaultCertificateOwnerOverride bool                    `json:"DefaultCertificateOwnerOverride,omitempty"`
	PrimaryKeyAlgorithms            []AlgorithmDataResponse `json:"PrimaryKeyAlgorithms,omitempty"`
	AlternativeKeyAlgorithms        []AlgorithmDataResponse `json:"AlternativeKeyAlgorithms,omitempty"`
}

// EnrollmentPatternMetadataFieldRequest represents metadata field configuration for requests
type EnrollmentPatternMetadataFieldRequest struct {
	Id             int    `json:"Id,omitempty"`
	DefaultValue   string `json:"DefaultValue,omitempty"`
	Validation     string `json:"Validation,omitempty"`
	Enrollment     int    `json:"Enrollment,omitempty"`
	Message        string `json:"Message,omitempty"`
	Options        string `json:"Options,omitempty"`
	DependsOn      string `json:"DependsOn,omitempty"`
	DependsOnValue string `json:"DependsOnValue,omitempty"`
}

// EnrollmentPatternMetadataFieldResponse represents metadata field configuration in responses
type EnrollmentPatternMetadataFieldResponse struct {
	MetadataId    int    `json:"MetadataId,omitempty"`
	DefaultValue  string `json:"DefaultValue,omitempty"`
	Validation    string `json:"Validation,omitempty"`
	Enrollment    int    `json:"Enrollment,omitempty"`
	Message       string `json:"Message,omitempty"`
	CaseSensitive bool   `json:"CaseSensitive,omitempty"`
}

// EnrollmentPatternDefaultRequest represents default value settings for requests
type EnrollmentPatternDefaultRequest struct {
	SubjectPart  string `json:"SubjectPart"`
	DefaultValue string `json:"DefaultValue,omitempty"`
}

// EnrollmentPatternDefaultResponse represents default value settings in responses
type EnrollmentPatternDefaultResponse struct {
	SubjectPart string `json:"SubjectPart,omitempty"`
	Value       string `json:"Value,omitempty"`
}

// EnrollmentPatternFieldRequest represents enrollment field configuration for requests
type EnrollmentPatternFieldRequest struct {
	Id             int    `json:"Id,omitempty"`
	DefaultValue   string `json:"DefaultValue,omitempty"`
	Validation     string `json:"Validation,omitempty"`
	Enrollment     int    `json:"Enrollment,omitempty"`
	Message        string `json:"Message,omitempty"`
	Options        string `json:"Options,omitempty"`
	DependsOn      string `json:"DependsOn,omitempty"`
	DependsOnValue string `json:"DependsOnValue,omitempty"`
}

// EnrollmentPatternFieldResponse represents enrollment field configuration in responses
type EnrollmentPatternFieldResponse struct {
	Id             int      `json:"Id,omitempty"`
	Name           string   `json:"Name,omitempty"`
	DefaultValue   string   `json:"DefaultValue,omitempty"`
	Validation     string   `json:"Validation,omitempty"`
	Enrollment     int      `json:"Enrollment,omitempty"`
	Message        string   `json:"Message,omitempty"`
	Options        []string `json:"Options,omitempty"`
	DependsOn      string   `json:"DependsOn,omitempty"`
	DependsOnValue string   `json:"DependsOnValue,omitempty"`
	DataType       int      `json:"DataType,omitempty"`
	Hint           string   `json:"Hint,omitempty"`
}

// EnrollmentPatternTemplateResponse represents template information in responses
type EnrollmentPatternTemplateResponse struct {
	Id                  int    `json:"Id,omitempty"`
	TemplateName        string `json:"TemplateName,omitempty"`
	CommonName          string `json:"CommonName,omitempty"`
	ConfigurationTenant string `json:"ConfigurationTenant,omitempty"`
	RequiresApproval    bool   `json:"RequiresApproval,omitempty"`
	FriendlyName        string `json:"FriendlyName,omitempty"`
}

// EnrollmentPatternAssociatedRoleResponse represents associated role information in responses
type EnrollmentPatternAssociatedRoleResponse struct {
	Id   int    `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
}

// EnrollmentPatternCAResponse represents certificate authority information in responses
type EnrollmentPatternCAResponse struct {
	Id                  int    `json:"Id,omitempty"`
	LogicalName         string `json:"LogicalName,omitempty"`
	HostName            string `json:"HostName,omitempty"`
	ConfigurationTenant string `json:"ConfigurationTenant,omitempty"`
}

// AlgorithmDataRequestV2 represents algorithm configuration for requests
type AlgorithmDataRequestV2 struct {
	KeyType   *string `json:"KeyType,omitempty"`
	KeySize   *int    `json:"KeySize,omitempty"`
	CurveName *string `json:"CurveName,omitempty"`
}

// AlgorithmDataResponse represents algorithm configuration in responses
type AlgorithmDataResponse struct {
	Name       string   `json:"Name,omitempty"`
	BitLengths []int    `json:"bit_lengths,omitempty"`
	Curves     []string `json:"curves,omitempty"`
}

// EnrollmentPatternsQueryParams represents query parameters for listing enrollment patterns
type EnrollmentPatternsQueryParams struct {
	QueryString   string `json:"queryString,omitempty"`
	PageReturned  int    `json:"pageReturned,omitempty"`
	ReturnLimit   int    `json:"returnLimit,omitempty"`
	SortField     string `json:"sortField,omitempty"`
	SortAscending *int   `json:"sortAscending,omitempty"` // 0=ascending, 1=descending
}
