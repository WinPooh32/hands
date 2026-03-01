package tools

import (
	"fmt"

	"github.com/WinPooh32/hands/pkg/i18n"
	"github.com/WinPooh32/hands/tools/bash"
	"github.com/WinPooh32/hands/tools/edit"
	"github.com/WinPooh32/hands/tools/glob"
	"github.com/WinPooh32/hands/tools/grep"
	"github.com/WinPooh32/hands/tools/read"
	"github.com/WinPooh32/hands/tools/write"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type CallFilter struct{}

type Kit struct {
	server *mcp.Server
	tools  []mcp.Tool
}

func (kit *Kit) Add(t mcp.Tool) error {
	if t.Name == "" {
		return fmt.Errorf("tool name cannot be empty")
	}

	kit.tools = append(kit.tools, t)

	return nil
}

func NewDefault(filters []CallFilter) *Kit {
	kit := New(filters)

	// Add default tools
	mcpToolRead := &mcp.Tool{Name: "read", Description: i18n.Tr(i18n.ReadDescription), Meta: nil, Annotations: nil, InputSchema: nil, OutputSchema: nil, Title: "", Icons: nil}
	mcp.AddTool(kit.server, mcpToolRead, read.Read)

	mcpToolWrite := &mcp.Tool{Name: "write", Description: i18n.Tr(i18n.WriteDescription), Meta: nil, Annotations: nil, InputSchema: nil, OutputSchema: nil, Title: "", Icons: nil}
	mcp.AddTool(kit.server, mcpToolWrite, write.Write)

	mcpToolEdit := &mcp.Tool{Name: "edit", Description: i18n.Tr(i18n.EditDescription), Meta: nil, Annotations: nil, InputSchema: nil, OutputSchema: nil, Title: "", Icons: nil}
	mcp.AddTool(kit.server, mcpToolEdit, edit.Edit)

	mcpToolGlob := &mcp.Tool{Name: "glob", Description: i18n.Tr(i18n.GlobDescription), Meta: nil, Annotations: nil, InputSchema: nil, OutputSchema: nil, Title: "", Icons: nil}
	mcp.AddTool(kit.server, mcpToolGlob, glob.Glob)

	mcpToolGrep := &mcp.Tool{Name: "grep", Description: i18n.Tr(i18n.GrepDescription), Meta: nil, Annotations: nil, InputSchema: nil, OutputSchema: nil, Title: "", Icons: nil}
	mcp.AddTool(kit.server, mcpToolGrep, grep.Grep)

	mcpToolBash := &mcp.Tool{Name: "bash", Description: i18n.Tr(i18n.BashDescription), Meta: nil, Annotations: nil, InputSchema: nil, OutputSchema: nil, Title: "", Icons: nil}
	mcp.AddTool(kit.server, mcpToolBash, bash.Bash)

	return kit
}

func New(filters []CallFilter) *Kit {
	return &Kit{
		server: mcp.NewServer(&mcp.Implementation{Name: "hands", Version: "0.1.0", Title: "", WebsiteURL: "", Icons: nil}, nil),
		tools:  []mcp.Tool{},
	}
}
