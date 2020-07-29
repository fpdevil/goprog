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
	Print the structure of HTML NodeTree outline by recursively
	looping over the html document structure
	***********************************************************
	`
)

var (
	depth int
	// depth = 2
)

func main() {
	fmt.Println(msg)

	body, err := fetch(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		return
	}

	outline([]string(nil), doc)

	fmt.Println("---------------------------------")
	for _, url := range os.Args[1:] {
		toutline(url)
	}
}

// forEachNode function calls the functions  pre(x) and post(x) for every
// node  x in  the  DOM tree  rooted  at node.  Both  of these  functions
// are  optional; pre  is  called  prior to  the  children being  visited
// (pre-prder) and post is called after being visited (post-order)
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

// startElem function will print the start html tag
func startElem(node *html.Node) {
	if node.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", node.Data)
		depth++
	}
}

// endElem function will print the end html tag
func endElem(node *html.Node) {
	if node.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", node.Data)
	}
}

func toutline(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElem, endElem)

	return nil
}

func outline(stack []string, node *html.Node) {
	if node.Type == html.ElementNode {
		stack = append(stack, node.Data)
		fmt.Println(stack)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
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
