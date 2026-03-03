package read_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/WinPooh32/hands/pkg/testutil"
	"github.com/WinPooh32/hands/tools/read"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	content := "Hello, World!\nThis is a test file.\nLine 3"
	require.NoError(t, os.WriteFile(tmpFile, []byte(content), 0600))

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	handler := func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		path, ok := params["path"].(string)
		if !ok {
			return mcputil.ErrorResult("path must be string"), nil
		}

		result, _, err := read.Read(ctx, request, read.ReadInput{Path: path})

		if err != nil {
			return result, fmt.Errorf("read.Read: %w", err)
		}
		return result, nil
	}

	mcpToolRead := &mcp.Tool{
		Name:        "Read",
		Description: "Read file",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"path": map[string]any{"type": "string"}},
			"required":   []any{"path"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolRead, handler)

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Read",
		Arguments: map[string]any{"path": tmpFile},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolSuccess(t, result, "Hello, World!")
	assert.Contains(t, result.Content[0].(*mcp.TextContent).Text, "Line 3")
}

func TestRead_FileNotFound(t *testing.T) {
	t.Parallel()

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolRead := &mcp.Tool{
		Name:        "Read",
		Description: "Read file",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"path": map[string]any{"type": "string"}},
			"required":   []any{"path"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolRead, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		path, ok := params["path"].(string)
		if !ok {
			return mcputil.ErrorResult("path must be string"), nil
		}

		result, _, err := read.Read(ctx, request, read.ReadInput{Path: path})

		if err != nil {
			return result, fmt.Errorf("read.Read: %w", err)
		}
		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Read",
		Arguments: map[string]any{"path": "/nonexistent/file.txt"},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolError(t, result, "open file")
}

func TestRead_EmptyPath(t *testing.T) {
	t.Parallel()

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolRead := &mcp.Tool{
		Name:        "Read",
		Description: "Read file",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"path": map[string]any{"type": "string"}},
			"required":   []any{"path"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolRead, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		path, ok := params["path"].(string)
		if !ok {
			return mcputil.ErrorResult("path must be string"), nil
		}

		result, _, err := read.Read(ctx, request, read.ReadInput{Path: path})

		if err != nil {
			return result, fmt.Errorf("read.Read: %w", err)
		}
		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Read",
		Arguments: map[string]any{"path": ""},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolError(t, result, "path is required")
}
