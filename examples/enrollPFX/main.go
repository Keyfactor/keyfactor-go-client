package main

import (
	"fmt"
	"keyfactor-go-client/pkg/keyfactor"
)

func main() {
	demoPFXEnrollment()
}

func demoPFXEnrollment() {
	// The base functionality of the Go client revolves around a credentials struct.
	// Create a pointer to this struct and populate it with example values

	clientConfig := &keyfactor.APIClient{
		Hostname: "keyfactor.example.com",
		Username: "username",
		Password: "password",
	}

	// To enroll a PFX certificate with the Keyfactor Go client,
	// first create a pointer to an Enroll PFX function Argument struct,
	// and populate the required fields. The below fields are the bare minimum

	PFXArgs := &keyfactor.EnrollPFXFctArgs{
		CertificateSubject: keyfactor.CertificateSubject{
			SubjectCommonName: "keyfactor-go-example",
		},
		KeyPassword:          "SuperSecurePassword55!",
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

	fmt.Printf("New certificate serial number: %s", response.CertificateInformation.SerialNumber)
}
