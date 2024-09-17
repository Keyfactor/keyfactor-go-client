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
	"os"

	main2 "github.com/Keyfactor/keyfactor-go-client/api"
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
