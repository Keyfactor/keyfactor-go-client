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

package examples

import (
	"fmt"
	"log"

	"github.com/Keyfactor/keyfactor-go-client/api"
)

var HOSTNAME = ""
var USERNAME = ""
var PASSWORD = ""

func main() {
	// Create a new Keyfactor client
	clientConfig := &api.AuthConfig{
		Hostname: HOSTNAME,
		Username: USERNAME,
		Password: PASSWORD,
	}
	client, err := api.NewKeyfactorClient(clientConfig)
	if err != nil {
		log.Fatal(err)
	}

	// ---------------------------------------------------------------------------------------------------------------//
	// Get Keyfactor security identity list
	//
	identities, err := client.GetSecurityIdentities()
	if err != nil {
		return
	}
	for _, identity := range identities {
		fmt.Printf("Id: %d AccountName: %s\n", identity.Id, identity.AccountName)
		fmt.Printf("    Roles IDs: ")
		for _, i := range identity.Roles {
			fmt.Printf("%d, ", i.Id)
		}
		fmt.Println("")
	}

	// ---------------------------------------------------------------------------------------------------------------//
	// Get Keyfactor security role list
	//
	role, err := client.GetSecurityRole(88)
	if err != nil {
		return
	}
	fmt.Printf("Role name: %s Role ID: %d\n", role.Name, role.Id)
}
