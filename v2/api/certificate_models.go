package api

// SANs holds arrays of strings associated with IPv4 (IP4), IPv6 (IP6), DNS, and URI SANs.
type SANs struct {
	IP4 []string `json:"ip4,omitempty"`
	IP6 []string `json:"ip6,omitempty"`
	DNS []string `json:"dns,omitempty"`
	URI []string `json:"uri,omitempty"`
}

// EnrollPFXFctArgs holds the function arguments used for calling the EnrollPFX method.
type EnrollPFXFctArgs struct {
	CustomFriendlyName          string `json:"CustomFriendlyName,omitempty"`
	Password                    string `json:"Password"`
	PopulateMissingValuesFromAD bool   `json:"PopulateMissingValuesFromAD"`
	// Configure the SubjectString field as the full string subject for the certificate. For example, if you don't have
	// subject fields individually separated, and the subject is already in the format required by RFC5280, use the SubjectString field.
	SubjectString string `json:"Subject"`

	// If the certificate subject is not already in the format required by RFC5280, configure the subject fields using a CertificateSubject
	// struct, and EnrollPFX will automatically compile this information into a proper subject.
	Subject              *CertificateSubject    `json:"-"`
	IncludeChain         bool                   `json:"IncludeChain"`
	RenewalCertificateId int                    `json:"RenewalCertificateId,omitempty"`
	CertificateAuthority string                 `json:"CertificateAuthority"`
	Timestamp            string                 `json:"Timestamp"`
	Template             string                 `json:"Template"`
	SANs                 *SANs                  `json:"SANs,omitempty"`
	Metadata             map[string]interface{} `json:"Metadata,omitempty"`
	CertFormat           string                 `json:"-"`
}

type EnrollPFXFctArgsV2 struct {
	Stores                      []CertificateStore `json:"Stores,omitempty"`
	CustomFriendlyName          string             `json:"CustomFriendlyName,omitempty"`
	Password                    string             `json:"Password"`
	PopulateMissingValuesFromAD bool               `json:"PopulateMissingValuesFromAD"`
	// Configure the SubjectString field as the full string subject for the certificate. For example, if you don't have
	// subject fields individually separated, and the subject is already in the format required by RFC5280, use the SubjectString field.
	SubjectString string `json:"Subject"`

	// If the certificate subject is not already in the format required by RFC5280, configure the subject fields using a CertificateSubject
	// struct, and EnrollPFX will automatically compile this information into a proper subject.
	Subject                              *CertificateSubject    `json:"-"`
	IncludeChain                         bool                   `json:"IncludeChain"`
	RenewalCertificateId                 int                    `json:"RenewalCertificateId,omitempty"`
	CertificateAuthority                 string                 `json:"CertificateAuthority"`
	Timestamp                            string                 `json:"Timestamp"`
	Template                             string                 `json:"Template"`
	SANs                                 *SANs                  `json:"SANs,omitempty"`
	Metadata                             map[string]interface{} `json:"Metadata,omitempty"`
	CertFormat                           string                 `json:"-"`
	InstallIntoExistingCertificateStores bool                   `json:"InstallIntoExistingCertificateStores,omitempty"`
	ChainOrder                           string                 `json:"ChainOrder,omitempty"`
	KeyType                              string                 `json:"KeyType,omitempty"`
	KeyLength                            int                    `json:"KeyLength,omitempty"`
}

// EnrollCSRFctArgs holds the function arguments used for calling the EnrollCSR method.
type EnrollCSRFctArgs struct {
	CSR                  string
	Timestamp            string                 `json:"Timestamp"`
	Template             string                 `json:"Template"`
	CertFormat           string                 `json:"-"`
	CertificateAuthority string                 `json:"CertificateAuthority"`
	IncludeChain         bool                   `json:"IncludeChain"`
	SANs                 *SANs                  `json:"SANs"`
	Metadata             map[string]interface{} `json:"Metadata"`
}

