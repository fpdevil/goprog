package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// visit function appends to the links slice each link that is found
// in the node and return the results
func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

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

func main() {
	body, err := fetch(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		return
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		return
	}

	for _, link := range visit([]string{}, doc) {
		fmt.Printf("%v\n", link)
	}
}
