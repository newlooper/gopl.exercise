package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	if node := ElementByID(doc, "page"); node != nil {
		fmt.Printf("%T\t<%s>", node.Type, node.Data)
	}

	return nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	pre := func(doc *html.Node) bool {
		for _, attr := range doc.Attr {
			if attr.Key == "id" && attr.Val == id {
				return true
			}
		}
		return false
	}

	if node, found := forEachNode(doc, pre, nil); found {
		return node
	}
	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) (*html.Node, bool) {
	if pre != nil {
		if pre(n) {
			return n, true
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if node, ok := forEachNode(c, pre, post); ok {
			return node, true
		}
	}

	if post != nil {
		if post(n) {
			return n, true
		}
	}
	return nil, false
}
