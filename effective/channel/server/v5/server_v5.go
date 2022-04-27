package v5

import (
	"github.com/youngzhu/golab/effective/channel/server"
	"net/http"
)

// 没看懂这个逻辑。。

var sem = make(chan int, server.MaxOutstanding)

func handle(queue chan *http.Request) {
	for r := range queue {
		process(r)
	}
}

func Serve(clientRequests chan *http.Request, quit chan bool) {
	// start handlers
	for i := 0; i < server.MaxOutstanding; i++ {
		go handle(clientRequests)
	}
	<-quit // wait to be told to exit
}

func process(r *http.Request) {

}
