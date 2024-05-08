package api

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/spbsoluble/go-pkcs12"
	"go.mozilla.org/pkcs7"
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
	//if ea.Password == "" {
	//	missingFields = append(missingFields, "Password")
	//}

	if len(missingFields) > 0 {
		return nil, errors.New("Required field(s) missing: " + strings.Join(missingFields, ", "))
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"x-certificateformat", ea.CertFormat},
		},
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

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "Enrollment/PFX",
		Headers:  headers,
		Payload:  &ea,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

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

func (c *Client) EnrollPFXV2(ea *EnrollPFXFctArgsV2) (*EnrollResponseV2, error) {
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
	//if ea.Password == "" {
	//	missingFields = append(missingFields, "Password")
	//}

	if len(missingFields) > 0 {
		return nil, errors.New("Required field(s) missing: " + strings.Join(missingFields, ", "))
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "2"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"x-certificateformat", ea.CertFormat},
		},
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

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "Enrollment/PFX",
		Headers:  headers,
		Payload:  &ea,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &EnrollResponseV2{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	//err = decodePKCS12Blob(jsonResp)
	//if err != nil {
	//	return nil, err
	//}
	return jsonResp, nil
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
func (c *Client) DownloadCertificate(
	certId int,
	thumbprint string,
	serialNumber string,
	issuerDn string,
) (*x509.Certificate, []*x509.Certificate, error) {
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

	payload := &downloadCertificateBody{
		CertID:       certId,
		SerialNumber: serialNumber,
		IssuerDN:     issuerDn,
		Thumbprint:   thumbprint,
		IncludeChain: true,
		ChainOrder:   "EndEntityFirst",
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"x-certificateformat", "P7B"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "Certificates/Download",
		Headers:  headers,
		Payload:  payload,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, nil, err
	}

	jsonResp := &downloadCertificateResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, nil, err
	}
	//buf, err := base64.StdEncoding.DecodeString(jsonResp.Content)
	//if err != nil {
	//	return nil, nil, err
	//}
	//
	//certs, err := pkcs7.Parse(buf)
	//if err != nil {
	//	return nil, nil, err
	//}

	certs, p7bErr := ConvertBase64P7BtoCertificates(jsonResp.Content)
	if p7bErr != nil {
		return nil, nil, p7bErr
	}

	var leaf *x509.Certificate
	if len(certs) > 1 {
		//leaf is last cert in chain
		leaf = certs[0] // First cert in chain is the leaf
		return leaf, certs, nil
	}

	return certs[0], nil, nil
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

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"x-certificateformat", ea.CertFormat},
		},
	}

	if ea.Timestamp == "" {
		ea.Timestamp = getTimestamp()
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "Enrollment/CSR",
		Headers:  headers,
		Payload:  &ea,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &EnrollResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	jsonResp.Certificates = jsonResp.CertificateInformation.Certificates
	return jsonResp, nil
}

