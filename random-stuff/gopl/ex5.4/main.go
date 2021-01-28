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
// attribute, images, scripts and stylesheets for each anchor and appends
// them to a string slice and returns the same
func visit(data map[string][]string, node *html.Node) {
	if node.Type == html.ElementNode {
		switch node.Data {
		case "a":
			for _, a := range node.Attr {
				if a.Key == "href" {
					data["links"] = append(data["links"], a.Val)
				}
			}
		case "img", "scripts":
			for _, i := range node.Attr {
				if i.Key == "src" {
					data["imgscr"] = append(data["imgscr"], i.Val)
				}
			}
		case "link":
			for _, l := range node.Attr {
				if l.Key == "media" {
					data["styles"] = append(data["styles"], l.Val)
				}
			}
		}
	}

	// call recursively for the rest of html nodes
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		visit(data, c)
	}
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

	data := make(map[string][]string)
	visit(data, doc)
	for key, val := range data {
		fmt.Println("----------------------------------------")
		fmt.Printf("[[%s]]\n", key)
		fmt.Println("----------------------------------------")
		for i, v := range val {
			fmt.Printf("%3d %s\n", i, v)
		}
	}
}

//!-exec
