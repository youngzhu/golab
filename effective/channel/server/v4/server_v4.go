package v4

import (
	"github.com/youngzhu/golab/effective/channel/server"
	"net/http"
)

var sem = make(chan int, server.MaxOutstanding)

func Serve(queue chan *http.Request) {
	for req := range queue {
		// 虽然奇怪，但这是有效的
		// 跟v3相似，虽然变量名一样，但表示不同的实例
		req := req // create new instance of req for the goroutine
		sem <- 1
		go func() {
			process(req)
			<-sem
		}()
	}
}

func process(r *http.Request) {

}
