package keyfactor

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type SANs struct {
	IP4 []string `json:"ip4,omitempty"`
	IP6 []string `json:"ip6,omitempty"`
	DNS []string `json:"dns,omitempty"`
	URI []string `json:"uri,omitempty"`
}

type EnrollPFXFctArgs struct {
	CustomFriendlyName          string
	KeyPassword                 string
	PopulateMissingValuesFromAD bool
	CertificateAuthority        string
	Template                    string
	IncludeChain                bool
	RenewalCertificateId        int
	CertFormat                  string
	CertificateSubject          CertificateSubject
	CertificateMetadata         []StringTuple
	CertificateSANs             *SANs
}

type EnrollCSRFctArgs struct {
	CSR                  string
	CertificateAuthority string
	Template             string
	IncludeChain         bool
	CertFormat           string
	CertificateMetadata  []StringTuple
	CertificateSANs      *SANs
}

type RevokeCertArgs struct {
	CertificateIds []int
	Reason         int
	Comment        string
	CollectionId   int
}

type DownloadCertArgs struct {
	CertID       int
	SerialNumber string
	IssuerDN     string
	Thumbprint   string
	IncludeChain bool
	CollectionId int
	CertFormat   string
}

type UpdateMetadataArgs struct {
	CertID              int                    `json:"Id"`
	CertificateMetadata []StringTuple          `json:"-"`
	Metadata            map[string]interface{} `json:"Metadata"`
	CollectionId        int                    `json:"CollectionId"`
}

type CertificateSubject struct {
	SubjectCommonName         string
	SubjectLocality           string
	SubjectOrganization       string
	SubjectCountry            string
	SubjectOrganizationalUnit string
	SubjectState              string
}

type enrollPFXBody struct {
	CustomFriendlyName          string                 `json:"CustomFriendlyName,omitempty"`
	Password                    string                 `json:"Password"`
	PopulateMissingValuesFromAD bool                   `json:"PopulateMissingValuesFromAD"`
	Subject                     string                 `json:"Subject"`
	IncludeChain                bool                   `json:"IncludeChain"`
	RenewalCertificateId        int                    `json:"RenewalCertificateId,omitempty"`
	CertificateAuthority        string                 `json:"CertificateAuthority"`
	Timestamp                   string                 `json:"Timestamp"`
	Template                    string                 `json:"Template"`
	SANs                        *SANs                  `json:"SANs"`
	Metadata                    map[string]interface{} `json:"Metadata"`
}

type enrollCSRBody struct {
	CSR                  string
	Timestamp            string                 `json:"Timestamp"`
	Template             string                 `json:"Template"`
	CertificateAuthority string                 `json:"CertificateAuthority"`
	IncludeChain         bool                   `json:"IncludeChain"`
	SANs                 *SANs                  `json:"SANs"`
	Metadata             map[string]interface{} `json:"Metadata"`
}

type downloadCertificateBody struct {
	CertID       int
	SerialNumber string
	IssuerDN     string
	Thumbprint   string
	IncludeChain bool
}

type revokeCertBody struct {
	CertificateIds []int  `json:"CertificateIds"`
	Reason         int    `json:"Reason"`
	Comment        string `json:"Comment"`
	EffectiveDate  string `json:"EffectiveDate"`
	CollectionId   int    `json:"CollectionId,omitempty"`
}

type EnrollResponse struct {
	Certificates           []string
	CertificateInformation CertificateInformation `json:"CertificateInformation"`
}

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

type DownloadCertificateResponse struct {
	Content string `json:"Content"`
}

func EnrollPFX(ea *EnrollPFXFctArgs, api *APIClient) (*EnrollResponse, error) {
	log.Println("[INFO] Enrolling PFX certificate with Keyfactor")

	/* Ensure required inputs exist */
	if (ea.Template == "") || (ea.KeyPassword == "") || (ea.CertificateAuthority == "") {
		return nil, errors.New("invalid or nonexistent values required for pfx enrollment")
	}

	/* Build subject */
	subject, err := createSubject(ea.CertificateSubject)
	if err != nil {
		return nil, err
	}

	// Once required fields are found, create request body
	log.Println("[TRACE] Creating request body")
	payload := &enrollPFXBody{
		CustomFriendlyName:          ea.CustomFriendlyName,
		Password:                    ea.KeyPassword,
		PopulateMissingValuesFromAD: ea.PopulateMissingValuesFromAD, // this field has default value
		Subject:                     subject,
		IncludeChain:                ea.IncludeChain, // IncludeChain has default value
		CertificateAuthority:        ea.CertificateAuthority,
		Timestamp:                   getTimestamp(),
		Template:                    ea.Template,
		SANs:                        ea.CertificateSANs,
		Metadata:                    mapTupleArrayToInterface(ea.CertificateMetadata),
	}

	// Set Keyfactor-specific headers
	log.Println("[TRACE] Setting request headers")
	headers := &APIHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"x-certificateformat", ea.CertFormat},
		},
	}

	log.Println("[TRACE] Creating request struct for Keyfactor client")
	keyfactorAPIStruct := &APIRequest{
		KFClient: api,
		Method:   "POST",
		Endpoint: "/KeyfactorAPI/Enrollment/PFX",
		Headers:  headers,
		Payload:  &payload,
	}

	resp, err := SendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		log.Printf("[DEBUG] POST succeeded with response code %d", resp.StatusCode)
		jsonResp := &EnrollResponse{}
		err = json.NewDecoder(resp.Body).Decode(&jsonResp)
		if err != nil {
			return nil, err
		}
		err = decodePKCS12Blob(jsonResp)
		if err != nil {
			return nil, err
		}
		return jsonResp, nil
	}

	var errorMessage interface{} // Decode JSON body to handle issue
	err = json.NewDecoder(resp.Body).Decode(&errorMessage)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Request failed with code %d and message %v", resp.StatusCode, errorMessage)
	stringMessage := fmt.Sprintf("%v", errorMessage)
	return nil, errors.New(stringMessage)
}

