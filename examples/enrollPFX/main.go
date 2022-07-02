package main

import (
	"fmt"
	"github.com/Keyfactor/keyfactor-go-client/pkg/keyfactor"
	"log"
)

var HOSTNAME = "example.com"
var USERNAME = "username"
var PASSWORD = "password"

func main() {
	// Create a new Keyfactor client
	clientConfig := &keyfactor.AuthConfig{
		Hostname: HOSTNAME,
		Username: USERNAME,
		Password: PASSWORD,
	}
	client, err := keyfactor.NewKeyfactorClient(clientConfig)
	if err != nil {
		log.Fatal(err)
	}

	// ---------------------------------------------------------------------------------------------------------------//
	// PFX Certificate Enrollment
	//
	// To enroll a PFX certificate with the Keyfactor Go client,
	// first create a pointer to an EnrollPFXFctArgs struct,
	// and populate the required fields. The below fields are the bare minimum
	PFXArgs := &keyfactor.EnrollPFXFctArgs{
		Subject: &keyfactor.CertificateSubject{
			SubjectCommonName: "keyfactor-go-example",
		},
		Password:             "#SuperSecurePassword55!",
		CertificateAuthority: "<Keyfactor certificate authority>",
		Template:             "<Keyfactor certificate template>",
	}
	// Then, call the Enroll PFX certificate function with the request arguments.
	enrollResponse, err := client.EnrollPFX(PFXArgs)
	if err != nil {
		return
	}

	// Enrolling a PFX certificate returns a pointer to an EnrollPFXResponse struct.

	fmt.Printf("%#v", enrollResponse.CertificateInformation)

	// ---------------------------------------------------------------------------------------------------------------//
	// Get Certificate Context
	//
	// To retrieve the context of a certificate stored by Keyfactor, first create a pointer to a
	// GetCertificateContextArgs struct and populate the required fields.
	getCertContextArgs := &keyfactor.GetCertificateContextArgs{
		Id:               enrollResponse.CertificateInformation.KeyfactorID, // Just for fun, get context of previously created certificate
		IncludeMetadata:  boolToPointer(true),
		IncludeLocations: boolToPointer(true),
	}
	// Then, call the get certificate context method with the request arguments.
	context, err := client.GetCertificateContext(getCertContextArgs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", context)

	// ---------------------------------------------------------------------------------------------------------------//
	// Download Certificate
	//
	// To download a persisted certificate from Keyfactor, call the
	// download certificate method with the requested arguments.
	leaf, chain, err := client.DownloadCertificate(enrollResponse.CertificateInformation.KeyfactorID, "", "", "")
	if err != nil {
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(leaf, chain)

	// ---------------------------------------------------------------------------------------------------------------//
	// Update Certificate metadata
	//
	// To update the metadata associated with a certificate in Keyfactor, first create a pointer to a
	// UpdateMetadataArgs struct and populate the required fields.
	metadata := &keyfactor.UpdateMetadataArgs{
		CertID: 1860,
		CertificateMetadata: []keyfactor.StringTuple{
			{"Department", "IT"},
			{"Email-Contact", "email@example.com"},
		},
	}
	// Then, call the update certificate metadata method with the request arguments.
	err = client.UpdateMetadata(metadata)
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
