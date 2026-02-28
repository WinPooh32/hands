package tools

import (
	"context"
	"flag"
)

type ServerConfig struct{}

type Server struct {
}

func NewServer() *Server {
	panic("TODO")
}

func (srv *Server) ParseArgsConfig() (*ServerConfig, error) {
	httpAddr := flag.String("http", "", "if set, use streamable HTTP at this address, instead of stdin/stdout")

	flag.Parse()

	_ = httpAddr

	panic("TODO")
}

func (srv *Server) Run(ctx context.Context, c *ServerConfig, kit *Kit) error {
	panic("TODO")
}
