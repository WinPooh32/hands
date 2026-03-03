package write_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/WinPooh32/hands/pkg/testutil"
	"github.com/WinPooh32/hands/tools/write"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrite(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolWrite := &mcp.Tool{
		Name:        "Write",
		Description: "Write file",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"path": map[string]any{"type": "string"}, "content": map[string]any{"type": "string"}},
			"required":   []any{"path", "content"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolWrite, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		path, _ := params["path"].(string)
		content, _ := params["content"].(string)

		result, _, err := write.Write(ctx, request, write.WriteInput{Path: path, Content: content})

		if err != nil {
			return result, fmt.Errorf("write.Write: %w", err)
		}

		return result, nil
	})

	testFile := filepath.Join(tmpDir, "test.txt")
	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Write",
		Arguments: map[string]any{"path": testFile, "content": "Hello, World!"},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolSuccess(t, result, "Successfully wrote")

	content, err := os.ReadFile(testFile)
	require.NoError(t, err)
	assert.Equal(t, "Hello, World!", string(content))
}

func TestWrite_NoPath(t *testing.T) {
	t.Parallel()

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolWrite := &mcp.Tool{
		Name:        "Write",
		Description: "Write file",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"path": map[string]any{"type": "string"}, "content": map[string]any{"type": "string"}},
			"required":   []any{"path", "content"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolWrite, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		path, _ := params["path"].(string)
		content, _ := params["content"].(string)
		result, _, err := write.Write(ctx, request, write.WriteInput{Path: path, Content: content})

		if err != nil {
			return result, fmt.Errorf("write.Write: %w", err)
		}

		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Write",
		Arguments: map[string]any{"content": "test"},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolError(t, result, "path is required")
}

func TestWrite_NoContent(t *testing.T) {
	t.Parallel()

	_, server, clientSession, serverSession := testutil.CreateSessionPair(t)
	t.Cleanup(func() {
		_ = clientSession.Close()
		_ = serverSession.Close()
	})

	mcpToolWrite := &mcp.Tool{
		Name:        "Write",
		Description: "Write file",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"path": map[string]any{"type": "string"}, "content": map[string]any{"type": "string"}},
			"required":   []any{"path", "content"},
		},
		Meta:         nil,
		Annotations:  nil,
		OutputSchema: nil,
		Title:        "",
		Icons:        nil,
	}
	server.AddTool(mcpToolWrite, func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var params map[string]any
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return mcputil.ErrorResult("invalid arguments"), nil
		}

		path, _ := params["path"].(string)
		content, _ := params["content"].(string)
		result, _, err := write.Write(ctx, request, write.WriteInput{Path: path, Content: content})

		if err != nil {
			return result, fmt.Errorf("write.Write: %w", err)
		}

		return result, nil
	})

	result, err := clientSession.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "Write",
		Arguments: map[string]any{"path": "/tmp/test.txt"},
		Meta:      nil,
	})
	require.NoError(t, err)

	testutil.AssertToolError(t, result, "content is required")
}
