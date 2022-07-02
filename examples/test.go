package main

import (
	"fmt"
	"github.com/Keyfactor/keyfactor-go-client/pkg/keyfactor"
	"log"
	"os"
)

func main() {
	// Create a new Keyfactor client
	clientConfig := &keyfactor.AuthConfig{
		Hostname: os.Getenv("HOSTNAME"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	client, err := keyfactor.NewKeyfactorClient(clientConfig)
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
}

// Helper functions
func boolToPointer(b bool) *bool {
	return &b
}

func intToPointer(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

func stringToPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
