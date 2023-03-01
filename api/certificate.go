package api

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spbsoluble/go-pkcs12"
	"github.com/spbsoluble/keyfactor-go-client-sdk"
	"go.mozilla.org/pkcs7"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// EnrollPFX takes arguments for EnrollPFXFctArgs to facilitate a call to Keyfactor
// that enrolls a PFX certificate with the supplied arguments.
func (c *Client) EnrollPFX(ea *EnrollPFXFctArgs) (*EnrollResponse, error) {
	log.Println("[INFO] Enrolling PFX certificate with Keyfactor")

	/* Ensure required inputs exist */
	var missingFields []string

	// TODO: Probably a better way to express these if blocks
	if ea.Template == "" {
		missingFields = append(missingFields, "Template")
	}
	if ea.CertificateAuthority == "" {
		missingFields = append(missingFields, "CertificateAuthority")
	}
	if ea.CertFormat == "" {
		missingFields = append(missingFields, "CertFormat")
	}

	if len(missingFields) > 0 {
		return nil, errors.New("Required field(s) missing: " + strings.Join(missingFields, ", "))
	}

	if ea.Timestamp == "" {
		ea.Timestamp = getTimestamp()
	}

	if ea.SubjectString == "" {
		if ea.Subject != nil {
			subject, err := createSubject(*ea.Subject)
			if err != nil {
				return nil, err
			}
			ea.SubjectString = subject
		} else {
			return nil, fmt.Errorf("subject is required to use enrollpfx(). Please configure either SubjectString or Subject")
		}
	}

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"
	xCertificateFormat := ea.CertFormat

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	newRenewalCertId := int32(ea.RenewalCertificateId)
	newTimestamp, err := time.Parse(ea.Timestamp, ea.Timestamp)

	var newSANs map[string][]string
	data, _ := json.Marshal(ea.SANs)
	json.Unmarshal(data, &newSANs)

	req := keyfactor_command_client_api.ModelsEnrollmentPFXEnrollmentRequest{
		CustomFriendlyName:          &ea.CustomFriendlyName,
		Password:                    &ea.Password,
		PopulateMissingValuesFromAD: &ea.PopulateMissingValuesFromAD,
		Subject:                     &ea.SubjectString, //TODO
		IncludeChain:                &ea.IncludeChain,
		RenewalCertificateId:        &newRenewalCertId,
		CertificateAuthority:        &ea.CertificateAuthority,
		Metadata:                    ea.Metadata,
		AdditionalEnrollmentFields:  nil,
		Timestamp:                   &newTimestamp,
		Template:                    &ea.Template,
		SANs:                        &newSANs,
	}

	resp, _, err := apiClient.EnrollmentApi.EnrollmentPostPFXEnroll(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XCertificateformat(xCertificateFormat).XKeyfactorApiVersion(xKeyfactorApiVersion).Request(req).Execute()

	//newIssuerDN := resp.CertificateInformation.IssuerDN.Get()

	mapResp, _ := resp.ToMap()
	jsonData, _ := json.Marshal(mapResp)
	var newResp EnrollResponse
	json.Unmarshal(jsonData, &newResp)

	//newResp := EnrollResponse{
	//	Certificates: nil, //TODO
	//	CertificateInformation: CertificateInformation{
	//		SerialNumber:       *resp.CertificateInformation.SerialNumber,
	//		IssuerDN:           *newIssuerDN,
	//		Thumbprint:         *resp.CertificateInformation.Thumbprint,
	//		KeyfactorID:        int(*resp.CertificateInformation.KeyfactorId),
	//		KeyfactorRequestID: int(*resp.CertificateInformation.KeyfactorRequestId),
	//		PKCS12Blob:         *resp.CertificateInformation.Pkcs12Blob,
	//		Certificates:       nil, //TODO
	//		RequestDisposition: *resp.CertificateInformation.RequestDisposition,
	//		DispositionMessage: *resp.CertificateInformation.DispositionMessage,
	//		EnrollmentContext:  &resp.CertificateInformation.EnrollmentContext,
	//	},
	//}

	if err != nil {
		return nil, err
	}

	return &newResp, nil
}

// DownloadCertificate takes arguments for DownloadCertArgs to facilitate a call to Keyfactor
// that downloads a certificate from Keyfactor.
// The download certificate endpoint requires one of the following to retrieve a cert:
//   - CertID
//   - Thumbprint
//   - SerialNumber AND IssuerDN
//
// Returns:
//   - Leaf certificate
//   - Certificate chain
func (c *Client) DownloadCertificate(certId int, thumbprint string, serialNumber string, issuerDn string) (*x509.Certificate, []*x509.Certificate, error) {
	log.Println("[INFO] Downloading certificate")

	/* The download certificate endpoint requires one of the following to retrieve a cert:
		- CertID
		- Thumbprint
		- SerialNumber AND IssuerDN

	Check for this input
	*/
	validInput := false
	if certId != 0 {
		validInput = true
	} else if thumbprint != "" {
		validInput = true
	} else if serialNumber != "" && issuerDn != "" {
		validInput = true
	}

	if !validInput {
		return nil, nil, fmt.Errorf("certID, thumbprint, or serial number AND issuer DN required to dowload certificate")
	}

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	newCertId := int32(certId)
	newIssuerDN := keyfactor_command_client_api.NullableString{}
	newIssuerDN.Set(&issuerDn)

	rq := keyfactor_command_client_api.ModelsCertificateDownloadRequest{
		CertID:       &newCertId,
		SerialNumber: &serialNumber,
		IssuerDN:     newIssuerDN,
		Thumbprint:   &thumbprint,
		IncludeChain: nil,
	}

	resp, _, err := apiClient.CertificateApi.CertificateDownloadCertificateAsync(context.Background()).Rq(rq).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	mapResp, _ := resp.ToMap()
	jsonData, _ := json.Marshal(mapResp)
	var newResp downloadCertificateResponse
	json.Unmarshal(jsonData, &newResp)

	if err != nil {
		return nil, nil, err
	}
	buf, err := base64.StdEncoding.DecodeString(newResp.Content)
	if err != nil {
		return nil, nil, err
	}

	certs, err := pkcs7.Parse(buf)
	if err != nil {
		return nil, nil, err
	}

	//todo: review this as it seems to be returning the wrong cert
	leaf := certs.Certificates[1]

	if len(certs.Certificates) > 1 {
		return leaf, certs.Certificates, nil
	}

	return leaf, nil, nil
}

// EnrollCSR takes arguments for EnrollCSRFctArgs to enroll a passed Certificate Signing
// Request with Keyfactor. An EnrollResponse containing a signed certificate is returned upon successful
// enrollment. Required fields to complete a CSR enrollment are:
//   - CSR                  : string
//   - Template             : string
//   - CertificateAuthority : string
func (c *Client) EnrollCSR(ea *EnrollCSRFctArgs) (*EnrollResponse, error) {
	log.Println("[INFO] Signing CSR with Keyfactor")

	/* Ensure required inputs exist */
	if (ea.Template == "") || (ea.CertificateAuthority == "") {
		return nil, errors.New("invalid or nonexistent values required for csr enrollment")
	}

	if ea.Timestamp == "" {
		ea.Timestamp = getTimestamp()
	}

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"
	xCertificateFormat := ea.CertFormat

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	eaJson, _ := json.Marshal(ea)
	var req keyfactor_command_client_api.ModelsEnrollmentCSREnrollmentRequest
	json.Unmarshal(eaJson, &req)

	resp, _, err := apiClient.EnrollmentApi.EnrollmentPostCSREnroll(context.Background()).XCertificateformat(xCertificateFormat).Request(req).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	mapResp, _ := resp.ToMap()
	jsonData, _ := json.Marshal(mapResp)
	var newResp EnrollResponse
	json.Unmarshal(jsonData, &newResp)

	if err != nil {
		return nil, err
	}

	newResp.Certificates = newResp.CertificateInformation.Certificates

	return &newResp, nil
}

// RevokeCert takes arguments for RevokeCertArgs to facilitate the revocation of
// all specified certificate IDs. It returns nil upon successful revocation, and an error if not.
// Required fields to revoke a list of certificates in Keyfactor are:
//   - CertificateIds : []int
//   - Comment        : string
func (c *Client) RevokeCert(ra *RevokeCertArgs) error {
	log.Println("[INFO] Revoking certificates")
	for _, certs := range ra.CertificateIds {
		log.Printf("[TRACE] Revoking ID %d", certs)
	}

	// Fields required by revoke cert API request are cert ID & comment
	// Go initializes integers to 0, check for zero input
	if (ra.CertificateIds[0] == 0) && (ra.Comment == "") {
		return errors.New("invalid or nonexistent values required for certificate revocation")
	}

	if ra.EffectiveDate != "" {
		ra.EffectiveDate = getTimestamp()
	}

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	raJson, _ := json.Marshal(ra)
	var req keyfactor_command_client_api.ModelsRevokeCertificateRequest
	json.Unmarshal(raJson, &req)

	_, httpResp, err := apiClient.CertificateApi.CertificateRevoke(context.Background()).Request(req).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return err
	}

	if httpResp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("[ERROR] Something unexpected happened, POST call to /Certificates/Revoke returned status %d", httpResp.StatusCode)
	}
	return nil
}

