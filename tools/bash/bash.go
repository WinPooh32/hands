package bash

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type BashInput struct {
	WorkingDir string `json:"workingDir" jsonschema:"working directory (default: current directory)"`
	Command    string `json:"command" jsonschema:"bash command to execute"`
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

	// Any user inputs is allowed to exec.
	//nolint:gosec
	cmd := exec.CommandContext(ctx, "/bin/bash", "-c", input.Command)
	cmd.Dir = workingDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return mcputil.ErrorResult(fmt.Sprintf("command failed: %v\nOutput: %s", err, strings.TrimSpace(string(output)))), nil, nil
	}

	return mcputil.TextResult(strings.TrimSpace(string(output))), nil, nil
}
