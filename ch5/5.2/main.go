package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			continue
		}

		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatalf("error parsing %s: %v", url, err)
		}

		var m = make(map[string]int)
		tagCount(m, doc)
		fmt.Printf("%s\n", url)
		for k, v := range m {
			fmt.Printf("%s\t%d\n", k, v)
		}
	}
}

func tagCount(m map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tagCount(m, c)
	}
}
