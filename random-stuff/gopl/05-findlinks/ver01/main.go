package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/net/html"
)

//#+!fetch
// fetch function will loop over a list of http url links provided
// as argument and makes a `GET` call to each to fetch the contetnt
// and store them into a byte slice to return the same
func fetch(urls []string) ([]byte, error) {
	var sb []byte
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stdout, "error calling %s: %s", url, err.Error())
			return nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stdout, "error reading response: %s", err.Error())
			return nil, err
		}
		resp.Body.Close()

		sb = append(sb, body...)
	}
	return sb, nil
}

//-!fetch

//!+visit
// visit function traverses a html node tree, extracts the links from
// the href attribute of eacg anchor element `<a href=...` and appends
// them to a string slice and returns the same
func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// call recursively for the rest of html nodes
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

//!-visit

//!+exec
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage: go run %s [<url1> <url2>...]\n", filepath.Base(os.Args[0]))
		return
	}

	sb, err := fetch(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error from fetch: %s", err.Error())
		return
	}

	doc, err := html.Parse(bytes.NewReader(sb))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks error: %s\n", err.Error())
		return
	}

	for i, link := range visit([]string{}, doc) {
		fmt.Printf("%3d: %s\n", i, link)
	}
}

//!-exec
