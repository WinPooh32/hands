package glob_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/WinPooh32/hands/pkg/testutil"
	"github.com/WinPooh32/hands/tools/glob"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGlob(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(tmpDir, "test1.go"), []byte("package main"), 0600))
	require.NoError(t, os.WriteFile(filepath.Join(tmpDir, "test2.go"), []byte("package main"), 0600))
	require.NoError(t, os.WriteFile(filepath.Join(tmpDir, "README.md"), []byte("# Test"), 0600))
	require.NoError(t, os.MkdirAll(filepath.Join(tmpDir, "subdir"), 0700))
	require.NoError(t, os.WriteFile(filepath.Join(tmpDir, "subdir", "test3.go"), []byte("package main"), 0600))

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolGlob := &mcp.Tool{
		Name:        "Glob",
		Description: "Find files by pattern",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"pattern": map[string]any{"type": "string"}, "dir": map[string]any{"type": "string"}},
			"required":   []any{"pattern"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolGlob, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		pattern, _ := params["pattern"].(string)

		dir, _ := params["dir"].(string)
		if dir == "" {
			dir = "."
		}

		result, _, err := glob.Glob(ctx, request, glob.GlobInput{Pattern: pattern, Dir: dir})

		if err != nil {
			return result, fmt.Errorf("glob.Glob: %w", err)
		}

		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Glob",
		Arguments: map[string]any{"pattern": "*.go", "dir": tmpDir},
		Meta:      nil,
	})
	require.NoError(t, err)

	assert.False(t, result.IsError)
	assert.Len(t, result.Content, 1)
	textContent := result.Content[0].(*mcp.TextContent)
	assert.Contains(t, textContent.Text, "Found 3 files:")
	assert.Contains(t, textContent.Text, "test1.go")
	assert.Contains(t, textContent.Text, "test2.go")
	assert.Contains(t, textContent.Text, "subdir/test3.go")
}

func TestGlob_NoPattern(t *testing.T) {
	t.Parallel()

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolGlob := &mcp.Tool{
		Name:        "Glob",
		Description: "Find files by pattern",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"pattern": map[string]any{"type": "string"}, "dir": map[string]any{"type": "string"}},
			"required":   []any{"pattern"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolGlob, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		pattern, _ := params["pattern"].(string)

		dir, _ := params["dir"].(string)
		if dir == "" {
			dir = "."
		}

		result, _, err := glob.Glob(ctx, request, glob.GlobInput{Pattern: pattern, Dir: dir})

		if err != nil {
			return result, fmt.Errorf("glob.Glob: %w", err)
		}

		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Glob",
		Arguments: map[string]any{"dir": "/tmp"},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolError(t, result, "pattern is required")
}

func TestGlob_CurrentDir(t *testing.T) {
	t.Parallel()

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolGlob := &mcp.Tool{
		Name:        "Glob",
		Description: "Find files by pattern",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"pattern": map[string]any{"type": "string"}, "dir": map[string]any{"type": "string"}},
			"required":   []any{"pattern"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolGlob, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		pattern, _ := params["pattern"].(string)

		dir, _ := params["dir"].(string)
		if dir == "" {
			dir = "."
		}

		result, _, err := glob.Glob(ctx, request, glob.GlobInput{Pattern: pattern, Dir: dir})

		if err != nil {
			return result, fmt.Errorf("glob.Glob: %w", err)
		}

		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Glob",
		Arguments: map[string]any{"pattern": "*.go"},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolSuccess(t, result, "Found")
}
