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
	// Get Certificate Store Types
	//
	// To retrieve data on a specific certificate store type supported by the Keyfactor instance referenced by the
	// Keyfactor Go client, use the GetCertStoreType method.
	storeType, err := client.GetCertStoreType(106)
	if err != nil {
		return
	}
	log.Printf("Short Name: %s", storeType.ShortName)

	// ---------------------------------------------------------------------------------------------------------------//
	// Create New AKV Certificate Store
	//
	// To create a new certificate store using the Keyfactor Go Client,
	// first create a pointer to an CreateStoreFctArgs struct,
	// and populate the required fields. The below fields are the bare minimum. Note that the properties required vary
	// between different store types. Use the GetCertStoreType method to determine the required fields.
	createStore := &keyfactor.CreateStoreFctArgs{
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
	// Then, use the CreateStore method to create the new certificate store.
	store, err := client.CreateStore(createStore)
	if err != nil {
		return
	}
	fmt.Printf("%#v", store)

	// ---------------------------------------------------------------------------------------------------------------//
	// Get Certificate Store Context
	//
	// To retrieve the context of a certificate store in Keyfactor, make a call to the GetCertificateStoreByID method.
	// Pass a string as the certificate store ID.
	storeContext, err := client.GetCertificateStoreByID(store.Id)
	if err != nil {
		return
	}
	fmt.Printf("%#v", storeContext)
}
