package api

import (
	"context"
	"fmt"
	"github.com/Keyfactor/keyfactor-go-client-sdk"
)

// GetAgentList returns a list of orchestrators registered in the Keyfactor instance
func (c *Client) GetAgentList() ([]Agent, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	resp, _, err := apiClient.AgentApi.AgentGetAgents(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	var revResp []Agent

	if err != nil {
		return revResp, err
	}

	for i := range resp {
		newAgent := Agent{
			AgentId:          *resp[i].AgentId,
			AgentPoolId:      "", //TODO
			ClientMachine:    *resp[i].ClientMachine,
			Username:         *resp[i].Username,
			AgentPlatform:    int(*resp[i].AgentPlatform),
			Status:           int(*resp[i].Status),
			EnableDiscover:   true, //TODO
			EnableMonitor:    true, //TODO
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

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	resp, _, err := apiClient.AgentApi.AgentGetAgentDetail(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	var revResp []Agent

	if err != nil {
		return revResp, err
	}

	revResp = []Agent{
		{
			AgentId:          *resp.AgentId,
			AgentPoolId:      "", //TODO
			ClientMachine:    *resp.ClientMachine,
			Username:         *resp.Username,
			AgentPlatform:    int(*resp.AgentPlatform),
			Status:           int(*resp.Status),
			EnableDiscover:   true, //TODO
			EnableMonitor:    true, //TODO
			Version:          *resp.Version,
			LastSeen:         "",
			Thumbprint:       *resp.Thumbprint,
			LegacyThumbprint: *resp.LegacyThumbprint,
		},
	}

	//TODO: error check

	return revResp, nil
}

func (c *Client) ApproveAgent(id string) (string, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

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

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

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

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	resp, err := apiClient.AgentApi.AgentReset1(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	resp, err := apiClient.AgentApi.AgentFetchLogs(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if resp.StatusCode == 204 {
		return "Reset agent successful.", nil
	}

	if err != nil {
		return "", err
	}

	return resp.Status, nil
}
