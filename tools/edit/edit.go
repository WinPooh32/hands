package edit

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/WinPooh32/hands/pkg/i18n"
	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const fileMode os.FileMode = 0600

type EditInput struct {
	Path    string `json:"path"`
	Search  string `json:"search"`
	Replace string `json:"replace"`
}

func Edit(ctx context.Context, _ *mcp.CallToolRequest, input EditInput) (*mcp.CallToolResult, any, error) {
	if input.Path == "" {
		return mcputil.ErrorResult("path is required"), nil, nil
	}

	if input.Search == "" {
		return mcputil.ErrorResult("search string is required"), nil, nil
	}

	data, err := os.ReadFile(input.Path)
	if err != nil {
		return mcputil.ErrorResult(fmt.Sprintf("failed to read file: %v", err)), nil, nil
	}

	content := string(data)

	if !strings.Contains(content, input.Search) {
		return mcputil.ErrorResult(fmt.Sprintf("search string '%s' not found in file", input.Search)), nil, nil
	}

	newContent := strings.ReplaceAll(content, input.Search, input.Replace)

	err = os.WriteFile(input.Path, []byte(newContent), fileMode)
	if err != nil {
		return mcputil.ErrorResult(fmt.Sprintf("failed to write file: %v", err)), nil, nil
	}

	return mcputil.TextResult(fmt.Sprintf("Successfully edited %s", input.Path)), nil, nil
}

// Schema returns the JSON schema for EditInput with translated descriptions
func Schema() any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"path": map[string]any{
				"description": i18n.Tr(i18n.EditArgPath),
			},
			"search": map[string]any{
				"description": i18n.Tr(i18n.EditArgSearch),
			},
			"replace": map[string]any{
				"description": i18n.Tr(i18n.EditArgReplace),
			},
		},
		"required": []string{"path", "search", "replace"},
	}
}
