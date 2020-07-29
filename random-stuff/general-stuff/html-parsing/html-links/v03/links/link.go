package links

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

// forEachNode function  calls the functions pre(x)  and post(x) for
// each  node x  in the  tree  rooted at  n. But  the functions  are
// optional
// pre(x) is called before the child nodes are visited (preorder)
// post(x) is called after (postorder)
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

func startElement(node *html.Node) {
	if node.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", node.Data)
		depth++
	}
}

func endElement(node *html.Node) {
	if node.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", node.Data)
	}
}

func outline(url string) error {
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stdout, "http.Get error %v\n", err)
		return err
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stdout, "html.Parse error %v\n", err)
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

// Extract function makes a HTTP Get call over the specified url
// and parses the html response to return the links in html doc
func Extract(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stdout, "http.Get error %v\n", err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, res.Status)
	}

	doc, err := html.Parse(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("[HTML] parsing %s: %v", url, err)
	}

	var links []string
	visitedNode := func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := res.Request.URL.Parse(a.Val)
				if err != nil {
					continue // we will ignore malformed url's
				}
				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visitedNode, nil)
	return links, nil
}
