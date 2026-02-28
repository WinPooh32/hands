package read

import (
	"context"

	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ReadInput struct {
}

func Read(ctx context.Context, _ *mcp.CallToolRequest, input ReadInput) (*mcp.CallToolResult, any, error) {
	return mcputil.TextResult("TODO"), nil, nil
}

func read() (string, error) {
	panic("TODO")
}
