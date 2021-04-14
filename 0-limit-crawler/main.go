//////////////////////////////////////////////////////////////////////
//
// Your task is to change the code to limit the crawler to at most one
// page per second, while maintaining concurrency (in other words,
// Crawl() must be called concurrently)
//
// @hint: you can achieve this by adding 3 lines
//

package main

import (
	"fmt"
	"sync"
	"time"
)

// Crawl uses `fetcher` from the `mockfetcher.go` file to imitate a
// real crawler. It crawls until the maximum depth has reached.
func Crawl(url string, depth int, wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
	var urls []string
	if depth <= 0 {
		return
	}
LOOP:
	for {
		select {
		case <-ch:
			body, urlsData, err := fetcher.Fetch(url)
			time.Sleep(time.Second)
			ch <- struct{}{}
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("found: %s %q\n", url, body)
			urls = urlsData
			break LOOP
		default:
			time.Sleep(time.Second)
		}
	}
	wg.Add(len(urls))
	for _, u := range urls {
		// Do not remove the `go` keyword, as Crawl() must be
		// called concurrently
		go Crawl(u, depth-1, wg, ch)
	}
	return
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	wg.Add(1)
	Crawl("http://golang.org/", 4, &wg, ch)
	wg.Wait()
}
