package main

import (
	"fmt"
	"os"

	"github.com/fpdevil/goprog/random-stuff/general-stuff/html-parsing/html-links/v03/links"

	log "github.com/sirupsen/logrus"
)

const (
	msg = `
	Getting all the html links recursively from the
	listed command line argument url's...
	************************************************

    `
)

func main() {
	fmt.Println(msg)
	args := os.Args

	// queue of items which need processing
	// list if URL's to crawl including duplicates if any
	workList := make(chan []string)
	// unique url's
	uniqueLinks := make(chan string)

	go func() {
		workList <- args[1:]
	}()

	// create 20 crawler goroutines to fetch each un visited link
	for i := 0; i < 20; i++ {
		go func() {
			for link := range uniqueLinks {
				foundLinks := crawl(link)
				go func() {
					workList <- foundLinks
				}()
			}
		}()
	}

	// now crawl the web concurrently
	visited := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !visited[link] {
				visited[link] = true
				uniqueLinks <- link
			}
		}
	}

	// perform web crawling using the Breadth First Search
	// using the command line args
	// bfs(crawl, os.Args[1:])
}

// bfs function  uses BreadFirstSearch  to call  f for  each item  in the
// worklist; any items returned  by f will be added to  worklist and f is
// called at the most once for each item
func bfs(fn func(item string) []string, worklist []string) {
	covered := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !covered[item] {
				covered[item] = true
				worklist = append(worklist, fn(item)...)
			}
		}
	}
}

// crawl function crawls over the html to get information
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}
