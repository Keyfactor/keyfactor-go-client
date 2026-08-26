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

type GetTemplateResponse struct {
	Id                     int                        `json:"Id,omitempty"`
	CommonName             string                     `json:"CommonName,omitempty"`
	TemplateName           string                     `json:"TemplateName,omitempty"`
	Oid                    string                     `json:"Oid,omitempty"`
	KeySize                string                     `json:"KeySize,omitempty"`
	KeyType                string                     `json:"KeyType,omitempty"`
	ForestRoot             string                     `json:"ForestRoot,omitempty"`
	FriendlyName           string                     `json:"FriendlyName,omitempty"`
	KeyRetention           string                     `json:"KeyRetention,omitempty"`
	KeyRetentionDays       int                        `json:"KeyRetentionDays,omitempty"`
	KeyArchival            bool                       `json:"KeyArchival,omitempty"`
	EnrollmentFields       []TemplateEnrollmentFields `json:"EnrollmentFields,omitempty"`
	MetadataFields         []TemplateMetadataFields   `json:"MetadataFields,omitempty"`
	AllowedEnrollmentTypes int                        `json:"AllowedEnrollmentTypes,omitempty"`
	TemplateRegexes        []TemplateRegex            `json:"TemplateRegexes,omitempty"`
	UseAllowedRequesters   bool                       `json:"UseAllowedRequesters,omitempty"`
	AllowedRequesters      []string                   `json:"AllowedRequesters,omitempty"`
	RFCEnforcement         bool                       `json:"RFCEnforcement,omitempty"`
	RequiresApproval       bool                       `json:"RequiresApproval,omitempty"`
	KeyUsage               int                        `json:"KeyUsage,omitempty"`
	// TemplatePolicy carries the template's key-algorithm policy (PrimaryKeyAlgorithms /
	// AlternativeKeyAlgorithms, wildcard/key-reuse flags, certificate owner role, etc).
	// It is null for templates that predate this policy model or have never had it
	// configured. Command's PUT /Templates full-replace validation derives an internal
	// "Policies" set from TemplatePolicy.PrimaryKeyAlgorithms/AlternativeKeyAlgorithms;
	// for templates linked to an enrollment pattern, omitting TemplatePolicy on update
	// collapses that set to empty and Command rejects the request with
	// "'Policies' cannot be empty" (confirmed against a live Command 25.4.1 instance).
	TemplatePolicy *TemplatePolicy `json:"TemplatePolicy,omitempty"`
}

// TemplateKeyAlgorithm describes one allowed key algorithm entry within a
// TemplatePolicy's PrimaryKeyAlgorithms/AlternativeKeyAlgorithms list. Field
// names intentionally match Command's lowercase/snake_case wire format for
// this nested object (unlike the rest of the Templates API, which is
// PascalCase).
type TemplateKeyAlgorithm struct {
	Name       string   `json:"name,omitempty"`
	BitLengths []int    `json:"bit_lengths,omitempty"`
	Curves     []string `json:"curves,omitempty"`
}

// TemplatePolicy models Command's per-template key/enrollment policy object,
// returned under GetTemplateResponse.TemplatePolicy and required (when the
// template has one configured) on UpdateTemplateArg.TemplatePolicy to avoid
// Command's full-replace PUT /Templates clearing it. See the comment on
// GetTemplateResponse.TemplatePolicy for the "'Policies' cannot be empty"
// validation error this addresses.
type TemplatePolicy struct {
	TemplateId                      int                    `json:"TemplateId,omitempty"`
	AllowKeyReuse                   *bool                  `json:"AllowKeyReuse,omitempty"`
	AllowWildcards                  *bool                  `json:"AllowWildcards,omitempty"`
	RFCEnforcement                  *bool                  `json:"RFCEnforcement,omitempty"`
	CertificateOwnerRole            *int                   `json:"CertificateOwnerRole,omitempty"`
	DefaultCertificateOwnerRoleId   *int                   `json:"DefaultCertificateOwnerRoleId,omitempty"`
	DefaultCertificateOwnerRoleName *string                `json:"DefaultCertificateOwnerRoleName,omitempty"`
	PrimaryKeyAlgorithms            []TemplateKeyAlgorithm `json:"PrimaryKeyAlgorithms,omitempty"`
	AlternativeKeyAlgorithms        []TemplateKeyAlgorithm `json:"AlternativeKeyAlgorithms,omitempty"`
}

type TemplateEnrollmentFields struct {
	Id       int
	Name     string
	Options  []string
	DataType int
}

type TemplateMetadataFields struct {
	Id           int
	DefaultValue string
	MetadataId   int
	Validation   string
	Enrollment   int
	Message      string
	Options      string
}

type TemplateRegex struct {
	TemplateId  int
	SubjectPart string
	RegEx       string
	Error       string
}

type UpdateTemplateArg struct {
	Id                     int                         `json:"Id,omitempty"`
	CommonName             string                      `json:"CommonName,omitempty"`
	TemplateName           string                      `json:"TemplateName,omitempty"`
	Oid                    string                      `json:"Oid,omitempty"`
	KeySize                string                      `json:"KeySize,omitempty"`
	KeyType                *string                     `json:"KeyType,omitempty"`
	ForestRoot             string                      `json:"ForestRoot,omitempty"`
	FriendlyName           *string                     `json:"FriendlyName,omitempty"`
	KeyRetention           *string                     `json:"KeyRetention,omitempty"`
	KeyRetentionDays       *int                        `json:"KeyRetentionDays,omitempty"`
	KeyArchival            *bool                       `json:"KeyArchival,omitempty"`
	EnrollmentFields       *[]TemplateEnrollmentFields `json:"EnrollmentFields,omitempty"`
	MetadataFields         *[]TemplateMetadataFields   `json:"MetadataFields,omitempty"`
	AllowedEnrollmentTypes *int                        `json:"AllowedEnrollmentTypes,omitempty"`
	TemplateRegexes        *[]TemplateRegex            `json:"TemplateRegexes,omitempty"`
	UseAllowedRequesters   *bool                       `json:"UseAllowedRequesters,omitempty"`
	AllowedRequesters      *[]string                   `json:"AllowedRequesters,omitempty"`
	RFCEnforcement         *bool                       `json:"RFCEnforcement,omitempty"`
	RequiresApproval       *bool                       `json:"RequiresApproval,omitempty"`
	// KeyUsage is an int32 bitmask on Command's wire format (e.g. 160 =
	// digitalSignature|keyEncipherment), matching GetTemplateResponse.KeyUsage and
	// Command's TemplateUpdateRequest/TemplateRetrievalResponse swagger schema
	// (both typed "integer"/"int32"). A *bool here previously produced a live
	// HTTP 400 ("Unexpected character encountered while parsing value: t. Path
	// 'KeyUsage'") since Command rejects a JSON boolean for an integer field.
	KeyUsage *int `json:"KeyUsage,omitempty"`
	// TemplatePolicy must be round-tripped from the corresponding GetTemplateResponse
	// on every update; see the field comment on GetTemplateResponse.TemplatePolicy.
	TemplatePolicy *TemplatePolicy `json:"TemplatePolicy,omitempty"`
}

type UpdateTemplateResponse struct{ GetTemplateResponse }
