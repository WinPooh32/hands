package read_test

import (
	"context"
	"testing"

	"github.com/WinPooh32/hands/tools/read"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input read.ReadInput
		want  *mcp.CallToolResult
		want2 any
	}{
		// TODO: Add success test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, got2, err := read.Read(context.Background(), nil, tt.input)
			require.NoError(t, err)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want2, got2)
		})
	}
}

func TestRead_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input read.ReadInput
		want  *mcp.CallToolResult
		want2 any
	}{
		// TODO: Add error test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, got2, err := read.Read(context.Background(), nil, tt.input)

			assert.Error(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want2, got2)
		})
	}
}
