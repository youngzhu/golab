package worker

import (
	"github.com/youngzhu/golab/crawler/fetch"
	"sync"
)

// Concurrent crawler with shared state and Mutex

type fetchState struct {
	mu      sync.Mutex
	fetched map[string]bool
}

func NewFetchState() *fetchState {
	fetched := make(map[string]bool)
	return &fetchState{fetched: fetched}
}

func ConcurrentMutex(url string, fetcher fetch.Fetcher, f *fetchState) {
	// 注释掉 Lock 和 Unlock
	// 执行  go run -race main.go
	// 会报错
	f.mu.Lock()
	already := f.fetched[url]
	f.fetched[url] = true
	f.mu.Unlock()

	if already {
		return
	}

	urls, err := fetcher.Fetch(url)
	if err != nil {
		return
	}
	var done sync.WaitGroup
	for _, u := range urls {
		done.Add(1)
		go func(uu string) {
			defer done.Done()
			ConcurrentMutex(uu, fetcher, f)
		}(u)
	}
	done.Wait()
	return
}
