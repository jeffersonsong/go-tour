package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Set[T comparable] struct {
	mu sync.Mutex
	m  map[T]bool
}

func (s *Set[T]) Add(val T) {
	s.mu.Lock()
	s.m[val] = true
	s.mu.Unlock()
}

func (s *Set[T]) Contains(val T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.m[val]
	return exists
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, visited *Set[string], bad *Set[string]) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		bad.Add(url)
		return
	}
	visited.Add(url)
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if !visited.Contains(u) && !bad.Contains(u) {
			go Crawl(u, depth-1, fetcher, visited, bad)
		}
	}
	return
}

func main() {
	visited := Set[string]{m: make(map[string]bool)}
	bad := Set[string]{m: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, &visited, &bad)
	time.Sleep(time.Duration(1) * time.Second)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
