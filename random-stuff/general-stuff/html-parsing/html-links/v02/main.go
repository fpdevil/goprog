package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

const (
	msg = `
	Getting all the html links from the listed url's
	************************************************

	`
)

func main() {
	fmt.Println(msg)

	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// visit function traverses a html document from a link get the href
// attributes of each enchor element
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

// findLinks function does all at one place, like making a HTTP Get call
// request for url, parse the response as a HTML and extract and then
// finally return the individual links
func findLinks(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("fetching %s: %s", url, res.Status)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

// fetch function makes a HTTP Get call to a list of urls and fetches
// the html content and stores in a byte list
func fetch(urls []string) ([]byte, error) {
	var byteb []byte
	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			return nil, err
		}

		// _, err = io.Copy(body, res.Body)
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %s: %v\n", url, err)
			return nil, err
		}

		byteb = append([]byte(nil), body...)
	}
	return byteb, nil
}
