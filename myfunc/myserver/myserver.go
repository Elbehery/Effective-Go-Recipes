package main

import (
	"fmt"
	"log"
)

type Server struct {
	verbose bool
	port    int
}

func NewServer(options ...func(*Server) error) (*Server, error) {
	srv := &Server{
		port: 8080,
	}

	for _, opt := range options {
		if err := opt(srv); err != nil {
			return nil, err
		}
	}
	return srv, nil
}

func WithVerbose(srv *Server) error {
	srv.verbose = true
	return nil
}

const portErrFmt = "port must be between 0 and %d, got %d"

func WithPort(port int) func(*Server) error {
	const maxPort = 0xFFFF
	return func(server *Server) error {
		if port <= 0 || port > maxPort {
			return fmt.Errorf(portErrFmt, maxPort, port)
		}
		server.port = port
		return nil
	}
}

func main() {
	srv, err := NewServer(WithPort(9999), WithVerbose)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("%#v\n", srv)
}
