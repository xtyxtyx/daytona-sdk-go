package daytona

import (
	"context"
	"fmt"
	"io"
	"os"

	api "github.com/PhilippBuschhaus/daytona-sdk-go/apiclient"
)

// ReadFile reads a file from a sandbox
func (c *Client) ReadFile(ctx context.Context, sandboxID, path string) ([]byte, error) {
	if sandboxID == "" || path == "" {
		return nil, fmt.Errorf("sandbox ID and path are required")
	}
	
	file, _, err := c.apiClient.ToolboxAPI.DownloadFile(ctx, sandboxID).Path(path).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	defer file.Close()
	
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %w", err)
	}
	
	return content, nil
}

// WriteFile writes content to a file in the sandbox
func (c *Client) WriteFile(ctx context.Context, sandboxID, path string, content []byte) error {
	if sandboxID == "" || path == "" {
		return fmt.Errorf("sandbox ID and path are required")
	}
	
	// Create a temporary file with the content
	tmpFile, err := os.CreateTemp("", "upload-*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()
	
	if _, err := tmpFile.Write(content); err != nil {
		return fmt.Errorf("failed to write temp file: %w", err)
	}
	
	// Reset file pointer for reading
	if _, err := tmpFile.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek temp file: %w", err)
	}
	
	// Upload the file
	_, err = c.apiClient.ToolboxAPI.UploadFile(ctx, sandboxID).
		Path(path).
		File(tmpFile).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	
	return nil
}

// WriteFiles writes multiple files to the sandbox
func (c *Client) WriteFiles(ctx context.Context, sandboxID string, files map[string]io.Reader) error {
	if sandboxID == "" {
		return fmt.Errorf("sandbox ID is required")
	}

	if len(files) == 0 {
		return fmt.Errorf("files map cannot be empty")
	}

	// Upload the files using the bulk upload API
	_, err := c.apiClient.ToolboxAPI.UploadFiles(ctx, sandboxID).
		Files(files).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to upload files: %w", err)
	}

	return nil
}

// CreateFolder creates a directory in the sandbox
func (c *Client) CreateFolder(ctx context.Context, sandboxID, path string) error {
	if sandboxID == "" || path == "" {
		return fmt.Errorf("sandbox ID and path are required")
	}
	
	_, err := c.apiClient.ToolboxAPI.CreateFolder(ctx, sandboxID).
		Path(path).
		Mode("0755"). // Default directory permissions
		Execute()
	if err != nil {
		return fmt.Errorf("failed to create folder: %w", err)
	}
	
	return nil
}

// DeleteFile deletes a file or directory in the sandbox
func (c *Client) DeleteFile(ctx context.Context, sandboxID, path string) error {
	if sandboxID == "" || path == "" {
		return fmt.Errorf("sandbox ID and path are required")
	}
	
	_, err := c.apiClient.ToolboxAPI.DeleteFile(ctx, sandboxID).Path(path).Execute()
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	
	return nil
}

// GetFileInfo gets information about a file or directory
func (c *Client) GetFileInfo(ctx context.Context, sandboxID, path string) (*api.FileInfo, error) {
	if sandboxID == "" || path == "" {
		return nil, fmt.Errorf("sandbox ID and path are required")
	}
	
	info, _, err := c.apiClient.ToolboxAPI.GetFileInfo(ctx, sandboxID).Path(path).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}
	
	return info, nil
}

// MoveFile moves or renames a file
func (c *Client) MoveFile(ctx context.Context, sandboxID, source, destination string) error {
	if sandboxID == "" || source == "" || destination == "" {
		return fmt.Errorf("sandbox ID, source, and destination are required")
	}
	
	_, err := c.apiClient.ToolboxAPI.MoveFile(ctx, sandboxID).
		Source(source).
		Destination(destination).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to move file: %w", err)
	}
	
	return nil
}

// ListFiles lists files in a directory
func (c *Client) ListFiles(ctx context.Context, sandboxID, path string) ([]api.FileInfo, error) {
	if sandboxID == "" {
		return nil, fmt.Errorf("sandbox ID is required")
	}
	
	if path == "" {
		path = "/"
	}
	
	files, _, err := c.apiClient.ToolboxAPI.ListFiles(ctx, sandboxID).Path(path).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to list files: %w", err)
	}
	
	return files, nil
}

// FindInFiles searches for text in files
func (c *Client) FindInFiles(ctx context.Context, sandboxID, pattern, path string) ([]api.Match, error) {
	if sandboxID == "" || pattern == "" {
		return nil, fmt.Errorf("sandbox ID and pattern are required")
	}
	
	req := c.apiClient.ToolboxAPI.FindInFiles(ctx, sandboxID).Pattern(pattern)
	if path != "" {
		req = req.Path(path)
	}
	
	matches, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to find in files: %w", err)
	}
	
	return matches, nil
}

// ReplaceInFiles replaces text in files
func (c *Client) ReplaceInFiles(ctx context.Context, sandboxID string, replaceReq api.ReplaceRequest) ([]api.ReplaceResult, error) {
	if sandboxID == "" {
		return nil, fmt.Errorf("sandbox ID is required")
	}
	
	results, _, err := c.apiClient.ToolboxAPI.ReplaceInFiles(ctx, sandboxID).
		ReplaceRequest(replaceReq).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to replace in files: %w", err)
	}
	
	return results, nil
}