// DeployPFXCertificate takes pointers to DeployPFXArgs structs holding
// configuration data required for the deployment of a newly enrolled PFX certificate.
// It returns a pointer to a DeployPFXResp struct if successful, and an error message
// if not. Required fields to deploy a certificate to a store maintained by Keyfactor are:
//   - StoreIds      : []string
//   - Password      : string
//   - CertificateId : int
//   - RequestId     : int
func (c *Client) DeployPFXCertificate(args *DeployPFXArgs) (*DeployPFXResp, error) {
	err := validateDeployPFXArgs(args)
	if err != nil {
		return nil, err
	}

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	argsJson, _ := json.Marshal(args)
	var req keyfactor_command_client_api.KeyfactorApiModelsEnrollmentEnrollmentManagementRequest
	json.Unmarshal(argsJson, &req)

	resp, _, err := apiClient.EnrollmentApi.EnrollmentInstallPFXToCertStore(context.Background()).Request(req).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	mapResp, _ := resp.ToMap()
	jsonData, _ := json.Marshal(mapResp)
	var newResp DeployPFXResp
	json.Unmarshal(jsonData, &newResp)

	return &newResp, nil
}

// GetCertificateContext takes arguments for GetCertificateContextArgs used to facilitate the retrieval
// of certificate context. The primary query required to get certificate context is the certificate ID. Include metadata
// and include locations add additional data, but can be set to false if they are unneeded. A pointer to a
// GetCertificateResponse structure is returned, containing the certificate context.