// RevokeCertArgs holds the function arguments used for calling the RevokeCert method.
type RevokeCertArgs struct {
	CertificateIds []int  `json:"CertificateIds"`
	Reason         int    `json:"Reason"`
	Comment        string `json:"Comment"`
	EffectiveDate  string `json:"EffectiveDate"`
	CollectionId   int    `json:"CollectionId,omitempty"`
}

// GetCertificateContextArgs holds the function arguments used for calling the GetCertificateContext method.
type GetCertificateContextArgs struct {
	IncludeMetadata      *bool  `json:"include_metadata,omitempty"`        // Query
	IncludeLocations     *bool  `json:"include_locations,omitempty"`       // Query
	CollectionId         *int   `json:"collection_id,omitempty"`           // Query
	Thumbprint           string `json:"thumbprint,omitempty"`              // Query
	CommonName           string `json:"common_name,omitempty"`             // Query
	Id                   int    `json:"id"`                                // Query
	IncludeHasPrivateKey *bool  `json:"include_has_private_key,omitempty"` // Query
	RequestId            int    `json:"request_id,omitempty"`              // Query
}

// DeployPFXArgs holds the function arguments used for calling the DeployPFXCertificate method.
type DeployPFXArgs struct {
	StoreIds      []string     `json:"StoreIds"`
	Password      string       `json:"Password"`
	StoreTypes    []StoreTypes `json:"StoreTypes"`
	CertificateId int          `json:"CertificateId"`
	RequestId     int          `json:"RequestId"`
	JobTime       *string      `json:"JobTime,omitempty"`
}

// recoverCertArgs holds the function arguments used for calling the RevokeCert method.
type recoverCertArgs struct {
	CertId       int    `json:"CertId,omitempty"`
	Password     string `json:"Password,omitempty"`
	SerialNumber string `json:"SerialNumber,omitempty"`
	IssuerDN     string `json:"IssuerDN,omitempty"`
	Thumbprint   string `json:"Thumbprint,omitempty"`
	IncludeChain bool   `json:"IncludeChain,omitempty"`
	ChainOrder   string `json:"ChainOrder,omitempty"`
}

// StoreTypes holds necessary store type metadata for creating and deploying certificates.
type StoreTypes struct {
	StoreTypeId int       `json:"StoreTypeId"`
	Alias       *string   `json:"Alias,omitempty"`
	Overwrite   *bool     `json:"Overwrite,omitempty"`
	Properties  *[]string `json:"Properties,omitempty"`
}

// DeployPFXResp holds response data from the DeployPFXCertificate method.
type DeployPFXResp struct {
	SuccessfulStores []string `json:"SuccessfulStores"`
	FailedStores     []string `json:"FailedStores"`
}

// CertificateSubject contains string elements for X.509V3 certificate distinguished name (subject)
type CertificateSubject struct {
	SubjectCommonName         string `json:"common_name,omitempty"`
	SubjectLocality           string `json:"locality,omitempty"`
	SubjectOrganization       string `json:"organization,omitempty"`
	SubjectCountry            string `json:"country,omitempty"`
	SubjectOrganizationalUnit string `json:"organizational_unit,omitempty"`
	SubjectState              string `json:"state,omitempty"`
}

// downloadCertificateBody is the API POST request body for PFX enrollment (DownloadCertificate).
type downloadCertificateBody struct {
	CertID       int    `json:"id,omitempty"`
	SerialNumber string `json:"serial_number,omitempty"`
	IssuerDN     string `json:"issuer_dn,omitempty"`
	Thumbprint   string `json:"thumbprint,omitempty"`
	IncludeChain bool   `json:"include_chain,omitempty"`
	ChainOrder   string `json:"chain_order,omitempty"`
}

// EnrollResponse is the outer certificate enrollment response. When Enroll functions are called, the certificates are
// placed inside the Certificates element, and certificate information is placed inside CertificateInformation
type EnrollResponse struct {
	Certificates           []string
	CertificateInformation CertificateInformation `json:"CertificateInformation"`
}

