package worker

import "github.com/youngzhu/golab/crawler/fetch"

// Concurrent crawler with channels

func ConcurrentChannel(url string, fetcher fetch.Fetcher) {
	ch := make(chan []string)
	go func() {
		ch <- []string{url}
	}()
	coordinator(ch, fetcher)
}

func coordinator(ch chan []string, fetcher fetch.Fetcher) {
	n := 1
	fetched := make(map[string]bool)
	for urls := range ch {
		for _, u := range urls {
			if !fetched[u] {
				fetched[u] = true
				n++
				go worker(u, ch, fetcher)
			}
		}
		n--
		if n == 0 {
			break
		}
	}
}

func worker(url string, ch chan []string, fetcher fetch.Fetcher) {
	urls, err := fetcher.Fetch(url)
	if err != nil {
		ch <- []string{}
	} else {
		ch <- urls
	}
}
