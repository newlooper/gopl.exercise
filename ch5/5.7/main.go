package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
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

	forEachNode(doc, startElement, endElement)

	return nil
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

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		strAttr := ""
		for _, attr := range n.Attr {
			strAttr += fmt.Sprintf("%s=%q ", attr.Key, attr.Val)
		}

		emptyElmt := ""
		if n.Data == "img" && n.FirstChild == nil {
			emptyElmt = " /"
		}

		if strAttr != "" {
			fmt.Printf("%*s<%s %s%s>\n", depth*2, "", n.Data, strAttr[:len(strAttr)-1], emptyElmt)
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, emptyElmt)
		}
		depth++
	case html.TextNode:
		for _, line := range strings.Split(n.Data, "\n") {
			line = strings.Trim(line, " \r\n")
			if line != "" {
				fmt.Printf("%*s%s\n", depth*2, "", line)
			}
		}
	case html.CommentNode:
		fmt.Printf("%*s<!-- %s -->\n", depth*1, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.Data != "img" || n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
