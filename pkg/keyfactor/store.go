package keyfactor

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type CreateStoreFctArgs struct {
	ContainerId             int
	ClientMachine           string
	StorePath               string
	CertStoreInventoryJobId int
	CertStoreType           int
	Approved                bool
	CreateIfMissing         bool
	Properties              interface{}
	AgentId                 string
	AgentAssigned           bool
	ContainerName           string
	InventorySchedule       interface{}
	ReEnrollmentStatus      *ReEnrollmnentConfig
	SetNewPasswordAllowed   bool
	Password                string
}

type ReEnrollmnentConfig struct {
	Data               bool
	AgentId            string
	Message            string
	JobProperties      []StringTuple
	CustomAliasAllowed int
}

type StorePasswordConfig struct {
	Value                       string
	SecretTypeGuid              string
	ProviderTypeParameterValues *ProviderTypeParams // Not yet implemented
}

/* Future non-critical functionality */

type ProviderTypeParams struct {
	Id           string
	Value        string
	InstanceId   string
	InstanceGuid string
	Provider     *ProviderParams
}

type ProviderParams struct {
	Id           int
	Name         string
	Area         int
	ProviderType *ProviderType
}

type ProviderType struct {
	Id   string
	Name string
}

type CertStoreTypeResponse struct {
	Name                string `json:"Name"`
	ShortName           string `json:"ShortName"`
	Capability          string `json:"Capability"`
	StoreType           int    `json:"StoreType"`
	ImportType          int    `json:"ImportType"`
	LocalStore          bool   `json:"LocalStore"`
	SupportedOperations struct {
		Add        bool `json:"Add"`
		Create     bool `json:"Create"`
		Discovery  bool `json:"Discovery"`
		Enrollment bool `json:"Enrollment"`
		Remove     bool `json:"Remove"`
	} `json:"SupportedOperations"`
	Properties []struct {
		StoreTypeID  int    `json:"StoreTypeID"`
		Name         string `json:"Name"`
		DisplayName  string `json:"DisplayName"`
		Type         string `json:"Type"`
		DependsOn    string `json:"DependsOn"`
		DefaultValue string `json:"DefaultValue"`
		Required     bool   `json:"Required"`
	} `json:"Properties"`
	PasswordOptions struct {
		EntrySupported bool   `json:"EntrySupported"`
		StoreRequired  bool   `json:"StoreRequired"`
		Style          string `json:"Style"`
	} `json:"PasswordOptions"`
	StorePathValue     []string `json:"store_path_value"`
	PrivateKeyAllowed  string   `json:"private_key_allowed"`
	JobProperties      []string `json:"job_properties"`
	ServerRequired     bool     `json:"ServerRequired"`
	PowerShell         bool     `json:"PowerShell"`
	BlueprintAllowed   bool     `json:"BlueprintAllowed"`
	CustomAliasAllowed string   `json:"CustomAliasAllowed"`
	ServerRegistration int      `json:"ServerRegistration"`
	InventoryEndpoint  string   `json:"InventoryEndpoint"`
	InventoryJobType   string   `json:"InventoryJobType"`
	ManagementJobType  string   `json:"ManagementJobType"`
	DiscoveryJobType   string   `json:"DiscoveryJobType"`
	EnrollmentJobType  string   `json:"EnrollmentJobType"`
}

func CreateStore(ca *CreateStoreFctArgs, api *APIClient) error {
	log.Println("[INFO] Creating new certificate store with Keyfactor")

	/* Validate input */
	err := validateCreateStoreArgs(ca)
	if err != nil {
		return err
	}

	return nil
}

func validateCreateStoreArgs(ca *CreateStoreFctArgs) error {
	if ca.ClientMachine == "" {
		return errors.New("client machine is required for creation of new certificate store")
	}
	if ca.StorePath == "" {
		return errors.New("store path is required for creation of new certificate store")
	}
	if ca.Properties == nil {
		return errors.New("properties attribute is required for creation of new certificate store")
	}
	if ca.AgentId == "" {
		return errors.New("orchestrator agent id is required for creation of new certificate store")
	}
	return nil
}

func GetCertStoreType(id int, api *APIClient) (*CertStoreTypeResponse, error) {

	// Set Keyfactor-specific headers
	log.Println("[TRACE] Setting request headers")
	headers := &APIHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	log.Println("[TRACE] Creating request struct for Keyfactor client")
	endpoint := fmt.Sprintf("/KeyfactorAPI/CertificateStoreTypes/%d", id)
	keyfactorAPIStruct := &APIRequest{
		KFClient: api,
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  headers,
		Payload:  nil,
	}

	resp, err := SendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		log.Printf("[DEBUG] GET succeeded with response code %d", resp.StatusCode)
		jsonResp := &CertStoreTypeResponse{}
		err = json.NewDecoder(resp.Body).Decode(&jsonResp)
		if err != nil {
			return nil, err
		}
		return jsonResp, nil
	}

	var errorMessage interface{} // Decode JSON body to handle issue
	err = json.NewDecoder(resp.Body).Decode(&errorMessage)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Request failed with code %d and message %v", resp.StatusCode, errorMessage)
	stringMessage := fmt.Sprintf("%v", errorMessage)
	return nil, errors.New(stringMessage)
}
