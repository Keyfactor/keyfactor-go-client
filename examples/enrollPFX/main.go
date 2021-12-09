package main

import (
	"fmt"
	"keyfactor-go-client/pkg/keyfactor"
	"log"
)

var HOSTNAME = "example.com"
var USERNAME = "username"
var PASSWORD = "password"

func main() {
	demoPFXEnrollment()
	demoDownloadCert()
	demoUpdateMetadata()
	demoGetCertificateContext()
}

func demoPFXEnrollment() {
	// Create a pointer to a credentials struct and populate it with example values
	clientConfig := &keyfactor.APIClient{
		Hostname: HOSTNAME,
		Username: USERNAME,
		Password: PASSWORD,
	}

	// To enroll a PFX certificate with the Keyfactor Go client,
	// first create a pointer to an Enroll PFX function Argument struct,
	// and populate the required fields. The below fields are the bare minimum
	PFXArgs := &keyfactor.EnrollPFXFctArgs{
		CertificateSubject: keyfactor.CertificateSubject{
			SubjectCommonName: "keyfactor-go-example",
		},
		KeyPassword:          "#SuperSecurePassword55!",
		CertificateAuthority: "<Keyfactor certificate authority>",
		Template:             "<Keyfactor certificate template>",
	}

	// Then, call the Enroll PFX certificate function with the authentication and request
	// arguments.
	response, err := keyfactor.EnrollPFX(PFXArgs, clientConfig)
	if err != nil {
		return
	}

	// Enrolling a PFX certificate returns a pointer to an EnrollPFXResponse struct.

	// By default, this method returns a PKCS#12 certificate. This is ugly to print,
	// so we'll print the new certificate's serial number.

	fmt.Printf("%#v", response.CertificateInformation)
	fmt.Printf("CertId: %d", response.CertificateInformation.KeyfactorID)
	fmt.Printf("RequestId: %d", response.CertificateInformation.KeyfactorRequestID)
}

func demoGetCertificateContext() {
	// Step 1: Create authentication structure
	clientConfig := &keyfactor.APIClient{
		Hostname: HOSTNAME,
		Username: USERNAME,
		Password: PASSWORD,
	}
	// Step 2: Create arguments structure
	getCertContextArgs := &keyfactor.GetCertificateContextArgs{
		Id:               2140,
		IncludeMetadata:  boolToPointer(true),
		IncludeLocations: boolToPointer(true),
	}
	// Step 3: Call associated function
	response, err := keyfactor.GetCertificateContext(getCertContextArgs, clientConfig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", response)
}

func demoDownloadCert() {
	// Step 1: Create authentication structure
	clientConfig := &keyfactor.APIClient{
		Hostname: HOSTNAME,
		Username: USERNAME,
		Password: PASSWORD,
	}
	// Step 2: Create arguments structure
	downloadArgs := &keyfactor.DownloadCertArgs{
		CertID:       1860,
		IncludeChain: true,
		CertFormat:   "PEM",
	}
	// Step 3: Call associated function
	resp, err := keyfactor.DownloadCertificate(downloadArgs, clientConfig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("bruh: %s", resp.Content)
}

func demoUpdateMetadata() {
	// Step 1: Create authentication structure
	clientConfig := &keyfactor.APIClient{
		Hostname: HOSTNAME,
		Username: USERNAME,
		Password: PASSWORD,
	}
	// Step 2: Create arguments structure
	metadata := &keyfactor.UpdateMetadataArgs{
		CertID: 1860,
		CertificateMetadata: []keyfactor.StringTuple{
			{"Department", "IT"},
			{"Email-Contact", "email@example.com"},
		},
	}
	// Step 3: Call associated function
	err := keyfactor.UpdateMetadata(metadata, clientConfig)
	if err != nil {
		log.Fatal(err)
	}
}

// Helper functions
func boolToPointer(b bool) *bool {
	return &b
}

func intToPointer(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

func stringToPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
