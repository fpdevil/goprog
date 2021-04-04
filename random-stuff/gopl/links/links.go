package links

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

//!+Extract

// Extract function makes an HTTP GET request to the specified URL, parses
// the response as HTML and returns the links contained in the html anchor
// nodes `<a href=` of the HTML document as a string slice.
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("url: %s fault: %v", url, err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Errorf("%s received %v", url, resp.StatusCode)
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %d", url, resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Errorf("parsing %s: %s", url, err)
		return nil, fmt.Errorf("parsing %s as html: %v", url, err)
	}

	// html links within the main parsed html page
	var links []string
	// anonymous function for visited nodes
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				// parse the URL relative to the base URL of document
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					log.Errorf("ignoring bad url: %s", url)
					continue
				}
				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visitNode, nil)
	return links, nil
}

//!-Extract

//+ forEachNode
//

// forEachNode function calls  the functions pre(x) and  post(x) for each
// node x in the tree rooted at  node n. Both these functons are optional
// pre is caled before children are visited (preorder)
// post is called after children are visited (postorder)
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

//!- forEachNode
