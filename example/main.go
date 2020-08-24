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
	users, response, err := client.DefaultApi.ListUsers(ctx, nil)
	if err != nil {
		log.Fatalf("list users failed. %v\n", err)
	}

	fmt.Printf("response: %v\n", response)
	fmt.Printf("ret type: %T\n", users)
	fmt.Printf("ret: %v\n", users)

	user, response, err := client.DefaultApi.AddUser(ctx, apiClient.NewUser{
		Name:   "Tom",
		Active: true,
	})
	if err != nil {
		log.Fatalf("add user failed. %v\n", err)
	}

	fmt.Printf("add response: %v\n", response)
	fmt.Printf("user: %v\n", user)

	user, response, err = client.DefaultApi.GetUser(ctx, 1)
	if err != nil {
		log.Fatalf("get user failed. %v\n", err)
	}

	fmt.Printf("get response: %v\n", response)
	fmt.Printf("user: %v\n", user)

	user, response, err = client.DefaultApi.UpdateUser(ctx, user.Id, apiClient.User{
		Name:   "Tom Hunks",
		Active: true,
	})
	if err != nil {
		log.Fatalf("update user failed. %v\n", err)
	}

	fmt.Printf("update response: %v\n", response)
	fmt.Printf("user: %v\n", user)

	response, err = client.DefaultApi.DeleteUser(ctx, user.Id)
	if err != nil {
		log.Fatalf("delete user failed. %v\n", err)
	}

	fmt.Printf("delete response: %v\n", response)
}
