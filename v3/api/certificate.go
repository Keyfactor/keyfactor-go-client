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
	if ea.Template == "" && ea.EnrollmentPatternId == 0 {
		missingFields = append(missingFields, "Template or EnrollmentPatternId")
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
			log.Println("[DEBUG] EnrollPFXV2: Subject is nil checks if there are SANs")
			if ea.SANs == nil || (len(ea.SANs.DNS) == 0 && len(ea.SANs.URI) == 0 && len(ea.SANs.IP4) == 0 &&
				len(ea.SANs.IP6) == 0) {
				return nil, fmt.Errorf("subject or subject alternative names are required to use enrollpfx(). Please configure either SubjectString or Subject or SANs")
			}
		}
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "Enrollment/PFX",
		Headers:  headers,
		Payload:  &ea,
	}

	log.Println("[TRACE] Request: ", keyfactorAPIStruct)

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
//   - Raw certificate data (as base64 string, if applicable)
//   - Error
func (c *Client) DownloadCertificate(
	certId int,
	thumbprint string,
	serialNumber string,
	issuerDn string,
	collectionId int,
	certificateFormat string,
) (*x509.Certificate, []*x509.Certificate, *string, error) {
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
		return nil, nil, nil, fmt.Errorf("certID, thumbprint, or serial number AND issuer DN required to dowload certificate")
	}

	payload := &downloadCertificateBody{
		CertID:       certId,
		SerialNumber: serialNumber,
		IssuerDN:     issuerDn,
		Thumbprint:   thumbprint,
		IncludeChain: true,
		ChainOrder:   "EndEntityFirst",
	}

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

	// Set Keyfactor-specific headers
	switch certificateFormat {
	case "CER", "CRT", "DER", "PEM":
		// do nothing these are valid formats
		break
	default:
		// if not specified or invalid format then default to P7B
		certificateFormat = "P7B"
	}
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"x-certificateformat", certificateFormat},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "Certificates/Download",
		Headers:  headers,
		Payload:  payload,
		Query:    &query,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, nil, nil, err
	}

	jsonResp := &downloadCertificateResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, nil, nil, err
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
		return nil, nil, &jsonResp.Content, p7bErr
	}

	leaf := findLeafCert(certs)
	if len(certs) > 1 {
		return leaf, certs, &jsonResp.Content, nil
	}
	return leaf, nil, &jsonResp.Content, nil
}

// findLeafCert returns the end-entity (leaf) certificate from a set of
// certificates. It identifies the leaf as the cert whose Subject is not used
// as an Issuer by any other cert in the set — i.e. nothing is signed by it.
// This is order-independent and handles both root-first and leaf-first P7Bs.
//
// When the set contains only one cert, or when the algorithm cannot determine
// a unique leaf (e.g. all certs are self-signed), certs[0] is returned as a
// safe fallback.
func findLeafCert(certs []*x509.Certificate) *x509.Certificate {
	if len(certs) == 0 {
		return nil
	}
	if len(certs) == 1 {
		return certs[0]
	}

	// Build a set of all RawIssuer values (subjects that issued something).
	issuers := make(map[string]bool, len(certs))
	for _, c := range certs {
		issuers[string(c.RawIssuer)] = true
	}

	// The leaf's Subject is not in the issuers set.
	for _, c := range certs {
		if !issuers[string(c.RawSubject)] {
			return c
		}
	}

	// Fallback: cannot distinguish (e.g. single self-signed cert in multi-cert set).
	return certs[0]
}

