package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"gopl.io/ch5/links"
)

var maxDepth = flag.Int("depth", 1, "max depth to fetched")

var tokens = make(chan struct{}, 20) // limit of parallel
var seen = make(map[string]bool)     // url records
var seenLock = sync.Mutex{}          // avoid race condition
var wg = sync.WaitGroup{}            // avoid main goroutine quit earlier

func main() {
	flag.Parse()
	for _, link := range flag.Args() {
		wg.Add(1)
		go crawl(link, 0)
	}
	wg.Wait()
}

func crawl(url string, depth int) { // via recursion, could get depth naturally
	defer wg.Done()
	fmt.Printf("%d\t%s\n", depth, url)

	if depth >= *maxDepth {
		return
	}

	tokens <- struct{}{} // will block when parallel limit reached
	list, err := links.Extract(url)
	<-tokens

	if err != nil {
		log.Print(err)
	}

	for _, link := range list {
		seenLock.Lock() // protect critical section
		if seen[link] {
			seenLock.Unlock()
			continue
		}
		seen[link] = true
		seenLock.Unlock()

		wg.Add(1)
		go crawl(link, depth+1) // via params, every goroutine has their own depth var, no need to sync
	}
}
