// Copyright 2024 Keyfactor
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"context"
	"encoding/json"

	"github.com/Keyfactor/keyfactor-go-client-sdk/api/keyfactor"
)

// GetCAList returns a list of certificate authorities supported by the Keyfactor instance
func (c *Client) GetCAList() ([]CA, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor.NewConfiguration(make(map[string]string))
	apiClient := keyfactor.NewAPIClient(configuration)

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
