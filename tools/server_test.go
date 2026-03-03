package tools_test

import (
	"context"
	"testing"

	"github.com/WinPooh32/hands/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServerRunStdio(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := tools.NewServer()
	kit := tools.New(nil)

	err := srv.Run(ctx, &tools.ServerConfig{HTTPAddr: "", StdioMode: true, HTTPMode: false}, kit)
	require.NoError(t, err)
}

func TestNewDefaultKit(t *testing.T) {
	t.Parallel()

	kit := tools.NewDefault(nil)
	assert.NotNil(t, kit, "NewDefault should return a non-nil kit")
}

func TestNewEmptyKit(t *testing.T) {
	t.Parallel()

	kit := tools.New(nil)
	assert.NotNil(t, kit, "New should return a non-nil kit")
}

func TestAddTool(t *testing.T) {
	t.Parallel()

	kit := tools.New(nil)

	// Test with empty tool name
	err := kit.Add(mcp.Tool{Name: "", Description: "", Meta: nil, Annotations: nil, InputSchema: nil, OutputSchema: nil, Title: "", Icons: nil})
	assert.Error(t, err, "adding tool with empty name should fail")

	// Test with valid tool
	err = kit.Add(mcp.Tool{Name: "test", Description: "test description", Meta: nil, Annotations: nil, InputSchema: nil, OutputSchema: nil, Title: "", Icons: nil})
	assert.NoError(t, err, "adding tool should succeed")
}
