package main

import (
	"context"
	"fmt"
	"log"

	apiClient "github.com/yhidai/openapi-getting-started/client"
)

func main() {
	client := apiClient.NewAPIClient(apiClient.NewConfiguration())

	ctx := context.Background()
	ret, response, err := client.DefaultApi.ListUsers(ctx, nil)
	if err != nil {
		log.Fatalf("listUsers failed. %v\n", err)
	}

	fmt.Printf("response: %v\n", response)
	fmt.Printf("ret: %v\n", ret)
}
