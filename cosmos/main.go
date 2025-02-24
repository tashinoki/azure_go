package main

import (
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
		return
	}

	credential, err := azidentity.NewDefaultAzureCredential(nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	clientOptions := azcosmos.ClientOptions{
		EnableContentResponseOnWrite: true,
	}

	cosmosDbEndpoint, ok := os.LookupEnv("COSMOS_DB_ENDPOINT")
	if !ok {
		fmt.Println("COSMOS_DB_ENDPOINT is not set")
		return
	}

	client, err := azcosmos.NewClient(cosmosDbEndpoint, credential, &clientOptions)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(client.Endpoint())
}
