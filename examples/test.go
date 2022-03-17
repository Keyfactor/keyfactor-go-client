package main

import (
	"github.com/Keyfactor/keyfactor-go-client/pkg/keyfactor"
	"log"
)

var HOSTNAME = "sedemo.thedemodrive.com"
var USERNAME = "HRoszell"
var PASSWORD = "Ferrari1!"

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
	// 6175d9f2-b7e4-40a2-a3c3-9acb91cdeae5
	//store := "6175d9f2-b7e4-40a2-a3c3-9acb91cdeae5"
	store := "db21f6da-fec5-418b-a093-3c6eb0a5970c"
	_, err = client.GetCertificateStoreByID(store)
	if err != nil {
		return
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
