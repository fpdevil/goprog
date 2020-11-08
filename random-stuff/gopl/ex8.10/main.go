package main

import (
	"fmt"
	"os"

	"github.com/fpdevil/goprog/random-stuff/gopl/ex8.10/links"
)

var (
	cancel = make(chan struct{})
	tokens = make(chan struct{}, 20) // 20 concurrent requests
)

func main() {
	fmt.Println("Ex: 8.10 Web Crawler")

	// list of urls
	worklist := make(chan []string)
	// number of pending sends to worklist
	var n int

	// start with the command line arguments
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	// concurrent web crawling
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
}

func crawl(url string) []string {
	fmt.Printf("url: %s\n", url)
	tokens <- struct{}{} // acquire token
	list, err := links.Extract(url, cancel)
	<-tokens // release the token

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil
	}
	return list
}
