package main

import (
	"fmt"
	"os"

	"github.com/fpdevil/goprog/random-stuff/gopl/links"
	log "github.com/sirupsen/logrus"
)

const (
	msg = `
	* Finding links in pages recursively *

	`
)

//!+breadthFirst
// breadthFirst function  calls anonymous function  f for each item  in the
// worklist. Any items returned by f are then added to the worklist. f is
// called at the most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	visited := make(map[string]bool) // keep track of all the visited links
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !visited[item] {
				visited[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//+crawl
// crawl function will be supplied to the breadthFirst function as argument
// the function prints the URL, extracts its links and returns them so
// that they too are visited...
func crawl(url string) []string {
	log.Printf("invoking the url %v", url)
	list, err := links.Extract(url)
	if err != nil {
		log.Infof("extracting from %s: %v", url, err)
	}
	return list
}

//!-crawl

func main() {
	// crawl the web in breadth first by starting from the
	// command line arguments supplied
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: <program> <link1> [<link2>...]")
		return
	}
	breadthFirst(crawl, args[1:])
}
