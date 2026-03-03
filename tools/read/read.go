package read

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/WinPooh32/hands/pkg/i18n"
	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ReadInput struct {
	Path string `json:"path" jsonschema:"path to the file to read"`
}

func Read(ctx context.Context, _ *mcp.CallToolRequest, input ReadInput) (*mcp.CallToolResult, any, error) {
	if input.Path == "" {
		return mcputil.ErrorResult("path is required"), nil, nil
	}

	file, err := os.Open(input.Path)
	if err != nil {
		return mcputil.ErrorResult(fmt.Sprintf("open file: %v", err)), nil, nil
	}

	defer func() { _ = file.Close() }()

	text, err := readWithLineNumbers(file)
	if err != nil {
		return mcputil.ErrorResult(fmt.Sprintf("read file: %v", err)), nil, nil
	}

	return mcputil.TextResult(text), nil, nil
}

func readWithLineNumbers(r io.Reader) (string, error) {
	sb := strings.Builder{}
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		fmt.Fprintf(&sb, "%6d→%s\n", lineNum, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scan line: %w", err)
	}

	return sb.String(), nil
}

// Schema returns the JSON schema for ReadInput with translated descriptions
func Schema() any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"path": map[string]any{
				"description": i18n.Tr(i18n.ReadArgPath),
			},
		},
		"required": []string{"path"},
	}
}
