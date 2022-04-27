package v1

import (
	"github.com/youngzhu/golab/effective/channel/server"
	"net/http"
)

// 问题：
// 每来一个请求就会创建一个新的协程，但能同时处理的请求有一个上限（MaxOutstanding）

var sem = make(chan int, server.MaxOutstanding)

func handle(r *http.Request) {
	sem <- 1   // wait for active queue to drain
	process(r) // may take a long time
	<-sem      // done; enable next request to run
}

func Serve(queue chan *http.Request) {
	for {
		req := <-queue
		go handle(req)
	}
}

func process(r *http.Request) {

}
