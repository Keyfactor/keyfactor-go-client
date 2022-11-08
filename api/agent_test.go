package api_test

import (
	"fmt"
	"github.com/Keyfactor/keyfactor-go-client/api"
	"io"
	"log"
	"os"
	"testing"
)

const (
	ApprovedAgentStatus   = 2
	AgentActionApprove    = "approve"
	AgentActionDisApprove = "disapprove"
	AgentActionReset      = "reset"
	AgentActionLogs       = "logs"
)

func TestClient_ApproveAgent(t *testing.T) {
	log.SetOutput(io.Discard)
	c, kfcErr := api.NewKeyfactorClient(&api.AuthConfig{})
	agentID := os.Getenv("TEST_KEYFACTOR_AGENT_ID")
	agentClientName := os.Getenv("TEST_KEYFACTOR_AGENT_NAME")

	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}

	type fields struct{}
	type args struct {
		id         string
		clientName string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "InvertApprovalStatus",
			fields: fields{},
			args: args{
				id:         agentID,
				clientName: agentClientName,
			},
			want:    "",
			wantErr: false,
		},
		{
			name:   "RevertApprovalStatus",
			fields: fields{},
			args: args{
				id:         agentID,
				clientName: agentClientName,
			},
			want:    "",
			wantErr: false, // TODO: Currently always errors out
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agents, lErr := c.GetAgentList()
			if lErr != nil {
				t.Errorf("unable to list agents. %s", lErr)
				return
			}
			foundAgent := false
			var orchestrator api.Agent
			for _, agent := range agents {
				if agent.AgentId == tt.args.id || agent.ClientMachine == tt.args.clientName {
					orchestrator = agent
					foundAgent = true
					break
				}
			}
			if !foundAgent {
				t.Errorf("unable to find agent with id %s or client name %s", tt.args.id, tt.args.clientName)
				return
			}
			if orchestrator.Status == ApprovedAgentStatus {
				_, err := c.DisApproveAgent(orchestrator.AgentId)
				if (err != nil) != tt.wantErr {
					t.Errorf("DisApproveAgent() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else {
				_, err := c.ApproveAgent(orchestrator.AgentId)
				if (err != nil) != tt.wantErr {
					t.Errorf("ApproveAgent() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}

func TestClient_FetchAgentLogs(t *testing.T) {
	log.SetOutput(io.Discard)
	c, kfcErr := api.NewKeyfactorClient(&api.AuthConfig{})
	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}
	agentID := os.Getenv("TEST_KEYFACTOR_AGENT_ID")
	agentClientName := os.Getenv("TEST_KEYFACTOR_AGENT_NAME")
	type fields struct{}
	type args struct {
		id         string
		clientName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "agent_fetchlogs_failure",
			fields: fields{}, args: args{
				id:         "",
				clientName: "",
			},
			want:    "invalid-agent-name",
			wantErr: true,
		},
		{
			name:   "agent_fetchlogs_success",
			fields: fields{}, args: args{
				id:         agentID,
				clientName: agentClientName,
			},
			want:    "Fetch logs successful.",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.id == "" && tt.args.clientName == "" && !tt.wantErr {
				t.Errorf("unable to find agent with id %s or client name %s", tt.args.id, agentClientName)
				return
			} else if tt.args.id == "" {
				agents, lErr := c.GetAgentList()
				if lErr != nil && !tt.wantErr {
					t.Errorf("unable to list agents. %s", lErr)
					return
				}
				foundAgent := false
				for _, agent := range agents {
					if agent.ClientMachine == tt.args.clientName {
						tt.args.id = agent.AgentId
						foundAgent = true
						break
					}
				}
				if !foundAgent && !tt.wantErr {
					t.Errorf("unable to find agent with id %s or client name %s", tt.args.id, tt.args.clientName)
					return
				}
			}
			got, err := c.FetchAgentLogs(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchAgentLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Check returned message
				if got != tt.want {
					t.Errorf("FetchAgentLogs() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestClient_GetAgent(t *testing.T) {
	log.SetOutput(io.Discard)
	c, kfcErr := api.NewKeyfactorClient(&api.AuthConfig{})
	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}
	agentID := os.Getenv("TEST_KEYFACTOR_AGENT_NAME")

	type fields struct{}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []api.Agent
		wantErr bool
	}{
		{
			name:   "agent_get_failure",
			fields: fields{}, args: args{
				id: "invalid-agent-name",
			},
			want:    []api.Agent{},
			wantErr: true,
		},
		{
			name:   "agent_get_success",
			fields: fields{}, args: args{
				id: agentID,
			},
			want:    []api.Agent{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := c.GetAgent(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAgent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Check that length of the array is greater than 0
			if !tt.wantErr {
				if len(got) == 0 {
					t.Errorf("GetAgent() got %d records, want %d record(s).\n", len(got), 1)
					return
				}
				// Check returned agent name
				if got[0].ClientMachine != tt.args.id {
					t.Errorf("GetAgent() got %s, want %s.\n", got[0].ClientMachine, tt.args.id)
					return
				}
			}
		})
	}
}

func TestClient_GetAgentList(t *testing.T) {
	log.SetOutput(io.Discard)
	log.SetOutput(io.Discard)
	c, kfcErr := api.NewKeyfactorClient(&api.AuthConfig{})
	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}

	type fields struct{}
	tests := []struct {
		name    string
		fields  fields
		want    []api.Agent
		wantErr bool
	}{
		{
			name:    "GetAgentList",
			fields:  fields{},
			want:    []api.Agent{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := c.GetAgentList()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAgentList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Check that length of the array is greater than 0
				if len(got) == 0 {
					t.Errorf("GetAgentList() got %d records, want at least %d record(s).\n", len(got), 1)
					return
				}
				// Check returned agent name
				if got[0].ClientMachine == "" {
					t.Errorf("GetAgentList() got empty agent name, want non-empty agent name.\n")
					return
				}
			}
		})
	}
}

func TestClient_ResetAgent(t *testing.T) {
	log.SetOutput(io.Discard)
	c, kfcErr := api.NewKeyfactorClient(&api.AuthConfig{})
	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}
	agentID := os.Getenv("TEST_KEYFACTOR_AGENT_ID")
	agentClientName := os.Getenv("TEST_KEYFACTOR_AGENT_NAME")

	type fields struct{}
	type args struct {
		id         string
		clientName string
		action     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "agent_reset_failure",
			fields: fields{},
			args: args{
				id:         "invalid-agent-id",
				clientName: "invalid-agent-name",
				action:     AgentActionReset,
			},
			want:    "Error 404 - the requested resource was not found. Please check the request and try again.",
			wantErr: true,
		},
		{
			name:   "agent_reset_success",
			fields: fields{},
			args: args{
				id:         agentID,
				clientName: agentClientName,
				action:     AgentActionReset,
			},
			want:    "Reset agent successful.",
			wantErr: false,
		},
		{
			name:   "approve_reset_agent",
			fields: fields{},
			args: args{
				id:         agentID,
				clientName: agentClientName,
				action:     AgentActionApprove,
			},
			want:    "Approve agent successful.",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.id == "" && tt.args.clientName == "" {
				t.Errorf("TestClient_ResetAgent() missing agent ID or agent name. Please set TEST_KEYFACTOR_AGENT_ID and TEST_KEYFACTOR_AGENT_NAME environment variables and try again.")
				return
			}

			aID, aErr := c.GetAgent(tt.args.clientName)
			if aErr != nil && !tt.wantErr {
				t.Errorf("ResetAgent() error = %v, wantErr %v", aErr, tt.wantErr)
				return
			}
			var got string
			var err error
			if tt.args.action == AgentActionReset {
				if aID != nil && len(aID) > 0 {
					got, err = c.ResetAgent(aID[0].AgentId)
				} else {
					got, err = c.ResetAgent(tt.args.id)
				}
			} else if tt.args.action == AgentActionApprove {
				if aID != nil && len(aID) > 0 {
					got, err = c.ApproveAgent(aID[0].AgentId)
				} else {
					got, err = c.ApproveAgent(tt.args.id)
				}
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ResetAgent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Check that length of the array is greater than 0
				if got != tt.want {
					t.Errorf("ResetAgent() got %s, want %s.\n", got, tt.want)
					return
				}
			} else {
				errStr := fmt.Sprintf("%s", err)
				if errStr != tt.want {
					t.Errorf("ResetAgent() got %s, want %s.\n", err, tt.want)
					return
				}
			}
		})
	}
}
