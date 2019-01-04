package main

import (
	"fmt"
	"sync"			// for sync.Mutex
	"sync/atomic"  	// For atomic counters
	"time"
)

// Safe HashSet Version
type SafeHashSet struct {
	sync.Mutex
	urls map[string]bool //Primarily we wanted use this as an hashset, so the value of map is not significant to us
}

var (
	urlSet     SafeHashSet
	urlCounter int64
)

// Adds an URL to the Set, returns true if new url was added (if not present already)
func (m *SafeHashSet) add(newUrl string) bool {
	m.Lock()
	defer m.Unlock()
	_, ok := m.urls[newUrl]
	if !ok {
		m.urls[newUrl] = true
		return true
	}
	return false
}


type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {

	// Decrement the atomic url counter, when this crawl function exits
	defer atomic.AddInt64(&urlCounter, -1)

	if depth <= 0 {
		return
	}

	// Don't Process a url if it is already processed
	isNewUrl := urlSet.add(url)

	if !isNewUrl {
		fmt.Printf("skip: \t%s\n", url)
		return
	}


	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: \t%s %q\n", url, body)

	for _, u := range urls {
		atomic.AddInt64(&urlCounter, 1)
		// Crawl parallely
		go Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	urlSet = SafeHashSet{urls: make(map[string]bool)}

	atomic.AddInt64(&urlCounter, 1)
	go Crawl("https://golang.org/", 4, fetcher)

	for atomic.LoadInt64(&urlCounter) > 0 {
		time.Sleep(100 * time.Microsecond)
	}
	fmt.Println("Exiting")
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
	return "", nil, fmt.Errorf("not found: \t%s", url)
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
