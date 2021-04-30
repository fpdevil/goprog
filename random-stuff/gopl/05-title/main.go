package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "usage: go run %s <input url>\n", filepath.Base(os.Args[0]))
		return
	}
	for _, arg := range os.Args[1:] {
		if err := title(arg); err != nil {
			fmt.Printf("title: %v\n", err)
		}
	}
}

//!+

func forEachNode(node *html.Node, pre, post func(node *html.Node)) {
	if pre != nil {
		pre(node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(node)
	}
}

//!-

//!+ title
//
// title function takes a url and returns the title from the html document
func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}

	forEachNode(doc, visitNode, nil)
	return nil
}

//!- title
