package read

import (
	"context"
	"fmt"
	"os"

	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ReadInput struct {
	Path string `json:"path" jsonschema:"path to the file to read"`
}

func Read(ctx context.Context, _ *mcp.CallToolRequest, input ReadInput) (*mcp.CallToolResult, any, error) {
	if input.Path == "" {
		return mcputil.ErrorResult("path is required"), nil, nil
	}

	data, err := os.ReadFile(input.Path)
	if err != nil {
		return mcputil.ErrorResult(fmt.Sprintf("failed to read file: %v", err)), nil, nil
	}

	return mcputil.TextResult(string(data)), nil, nil
}