// EnrollCSR takes arguments for EnrollCSRFctArgs to enroll a passed Certificate Signing
// Request with Keyfactor. An EnrollResponse containing a signed certificate is returned upon successful
// enrollment. Required fields to complete a CSR enrollment are:
//   - CSR                  : string
//   - Template             : string  (or EnrollmentPatternId on Command v25+)
//   - CertificateAuthority : string  (optional when using a template or enrollment pattern)
func (c *Client) EnrollCSR(ea *EnrollCSRFctArgs) (*EnrollResponse, error) {
	log.Println("[INFO] Signing CSR with Keyfactor")

	/* Ensure required inputs exist.
	   On Command v25+ an EnrollmentPatternId can substitute for Template.
	   CertificateAuthority is optional when a template or enrollment pattern is provided;
	   it is only required when enrolling against a standalone CA. */
	if ea.Template == "" && ea.EnrollmentPatternId == 0 {
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
	//for _, certs := range rvargs.CertificateIds {
	//	log.Printf("[TRACE] Revoking ID %d", certs)
	//}

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
	certificateFormat string,
) (interface{}, *x509.Certificate, []*x509.Certificate, *string, error) {
	log.Println("[DEBUG] Enter RecoverCertificate")
	log.Println("[INFO] Recovering certificate ID:", certId)
	/* The download certificate endpoint requires one of the following to retrieve a cert:
		- CertID
		- Thumbprint
		- SerialNumber AND IssuerDN

	Check for this input
	*/
	if certificateFormat == "" {
		certificateFormat = "PFX"
	}
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
		return nil, nil, nil, nil, fmt.Errorf("certID, thumbprint, or serial number AND issuer DN required to download certificate")
	}
	log.Println("[DEBUG] RecoverCertificate: Valid input")

	if password == "" {
		return nil, nil, nil, nil, fmt.Errorf("password required to recover private key with certificate")
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
			{"x-certificateformat", certificateFormat},
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
		return nil, nil, nil, nil, err
	}

	jsonResp := &recoverCertResponse{}
	log.Println("[DEBUG] RecoverCertificate: Decoding response")
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		log.Println("[ERROR] RecoverCertificate: Error decoding response from Keyfactor Command", err.Error())
		return nil, nil, nil, nil, err
	}

	switch certificateFormat {
	case "PFX", "pfx", "pkcs12", "p12", "jks", "JKS":
		log.Println("[DEBUG] RecoverCertificate: decoding `PFX` response field")
		pfxDer := jsonResp.PFX
		if pfxDer == "" {
			log.Println("[ERROR] RecoverCertificate: Error decoding PFX", err.Error())
			return nil, nil, nil, &pfxDer, fmt.Errorf("pfx field in response is empty")
		}
		log.Println("[INFO] Recovered certificate successfully")
		log.Println("[DEBUG] RecoverCertificate returning in PFX format")
		return nil, nil, nil, &pfxDer, nil
	case "PEM", "pem":
		log.Println("[DEBUG] RecoverCertificate: Decoding PFX")
		pfxDer, dErr := base64.StdEncoding.DecodeString(jsonResp.PFX)
		if dErr != nil {
			log.Println("[ERROR] RecoverCertificate: Error decoding PFX", dErr.Error())
			return nil, nil, nil, &jsonResp.PFX, dErr
		}

		log.Println("[DEBUG] RecoverCertificate: Decoding PFX chain")
		priv, leaf, chain, pErr := pkcs12.DecodeChain(
			pfxDer,
			rca.Password,
		) // TODO: Attempt to parse as PKCS12 because that used to be the "default" export format.
		if pErr != nil {
			log.Println("[ERROR] RecoverCertificate: Error decoding PFX chain", pErr.Error())
			return nil, nil, nil, &jsonResp.PFX, nil //TODO: Don't return error because it's probably actually a PEM
		}

		log.Println("[INFO] Recovered certificate successfully")
		log.Println("[DEBUG] RecoverCertificate: ", leaf, chain)
		return priv, leaf, chain, &jsonResp.PFX, nil
	default:
		log.Println("[DEBUG] RecoverCertificate: Decoding PFX")
		pfxDer, dErr := base64.StdEncoding.DecodeString(jsonResp.PFX)
		if dErr != nil {
			log.Println("[ERROR] RecoverCertificate: Error decoding PFX", dErr.Error())
			return nil, nil, nil, &jsonResp.PFX, dErr
		}

		log.Println("[DEBUG] RecoverCertificate: Decoding PFX chain")
		priv, leaf, chain, pErr := pkcs12.DecodeChain(pfxDer, rca.Password)
		if pErr != nil {
			log.Println("[ERROR] RecoverCertificate: Error decoding PFX chain", pErr.Error())
			return nil, nil, nil, &jsonResp.PFX, pErr
		}

		log.Println("[INFO] Recovered certificate successfully")
		log.Println("[DEBUG] RecoverCertificate returning in PEM format")

		var pemCerts []string

		// Encode leaf certificate to PEM
		pemLeaf := pem.EncodeToMemory(
			&pem.Block{
				Type:  "CERTIFICATE",
				Bytes: leaf.Raw,
			},
		)
		pemCerts = append(pemCerts, string(pemLeaf))

		// Encode chain certificates to PEM
		for _, cert := range chain {
			pemCert := pem.EncodeToMemory(
				&pem.Block{
					Type:  "CERTIFICATE",
					Bytes: cert.Raw,
				},
			)
			pemCerts = append(pemCerts, string(pemCert))
		}

		pemData := strings.Join(pemCerts, "\n")
		return priv, leaf, chain, &pemData, nil
	}

}

