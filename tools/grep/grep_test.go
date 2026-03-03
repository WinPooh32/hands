package grep_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/WinPooh32/hands/pkg/testutil"
	"github.com/WinPooh32/hands/tools/grep"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGrep(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	require.NoError(t, os.WriteFile(testFile, []byte("Hello, World!\nThis is a test file.\nLine 3 with test"), 0600))

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolGrep := &mcp.Tool{
		Name:        "Grep",
		Description: "Search content in files",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"pattern": map[string]any{"type": "string"}, "path": map[string]any{"type": "string"}, "ignoreCase": map[string]any{"type": "boolean"}},
			"required":   []any{"pattern", "path"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolGrep, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		pattern, _ := params["pattern"].(string)
		path, _ := params["path"].(string)
		ignoreCase, _ := params["ignoreCase"].(bool)

		result, _, err := grep.Grep(ctx, request, grep.GrepInput{Pattern: pattern, Path: path, IgnoreCase: ignoreCase})

		if err != nil {
			return result, fmt.Errorf("grep.Grep: %w", err)
		}
		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Grep",
		Arguments: map[string]any{"pattern": "test", "path": testFile},
		Meta:      nil,
	})
	require.NoError(t, err)

	assert.False(t, result.IsError)
	assert.Len(t, result.Content, 1)
	textContent := result.Content[0].(*mcp.TextContent)
	assert.Contains(t, textContent.Text, "Found 2 matches:")
	assert.Contains(t, textContent.Text, "test.txt:2: This is a test file.")
	assert.Contains(t, textContent.Text, "test.txt:3: Line 3 with test")
}

func TestGrep_IgnoreCase(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	require.NoError(t, os.WriteFile(testFile, []byte("HELLO, WORLD!\nhello world\nTest"), 0600))

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolGrep := &mcp.Tool{
		Name:        "Grep",
		Description: "Search content in files",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"pattern": map[string]any{"type": "string"}, "path": map[string]any{"type": "string"}, "ignoreCase": map[string]any{"type": "boolean"}},
			"required":   []any{"pattern", "path"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolGrep, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		pattern, _ := params["pattern"].(string)
		path, _ := params["path"].(string)
		ignoreCase, _ := params["ignoreCase"].(bool)
		result, _, err := grep.Grep(ctx, request, grep.GrepInput{Pattern: pattern, Path: path, IgnoreCase: ignoreCase})

		if err != nil {
			return result, fmt.Errorf("grep.Grep: %w", err)
		}
		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Grep",
		Arguments: map[string]any{"pattern": "hello", "path": testFile, "ignoreCase": true},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolSuccess(t, result, "Found 2 matches:")
}

func TestGrep_NoMatches(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	require.NoError(t, os.WriteFile(testFile, []byte("Hello, World!\nThis is a test file."), 0600))

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolGrep := &mcp.Tool{
		Name:        "Grep",
		Description: "Search content in files",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"pattern": map[string]any{"type": "string"}, "path": map[string]any{"type": "string"}, "ignoreCase": map[string]any{"type": "boolean"}},
			"required":   []any{"pattern", "path"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolGrep, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		pattern, _ := params["pattern"].(string)
		path, _ := params["path"].(string)
		ignoreCase, _ := params["ignoreCase"].(bool)
		result, _, err := grep.Grep(ctx, request, grep.GrepInput{Pattern: pattern, Path: path, IgnoreCase: ignoreCase})

		if err != nil {
			return result, fmt.Errorf("grep.Grep: %w", err)
		}
		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Grep",
		Arguments: map[string]any{"pattern": "xyz123notfound", "path": testFile},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolSuccess(t, result, "No matches found")
}

func TestGrep_NoPattern(t *testing.T) {
	t.Parallel()

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolGrep := &mcp.Tool{
		Name:        "Grep",
		Description: "Search content in files",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"pattern": map[string]any{"type": "string"}, "path": map[string]any{"type": "string"}, "ignoreCase": map[string]any{"type": "boolean"}},
			"required":   []any{"pattern", "path"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolGrep, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		pattern, _ := params["pattern"].(string)
		path, _ := params["path"].(string)
		ignoreCase, _ := params["ignoreCase"].(bool)
		result, _, err := grep.Grep(ctx, request, grep.GrepInput{Pattern: pattern, Path: path, IgnoreCase: ignoreCase})

		if err != nil {
			return result, fmt.Errorf("grep.Grep: %w", err)
		}
		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Grep",
		Arguments: map[string]any{"path": "/tmp/test.txt"},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolError(t, result, "pattern is required")
}
