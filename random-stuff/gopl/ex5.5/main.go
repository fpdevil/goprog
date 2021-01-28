package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/net/html"
)

const (
	timeout = time.Duration(2 * time.Second)
)

//!+initClient

// initClient function creates a client context with specified
// timeout in seconds explicitly
func initClient() http.Client {
	transport := http.Transport{
		DialContext: (&net.Dialer{
			Timeout: timeout,
		}).DialContext,
	}

	client := http.Client{
		Transport: &transport,
	}

	return client
}

//!-initClient

//!+CountWordsAndImages

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	// resp, err := http.Get(url)
	// specify a timeout of 2 seconds
	client := initClient()
	resp, err := client.Get(url)
	if err != nil {
		err = fmt.Errorf("http GET error %s: %s", url, err.Error())
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("error parsing html: %s", err.Error())
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

//!-CountWordsAndImages

//!+countWordsAndImages

// countWordsAndImages function takes a html node, parses and
// returns the number of words and images
func countWordsAndImages(node *html.Node) (words, images int) {
	if node.Type == html.ElementNode {
		if node.Data == "img" {
			images++
		} else if node.Data == "style" || node.Data == "script" {
			return
		}
	} else if node.Type == html.TextNode {
		text := strings.TrimSpace(node.Data)
		for _, line := range strings.Split(text, "\n") {
			if line != "" {
				words += len(strings.Fields(line))
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		word, image := countWordsAndImages(c)
		words += word
		images += image
	}

	return
}

//!-countWordsAndImages

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <url>", filepath.Base(os.Args[0]))
	}

	wiMap := make(map[string][]map[string]interface{})

	for _, url := range os.Args[1:] {
		m := make(map[string]interface{})
		w, i, err := CountWordsAndImages(url)
		if err != nil {
			m["error"] = err.Error()
		}
		m["words"] = w
		m["images"] = i
		wiMap[url] = append(wiMap[url], m)
	}

	for key, value := range wiMap {
		fmt.Printf("%s\n", strings.Repeat("-", 48))
		fmt.Printf("[url: %s]\n", key)
		fmt.Printf("%s\n", strings.Repeat("-", 48))
		for _, m := range value {
			for k, v := range m {
				fmt.Printf("# %-8s: %-3v\n", k, v)
			}
		}
	}
}
