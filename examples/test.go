package keyfactor

import (
	"fmt"
	main2 "github.com/Keyfactor/keyfactor-go-client"
	"log"
	"os"
)

func main() {
	// Create a new Keyfactor client
	clientConfig := &main2.AuthConfig{
		Hostname: os.Getenv("HOSTNAME"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	client, err := main2.NewKeyfactorClient(clientConfig)
	if err != nil {
		log.Fatal(err)
	}

	leaf, chain, err := client.DownloadCertificate(3786, "", "", "")
	if err != nil {
		return
	}
	fmt.Println(leaf)
	for _, i := range chain {
		fmt.Println(i)
	}

	priv, leaf, chain, err := client.RecoverCertificate(3786, "", "", "", "TerraformAccTestBasic")
	if err != nil {
		return
	}
	fmt.Println(leaf)
	fmt.Println(priv)
	for _, i := range chain {
		fmt.Println(i)
	}
}