type EnrollResponseV2 struct {
	SuccessfulStores       []string               `json:"SuccessfulStores"`
	CertificateInformation CertificateInformation `json:"CertificateInformation"`
	Metadata               interface{}            `json:"Metadata,omitempty"`
}

// CertificateInformation contains response data from the Enroll methods.
type CertificateInformation struct {
	SerialNumber       string      `json:"SerialNumber"`
	IssuerDN           string      `json:"IssuerDN"`
	Thumbprint         string      `json:"Thumbprint"`
	KeyfactorID        int         `json:"KeyfactorID"`
	KeyfactorRequestID int         `json:"KeyfactorRequestId"`
	PKCS12Blob         string      `json:"PKCS12Blob"`
	Certificates       []string    `json:"Certificates"`
	RequestDisposition string      `json:"RequestDisposition"`
	DispositionMessage string      `json:"DispositionMessage"`
	EnrollmentContext  interface{} `json:"EnrollmentContext"`
}

type CertificateInformationV2 struct {
	SerialNumber              string        `json:"SerialNumber"`
	IssuerDN                  string        `json:"IssuerDN"`
	Thumbprint                string        `json:"Thumbprint"`
	KeyfactorId               int           `json:"KeyfactorId"`
	Pkcs12Blob                string        `json:"Pkcs12Blob"`
	Password                  interface{}   `json:"Password"`
	WorkflowInstanceId        string        `json:"WorkflowInstanceId"`
	WorkflowReferenceId       int           `json:"WorkflowReferenceId"`
	StoreIdsInvalidForRenewal []interface{} `json:"StoreIdsInvalidForRenewal"`
	KeyfactorRequestId        int           `json:"KeyfactorRequestId"`
	RequestDisposition        string        `json:"RequestDisposition"`
	DispositionMessage        string        `json:"DispositionMessage"`
	EnrollmentContext         interface{}   `json:"EnrollmentContext"`
}

// GetCertificateResponse contains the response elements returned from the GetCertificateContext method.
type GetCertificateResponse struct {
	Id                       int    `json:"Id"`
	Thumbprint               string `json:"Thumbprint"`
	SerialNumber             string `json:"SerialNumber"`
	IssuedDN                 string `json:"IssuedDN"`
	IssuedCN                 string `json:"IssuedCN"`
	ImportDate               string `json:"ImportDate"`
	NotBefore                string `json:"NotBefore"`
	NotAfter                 string `json:"NotAfter"`
	IssuerDN                 string `json:"IssuerDN"`
	PrincipalId              string `json:"PrincipalId"`
	TemplateId               int    `json:"TemplateId"`
	CertState                int    `json:"CertState"`
	KeySizeInBits            int    `json:"KeySizeInBits"`
	KeyType                  int    `json:"KeyType"`
	RequesterId              int    `json:"RequesterId"`
	IssuedOU                 string `json:"IssuedOU"`
	KeyUsage                 int    `json:"KeyUsage"`
	SigningAlgorithm         string `json:"SigningAlgorithm"`
	CertStateString          string `json:"CertStateString"`
	KeyTypeString            string `json:"KeyTypeString"`
	RevocationEffDate        string `json:"RevocationEffDate"`
	RevocationReason         int    `json:"RevocationReason"`
	RevocationComment        string `json:"RevocationComment"`
	CertificateAuthorityId   int    `json:"CertificateAuthorityId"`
	CertificateAuthorityName string `json:"CertificateAuthorityName"`
	TemplateName             string `json:"TemplateName"`
	ArchivedKey              bool   `json:"ArchivedKey"`
	HasPrivateKey            bool   `json:"HasPrivateKey"`
	PrincipalName            string `json:"PrincipalName"`
	CertRequestId            int    `json:"CertRequestId"`
	RequesterName            string `json:"RequesterName"`
	ContentBytes             string `json:"ContentBytes"`
	ExtendedKeyUsages        []interface{}
	SubjectAltNameElements   []SubjectAltNameElements `json:"SubjectAltNameElements"`
	CRLDistributionPoints    []CRLDistributionPoints  `json:"CRLDistributionPoints"`
	LocationsCount           []LocationsCount         `json:"LocationsCount"`
	SSLLocations             []SSLLocations           `json:"SSLLocations"`
	Locations                []CertificateLocations   `json:"Locations"`
	Metadata                 interface{}              `json:"Metadata"`
	CertificateKeyId         int                      `json:"CertificateKeyId"`
	CARowIndex               int                      `json:"CARowIndex"`
	DetailedKeyUsage         []DetailedKeyUsage       `json:"detailed_key_usage"`
	KeyRecoverable           bool                     `json:"KeyRecoverable"`
}

