package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

//!+visit
// visit function traverses a html node tree, extracts the links from
// the href attribute of eacg anchor element `<a href=...` and appends
// the url value of href to a string slice and returns the same
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

//!-visit

//!+findLinks
// findLinks function does a HTTP GET request from url and parses the
// response as HTML and then extracts and returns the links
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
	return visit([]string{}, doc), nil
}

//!-findLinks

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findLinks: %v\n", err)
			continue
		}
		for i, link := range links {
			// fmt.Fprintf(os.Stdout, "%v\n", link)
			fmt.Printf("*%4d: %v\n", i, link)
		}
	}
}
