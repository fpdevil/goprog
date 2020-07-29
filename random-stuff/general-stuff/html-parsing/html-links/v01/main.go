package main

import (
	"bytes"
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

	body, err := fetch(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	// fmt.Printf("%s\n", body)

	doc, err := html.Parse(bytes.NewReader(body))
	// doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		return
	}

	for _, link := range visit([]string{}, doc) {
		fmt.Printf("%v\n", link)
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
