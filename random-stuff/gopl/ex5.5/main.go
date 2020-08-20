package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}

	if (n.Type == html.TextNode) && (n.Parent.Data != "script" && n.Parent.Data != "style") {
		data := strings.Fields(strings.TrimSpace(n.Data))
		words += len(data)
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	fwords, fimages := countWordsAndImages(n.FirstChild)
	nwords, nimages := countWordsAndImages(n.NextSibling)

	words += fwords + nwords
	images += fimages + nimages

	return
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: <prog> <url>")
	}
	url := os.Args[1]
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("words: %d\nimages: %d\n", words, images)
}