// RevokeCert takes arguments for RevokeCertArgs to facilitate the revocation of
// all specified certificate IDs. It returns nil upon successful revocation, and an error if not.
// Required fields to revoke a list of certificates in Keyfactor are:
//   - CertificateIds : []int
//   - Comment        : string
func (c *Client) RevokeCert(rvargs *RevokeCertArgs) error {
	log.Println("[INFO] Revoking certificates")
	for _, certs := range rvargs.CertificateIds {
		log.Printf("[TRACE] Revoking ID %d", certs)
	}

	// Fields required by revoke cert API request are cert ID & comment
	// Go initializes integers to 0, check for zero input
	if (rvargs.CertificateIds[0] == 0) && (rvargs.Comment == "") {
		return errors.New("invalid or nonexistent values required for certificate revocation")
	}

	if rvargs.EffectiveDate == "" || rvargs.EffectiveDate == "{null}" || rvargs.EffectiveDate == "null" {
		rvargs.EffectiveDate = getTimestamp()
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "Certificates/Revoke",
		Headers:  headers,
		Payload:  &rvargs,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"[ERROR] Something unexpected happened, %s call to %s returned status %d",
			keyfactorAPIStruct.Method,
			keyfactorAPIStruct.Endpoint,
			resp.StatusCode,
		)
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

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "Enrollment/PFX/Deploy",
		Headers:  headers,
		Payload:  &args,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &DeployPFXResp{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, nil
}

// GetCertificateContext takes arguments for GetCertificateContextArgs used to facilitate the retrieval
// of certificate context. The primary query required to get certificate context is the certificate ID. Include metadata
// and include locations add additional data, but can be set to false if they are unneeded. A pointer to a
// GetCertificateResponse structure is returned, containing the certificate context.
func (c *Client) GetCertificateContext(gca *GetCertificateContextArgs) (*GetCertificateResponse, error) {
	if gca.Id <= 0 && gca.Thumbprint == "" && gca.CommonName == "" && gca.RequestId <= 0 {
		return nil, errors.New("keyfactor certificate id, common name, or thumbprint are required to get certificate")
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	// Construct URL query for /Certificates/{ID} requests
	query := apiQuery{
		Query: []StringTuple{},
	}
	if gca.IncludeLocations != nil || gca.CollectionId != nil || gca.IncludeMetadata != nil || gca.IncludeHasPrivateKey != nil {
		if gca.IncludeLocations != nil {
			query.Query = append(
				query.Query, StringTuple{
					"includeLocations", strconv.FormatBool(*gca.IncludeLocations),
				},
			)
		}
		if gca.IncludeMetadata != nil {
			query.Query = append(
				query.Query, StringTuple{
					"includeMetadata", strconv.FormatBool(*gca.IncludeMetadata),
				},
			)
		}
		if gca.CollectionId != nil {
			query.Query = append(
				query.Query, StringTuple{
					"collectionId", fmt.Sprintf("%d", *gca.CollectionId),
				},
			)
		}
		if gca.IncludeHasPrivateKey != nil {
			query.Query = append(
				query.Query, StringTuple{
					"includeHasPrivateKey", strconv.FormatBool(*gca.IncludeHasPrivateKey),
				},
			)
		}
	}

	var endpoint string
	if gca.Id <= 0 && gca.Thumbprint != "" {
		query.Query = append(
			query.Query, StringTuple{
				"pq.queryString", fmt.Sprintf(`Thumbprint -eq "%s"`, gca.Thumbprint),
			},
		)
		endpoint = "Certificates"
	} else if gca.Id <= 0 && gca.CommonName != "" {
		query.Query = append(
			query.Query, StringTuple{
				"pq.queryString", fmt.Sprintf(`IssuedCN -eq "%s"`, gca.CommonName),
			},
		)
		endpoint = "Certificates"
	} else if (gca.Id <= 0 && gca.CommonName == "" && gca.Thumbprint == "") && gca.RequestId > 0 {
		query.Query = append(
			query.Query, StringTuple{
				"pq.queryString", fmt.Sprintf(`CertRequestId -eq %d`, gca.RequestId),
			},
		)
		endpoint = "Certificates"
	} else {
		endpoint = "Certificates/" + fmt.Sprintf("%d", gca.Id)
	}

	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  headers,
		Query:    &query,
		Payload:  nil,
	}

	//create string of request in cURL format

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := GetCertificateResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		var lCerts []GetCertificateResponse
		lResp, _ := c.sendRequest(keyfactorAPIStruct)
		lErr := json.NewDecoder(lResp.Body).Decode(&lCerts)
		if lErr != nil {
			return nil, lErr
		}

		//Check if there are multiple certs returned if ther are iterate and return the one with the most recent ImportDate
		if len(lCerts) > 1 {
			var newestCert GetCertificateResponse
			for _, cert := range lCerts {

				if gca.RequestId > 0 && cert.CertRequestId == gca.RequestId {
					return &cert, nil
				} else if gca.Thumbprint == cert.Thumbprint {
					return &cert, nil
				}

				importDate, _ := time.Parse(time.RFC3339, cert.ImportDate)
				// Check if newestCert is empty, if it is set it to the first cert in the list
				if newestCert.ImportDate == "" {
					newestCert = cert
					continue
				}
				currentNewestImportDate, _ := time.Parse(time.RFC3339, newestCert.ImportDate)
				if importDate.After(currentNewestImportDate) {
					newestCert = cert
				}
			}
			return &newestCert, nil
		} else if len(lCerts) == 0 {
			return nil, fmt.Errorf("no certificate found")
		}
		return &lCerts[0], nil // Return first cert in list
	}
	return &jsonResp, err
}

