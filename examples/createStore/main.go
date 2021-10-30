package main

import (
	"fmt"
	"keyfactor-go-client/pkg/keyfactor"
	"log"
)

var HOSTNAME = "hostname"
var USERNAME = "username"
var PASSWORD = "password"

func main() {
	getCertStoreTypeInfo()
}

func getCertStoreTypeInfo() {
	clientConfig := &keyfactor.APIClient{
		Hostname: HOSTNAME,
		Username: USERNAME,
		Password: PASSWORD,
	}

	response, err := keyfactor.GetCertStoreType(2, clientConfig)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Short Name: %s", response.ShortName)
}

func demoCreateStore() {
	clientConfig := &keyfactor.APIClient{
		Hostname: HOSTNAME,
		Username: USERNAME,
		Password: PASSWORD,
	}

	args := &keyfactor.CreateStoreFctArgs{
		ClientMachine: "machine",
	}

	resp, err := keyfactor.CreateStore(args, clientConfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("bruh: %s", resp.ClientMachine)
}
