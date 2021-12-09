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
	demoDeployStore()
	demoCreateStore()
	demoGetStoreById()
	getCertStoreTypeInfo()
}

func getCertStoreTypeInfo() {
	clientConfig := &keyfactor.APIClient{
		Hostname: HOSTNAME,
		Username: USERNAME,
		Password: PASSWORD,
	}

	response, err := keyfactor.GetCertStoreType(106, clientConfig)
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
		AgentId:       "keyfactor orchestrator agent ID",
		ClientMachine: "aks_demo",
		StorePath:     "https://coolvault.vault.azure.net/",
		CertStoreType: 106,
		Properties: []keyfactor.StringTuple{
			{"TenantID", "tenant"},
			{"ResourceGroupName", "resource group name"},
			{"ApplicationId", "app ID"},
			{"ClientSecret", "client secret"},
			{"SubscriptionId", "subscription ID"},
			{"APIObjectId", "object ID"},
			{"VaultName", "coolvault"},
		},
	}

	resp, err := keyfactor.CreateStore(args, clientConfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("bruh: %#v", resp)
}

func demoDeployStore() {
	clientConfig := &keyfactor.APIClient{
		Hostname: HOSTNAME,
		Username: USERNAME,
		Password: PASSWORD,
	}
	deployArgs := &keyfactor.DeployPFXArgs{
		StoreIds: []string{
			"store ID",
		},
		Password:      "password",
		CertificateId: 2136,
		RequestId:     1550,
		StoreTypes: []keyfactor.StoreTypes{
			{StoreTypeId: 106, Alias: stringToPointer("newCertificate")},
		},
	}

	resp, err := keyfactor.DeployPFXCertificate(deployArgs, clientConfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", resp)
}

func demoGetStoreById() {
	clientConfig := &keyfactor.APIClient{
		Hostname: HOSTNAME,
		Username: USERNAME,
		Password: PASSWORD,
	}

	store := "abcdefghijklmnopqrstuvwxyz"

	resp, err := keyfactor.GetCertificateStoreByID(store, clientConfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", resp)
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
