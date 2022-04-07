package v1

import "net"

type Server struct {
	listener net.Listener
}

func (s *Server) run() {

}

// NewServer returns a new Server listening on addr
func NewServer(addr string) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	srv := Server{listener: l}
	go srv.run()
	return &srv, nil
}
