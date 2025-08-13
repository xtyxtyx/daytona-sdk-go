package main

import (
	"context"
	"fmt"
	"log"

	daytona "github.com/daytonaio/daytona-sdk-go/pkg"
)

func main() {
	// Create SDK client
	// API key is loaded from DAYTONA_API_KEY environment variable
	client, err := daytona.NewClient(&daytona.Config{})
	if err != nil {
		log.Fatal("Failed to create client:", err)
	}

	ctx := context.Background()
	
	fmt.Println("=== Command Execution Example ===\n")

	// For demonstration, we'll use a sandbox ID
	// In real usage, you would create or get a sandbox first
	sandboxID := "example-sandbox-id"

	// Simple command
	fmt.Println("Executing simple command...")
	result, err := client.ExecuteCommand(ctx, sandboxID, &daytona.ExecuteCommandRequest{
		Command: "echo 'Hello, World!'",
	})
	if err != nil {
		log.Printf("Failed to execute command: %v\n", err)
	} else {
		fmt.Printf("Output: %s\n", result.Result)
		fmt.Printf("Exit code: %.0f\n", result.ExitCode)
	}

	// Command with working directory
	fmt.Println("\nExecuting command with working directory...")
	result, err = client.ExecuteCommand(ctx, sandboxID, &daytona.ExecuteCommandRequest{
		Command: "pwd",
		Cwd:     "/tmp",
	})
	if err != nil {
		log.Printf("Failed to execute command: %v\n", err)
	} else {
		fmt.Printf("Current directory: %s\n", result.Result)
	}

	// Command with timeout
	fmt.Println("\nExecuting command with timeout...")
	result, err = client.ExecuteCommand(ctx, sandboxID, &daytona.ExecuteCommandRequest{
		Command: "ls -la /",
		Timeout: 5.0, // 5 seconds
	})
	if err != nil {
		log.Printf("Failed to execute command: %v\n", err)
	} else {
		fmt.Printf("Directory listing:\n%s\n", result.Result)
	}

	// Python script execution
	fmt.Println("\nExecuting Python script...")
	pythonCode := `
import json
import sys

data = {"message": "Hello from Python", "status": "success"}
print(json.dumps(data, indent=2))
sys.exit(0)
`
	
	// First, create the Python file
	err = client.WriteFile(ctx, sandboxID, "/tmp/test.py", []byte(pythonCode))
	if err != nil {
		log.Printf("Failed to write Python file: %v\n", err)
	} else {
		// Execute the Python script
		result, err = client.ExecuteCommand(ctx, sandboxID, &daytona.ExecuteCommandRequest{
			Command: "python3 /tmp/test.py",
		})
		if err != nil {
			log.Printf("Failed to execute Python script: %v\n", err)
		} else {
			fmt.Printf("Python output:\n%s\n", result.Result)
		}
	}

	// Session-based commands
	fmt.Println("\nCreating session for interactive commands...")
	sessionID := "test-session-123"
	err = client.CreateSession(ctx, sandboxID, sessionID)
	if err != nil {
		log.Printf("Failed to create session: %v\n", err)
	} else {
		fmt.Println("✓ Session created")

		// Execute commands in session
		fmt.Println("\nExecuting commands in session...")
		
		// Set environment variable
		sessResult, err := client.ExecuteSessionCommand(ctx, sandboxID, sessionID, "export TEST_VAR='Hello from session'")
		if err != nil {
			log.Printf("Failed to set env var: %v\n", err)
		}

		// Use the environment variable
		sessResult, err = client.ExecuteSessionCommand(ctx, sandboxID, sessionID, "echo $TEST_VAR")
		if err != nil {
			log.Printf("Failed to execute session command: %v\n", err)
		} else if sessResult.Output != nil {
			fmt.Printf("Session output: %s\n", *sessResult.Output)
		}

		// Clean up session
		err = client.DeleteSession(ctx, sandboxID, sessionID)
		if err != nil {
			log.Printf("Failed to delete session: %v\n", err)
		} else {
			fmt.Println("✓ Session deleted")
		}
	}

	// Complex command example
	fmt.Println("\nExecuting complex command pipeline...")
	result, err = client.ExecuteCommand(ctx, sandboxID, &daytona.ExecuteCommandRequest{
		Command: "echo 'Line 1\nLine 2\nLine 3' | grep 'Line' | wc -l",
	})
	if err != nil {
		log.Printf("Failed to execute pipeline: %v\n", err)
	} else {
		fmt.Printf("Number of lines: %s", result.Result)
	}

	fmt.Println("\n=== Example Complete ===")
}