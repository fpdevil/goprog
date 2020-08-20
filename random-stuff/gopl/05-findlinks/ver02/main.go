package main

import (
	"fmt"
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

// findLinks function does a HTTP GET request from url ans parses the
// response as HTML and extracts and returns the links
func findLinks(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("received %v: %v", url, res.StatusCode)
	}

	doc, err := html.Parse(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("html parsing error %v: %v", url, err)
	}
	return visit(nil, doc), nil
}

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findLinks: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Fprintf(os.Stdout, "%v\n", link)
		}
	}
}
