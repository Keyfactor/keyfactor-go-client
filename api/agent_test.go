package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

func initClient() (*Client, error) {
	var clientAuth AuthConfig
	clientAuth.Username = os.Getenv("KEYFACTOR_USERNAME")
	log.Printf("[DEBUG] Username: %s", clientAuth.Username)
	clientAuth.Password = os.Getenv("KEYFACTOR_PASSWORD")
	log.Printf("[DEBUG] Password: %s", clientAuth.Password)
	clientAuth.Domain = os.Getenv("KEYFACTOR_DOMAIN")
	log.Printf("[DEBUG] Domain: %s", clientAuth.Domain)
	clientAuth.Hostname = os.Getenv("KEYFACTOR_HOSTNAME")
	log.Printf("[DEBUG] Hostname: %s", clientAuth.Hostname)

	c, err := NewKeyfactorClient(&clientAuth)

	if err != nil {
		log.Fatalf("Error creating Keyfactor client: %s", err)
	}
	return c, err
}

func TestClient_ApproveAgent(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	kfClient, _ := initClient()
	agentID := os.Getenv("TEST_KEYFACTOR_AGENT_ID")
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "agent_approve_failure",
			fields: fields{
				hostname:        kfClient.hostname,
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			}, args: args{
				id: agentID,
			},
			want:    "",
			wantErr: true, // TODO: Currently always errors out
		},
		//{
		//	name: "agent_approve_success",
		//	fields: fields{
		//		hostname:        kfClient.hostname,
		//		httpClient:      kfClient.httpClient,
		//		basicAuthString: kfClient.basicAuthString,
		//	}, args: args{
		//		id: "1",
		//	},
		//	want:    "",
		//	wantErr: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				hostname:        tt.fields.hostname,
				httpClient:      tt.fields.httpClient,
				basicAuthString: tt.fields.basicAuthString,
			}
			got, err := c.ApproveAgent(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApproveAgent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ApproveAgent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DisApproveAgent(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	kfClient, _ := initClient()
	agentID := os.Getenv("TEST_KEYFACTOR_AGENT_ID")
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "agent_disapprove_failure",
			fields: fields{
				hostname:        kfClient.hostname,
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			}, args: args{
				id: agentID,
			},
			want:    "",
			wantErr: true, // TODO: Currently always errors out
		},
		//{
		//	name: "agent_disapprove_success",
		//	fields: fields{
		//		hostname:        kfClient.hostname,
		//		httpClient:      kfClient.httpClient,
		//		basicAuthString: kfClient.basicAuthString,
		//	}, args: args{
		//		id: "1",
		//	},
		//	want:    "",
		//	wantErr: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				hostname:        tt.fields.hostname,
				httpClient:      tt.fields.httpClient,
				basicAuthString: tt.fields.basicAuthString,
			}
			got, err := c.DisApproveAgent(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DisApproveAgent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DisApproveAgent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_FetchAgentLogs(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	kfClient, _ := initClient()
	agentID := os.Getenv("TEST_KEYFACTOR_AGENT_ID")
	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "agent_fetchlogs_failure",
			fields: fields{
				hostname:        kfClient.hostname,
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			}, args: args{
				id: "",
			},
			want:    "invalid-agent-id",
			wantErr: true,
		},
		{
			name: "agent_fetchlogs_success",
			fields: fields{
				hostname:        kfClient.hostname,
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			}, args: args{
				id: agentID,
			},
			want:    "Fetch logs successful.",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				hostname:        tt.fields.hostname,
				httpClient:      tt.fields.httpClient,
				basicAuthString: tt.fields.basicAuthString,
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
	log.SetOutput(ioutil.Discard)
	kfClient, _ := initClient()
	agentID := os.Getenv("TEST_KEYFACTOR_AGENT_NAME")

	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
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
			name: "agent_get_failure",
			fields: fields{
				hostname:        kfClient.hostname,
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			}, args: args{
				id: "invalid-agent-id",
			},
			want:    []Agent{},
			wantErr: true,
		},
		{
			name: "agent_get_success",
			fields: fields{
				hostname:        kfClient.hostname,
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			}, args: args{
				id: agentID,
			},
			want:    []Agent{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				hostname:        tt.fields.hostname,
				httpClient:      tt.fields.httpClient,
				basicAuthString: tt.fields.basicAuthString,
			}
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
	log.SetOutput(ioutil.Discard)
	kfClient, _ := initClient()

	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Agent
		wantErr bool
	}{
		{
			name: "agent_list_failure",
			fields: fields{
				hostname:        "",
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			},
			want:    []Agent{},
			wantErr: true,
		},
		{
			name: "agent_list_success",
			fields: fields{
				hostname:        kfClient.hostname,
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			},
			want:    []Agent{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				hostname:        tt.fields.hostname,
				httpClient:      tt.fields.httpClient,
				basicAuthString: tt.fields.basicAuthString,
			}
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
	log.SetOutput(ioutil.Discard)
	kfClient, _ := initClient()
	agentID := os.Getenv("TEST_KEYFACTOR_AGENT_NAME")

	type fields struct {
		hostname        string
		httpClient      *http.Client
		basicAuthString string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "agent_reset_failure",
			fields: fields{
				hostname:        kfClient.hostname,
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			}, args: args{
				id: "invalid-agent-id",
			},
			want:    "Error 404 - the requested resource was not found. Please check the request and try again.",
			wantErr: true,
		},
		{
			name: "agent_reset_success",
			fields: fields{
				hostname:        kfClient.hostname,
				httpClient:      kfClient.httpClient,
				basicAuthString: kfClient.basicAuthString,
			}, args: args{
				id: agentID,
			},
			want:    "Reset agent successful.",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				hostname:        tt.fields.hostname,
				httpClient:      tt.fields.httpClient,
				basicAuthString: tt.fields.basicAuthString,
			}
			aID, aErr := c.GetAgent(tt.args.id)
			if aErr != nil && !tt.wantErr {
				t.Errorf("ResetAgent() error = %v, wantErr %v", aErr, tt.wantErr)
				return
			}
			var got string
			var err error
			if aID != nil && len(aID) > 0 {
				got, err = c.ResetAgent(aID[0].AgentId)
			} else {
				got, err = c.ResetAgent(agentID)
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
