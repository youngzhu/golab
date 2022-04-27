package v3

import (
	"github.com/youngzhu/golab/effective/channel/server"
	"net/http"
)

// 修复了v2的问题，协程函数加一个入参

var sem = make(chan int, server.MaxOutstanding)

func Serve(queue chan *http.Request) {
	for req := range queue {
		sem <- 1
		go func(req *http.Request) {
			process(req)
			<-sem
		}(req)
	}
}

func process(r *http.Request) {

}
