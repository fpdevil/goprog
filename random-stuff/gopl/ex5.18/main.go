package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		log.Errorln("provide input urls for parsing")
		return
	}

	url, tagnames := os.Args[1], os.Args[2:]
	body, err := fetch(url)
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	nodes := ElementsByTag(doc, tagnames...)

	for _, node := range nodes {
		fmt.Printf("%v\n", node)
	}
}

func ElementsByTag(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	pre := func(node *html.Node) {
		if node.Type == html.ElementNode {
			for _, n := range name {
				if node.Data == n {
					nodes = append(nodes, node)
				}
			}
		}
	}
	forEachNode(doc, pre, nil)
	return nodes
}

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

func fetch(urls ...string) ([]byte, error) {
	var b []byte
	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			log.Errorf("http GET %s: %v", url, err)
			return nil, err
		}

		if res.StatusCode != http.StatusOK {
			log.Errorf("http status %s: %d", url, res.StatusCode)
			return nil, fmt.Errorf("read failed: %s", res.Status)
		}

		body, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			log.Errorf("read fail %v", err)
			return nil, fmt.Errorf("error %v", err)
		}
		b = append([]byte(nil), body...)
	}
	return b, nil
}
