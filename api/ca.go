package api

import (
	"context"
	"encoding/json"
	"github.com/Keyfactor/keyfactor-go-client-sdk"
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

	for i := range resp {
		jsonParameters, _ := json.Marshal(*resp[i].ExplicitPassword)
		var newExplicitPassword ExplicitPassword
		json.Unmarshal(jsonParameters, &newExplicitPassword)
		newCA := CA{
			Id:                     int(*resp[i].Id),
			LogicalName:            *resp[i].LogicalName,
			HostName:               *resp[i].HostName,
			Delegate:               *resp[i].Delegate,
			ForestRoot:             *resp[i].ForestRoot,
			Remote:                 *resp[i].Remote,
			Agent:                  *resp[i].Agent,
			Standalone:             *resp[i].Standalone,
			MonitorThresholds:      *resp[i].MonitorThresholds,
			IssuanceMax:            int(*resp[i].IssuanceMax),
			IssuanceMin:            int(*resp[i].IssuanceMin),
			DenialMax:              int(*resp[i].DenialMax),
			FailureMax:             int(*resp[i].FailureMax),
			RFCEnforcement:         *resp[i].RFCEnforcement,
			Properties:             *resp[i].Properties,
			AllowedEnrollmentTypes: int(*resp[i].AllowedEnrollmentTypes),
			KeyRetention:           int(*resp[i].KeyRetention),
			KeyRetentionDays:       int(*resp[i].KeyRetentionDays),
			ExplicitCredentials:    *resp[i].ExplicitCredentials,
			SubscriberTerms:        *resp[i].SubscriberTerms,
			ExplicitUser:           *resp[i].ExplicitUser,
			ExplicitPassword:       newExplicitPassword,
			UseAllowedRequesters:   *resp[i].UseAllowedRequesters,
			AllowedRequesters:      resp[i].AllowedRequesters,
		}
		revResp = append(revResp, newCA)
	}

	return revResp, nil
}
