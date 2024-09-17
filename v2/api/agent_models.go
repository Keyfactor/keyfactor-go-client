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
