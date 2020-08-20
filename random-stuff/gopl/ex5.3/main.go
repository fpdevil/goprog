package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func fetch(urls []string) ([]byte, error) {
	var b []byte
	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error %v from %s\n", err, url)
			return nil, err
		}
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error %v\n", err)
		}
		b = append([]byte(nil), body...)
	}
	return b, nil
}

func textNodes(node *html.Node, stack []string) []string {
	if node == nil {
		return stack
	}

	if node.Type == html.TextNode {
		if node.Parent.Data != "script" && node.Parent.Data != "style" {
			data := strings.Split(strings.TrimSpace(node.Data), "\n")
			for _, line := range data {
				if len(line) != 0 {
					stack = append(stack, line)
				}
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		stack = textNodes(c, stack)
	}

	return stack
}

func main() {
	body, err := fetch(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "textNodes: %v\n", err)
		return
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		fmt.Fprintf(os.Stderr, "textNodes: %v\n", err)
	}
	for _, txt := range textNodes(doc, []string(nil)) {
		fmt.Fprintf(os.Stdout, "* %v\n", txt)
	}
}
