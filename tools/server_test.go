package tools_test

import (
	"context"
	"testing"
	"time"

	"github.com/WinPooh32/hands/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
)

func TestParseArgsConfig(t *testing.T) {
	t.Parallel()

	// Test that NewServer returns a valid server
	srv := tools.NewServer()
	assert.NotNil(t, srv, "NewServer should return a non-nil server")

	// ParseArgsConfig with empty args should return default config
	config, err := srv.ParseArgsConfig()
	assert.NoError(t, err, "ParseArgsConfig should not error with no args")
	assert.NotNil(t, config, "ParseArgsConfig should return a non-nil config")
	assert.True(t, config.StdioMode, "default mode should be stdio")
	assert.False(t, config.HTTPMode, "HTTP mode should be false by default")
}

func TestServerRunStdio(t *testing.T) {
	t.Parallel()

	// This test verifies the Run method handles stdio mode correctly
	// With no tools added, the server should still run (just won't have tools capability)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := tools.NewServer()
	kit := tools.New(nil)

	// Start Run in a goroutine
	done := make(chan error, 1)

	go func() {
		err := srv.Run(ctx, &tools.ServerConfig{HTTPAddr: "", StdioMode: true, HTTPMode: false}, kit)
		done <- err
	}()

	// Give it a moment to start
	select {
	case err := <-done:
		// Run returned immediately - this might happen if no tools registered
		assert.NoError(t, err, "Run should complete without error")
	case <-time.After(500 * time.Millisecond):
		// Run is still blocking - this is also valid behavior
	}
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
