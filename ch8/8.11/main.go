package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"sync"
)

func main() {

	flag.Parse()
	responses := make(chan *http.Response)
	wg := &sync.WaitGroup{}

	//////////////////////////////////////////////////////////
	// req.Cancel 被废弃，官方推荐使用 http.NewRequestWithContext
	ctx, cancel := context.WithCancel(context.Background())

	for _, url := range flag.Args() {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Printf("HEAD %s: %s", url, err)
				return
			}
			resp.Body.Close()

			select {
			case responses <- resp:
			case <-ctx.Done():
				log.Printf("cancel fetch [%s]: return from blocking", url)
			}

		}(url)
	}

	resp := <-responses // 第一个完成的 http 请求
	cancel()            // 取消所有 http 请求

	log.Printf("First response: [%s], all the others will be canceld...\n", resp.Request.URL)

	wg.Wait()
}
