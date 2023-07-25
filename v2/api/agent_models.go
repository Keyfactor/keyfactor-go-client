package api

type Agent9x struct {
	AgentId          string `json:"AgentId"`
	AgentPoolId      string `json:"AgentPoolId"`
	ClientMachine    string `json:"ClientMachine"`
	Username         string `json:"Username"`
	AgentPlatform    int    `json:"AgentPlatform"`
	Status           int    `json:"Status"`
	EnableDiscover   bool   `json:"EnableDiscover"`
	EnableMonitor    bool   `json:"EnableMonitor"`
	Version          string `json:"Version"`
	LastSeen         string `json:"LastSeen"`
	Thumbprint       string `json:"Thumbprint"`
	LegacyThumbprint string `json:"LegacyThumbprint"`
}

type Agent struct {
	AgentId                     string   `json:"AgentId"`
	ClientMachine               string   `json:"ClientMachine"`
	Username                    string   `json:"Username"`
	AgentPlatform               int      `json:"AgentPlatform"`
	Status                      int      `json:"Status"`
	Version                     string   `json:"Version"`
	LastSeen                    string   `json:"LastSeen"`
	Capabilities                []string `json:"Capabilities"`
	Blueprint                   string   `json:"Blueprint"`
	Thumbprint                  string   `json:"Thumbprint"`
	LegacyThumbprint            string   `json:"LegacyThumbprint"`
	AuthCertificateReenrollment string   `json:"AuthCertificateReenrollment"`
	LastThumbprintUsed          string   `json:"LastThumbprintUsed"`
	LastErrorCode               int      `json:"LastErrorCode"`
	LastErrorMessage            string   `json:"LastErrorMessage"`
}
