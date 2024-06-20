package main

import (
    "fmt"
    "log"
)


// Server is an HTTP server.
type Server struct {
    verbose bool
    port    int
}



// NewServer returns a Server with options.
func NewServer(options ...func(*Server) error) (*Server, error) {
    srv := &Server{
        port: 8080, // default port
    }
    for _, opt := range options {
        if err := opt(srv); err != nil {
            return nil, err
        }
    }

    return srv, nil
}



// WithVerbose sets the verbose option on s.
func WithVerbose(s *Server) error {
    s.verbose = true
    return nil
}


const portErrFmt = "port must be between 0 and %d, got %d"

func WithPort(port int) func(*Server) error {
    const maxPort = 0xFFFF
    return func(s *Server) error {
        if port <= 0 || port > maxPort {
            return fmt.Errorf(portErrFmt, maxPort, port)
        }
        s.port = port
        return nil
    }
}


func main() {
    srv, err := NewServer(WithPort(9999))
    if err != nil {
        log.Fatalf("error: %s", err)
    }

    fmt.Printf("%#v\n", srv)
}
