package v2

import (
	"crypto/tls"
	"net"
	"time"
)

type Server struct {
	listener net.Listener
}

func (s *Server) run() {

}

// v1显得单薄，没有超时限制、最大连接数等

// NewServer returns a new Server listening on addr
// clientTimeout defines the maximum length of an idle connection,
// or forever if not provided.
// maxConns limits the number of connections.
// maxConcurrent limits the number of concurrent connections
// from a single IP address.
// cert is the TLS certificate for the connection.
func NewServer(addr string,
	clientTimeout time.Duration,
	maxConns, maxConcurrent int,
	cert *tls.Certificate) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	srv := Server{listener: l}
	go srv.run()
	return &srv, nil
}
