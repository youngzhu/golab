package worker

import "github.com/youngzhu/golab/crawler/fetch"

// Serial crawler
func Serial(url string, fetcher fetch.FakeFetcher, fetched map[string]bool) {
	if fetched[url] {
		return
	}
	fetched[url] = true
	urls, err := fetcher.Fetch(url)
	if err != nil {
		return
	}
	for _, u := range urls {
		Serial(u, fetcher, fetched)
	}
	return
}
