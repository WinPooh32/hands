package grep

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/WinPooh32/hands/pkg/mcputil"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type GrepInput struct {
	Pattern    string `json:"pattern" jsonschema:"pattern to search for (regex)"`
	Path       string `json:"path" jsonschema:"file or directory to search in"`
	IgnoreCase bool   `json:"ignoreCase,omitzero" jsonschema:"ignore case in search (default: false)"`
}

func Grep(ctx context.Context, _ *mcp.CallToolRequest, input GrepInput) (*mcp.CallToolResult, any, error) {
	if input.Pattern == "" {
		return mcputil.ErrorResult("pattern is required"), nil, nil
	}

	if input.Path == "" {
		return mcputil.ErrorResult("path is required"), nil, nil
	}

	var re *regexp.Regexp
	if input.IgnoreCase {
		re = regexp.MustCompile("(?i)" + input.Pattern)
	} else {
		re = regexp.MustCompile(input.Pattern)
	}

	absPath, err := filepath.Abs(input.Path)
	if err != nil {
		return mcputil.ErrorResult(fmt.Sprintf("failed to get absolute path: %v", err)), nil, nil
	}

	info, err := os.Stat(absPath)
	if err != nil {
		return mcputil.ErrorResult(fmt.Sprintf("path not found: %v", err)), nil, nil
	}

	var results []string

	if info.IsDir() {
		err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			return searchFile(path, re, &results)
		})
	} else {
		err = searchFile(absPath, re, &results)
	}

	if err != nil {
		return mcputil.ErrorResult(fmt.Sprintf("search error: %v", err)), nil, nil
	}

	if len(results) == 0 {
		return mcputil.TextResult("No matches found"), nil, nil
	}

	result := fmt.Sprintf("Found %d matches:", len(results))
	for _, r := range results {
		result += fmt.Sprintf("\n%s", r)
	}

	return mcputil.TextResult(result), nil, nil
}

func searchFile(path string, re *regexp.Regexp, results *[]string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", path, err)
	}

	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)

	lineNum := 0
	for scanner.Scan() {
		lineNum++

		line := scanner.Text()
		if re.MatchString(line) {
			*results = append(*results, fmt.Sprintf("%s:%d: %s", path, lineNum, line))
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scan: %w", err)
	}

	return nil
}
