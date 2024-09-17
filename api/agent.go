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
	"fmt"

	"github.com/Keyfactor/keyfactor-go-client-sdk/api/keyfactor"
)

// GetAgentList returns a list of orchestrators registered in the Keyfactor instance
func (c *Client) GetAgentList() ([]Agent, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor.NewConfiguration(make(map[string]string))
	apiClient := keyfactor.NewAPIClient(configuration)

	resp, _, err := apiClient.AgentApi.AgentGetAgents(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	var revResp []Agent

	if err != nil {
		return revResp, err
	}

	for i := range resp {
		newAgent := Agent{
			AgentId:          *resp[i].AgentId,
			AgentPoolId:      "",
			ClientMachine:    *resp[i].ClientMachine,
			Username:         *resp[i].Username,
			AgentPlatform:    int(*resp[i].AgentPlatform),
			Status:           int(*resp[i].Status),
			EnableDiscover:   false, //TODO
			EnableMonitor:    false, //TODO
			Version:          *resp[i].Version,
			LastSeen:         resp[i].LastSeen.String(),
			Thumbprint:       *resp[i].Thumbprint,
			LegacyThumbprint: *resp[i].LegacyThumbprint,
		}
		revResp = append(revResp, newAgent)
	}

	return revResp, nil
}

func (c *Client) GetAgent(id string) ([]Agent, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor.NewConfiguration(make(map[string]string))
	apiClient := keyfactor.NewAPIClient(configuration)

	resp, _, err := apiClient.AgentApi.AgentGetAgentDetail(
		context.Background(),
		id,
	).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	var revResp []Agent

	if err != nil {
		return revResp, err
	}

	revResp = []Agent{
		{
			AgentId:          resp.GetAgentId(),
			AgentPoolId:      "",
			ClientMachine:    resp.GetClientMachine(),
			Username:         resp.GetUsername(),
			AgentPlatform:    int(resp.GetAgentPlatform()),
			Status:           int(resp.GetStatus()),
			EnableDiscover:   false, //TODO
			EnableMonitor:    false, //TODO
			Version:          resp.GetVersion(),
			LastSeen:         resp.GetLastSeen().String(),
			Thumbprint:       resp.GetThumbprint(),
			LegacyThumbprint: resp.GetLegacyThumbprint(),
		},
	}

	return revResp, nil
}

func (c *Client) ApproveAgent(id string) (string, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor.NewConfiguration(make(map[string]string))
	apiClient := keyfactor.NewAPIClient(configuration)

	var ids = []string{id}

	resp, err := apiClient.AgentApi.AgentApprove(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AgentIds(ids).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if resp.StatusCode == 204 {
		return "Approve agent successful.", nil
	}

	if err != nil {
		return "", err
	}

	return resp.Status, nil
}

func (c *Client) DisApproveAgent(id string) (string, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor.NewConfiguration(make(map[string]string))
	apiClient := keyfactor.NewAPIClient(configuration)

	var ids = []string{id}

	resp, err := apiClient.AgentApi.AgentDisapprove(context.Background()).AgentIds(ids).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if resp.StatusCode == 204 {
		return fmt.Sprintf("Disapproving %s successful.", id), nil
	}

	if err != nil {
		return "", err
	}

	return resp.Status, nil
}

func (c *Client) ResetAgent(id string) (string, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor.NewConfiguration(make(map[string]string))
	apiClient := keyfactor.NewAPIClient(configuration)

	resp, err := apiClient.AgentApi.AgentReset1(
		context.Background(),
		id,
	).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if resp.StatusCode == 204 {
		return "Reset agent successful.", nil
	}

	if err != nil {
		return "", err
	}

	return resp.Status, nil
}

func (c *Client) FetchAgentLogs(id string) (string, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor.NewConfiguration(make(map[string]string))
	apiClient := keyfactor.NewAPIClient(configuration)

	resp, err := apiClient.AgentApi.AgentFetchLogs(
		context.Background(),
		id,
	).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if resp.StatusCode == 204 {
		return "Reset agent successful.", nil
	}

	if err != nil {
		return "", err
	}

	return resp.Status, nil
}