// TODO: change to allow acception of Thumbprint in place of ID
func (c *Client) GetCertificateContext(gca *GetCertificateContextArgs) (*GetCertificateResponse, error) {

	if gca.Id == 0 {
		return nil, errors.New("keyfactor certificate id is required to get certificate")
	}

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	resp, _, err := apiClient.CertificateApi.CertificateGetCertificate(context.Background(), int32(gca.Id)).IncludeLocations(*gca.IncludeLocations).IncludeMetadata(*gca.IncludeMetadata).CollectionId(int32(*gca.CollectionId)).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	mapResp, _ := resp.ToMap()
	jsonData, _ := json.Marshal(mapResp)
	var newResp GetCertificateResponse
	json.Unmarshal(jsonData, &newResp)

	return &newResp, err
}

func (c *Client) ListCertificates(q map[string]string) ([]GetCertificateResponse, error) {

	type query struct {
		collectionId         int32
		pqQueryString        string
		includeMetadata      bool
		includeHasPrivateKey bool
		verbose              int32
		pqPageReturned       int32
		pqReturnLimit        int32
		pqSortField          string
		pqSortAscending      int32
		pqIncludeRevoked     bool
		pqIncludeExpired     bool
	}

	qry := apiQuery{
		Query: []StringTuple{},
	}

	newQuery := query{
		collectionId:         0,
		pqQueryString:        "",
		includeMetadata:      false,
		includeHasPrivateKey: false,
		verbose:              0,
		pqPageReturned:       0,
		pqReturnLimit:        0,
		pqSortField:          "",
		pqSortAscending:      0,
		pqIncludeRevoked:     false,
		pqIncludeExpired:     false,
	}

	searchCollection, ok := q["collection"]
	if ok {
		collectionIdInt, _ := strconv.ParseInt(searchCollection, 10, 64)
		newQuery.collectionId = int32(collectionIdInt)
	}
	subjectName, ok := q["subject"]
	if ok {
		newQuery.pqQueryString = fmt.Sprintf(`IssuedCN -eq "%s"`, subjectName)
	}
	tp, tpOk := q["thumbprint"]
	if tpOk {
		qry.Query = append(qry.Query, StringTuple{
			"pq.queryString", fmt.Sprintf(`Thumbprint -eq "%s"`, tp),
		})
	}

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	resp, _, err := apiClient.CertificateApi.CertificateQueryCertificates(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(newQuery.collectionId).IncludeLocations(true).IncludeMetadata(newQuery.includeMetadata).IncludeHasPrivateKey(newQuery.includeHasPrivateKey).Verbose(newQuery.verbose).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(newQuery.pqQueryString).PqPageReturned(newQuery.pqPageReturned).PqReturnLimit(newQuery.pqReturnLimit).PqSortField(newQuery.pqSortField).PqSortAscending(newQuery.pqSortAscending).PqIncludeRevoked(newQuery.pqIncludeRevoked).PqIncludeExpired(newQuery.pqIncludeExpired).Execute()

	if err != nil {
		return nil, err
	}

	var newResp []GetCertificateResponse

	for i := range resp {
		mapResp, _ := resp[i].ToMap()
		jsonData, _ := json.Marshal(mapResp)
		var newCert GetCertificateResponse
		json.Unmarshal(jsonData, &newCert)
		newResp = append(newResp, newCert)
	}

	return newResp, err
}

// RecoverCertificate takes arguments for RecoverCertArgs to facilitate a call to Keyfactor
// that recovers a certificate and associated private key (if retained) in the specified format.
// The download certificate endpoint requires one of the following to retrieve a cert:
//   - CertID
//   - Thumbprint
//   - SerialNumber AND IssuerDN
//
// Additionally, the certificate Password is required.
// Returns:
//   - Private key (*rsa.PrivateKey or *ecdsa.PrivateKey)
//   - Leaf certificate (*x509.Certificate)
//   - Certificate chain ([]*x509.Certificate)
func (c *Client) RecoverCertificate(certId int, thumbprint string, serialNumber string, issuerDn string, password string) (interface{}, *x509.Certificate, []*x509.Certificate, error) {
	log.Println("[INFO] Recovering certificate ID:", certId)
	/* The download certificate endpoint requires one of the following to retrieve a cert:
		- CertID
		- Thumbprint
		- SerialNumber AND IssuerDN

	Check for this input
	*/
	validInput := false
	if certId != 0 {
		validInput = true
	} else if thumbprint != "" {
		validInput = true
	} else if serialNumber != "" && issuerDn != "" {
		validInput = true
	}

	if !validInput {
		return nil, nil, nil, fmt.Errorf("certID, thumbprint, or serial number AND issuer DN required to dowload certificate")
	}

	if password == "" {
		return nil, nil, nil, fmt.Errorf("password required to recover private key with certificate")
	}

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	newCertId := int32(certId)
	newIssuerDN := keyfactor_command_client_api.NullableString{}
	newIssuerDN.Set(&issuerDn)
	newIncludeChain := true

	newReq := keyfactor_command_client_api.ModelsCertificateRecoveryRequest{
		Password:     password,
		CertID:       &newCertId,
		SerialNumber: &serialNumber,
		IssuerDN:     newIssuerDN,
		Thumbprint:   &thumbprint,
		IncludeChain: &newIncludeChain,
	}

	resp, _, err := apiClient.CertificateApi.CertificateRecoverCertificateAsync(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Rq(newReq).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, nil, nil, err
	}

	mapResp, _ := resp.ToMap()
	jsonData, _ := json.Marshal(mapResp)
	var newResp recoverCertResponse
	json.Unmarshal(jsonData, &newResp)

	pfxDer, err := base64.StdEncoding.DecodeString(newResp.PFX)

	priv, leaf, chain, err := pkcs12.DecodeChain(pfxDer, newReq.Password)
	if err != nil {
		return nil, nil, nil, err
	}

	return priv, leaf, chain, nil
}

// createSubject builds the certificate subject string from a passed CertificateSubject argument.
func createSubject(cs CertificateSubject) (string, error) {
	var subject string

	if cs.SubjectCommonName != "" && cs.SubjectCommonName != "<null>" {
		subject = "CN=" + cs.SubjectCommonName + ","
	} else {
		return "", errors.New("build subject: common name required") // Common name is required!
	}
	if cs.SubjectOrganizationalUnit != "" && cs.SubjectOrganizationalUnit != "<null>" {
		subject += "OU=" + cs.SubjectOrganizationalUnit + ","
	}
	if cs.SubjectOrganization != "" && cs.SubjectOrganization != "<null>" {
		subject += "O=" + cs.SubjectOrganization + ","
	}
	if cs.SubjectLocality != "" && cs.SubjectLocality != "<null>" {
		subject += "L=" + cs.SubjectLocality + ","
	}
	if cs.SubjectState != "" && cs.SubjectState != "<null>" {
		subject += "ST=" + cs.SubjectState + ","
	}
	if cs.SubjectCountry != "" && cs.SubjectCountry != "<null>" {
		subject += "C=" + cs.SubjectCountry + ","
	}
	subject = strings.TrimRight(subject, ",") // remove trailing comma
	log.Printf("[DEBUG] createSubject(): Certificate subject created: %s\n", subject)
	return subject, nil
}

// validateDeployPFXArgs validates the arguments required to deploy a PFX certificate.
func validateDeployPFXArgs(dpfxa *DeployPFXArgs) error {
	if dpfxa.StoreIds == nil {
		return errors.New("store id required for deployment of pfx certificate")
	}
	if dpfxa.Password == "" {
		return errors.New("password required for deployment of pfx certificate")
	}
	if dpfxa.StoreTypes == nil {
		return errors.New("store type required for deployment of pfx certificate")
	}
	if dpfxa.CertificateId == 0 {
		return errors.New("certificate id required for deployment of pfx certificate")
	}
	if dpfxa.RequestId == 0 {
		return errors.New("request id required for deployment of pfx certificate")
	}
	return nil
}

// decodePKCS12Blob decodes a PKCS12 blob.
func decodePKCS12Blob(resp *EnrollResponse) error {
	log.Println("[TRACE] Decoding certificate")
	// Keyfactor returns base-64 PFX (PKCS#12) or zipped certificate. Decode here.
	if resp.CertificateInformation.PKCS12Blob != "" {
		cert, err := base64.StdEncoding.DecodeString(resp.CertificateInformation.PKCS12Blob)
		if err != nil {
			return err
		}
		temp := make([]string, 1) // Create temp 1 wide string array to hold certificate
		temp = append(temp, string(cert))
		resp.Certificates = temp
		return nil
	} else {
		resp.Certificates = nil
	}
	return nil
}

// mapTupleArrayToInterface takes an array of StringTuple structs and maps each element to
// a map[string]interface{}
func mapTupleArrayToInterface(i []StringTuple) map[string]interface{} {
	temp := make(map[string]interface{}, len(i)) // Create string-index-able interface array from tuple struct
	for _, field := range i {
		temp[field.Elem1] = field.Elem2
	}
	return temp
}