func (c *Client) ListCertificates(q map[string]string) ([]GetCertificateResponse, error) {
	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	// Construct URL query for /Certificates/{ID} requests
	query := apiQuery{
		Query: []StringTuple{},
	}
	query.Query = append(
		query.Query, StringTuple{
			"includeLocations", "true",
		},
	)
	searchCollection, cOk := q["collection"]
	if cOk {
		query.Query = append(
			query.Query, StringTuple{
				"collectionId", searchCollection,
			},
		)
	}
	subjectName, sOk := q["subject"]
	if sOk {
		query.Query = append(
			query.Query, StringTuple{
				"pq.queryString", fmt.Sprintf(`IssuedCN -eq "%s"`, subjectName),
			},
		)
	}
	tp, tpOk := q["thumbprint"]
	if tpOk {
		query.Query = append(
			query.Query, StringTuple{
				"pq.queryString", fmt.Sprintf(`Thumbprint -eq "%s"`, tp),
			},
		)
	}

	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: "Certificates",
		Headers:  headers,
		Query:    &query,
		Payload:  nil,
	}

	cid, cidOk := q["id"]
	if cidOk {
		keyfactorAPIStruct.Endpoint = fmt.Sprintf("Certificates/%s", cid)
		keyfactorAPIStruct.Query = nil
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp []GetCertificateResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, err
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
func (c *Client) RecoverCertificate(
	certId int,
	thumbprint string,
	serialNumber string,
	issuerDn string,
	password string,
	collectionId int,
) (interface{}, *x509.Certificate, []*x509.Certificate, error) {
	log.Println("[DEBUG] Enter RecoverCertificate")
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
		log.Println("[ERROR] RecoverCertificate: certID, thumbprint, or serial number AND issuer DN required to download certificate")
		return nil, nil, nil, fmt.Errorf("certID, thumbprint, or serial number AND issuer DN required to download certificate")
	}
	log.Println("[DEBUG] RecoverCertificate: Valid input")

	if password == "" {
		return nil, nil, nil, fmt.Errorf("password required to recover private key with certificate")
	}

	rca := &recoverCertArgs{
		CertId:       certId,
		Password:     password,
		SerialNumber: serialNumber,
		IssuerDN:     issuerDn,
		Thumbprint:   thumbprint,
		IncludeChain: true,
	}

	log.Println("[DEBUG] RecoverCertificate: Recovering certificate with args:", rca)
	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"x-certificateformat", "PFX"},
		},
	}

	//if collectionId is > 0 then add to query params
	query := apiQuery{
		Query: []StringTuple{},
	}
	if collectionId > 0 {
		log.Println("[DEBUG] RecoverCertificate: Collection ID:", collectionId)
		query.Query = append(
			query.Query, StringTuple{
				"collectionId", fmt.Sprintf("%d", collectionId),
			},
		)
		log.Println("[DEBUG] RecoverCertificate: Query:", query)
	}

	log.Println("[DEBUG] RecoverCertificate: Creating recover certificate request")
	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "Certificates/Recover",
		Headers:  headers,
		Payload:  &rca,
		Query:    &query,
	}

	log.Println("[INFO] Attempting to recover certificate from Keyfactor Command")
	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		log.Println("[ERROR] RecoverCertificate: Error recovering certificate from Keyfactor Command", err.Error())
		return nil, nil, nil, err
	}

	jsonResp := &recoverCertResponse{}
	log.Println("[DEBUG] RecoverCertificate: Decoding response")
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		log.Println("[ERROR] RecoverCertificate: Error decoding response from Keyfactor Command", err.Error())
		return nil, nil, nil, err
	}

	log.Println("[DEBUG] RecoverCertificate: Decoding PFX")
	pfxDer, err := base64.StdEncoding.DecodeString(jsonResp.PFX)
	if err != nil {
		log.Println("[ERROR] RecoverCertificate: Error decoding PFX", err.Error())
		return nil, nil, nil, err
	}

	log.Println("[DEBUG] RecoverCertificate: Decoding PFX chain")
	priv, leaf, chain, err := pkcs12.DecodeChain(pfxDer, rca.Password)
	if err != nil {
		log.Println("[ERROR] RecoverCertificate: Error decoding PFX chain", err.Error())
		return nil, nil, nil, err
	}

	log.Println("[INFO] Recovered certificate successfully")
	//log.Println("[DEBUG] RecoverCertificate: ", leaf, chain)
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

// ConvertBase64P7BtoPEM takes a base64 encoded P7B certificate string and converts it to PEM format.
func ConvertBase64P7BtoPEM(base64P7B string) ([]string, error) {
	// Decode the base64 string to a byte slice.
	decodedBytes, err := base64.StdEncoding.DecodeString(base64P7B)
	if err != nil {
		return []string{}, fmt.Errorf("error decoding base64 string: %w", err)
	}

	// Parse the PKCS#7 structure.
	p7, err := pkcs7.Parse(decodedBytes)

	if err != nil {
		return []string{}, fmt.Errorf("error parsing PKCS#7 data: %w", err)
	}

	// Initialize an empty string to append the PEM encoded certificates.
	var pemEncodedCerts []string

	// Encode each certificate found in the PKCS#7 structure into PEM format.
	for _, cert := range p7.Certificates {
		pemBlock := &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: cert.Raw,
		}
		pemEncoded := pem.EncodeToMemory(pemBlock)
		pemEncodedCerts = append(pemEncodedCerts, string(pemEncoded))
	}

	return pemEncodedCerts, nil
}

// ConvertBase64P7BtoCertificates takes a base64 encoded P7B certificate string and returns a slice of *x509.Certificate.
func ConvertBase64P7BtoCertificates(base64P7B string) ([]*x509.Certificate, error) {
	// Decode the base64 string to a byte slice.
	decodedBytes, err := base64.StdEncoding.DecodeString(base64P7B)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 string: %w", err)
	}

	// Parse the PKCS#7 structure.
	p7, err := pkcs7.Parse(decodedBytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing PKCS#7 data: %w", err)
	}

	// Return the certificates.
	return p7.Certificates, nil
}
