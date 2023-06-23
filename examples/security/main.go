package examples

import (
	"fmt"
	"github.com/Keyfactor/keyfactor-go-client/api"
	"log"
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
