package v7

import (
	"crypto/tls"
	"net"
	"time"
)

type Server struct {
	listener net.Listener
	timeout  time.Duration
}

func (s *Server) run() {

}

// v1显得单薄，没有超时限制、最大连接数等
// v2是大而全了。但如果只想使用默认值呢？
// v3似乎可以各取所需了。但随着可选项的增多，函数也越多
// v6已经很好了，但还不够灵活。

func NewServer(addr string, options ...func(server *Server)) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	srv := &Server{listener: l}
	for _, opt := range options {
		opt(srv)
	}

	return srv, nil
}

func foo() {
	srv, _ := NewServer("localhost") // 使用默认值
	srv.run()

	timeout := func(s *Server) {
		s.timeout = 60 * time.Second
	}

	tls := func(s *Server) {
		config := loadTLSConfig()
		s.listener = tls.NewListener(s.listener, config)
	}

	// listen securely with a 60 second timeout
	srv2, _ := NewServer("localhost",
		timeout,
		tls)
	srv2.run()

}

func loadTLSConfig() *tls.Config {
	return nil
}
