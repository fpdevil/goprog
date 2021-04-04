package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fpdevil/goprog/random-stuff/gopl/links"
)

// The 02 version of Crawl crawls the web links starting with the
// command line arguments.
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract
//

//!+semaphore
// tokens is a counting semapore used to enforce a limit of 20
// concurrent requests at a time.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{} // acquire a token or a place
	list, err := links.Extract(url)

	// release the token now
	<-tokens

	if err != nil {
		log.Print(err)
	}

	return list
}

//!-semaphore

//!+main
func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// start with the command line arguments
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	// crawl the web concurrently
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

//!-main
