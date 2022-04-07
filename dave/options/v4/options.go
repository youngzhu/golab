package v4

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
// v3似乎可以各取所需了。但随着可选项的增多，函数也越多

// Config is used to configure the Server
type Config struct {
	// the port to listen on, if unset default to 8080
	Port int

	// sets the amount of time before closing idle connections,
	// or forever if not provided.
	Timeout time.Duration

	// the server will accept TLS connections that Cert provided.
	Cert *tls.Certificate
}

func NewServer(addr string, config Config) (*Server, error) {
	return nil, nil
}

// 好处：
// 再多的配置也不影响 NewServer 的API

// 如果要用默认值，怎么设值呢？
func foo() {
	srv, _ := NewServer("localhost", Config{
		Port: 0, // 0 不是空值，更不是默认值
	})
	srv.run()
}

// 这样是可以的。但。。。
func bar() {
	srv, _ := NewServer("localhost", Config{})
	srv.run()
}
