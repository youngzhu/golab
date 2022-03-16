package main

import (
	"fmt"
	"github.com/youngzhu/golab/crawler/fetch"
	"github.com/youngzhu/golab/crawler/worker"
)

func main() {
	fetcher := fetch.NewFakeFetcher()

	fmt.Printf("=== Serial ===\n")
	worker.Serial("https://golang.org/", fetcher, make(map[string]bool))

	fmt.Printf("=== Concurrent Mutex ===\n")
	worker.ConcurrentMutex("https://golang.org/", fetcher, worker.NewFetchState())

	fmt.Printf("=== Concurrent Channel ===\n")
	worker.ConcurrentChannel("https://golang.org/", fetcher)

}
