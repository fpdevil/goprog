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

const (
	msg = `
	Outline prints the outline of HTML document tree
	************************************************
	`
)

func main() {
	fmt.Println(msg)
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
		fmt.Fprintf(os.Stderr, "outline error: %s\n", err.Error())
		return
	}

	outline([]string{}, doc)
}

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

//+!outline
// outline function takes a url and pushes an element onto a
// stack, without popping out the element
func outline(stack []string, node *html.Node) {
	if node.Type == html.ElementNode {
		stack = append(stack, node.Data) // push the tag
		fmt.Printf("%v\n", stack)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

//!-outline
