package main

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

func main() {
	credential, err := azidentity.NewDefaultAzureCredential(nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	clientOptions := azcosmos.ClientOptions{
		EnableContentResponseOnWrite: true,
	}

	client, err := azcosmos.NewClient("https://sci-stg-cosmosdb-sql.documents.azure.com:443/", credential, &clientOptions)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(client.Endpoint())
}
