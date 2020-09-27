package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// Extract function makes GET call to url, parses the response as
// HTML and returns the links in the HTML document
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visited := func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visited, nil)
	return links, nil
}

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

// enforce a limit of 20 concurrent requests
var tokens = make(chan struct{}, 20)

// crawl function crawls over the web page using the submitted url
func crawl(url string) []string {
	fmt.Printf("url: %s\n", url)
	tokens <- struct{}{} // get a token
	list, err := Extract(url)
	<-tokens // release the token
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil
	}
	return list
}

func crawlOld(url string) []string {
	fmt.Printf("url: %s\n", url)
	list, err := Extract(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// start with the command line arguments
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}

	// for list := range worklist {
	// 	for _, link := range list {
	// 		if !seen[link] {
	// 			seen[link] = true
	// 			go func(link string) {
	// 				worklist <- crawl(link)
	// 			}(link)
	// 		}
	// 	}
	// }
}
