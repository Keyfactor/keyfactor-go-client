package api

import (
	"fmt"
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
	c, kfcErr := NewKeyfactorClient(&AuthConfig{})

	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}
	agents, lErr := c.GetAgentList()
	if lErr != nil {
		t.Errorf("unable to list agents. %s", lErr)
		return
	}
	if len(agents) == 0 {
		t.Errorf("no agents found in client.")
		return
	}
	agentID := agents[0].AgentId
	agentClientName := agents[0].ClientMachine
	agentStatus := agents[0].Status
	type fields struct{}
	type args struct {
		id         string
		clientName string
		status     int
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
				status:     agentStatus,
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
				status:     agentStatus,
			},
			want:    "",
			wantErr: false, // TODO: Currently always errors out
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.status == ApprovedAgentStatus {
				_, err := c.DisApproveAgent(tt.args.id)
				if (err != nil) != tt.wantErr {
					t.Errorf("DisApproveAgent() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else {
				_, err := c.ApproveAgent(tt.args.id)
				if (err != nil) != tt.wantErr {
					t.Errorf("ApproveAgent() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}

// TODO
func TestClient_FetchAgentLogs(t *testing.T) {
	log.SetOutput(io.Discard)
	c, kfcErr := NewKeyfactorClient(&AuthConfig{})
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
	c, kfcErr := NewKeyfactorClient(&AuthConfig{})
	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}
	agents, lErr := c.GetAgentList()
	if lErr != nil {
		t.Errorf("unable to list agents. %s", lErr)
		return
	}
	if len(agents) == 0 {
		t.Errorf("no agents found in client.")
		return
	}
	agentID := agents[0].AgentId
	type fields struct{}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Agent
		wantErr bool
	}{
		{
			name:   "agent_get_failure",
			fields: fields{}, args: args{
				id: "invalid-agent-name",
			},
			want:    []Agent{},
			wantErr: true,
		},
		{
			name:   "agent_get_success",
			fields: fields{}, args: args{
				id: agentID,
			},
			want:    []Agent{},
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
			}
		})
	}
}

func TestClient_GetAgentList(t *testing.T) {
	log.SetOutput(io.Discard)
	log.SetOutput(io.Discard)
	c, kfcErr := NewKeyfactorClient(&AuthConfig{})
	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}

	type fields struct{}
	tests := []struct {
		name    string
		fields  fields
		want    []Agent
		wantErr bool
	}{
		{
			name:    "GetAgentList",
			fields:  fields{},
			want:    []Agent{},
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
	c, kfcErr := NewKeyfactorClient(&AuthConfig{})
	if kfcErr != nil {
		t.Errorf("unable to connect to Keyfactor. Please check your credentials and try again. %s", kfcErr)
		return
	}
	agents, lErr := c.GetAgentList()
	if lErr != nil {
		t.Errorf("unable to list agents. %s", lErr)
		return
	}
	if len(agents) == 0 {
		t.Errorf("no agents found in client.")
		return
	}
	agentID := agents[0].AgentId
	agentClientName := agents[0].ClientMachine
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
			want:    "404 Not Found",
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
			var got string
			var err error
			if tt.args.action == AgentActionReset {
				got, err = c.ResetAgent(tt.args.id)
			} else if tt.args.action == AgentActionApprove {
				got, err = c.ApproveAgent(tt.args.id)
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
