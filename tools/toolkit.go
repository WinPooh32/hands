package tools

import "github.com/WinPooh32/hands/pkg/tool"

type CallFilter struct {
}

type Kit struct{}

func (kit *Kit) Add(t tool.Tool) error {
	panic("TODO")
}

type HTTPConfig struct{}

func (kit *Kit) serveMCPHTTP(c HTTPConfig) error {
	panic("TODO")
}

type StdioConfig struct{}

func (kit *Kit) serveMCPStdio(c StdioConfig) error {
	panic("TODO")
}

func NewDefault(filters []CallFilter) *Kit {
	// make a new toolkit with all tools
	panic("TODO")
}

func New(filters []CallFilter) *Kit {
	// make a new empty toolkit
	panic("TODO")
}