// ChangeCertificateOwnerRole changes the certificate's owner. Users must be in the current owner's role and the new owner's role.
// If removing the owner, leave both NewRoleId and NewRoleName empty in the request.
// Calls PUT /Certificates/{id}/Owner endpoint.
func (c *Client) ChangeCertificateOwnerRole(
	certificateId int,
	req *OwnerRequest,
	params ...*CertificateOwnerChangeParams,
) error {
	log.Printf("[INFO] Changing owner of certificate with ID %d in Keyfactor", certificateId)

	// Validate certificate ID
	if certificateId <= 0 {
		return errors.New("certificate ID must be a positive integer")
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"Content-Type", "application/json"},
		},
	}

	// Build URL with query parameters
	endpoint := fmt.Sprintf("Certificates/%d/Owner", certificateId)
	var queryParams []string

	if len(params) > 0 && params[0] != nil {
		param := params[0]
		if param.CollectionId != nil {
			queryParams = append(queryParams, fmt.Sprintf("collectionId=%d", *param.CollectionId))
		}
		if param.ContainerId != nil {
			queryParams = append(queryParams, fmt.Sprintf("containerId=%d", *param.ContainerId))
		}
	}

	if len(queryParams) > 0 {
		endpoint += "?" + strings.Join(queryParams, "&")
	}

	keyfactorAPIStruct := &request{
		Method:   "PUT",
		Endpoint: endpoint,
		Headers:  headers,
		Payload:  req,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return err
	}

	// Check if the response indicates success (204 No Content expected)
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to change certificate owner: HTTP %d", resp.StatusCode)
	}

	return nil
}

// createSubject builds the certificate subject string from a passed CertificateSubject argument.
func createSubject(cs CertificateSubject) (string, error) {
	var subject string

	if cs.SubjectCommonName != "" && cs.SubjectCommonName != "<null>" {
		subject = "CN=" + escapeDNValue(cs.SubjectCommonName) + ","
	} else {
		return "", errors.New("build subject: common name required") // Common name is required!
	}
	if cs.SubjectOrganizationalUnit != "" && cs.SubjectOrganizationalUnit != "<null>" {
		subject += "OU=" + escapeDNValue(cs.SubjectOrganizationalUnit) + ","
	}
	if cs.SubjectOrganization != "" && cs.SubjectOrganization != "<null>" {
		subject += "O=" + escapeDNValue(cs.SubjectOrganization) + ","
	}
	if cs.SubjectLocality != "" && cs.SubjectLocality != "<null>" {
		subject += "L=" + escapeDNValue(cs.SubjectLocality) + ","
	}
	if cs.SubjectState != "" && cs.SubjectState != "<null>" {
		subject += "ST=" + escapeDNValue(cs.SubjectState) + ","
	}
	if cs.SubjectCountry != "" && cs.SubjectCountry != "<null>" {
		subject += "C=" + escapeDNValue(cs.SubjectCountry) + ","
	}
	subject = strings.TrimRight(subject, ",") // remove trailing comma
	log.Printf("[DEBUG] createSubject(): Certificate subject created: %s\n", subject)
	return subject, nil
}

// escapeDNValue ensures that a value in a DN is properly escaped if it contains special characters.
func escapeDNValue(value string) string {
	// If the value contains a comma, quote it
	if strings.Contains(value, ",") {
		return `"` + value + `"`
	}
	return value
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
