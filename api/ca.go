package api

import (
	"context"
	"encoding/json"
)

// GetCAList returns a list of certificate authorities supported by the Keyfactor instance
func (c *Client) GetCAList() ([]CA, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	resp, _, err := apiClient.CertificateAuthorityApi.CertificateAuthorityGetCas(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	var revResp []CA

	if err != nil {
		return revResp, err
	}

	type ExplicitPassword struct {
		SecretValue string   `json:"SecretValue"`
		Parameters  struct{} `json:"Parameters"`
		Provider    int      `json:"Provider"`
	}

	for _, val := range resp {
		vJson, _ := json.Marshal(val)
		var newCA CA
		json.Unmarshal(vJson, &newCA)
		revResp = append(revResp, newCA)
	}

	return revResp, nil
}
