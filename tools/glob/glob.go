package glob

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/WinPooh32/hands/pkg/i18n"
	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type GlobInput struct {
	Pattern string `json:"pattern"`
	Dir     string `json:"dir,omitzero"`
}

func Glob(ctx context.Context, _ *mcp.CallToolRequest, input GlobInput) (*mcp.CallToolResult, any, error) {
	if input.Pattern == "" {
		return mcputil.ErrorResult("pattern is required"), nil, nil
	}

	dir := input.Dir
	if dir == "" {
		dir = "."
	}

	absDir, err := filepath.Abs(dir)
	if err != nil {
		return mcputil.ErrorResult(fmt.Sprintf("failed to get absolute path: %v", err)), nil, nil
	}

	var matches []string

	err = filepath.Walk(absDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories that match the pattern (we use Match on the basename)
		matched, _ := filepath.Match(input.Pattern, info.Name())
		if matched {
			matches = append(matches, path)
		}

		return nil
	})
	if err != nil {
		return mcputil.ErrorResult(fmt.Sprintf("failed to search: %v", err)), nil, nil
	}

	result := fmt.Sprintf("Found %d files:", len(matches))
	for _, m := range matches {
		result += fmt.Sprintf("\n- %s", m)
	}

	return mcputil.TextResult(result), nil, nil
}

// Schema returns the JSON schema for GlobInput with translated descriptions
func Schema() any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"pattern": map[string]any{
				"description": i18n.Tr(i18n.GlobArgPattern),
			},
			"dir": map[string]any{
				"description": i18n.Tr(i18n.GlobArgDir),
			},
		},
		"required": []string{"pattern"},
	}
}
