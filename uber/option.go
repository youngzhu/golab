package uber

import (
	"log"
	"time"
)

/************ Bad ********/

type Connection struct {
}

func Connect(
	addr string,
	timeout time.Duration,
	caching bool,
) (*Connection, error) {

	// ...
	return &Connection{}, nil
}

const (
	DefaultCaching = true
	DefaultTimeout = 60
)

func foo() {
	// 调用时三个参数都必须填写，即使用户希望使用默认值
	Connect("addr", DefaultTimeout, DefaultCaching)
	Connect("addr", 30, DefaultCaching)
	Connect("addr", DefaultTimeout, false)
	Connect("addr", 180, false)
}

/************ Good ********/
type options struct {
	cache  bool
	logger *log.Logger
}

type Option interface {
	apply(opts *options)
}

type cacheOption bool

func (c cacheOption) apply(opts *options) {
	opts.cache = bool(c)
}

func WithCache(c bool) Option {
	return cacheOption(c)
}

type loggerOption struct {
	logger *log.Logger
}

func (l loggerOption) apply(opts *options) {
	opts.logger = l.logger
}

func WithLogger(logger *log.Logger) Option {
	return loggerOption{logger: logger}
}

func Open(
	addr string,
	opts ...Option,
) (*Connection, error) {
	// 默认的选项
	defaultOpts := options{
		cache:  DefaultCaching,
		logger: log.Default(),
	}

	// 根据入参改变默认选项
	for _, o := range opts {
		o.apply(&defaultOpts)
	}

	// ...
	return &Connection{}, nil
}

func bar() {
	// 只有在需要时才增加参数
	Open("addr")
	Open("addr", WithLogger(log.Default()))
	Open("addr", WithCache(false))
	Open(
		"addr",
		WithCache(false),
		WithLogger(log.Default()),
	)
}