type ListCertificateResponse struct {
	Certificates []GetCertificateResponse `json:"Certificates"`
}

// recoverCertResponse contains the response elements returned from the RecoverCertificate method.
type recoverCertResponse struct {
	PFX      string `json:"PFX"`
	FileName string `json:"FileName"`
}

// DetailedKeyUsage contains key useage data returned by the GetCertificateContext method.
type DetailedKeyUsage struct {
	CrlSign          bool   `json:"CrlSign,omitempty"`
	DataEncipherment bool   `json:"DataEncipherment,omitempty"`
	DecipherOnly     bool   `json:"DecipherOnly,omitempty"`
	DigitalSignature bool   `json:"DigitalSignature,omitempty"`
	EncipherOnly     bool   `json:"EncipherOnly,omitempty"`
	KeyAgreement     bool   `json:"KeyAgreement,omitempty"`
	KeyCertSign      bool   `json:"KeyCertSign,omitempty"`
	KeyEncipherment  bool   `json:"KeyEncipherment,omitempty"`
	NonRepudiation   bool   `json:"NonRepudiation,omitempty"`
	HexCode          string `json:"HexCode,omitempty"`
}

// CertificateLocations contains response and request data for the GetCertificateContext and DeployPFXCertificate
// methods
type CertificateLocations struct {
	StoreMachine string `json:"StoreMachine,omitempty"`
	StorePath    string `json:"StorePath,omitempty"`
	StoreType    int    `json:"StoreType,omitempty"`
	Alias        string `json:"Alias,omitempty"`
	ChainLevel   int    `json:"ChainLevel,omitempty"`
	CertStoreId  string `json:"CertStoreId,omitempty"`
}

// SSLLocations contains detailed information on the locations that the certificate was found in a scan.
type SSLLocations struct {
	StorePath   string `json:"StorePath,omitempty"`
	AgentPool   string `json:"AgentPool,omitempty"`
	IPAddress   string `json:"IPAddress,omitempty"`
	Port        int    `json:"Port,omitempty"`
	NetworkName string `json:"NetworkName,omitempty"`
}

// LocationsCount contains details on what kind of and how many stores the certificate is deployed inside.
type LocationsCount struct {
	Type  string `json:"Type,omitempty"`
	Count int    `json:"Count,omitempty"`
}

// CRLDistributionPoints contains details on the CRL distribution and is returned inside GetCertificateResponse with
// the GetCertificateContext method.
type CRLDistributionPoints struct {
	Id      int    `json:"Id"`
	URL     string `json:"URL"`
	URLHash string `json:"URLHash"`
}

// SubjectAltNameElements contains detailed information on the SANs attached to a certificate, and is returned inside
// the GetCertificateContext method
type SubjectAltNameElements struct {
	Id        int    `json:"Id"`
	Value     string `json:"Value"`
	Type      int    `json:"Type"`
	ValueHash string `json:"ValueHash"`
}

// downloadCertificateResponse holds a raw string containing a Base64 encoded certificate.
type downloadCertificateResponse struct {
	Content string `json:"Content"`
}
