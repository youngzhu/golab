package v2

import (
	"github.com/youngzhu/golab/effective/channel/server"
	"net/http"
)

// v2
// 没有了handle
// 把信号量放在Serve里

var sem = make(chan int, server.MaxOutstanding)

func Serve(queue chan *http.Request) {
	for req := range queue {
		sem <- 1
		go func() {
			// 有问题，说是req在协程之间共享了
			process(req)
			<-sem
		}()
	}
}

func process(r *http.Request) {

}
