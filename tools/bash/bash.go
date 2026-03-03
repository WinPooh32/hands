package bash

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/WinPooh32/hands/pkg/i18n"
	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type BashInput struct {
	WorkingDir string `json:"workingDir,omitzero"`
	Command    string `json:"command"`
}

func Bash(ctx context.Context, _ *mcp.CallToolRequest, input BashInput) (*mcp.CallToolResult, any, error) {
	if input.Command == "" {
		return mcputil.ErrorResult("command is required"), nil, nil
	}

	workingDir := input.WorkingDir
	if workingDir == "" {
		workingDir = "."
	}

	info, err := os.Stat(workingDir)
	if err != nil {
		return mcputil.ErrorResult(fmt.Sprintf("working directory does not exist: %v", err)), nil, nil
	}

	if !info.IsDir() {
		return mcputil.ErrorResult("working directory is not a directory"), nil, nil
	}

	cmd := exec.CommandContext(ctx, "/bin/bash", "-c", input.Command)
	cmd.Dir = workingDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return mcputil.ErrorResult(fmt.Sprintf("command failed: %v\nOutput: %s", err, strings.TrimSpace(string(output)))), nil, nil
	}

	return mcputil.TextResult(strings.TrimSpace(string(output))), nil, nil
}

// Schema returns the JSON schema for BashInput with translated descriptions
func Schema() any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"command": map[string]any{
				"description": i18n.Tr(i18n.BashArgCommand),
			},
			"workingDir": map[string]any{
				"description": i18n.Tr(i18n.BashArgWorkingDir),
			},
		},
		"required": []string{"command"},
	}
}
