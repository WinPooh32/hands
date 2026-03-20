package write

import (
	"context"
	"fmt"
	"os"

	"github.com/WinPooh32/hands/pkg/i18n"
	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const fileMode os.FileMode = 0600

type WriteInput struct {
	Path    string `json:"path"`
	Content string `json:"content"`
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

// Schema returns the JSON schema for WriteInput with translated descriptions
func Schema() any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"path": map[string]any{
				"description": i18n.Tr(i18n.WriteArgPath),
				"type":        "string",
			},
			"content": map[string]any{
				"description": i18n.Tr(i18n.WriteArgContent),
				"type":        "string",
			},
		},
		"required": []string{"path", "content"},
	}
}
