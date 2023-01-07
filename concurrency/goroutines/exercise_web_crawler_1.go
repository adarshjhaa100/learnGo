package goroutines

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


type fakeResult struct {
	body string
	urls []string
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

// To store the results of 
type UrlCache struct {
	visited map[string] fakeResult
	allUrls fakeFetcher
	mu sync.Mutex
	counter [2]int
}

// Check if url is visited or not
func (cache *UrlCache) IsVisited(url string) bool{
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.counter[1]++
	_, ok := cache.visited[url]
	return ok

}  

// Insert new url into cache
func (cache *UrlCache) Insert(url string, res *fakeResult) {
	cache.mu.Lock()
	cache.visited[url] = *res
	cache.counter[0]++
	cache.mu.Unlock()
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}


// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, urlCache *UrlCache) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice. (DONE)
	// This implementation doesn't do either:
	time.Sleep(time.Second)
	fmt.Println("visiting: ", url, time.Now())

	if depth <= 0 {
		return
	}

	var (
		body string
		urls []string
		err error
	)

	// Fetch url if its not in cache, else pull the result from cache
	if(!urlCache.IsVisited(url)) {
		urlCache.mu.Lock()
			body, urls, err = fetcher.Fetch(url)
		urlCache.mu.Unlock()

		urlCache.Insert(url, &fakeResult{body, urls})

	} else {
		return
		// fmt.Println("pulled from cache")
		// urlBody, ok := urlCache.visited[url] 
		// if ok {
		// 	body, urls = urlBody.body, urlBody.urls
		// }
	}

	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Printf("found: %s %q \n%v\n", url, body, urls)
	for _, u := range urls {
		if(!urlCache.IsVisited(u)) {
			Crawl(u, depth-1, fetcher, urlCache)
		}
	}
	
}

func RunSolnWebCrawler() {

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

	urlCache := UrlCache{
		visited: make(map[string]fakeResult),
		allUrls: fetcher,
	}

	Crawl("https://golang.org/", 10, &urlCache.allUrls, &urlCache)
	
	urlCache.mu.Lock()
	fmt.Println("Times fetch called actually: ", urlCache.counter )
	urlCache.mu.Unlock()
}