func DownloadCertificate(da *DownloadCertArgs, api *APIClient) (*DownloadCertificateResponse, error) {
	log.Println("[INFO] Downloading certificate")

	/* The download certificate endpoint requires one of the following to retrieve a cert:
		- CertID
		- Thumbprint
		- SerialNumber AND IssuerDN

	Check for this input
	*/
	var validInput = false
	if da.CertID != 0 {
		validInput = true
	} else if da.Thumbprint != "" {
		validInput = true
	} else if (da.SerialNumber != "") && (da.IssuerDN != "") {
		validInput = true
	}
	if validInput != true {
		return nil, errors.New("invalid input received for cert download request")
	}

	log.Println("[TRACE] Creating request body")

	payload := &downloadCertificateBody{
		CertID:       da.CertID,
		SerialNumber: da.SerialNumber,
		IssuerDN:     da.IssuerDN,
		Thumbprint:   da.Thumbprint,
		IncludeChain: da.IncludeChain,
	}

	// Set Keyfactor-specific headers
	log.Println("[TRACE] Setting request headers")
	headers := &APIHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"x-certificateformat", da.CertFormat},
		},
	}

	log.Println("[TRACE] Creating request struct for Keyfactor client")
	keyfactorAPIStruct := &APIRequest{
		KFClient: api,
		Method:   "POST",
		Endpoint: "/KeyfactorAPI/Certificates/Download",
		Headers:  headers,
		Payload:  &payload,
	}

	resp, err := SendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		log.Printf("[DEBUG] POST succeeded with response code %d", resp.StatusCode)
		jsonResp := &DownloadCertificateResponse{}
		err = json.NewDecoder(resp.Body).Decode(&jsonResp)
		if err != nil {
			return nil, err
		}
		content, err := base64.StdEncoding.DecodeString(jsonResp.Content)
		if err != nil {
			return nil, err
		}
		jsonResp.Content = string(content)

		return jsonResp, nil
	}

	var errorMessage interface{} // Decode JSON body to handle issue
	err = json.NewDecoder(resp.Body).Decode(&errorMessage)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] Request failed with code %d and message %v", resp.StatusCode, errorMessage)
	stringMessage := fmt.Sprintf("%v", errorMessage)

	return nil, errors.New(stringMessage)
}

func EnrollCSR(ea *EnrollCSRFctArgs, api *APIClient) (*EnrollResponse, error) {
	log.Println("[INFO] Signing CSR with Keyfactor")

	/* Ensure required inputs exist */
	if (ea.Template == "") || (ea.CertificateAuthority == "") {
		return nil, errors.New("invalid or nonexistent values required for csr enrollment")
	}

	log.Println("[TRACE] Creating request body")
	payload := &enrollCSRBody{
		CSR:                  ea.CSR,
		Timestamp:            getTimestamp(),
		Template:             ea.Template,
		CertificateAuthority: ea.CertificateAuthority,
		IncludeChain:         ea.IncludeChain,
		SANs:                 ea.CertificateSANs,
		Metadata:             mapTupleArrayToInterface(ea.CertificateMetadata),
	}

	// Set Keyfactor-specific headers
	log.Println("[TRACE] Setting request headers")
	headers := &APIHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"x-certificateformat", ea.CertFormat},
		},
	}

	log.Println("[TRACE] Creating request struct for Keyfactor client")
	keyfactorAPIStruct := &APIRequest{
		KFClient: api,
		Method:   "POST",
		Endpoint: "/KeyfactorAPI/Enrollment/CSR",
		Headers:  headers,
		Payload:  &payload,
	}

	resp, err := SendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		log.Printf("[DEBUG] POST succeeded with response code %d", resp.StatusCode)
		jsonResp := &EnrollResponse{}
		err = json.NewDecoder(resp.Body).Decode(&jsonResp)
		if err != nil {
			return nil, err
		}
		jsonResp.Certificates = jsonResp.CertificateInformation.Certificates
		return jsonResp, nil
	}

	var errorMessage interface{} // Decode JSON body to handle issue
	err = json.NewDecoder(resp.Body).Decode(&errorMessage)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] Request failed with code %d and message %v", resp.StatusCode, errorMessage)
	stringMessage := fmt.Sprintf("%v", errorMessage)

	return nil, errors.New(stringMessage)
}

