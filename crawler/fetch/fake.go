package fetch

import "fmt"

type FakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func NewFakeFetcher() FakeFetcher {
	return FakeFetcher{
		"https://golang.org/": &fakeResult{
			body: "The Go Programming Language",
			urls: []string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": &fakeResult{
			body: "Packages",
			urls: []string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		},
		"https://golang.org/pkg/fmt/": &fakeResult{
			body: "Package fmt",
			urls: []string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
		"https://golang.org/pkg/os/": &fakeResult{
			body: "Package os",
			urls: []string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
	}
}

func (f FakeFetcher) Fetch(url string) ([]string, error) {
	if res, ok := f[url]; ok {
		fmt.Printf("found:   %s\n", url)
		return res.urls, nil
	}
	fmt.Printf("missing: %s\n", url)
	return nil, fmt.Errorf("not found: %s", url)
}
