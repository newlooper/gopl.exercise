package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("%v", err)
		}

		doc, err := html.Parse(resp.Body)
		if err != nil {
			resp.Body.Close()
			log.Fatalf("%v", err)
		}
		resp.Body.Close()

		headers := ElementsByTagName(doc, "h1", "h2", "h3", "h4", "h5")
		for _, n := range headers {
			fmt.Printf("<%s>%s</%[1]s>\n", strings.Trim(n.Data, "\t \r\n"), n.FirstChild.Data)
		}
	}
}

func ElementsByTagName(doc *html.Node, tags ...string) (matched []*html.Node) {
	pre := func(doc *html.Node) {
		for _, tag := range tags {
			if doc.Type == html.ElementNode &&
				doc.Data == tag &&
				doc.FirstChild != nil {
				matched = append(matched, doc)
			}
		}
	}
	forEachNode(doc, pre, nil)
	return
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
