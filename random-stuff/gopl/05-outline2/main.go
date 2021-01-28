package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/net/html"
)

const (
	msg = `
	Outline prints the outline of HTML document tree
	************************************************
	`
)

var depth int

func main() {
	fmt.Println(msg)
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage: go run %s [<url1> <url2>...]\n", filepath.Base(os.Args[0]))
		return
	}

	for _, url := range os.Args[1:] {
		outline(url)
	}
}

//!+ forEachNode

// forEachNode function  calls the function  pre(x) and post(x)  for each
// node x in the tree rooted at  n. Both functions are optional, with pre
// calles  before the  children are  visited (preorder)  and post  called
// after (postorder).
func forEachNode(node *html.Node, pre, post func(*html.Node)) {
	// pre is called before the node's children are visited
	if pre != nil {
		pre(node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	// post is called after the node's children are visited
	if post != nil {
		post(node)
	}
}

//!- forEachNode

//!+startend

func startElement(node *html.Node) {
	if node.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", node.Data)
		depth++
	}
}

func endElement(node *html.Node) {
	if node.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", node.Data)
	}
}

//!-startend

//!+outline

// outline function takes a url and makes a get call to the same
// parses the response body and processes the same
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

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!-outline
