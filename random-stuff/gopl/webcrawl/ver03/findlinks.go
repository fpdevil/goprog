// Crawl3 crawls the web links starting with command line arguments
//
// This version of crawl uses a bounded parallelism.
//
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fpdevil/goprog/random-stuff/gopl/links"
)

// crawl function crawls over the web links provded as argument
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string) // list of URL's may have duplicates
	unseen := make(chan string)     // de-duplicated url's

	// add command line arguments to worklist
	go func() {
		worklist <- os.Args[1:]
	}()

	// create 20 crawler goroutines to fetch each unseen link
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseen {
				found := crawl(link)
				go func() {
					worklist <- found
				}()
			}
		}()
	}

	// the main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseen <- link
			}
		}
	}
}
