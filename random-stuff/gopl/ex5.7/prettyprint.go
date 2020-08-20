package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

const (
	msg = `
	Ex 5.7| Outline prints the outline of HTML document tree
	********************************************************
	`
)

func main() {
	fmt.Println(msg)

	for _, url := range os.Args[1:] {
		outline(url)
	}
}

//!+forEachNode
// forEachNode function calls  the functions pre(x) and  post(x) for each
// node x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder)
// post is called after the children are visited (postorder)
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

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		// using * adverb in %*s prints a string padded with a variable number
		// of space. width and string are provided by depth*2 and ""
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s=%q", a.Key, a.Val)
		}
		if n.FirstChild == nil {
			fmt.Println("/>")
		} else {
			fmt.Println(">")
		}
		depth++
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
	} else if (n.Type == html.TextNode) && (n.Parent.Data != "script" && n.Parent.Data != "style") {
		trim := strings.TrimSpace(n.Data)
		if trim != "" {
			fmt.Printf("%*s%s\n", depth*2, "", trim)
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		// using * adverb in %*s prints a string padded with a variable number
		// of space. width and string are provided by depth*2 and ""
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

//!-startend

func outline(url string) error {
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "url get error %v: %v\n", url, err)
		return err
	}

	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "url parse error %v: %v\n", url, err)
		return err
	}

	//!+call func
	forEachNode(doc, startElement, endElement)
	//!-call func

	return nil
}
