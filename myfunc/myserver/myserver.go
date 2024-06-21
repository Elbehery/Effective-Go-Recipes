package myserver

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