func UpdateMetadata(um *UpdateMetadataArgs, api *APIClient) error {
	// Metadata in Keyfactor varies between deployments
	// Instead of hard coding possibilities, take array of string tuple types and create
	// string-indexed array of interfaces for JSON compilation
	um.Metadata = mapTupleArrayToInterface(um.CertificateMetadata)

	// Set Keyfactor-specific headers
	log.Println("[TRACE] Setting request headers")
	headers := &APIHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	log.Println("[TRACE] Creating request struct for Keyfactor client")
	keyfactorAPIStruct := &APIRequest{
		KFClient: api,
		Method:   "PUT",
		Endpoint: "/KeyfactorAPI/Certificates/Metadata",
		Headers:  headers,
		Payload:  um,
	}

	resp, err := SendRequest(keyfactorAPIStruct)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNoContent {
		log.Printf("[DEBUG] PUT succeeded with response code %d", resp.StatusCode)
		return nil
	}

	var errorMessage interface{} // Decode JSON body to handle issue
	err = json.NewDecoder(resp.Body).Decode(&errorMessage)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Request failed with code %d and message %v", resp.StatusCode, errorMessage)
	stringMessage := fmt.Sprintf("%v", errorMessage)
	return errors.New(stringMessage)
}

func mapTupleArrayToInterface(i []StringTuple) map[string]interface{} {
	temp := make(map[string]interface{}, len(i)) // Create string-index-able interface array from tuple struct
	for _, field := range i {
		temp[field.Elem1] = field.Elem2
	}
	return temp
}

func RevokeCert(ra *RevokeCertArgs, api *APIClient) error {
	log.Println("[INFO] Revoking certificates")
	for _, certs := range ra.CertificateIds {
		log.Printf("[TRACE] Revoking ID %d", certs)
	}

	// Fields required by revoke cert API request are cert ID & comment
	// Go initializes integers to 0, check for zero input
	if (ra.CertificateIds[0] == 0) && (ra.Comment == "") {
		return errors.New("invalid or nonexistent values required for certificate revocation")
	}

	log.Println("[TRACE] Creating request body")
	payload := &revokeCertBody{
		CertificateIds: ra.CertificateIds,
		Comment:        ra.Comment,
		Reason:         ra.Reason,
		EffectiveDate:  getTimestamp(),
	}
	if ra.CollectionId != 0 {
		payload.CollectionId = ra.CollectionId
	}

	// Set Keyfactor-specific headers
	log.Println("[TRACE] Setting request headers")
	headers := &APIHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	log.Println("[TRACE] Creating request struct for Keyfactor client")
	keyfactorAPIStruct := &APIRequest{
		KFClient: api,
		Method:   "POST",
		Endpoint: "/KeyfactorAPI/Certificates/Revoke",
		Headers:  headers,
		Payload:  &payload,
	}

	resp, err := SendRequest(keyfactorAPIStruct)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNoContent {
		log.Printf("[DEBUG] POST succeeded with response code %d", resp.StatusCode)
		return nil
	}

	var errorMessage interface{} // Decode JSON body to handle issue
	err = json.NewDecoder(resp.Body).Decode(&errorMessage)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Request failed with code %d and message %v", resp.StatusCode, errorMessage)
	stringMessage := fmt.Sprintf("%v", errorMessage)
	return errors.New(stringMessage)
}

func createSubject(cs CertificateSubject) (string, error) {
	var subject string

	if cs.SubjectCommonName != "" {
		subject = "CN=" + cs.SubjectCommonName + ","
	} else {
		return "", errors.New("build subject: common name required") // Common name is required!
	}

	if cs.SubjectOrganizationalUnit != "" {
		subject += "OU=" + cs.SubjectOrganizationalUnit + ","
	}

	if cs.SubjectOrganization != "" {
		subject += "O=" + cs.SubjectOrganization + ","
	}

	if cs.SubjectLocality != "" {
		subject += "L=" + cs.SubjectLocality + ","
	}

	if cs.SubjectState != "" {
		subject += "ST=" + cs.SubjectState + ","
	}

	if cs.SubjectCountry != "" {
		subject += "C=" + cs.SubjectCountry + ","
	}

	subject = strings.TrimRight(subject, ",") // remove trailing comma

	log.Printf("[DEBUG] createSubject(): Certificate subject created: %s\n", subject)

	return subject, nil
}

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
