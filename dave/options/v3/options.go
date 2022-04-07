package v3

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
// v2是大而全了。但如果只想使用默认值呢？

// NewServer returns a new Server listening on addr
func NewServer(addr string) (*Server, error) {
	return nil, nil
}

// NewTLSServer returns a secure server listening on addr
func NewTLSServer(addr string,
	cert *tls.Certificate) (*Server, error) {
	return nil, nil
}

// NewServerWithTimeout returns a server listening on addr
// that disconnects idle clients.
func NewServerWithTimeout(addr string,
	clientTimeout time.Duration) (*Server, error) {
	return nil, nil
}

// NewTLSServerWithTimeout returns a secure server listening on addr
// that disconnects idle clients.
func NewTLSServerWithTimeout(addr string,
	cert *tls.Certificate,
	timeout time.Duration) (*Server, error) {
	return nil, nil
}
