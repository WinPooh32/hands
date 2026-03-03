package bash_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/WinPooh32/hands/pkg/testutil"
	"github.com/WinPooh32/hands/tools/bash"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/require"
)

func TestBash(t *testing.T) {
	t.Parallel()

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolBash := &mcp.Tool{
		Name:        "Bash",
		Description: "Execute bash command",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"workingDir": map[string]any{"type": "string"}, "command": map[string]any{"type": "string"}},
			"required":   []any{"command"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolBash, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		command, _ := params["command"].(string)

		workingDir, _ := params["workingDir"].(string)
		if workingDir == "" {
			workingDir = "."
		}

		result, _, err := bash.Bash(ctx, request, bash.BashInput{Command: command, WorkingDir: workingDir})

		if err != nil {
			return result, fmt.Errorf("bash.Bash: %w", err)
		}

		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Bash",
		Arguments: map[string]any{"command": "echo Hello, World!"},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolSuccess(t, result, "Hello, World!")
}

func TestBash_WorkingDir(t *testing.T) {
	t.Parallel()

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolBash := &mcp.Tool{
		Name:        "Bash",
		Description: "Execute bash command",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"workingDir": map[string]any{"type": "string"}, "command": map[string]any{"type": "string"}},
			"required":   []any{"command"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolBash, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		command, _ := params["command"].(string)

		workingDir, _ := params["workingDir"].(string)
		if workingDir == "" {
			workingDir = "."
		}

		result, _, err := bash.Bash(ctx, request, bash.BashInput{Command: command, WorkingDir: workingDir})

		if err != nil {
			return result, fmt.Errorf("bash.Bash: %w", err)
		}

		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Bash",
		Arguments: map[string]any{"command": "pwd", "workingDir": "/tmp"},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolSuccess(t, result, "/tmp")
}

func TestBash_FailCommand(t *testing.T) {
	t.Parallel()

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolBash := &mcp.Tool{
		Name:        "Bash",
		Description: "Execute bash command",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"workingDir": map[string]any{"type": "string"}, "command": map[string]any{"type": "string"}},
			"required":   []any{"command"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolBash, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		command, _ := params["command"].(string)

		workingDir, _ := params["workingDir"].(string)
		if workingDir == "" {
			workingDir = "."
		}

		result, _, err := bash.Bash(ctx, request, bash.BashInput{Command: command, WorkingDir: workingDir})

		if err != nil {
			return result, fmt.Errorf("bash.Bash: %w", err)
		}

		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Bash",
		Arguments: map[string]any{"command": "exit 1"},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolError(t, result, "command failed")
}

func TestBash_NoCommand(t *testing.T) {
	t.Parallel()

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolBash := &mcp.Tool{
		Name:        "Bash",
		Description: "Execute bash command",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"workingDir": map[string]any{"type": "string"}, "command": map[string]any{"type": "string"}},
			"required":   []any{"command"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolBash, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		command, _ := params["command"].(string)

		workingDir, _ := params["workingDir"].(string)
		if workingDir == "" {
			workingDir = "."
		}

		result, _, err := bash.Bash(ctx, request, bash.BashInput{Command: command, WorkingDir: workingDir})

		if err != nil {
			return result, fmt.Errorf("bash.Bash: %w", err)
		}

		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Bash",
		Arguments: map[string]any{"workingDir": "/tmp"},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolError(t, result, "command is required")
}
