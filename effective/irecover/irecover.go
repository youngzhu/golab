package irecover

import "log"

type Work struct {
}

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work *Work) {
	// 即使遇到panic，日志也会被记录
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()

	do(work)
}

func do(work *Work) {

}
