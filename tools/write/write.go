package write

import (
	"context"
	"fmt"
	"os"

	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const fileMode os.FileMode = 0600

type WriteInput struct {
	Path    string `json:"path" jsonschema:"path to the file to write"`
	Content string `json:"content" jsonschema:"content to write"`
}

func Write(ctx context.Context, _ *mcp.CallToolRequest, input WriteInput) (*mcp.CallToolResult, any, error) {
	if input.Path == "" {
		return mcputil.ErrorResult("path is required"), nil, nil
	}

	if input.Content == "" {
		return mcputil.ErrorResult("content is required"), nil, nil
	}

	err := os.WriteFile(input.Path, []byte(input.Content), fileMode)
	if err != nil {
		return mcputil.ErrorResult(fmt.Sprintf("failed to write file: %v", err)), nil, nil
	}

	return mcputil.TextResult(fmt.Sprintf("Successfully wrote to %s", input.Path)), nil, nil
}
