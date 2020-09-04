package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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
		fmt.Println(url)
		text(doc)
	}
}

func text(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "style" || n.Data == "script") {
		return
	}

	if n.Type == html.TextNode {
		text := strings.Trim(n.Data, " \r\n")
		if len(text) > 0 {
			fmt.Println(text)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text(c)
	}
}
