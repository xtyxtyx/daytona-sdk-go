# Daytona Go SDK (Unofficial)

⚠️ **UNOFFICIAL SDK - IN DEVELOPMENT** ⚠️

This is an unofficial Go SDK for the [Daytona](https://www.daytona.io) platform. It is currently in active development and not affiliated with or endorsed by Daytona Platforms Inc.

## Overview

The Daytona Go SDK provides a Go interface for interacting with the Daytona API, enabling developers to manage sandboxes, execute commands, handle files, and more programmatically.

## Features

- 🏗️ **Sandbox Management** - Create, start, stop, archive, and delete sandboxes
- 📁 **File Operations** - Upload, download, create folders, and manage files
- 💻 **Command Execution** - Run commands and manage sessions
- 🔧 **Custom Images** - Build custom sandbox environments with pre-installed packages
- 🐍 **PySpark Support** - Examples for data processing with PySpark
- 📊 **Data Generation** - Generate and analyze CSV data

## Installation

```bash
go get github.com/daytonaio/daytona-sdk-go
```

The SDK includes the Daytona API client bundled within the `apiclient` directory, so no additional dependencies are needed.

### Updating the API Client

To update the bundled API client to the latest version from the Daytona repository:

```bash
./scripts/update-apiclient.sh
```

This script will fetch the latest API client from GitHub and update the local copy.

## Configuration

The SDK automatically loads environment variables from `.env` files using [godotenv](https://github.com/joho/godotenv). 

Create a `.env` file in your project root:

```bash
DAYTONA_API_KEY=your_api_key_here
DAYTONA_API_URL=https://app.daytona.io/api  # Optional, defaults to this URL
```

The SDK will automatically:
- Load `.env` from the current directory
- Load `../.env` from the parent directory (useful for examples)
- Use system environment variables if no `.env` file exists

No need to manually source the `.env` file or use `export` commands!

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    daytona "github.com/daytonaio/daytona-sdk-go"
)

func main() {
    // Create client (automatically loads .env file)
    client, err := daytona.NewClient(&daytona.Config{})
    if err != nil {
        log.Fatal(err)
    }
    
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
    
    // Execute a command
    result, err := client.ExecuteCommand(ctx, sandbox.GetId(), &daytona.ExecuteCommandRequest{
        Command: "echo 'Hello from Daytona!'",
    })
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Output: %s\n", result.Result)
    
    // Clean up
    err = client.DeleteSandbox(ctx, sandbox.GetId(), true)
    if err != nil {
        log.Fatal(err)
    }
}
```

## Examples

The SDK includes comprehensive examples in the `examples/` directory. All examples automatically load the `.env` file from the project root.

### Basic Operations
```bash
# No need to source .env - it's loaded automatically!
go run examples/basic/main.go
```
Creates a sandbox, executes commands, manages files, and cleans up.

### File Operations
```bash
go run examples/file_operations/main.go
```
Demonstrates file creation, reading, writing, moving, and deletion.

### CSV Data Generation
```bash
go run examples/csv_generation/main.go
```
Generates mock sales data, analyzes it, and exports results.

### PySpark with Declarative Builder
```bash
go run examples/pyspark/main.go
```
Creates a custom image with PySpark pre-installed using the declarative builder pattern.

### Comprehensive Example
```bash
go run examples/comprehensive/main.go
```
Demonstrates all SDK features including sessions, search operations, and lifecycle management.

## Custom Images

Build custom sandbox environments with pre-installed packages:

```go
// Build a custom Dockerfile
dockerfile := `FROM python:3.11-slim
RUN pip install pandas numpy scikit-learn
WORKDIR /workspace`

sandbox, err := client.CreateSandbox(ctx, &daytona.CreateSandboxRequest{
    DockerfileContent: daytona.StringPtr(dockerfile),
    // ... other options
})
```

## API Reference

### Client Creation
```go
client, err := daytona.NewClient(&daytona.Config{
    APIKey:  "your_api_key",     // Optional, uses DAYTONA_API_KEY env var
    BaseURL: "https://...",      // Optional, uses DAYTONA_API_URL env var
    Timeout: 30 * time.Second,   // Optional, defaults to 30s
})
```

### Sandbox Operations
- `CreateSandbox(ctx, request)` - Create a new sandbox
- `GetSandbox(ctx, sandboxID)` - Get sandbox details
- `ListSandboxes(ctx)` - List all sandboxes
- `StartSandbox(ctx, sandboxID)` - Start a sandbox
- `StopSandbox(ctx, sandboxID)` - Stop a sandbox
- `ArchiveSandbox(ctx, sandboxID)` - Archive a sandbox
- `DeleteSandbox(ctx, sandboxID, force)` - Delete a sandbox
- `WaitForSandboxReady(ctx, sandboxID, timeout)` - Wait for sandbox to be ready

### File Operations
- `WriteFile(ctx, sandboxID, path, content)` - Write a file
- `ReadFile(ctx, sandboxID, path)` - Read a file
- `CreateFolder(ctx, sandboxID, path)` - Create a directory
- `DeleteFile(ctx, sandboxID, path)` - Delete a file or directory
- `MoveFile(ctx, sandboxID, source, dest)` - Move/rename a file
- `ListFiles(ctx, sandboxID, path)` - List files in a directory
- `GetFileInfo(ctx, sandboxID, path)` - Get file information
- `FindInFiles(ctx, sandboxID, pattern, path)` - Search in files

### Command Execution
- `ExecuteCommand(ctx, sandboxID, request)` - Execute a command
- `CreateSession(ctx, sandboxID, sessionID)` - Create a session
- `ExecuteSessionCommand(ctx, sandboxID, sessionID, command)` - Execute in session
- `DeleteSession(ctx, sandboxID, sessionID)` - Delete a session

## Development Status

This SDK is under active development. Current limitations:

- Not all Daytona API endpoints are implemented
- API may change as the SDK evolves
- Limited testing coverage
- No official support from Daytona

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Disclaimer

This is an unofficial SDK and is not affiliated with, endorsed by, or supported by Daytona Platforms Inc. Use at your own risk.

## Resources

- [Official Daytona Documentation](https://www.daytona.io/docs)
- [Daytona API Reference](https://www.daytona.io/docs/api)
- [Official TypeScript SDK](https://github.com/daytonaio/daytona/tree/main/libs/sdk-typescript)