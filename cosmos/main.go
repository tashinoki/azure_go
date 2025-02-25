package main

import (
	"context"
	"encoding/json"
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

	cosmosDbName, ok := os.LookupEnv("COSMOS_DATABASE")
	if !ok {
		fmt.Println("COSMOS_DATABASE is not set")
		return
	}

	database, err := client.NewDatabase(cosmosDbName)

	if err != nil {
		fmt.Println(err)
		return
	}

	cosmosContainerName, ok := os.LookupEnv("COSMOS_CONTAINER")
	if !ok {
		fmt.Println("COSMOS_CONTAINER is not set")
		return
	}

	container, err := database.NewContainer(cosmosContainerName)

	if err != nil {
		fmt.Println(err)
		return
	}

	pk := azcosmos.NewPartitionKeyString("")
	newItem = Item{Id: "1", Pk: ""}
	newBytes := []byte{}

	encoding.encode(newItem, &newBytes)
	_, err := container.CreateItem(context.TODO(), pk, &newBytes, nil)

	partitionKey := azcosmos.NewPartitionKeyString("")
	query := "SELECT * FROM c WHERE c.pk = @partitionKey"

	queryOptions := azcosmos.QueryOptions{
		QueryParameters: []azcosmos.QueryParameter{
			{Name: "@partitionKey", Value: partitionKey},
		},
	}

	pager := container.NewQueryItemsPager(query, partitionKey, &queryOptions)
	for pager.More() {
		response, err := pager.NextPage(context.TODO())
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(response.Items)
		for _, bytes := range response.Items {
			item := Item{}
			err := json.Unmarshal(bytes, &item)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(item)
		}
	}
}

type Item struct {
	Id string `json:"id"`
	Pk string `json:"pk"`
}
