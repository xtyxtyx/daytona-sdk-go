package daytona_test

import (
	"context"
	"fmt"
	"log"

	daytona "github.com/daytonaio/daytona-sdk-go"
)

func ExampleNewClient() {
	// Create a new client (automatically loads .env file)
	client, err := daytona.NewClient(&daytona.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Use the client
	ctx := context.Background()
	sandboxes, err := client.ListSandboxes(ctx)
	if err != nil {
		log.Printf("Failed to list sandboxes: %v", err)
		return
	}

	fmt.Printf("Found %d sandboxes\n", len(sandboxes))
}

func ExampleClient_CreateSandbox() {
	client, _ := daytona.NewClient(&daytona.Config{})
	ctx := context.Background()

	// Create a sandbox
	sandbox, err := client.CreateSandbox(ctx, &daytona.CreateSandboxRequest{
		User:     daytona.StringPtr("daytona"),
		Target:   daytona.StringPtr("eu"),
		Snapshot: daytona.StringPtr("daytonaio/sandbox:0.4.3"),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created sandbox: %s\n", sandbox.GetId())
}

func ExampleClient_ExecuteCommand() {
	client, _ := daytona.NewClient(&daytona.Config{})
	ctx := context.Background()
	sandboxID := "your-sandbox-id"

	// Execute a command
	result, err := client.ExecuteCommand(ctx, sandboxID, &daytona.ExecuteCommandRequest{
		Command: "echo 'Hello, World!'",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output: %s\n", result.Result)
}