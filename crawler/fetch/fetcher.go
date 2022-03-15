package fetch

// Fetcher returns a slice of URLs found on that page.
type Fetcher interface {
	Fetch(url string) (urls []string, err error)
}
