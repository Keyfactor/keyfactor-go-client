package main

import (
	"fmt"
	"github.com/Keyfactor/keyfactor-go-client/api"
	"log"
	"os"
)

func main() {
	// Create a new Keyfactor client
	clientConfig := &api.AuthConfig{
		Hostname: os.Getenv("KEYFACTOR_HOSTNAME"),
		Username: os.Getenv("KEYFACTOR_USERNAME"),
		Password: os.Getenv("KEYFACTOR_PASSWORD"),
	}
	client, err := api.NewKeyfactorClient(clientConfig)
	if err != nil {
		log.Fatal(err)
	}

	// ---------------------------------------------------------------------------------------------------------------//
	// Get Certificate Store Types
	//
	// To retrieve data on a specific certificate store type supported by the Keyfactor instance referenced by the
	// Keyfactor Go client, use the GetCertificateStoreType method.
	storeType, err := client.GetCertificateStoreType(106)
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
	// between different store types. Use the GetCertificateStoreType method to determine the required fields.
	properties := make(map[string]string)
	properties["TenantID"] = "tenant"
	properties["ResourceGroupName"] = "resource group name"
	properties["ApplicationId"] = "app ID"
	properties["ClientSecret"] = "client secret"
	properties["SubscriptionId"] = "subscription ID"
	properties["APIObjectId"] = "object ID"
	properties["VaultName"] = "coolvault"

	createStore := &api.CreateStoreFctArgs{
		AgentId:       "keyfactor orchestrator agent ID",
		ClientMachine: "aks_demo",
		StorePath:     "https://coolvault.vault.azure.net/",
		CertStoreType: 106,
		Properties:    properties,
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
