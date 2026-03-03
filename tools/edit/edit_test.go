package edit_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/WinPooh32/hands/pkg/testutil"
	"github.com/WinPooh32/hands/tools/edit"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEdit(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	initialContent := "Hello, World!\nThis is a test file.\nLine 3"
	require.NoError(t, os.WriteFile(testFile, []byte(initialContent), 0600))

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolEdit := &mcp.Tool{
		Name:        "Edit",
		Description: "Edit file",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"path": map[string]any{"type": "string"}, "search": map[string]any{"type": "string"}, "replace": map[string]any{"type": "string"}},
			"required":   []any{"path", "search", "replace"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolEdit, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		path, _ := params["path"].(string)
		search, _ := params["search"].(string)
		replace, _ := params["replace"].(string)

		result, _, err := edit.Edit(ctx, request, edit.EditInput{Path: path, Search: search, Replace: replace})

		if err != nil {
			return result, fmt.Errorf("edit.Edit: %w", err)
		}

		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Edit",
		Arguments: map[string]any{"path": testFile, "search": "World", "replace": "Universe"},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolSuccess(t, result, "Successfully edited")

	content, err := os.ReadFile(testFile)
	require.NoError(t, err)
	assert.Equal(t, "Hello, Universe!\nThis is a test file.\nLine 3", string(content))
}

func TestEdit_SearchNotFound(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	initialContent := "Hello, World!\nThis is a test file."
	require.NoError(t, os.WriteFile(testFile, []byte(initialContent), 0600))

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolEdit := &mcp.Tool{
		Name:        "Edit",
		Description: "Edit file",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"path": map[string]any{"type": "string"}, "search": map[string]any{"type": "string"}, "replace": map[string]any{"type": "string"}},
			"required":   []any{"path", "search", "replace"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolEdit, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		path, _ := params["path"].(string)
		search, _ := params["search"].(string)
		replace, _ := params["replace"].(string)
		result, _, err := edit.Edit(ctx, request, edit.EditInput{Path: path, Search: search, Replace: replace})

		if err != nil {
			return result, fmt.Errorf("edit.Edit: %w", err)
		}

		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Edit",
		Arguments: map[string]any{"path": testFile, "search": "not found", "replace": "replacement"},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolError(t, result, "search string 'not found' not found in file")
}

func TestEdit_NoSearch(t *testing.T) {
	t.Parallel()

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolEdit := &mcp.Tool{
		Name:        "Edit",
		Description: "Edit file",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"path": map[string]any{"type": "string"}, "search": map[string]any{"type": "string"}, "replace": map[string]any{"type": "string"}},
			"required":   []any{"path", "search", "replace"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolEdit, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		path, _ := params["path"].(string)
		search, _ := params["search"].(string)
		replace, _ := params["replace"].(string)
		result, _, err := edit.Edit(ctx, request, edit.EditInput{Path: path, Search: search, Replace: replace})

		if err != nil {
			return result, fmt.Errorf("edit.Edit: %w", err)
		}

		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Edit",
		Arguments: map[string]any{"path": "/tmp/test.txt", "replace": "replacement"},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolError(t, result, "search string is required")
}
