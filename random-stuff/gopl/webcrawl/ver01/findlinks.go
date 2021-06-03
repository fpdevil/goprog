package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/fpdevil/goprog/random-stuff/gopl/links"
)

// ver01 of the crawl crawls web links starting with the
// command line arguments

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

//!+main
func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: go run %s <url1> <url2>...\n", path.Base(os.Args[0]))
		return
	}
	// Record queue of items that need processing with each item being
	// a list of URL's to crawl. After each call to the crawl function
	// being processed in it's own goroutine, the results are sent back
	// to the worklist
	worklist := make(chan []string)

	// start with the command line arguments
	go func() {
		worklist <- os.Args[1:]
	}()

	// crawl the web concurrently
	// perform a Breadth First Search
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

//!-main

/*
//!+output
⚡
⇒  go run findlinks.go https://heroku.com
https://heroku.com
https://www.heroku.com/products
https://www.heroku.com/#skip-link
https://www.heroku.com/elements
https://www.heroku.com/what
https://www.heroku.com/home
...
ERRO[0003] url: https://appexchange.salesforce.com/appxListingDetail?listingId=a0N3u00000OMemVEAT fault: Get "https://appexchange.salesforce.com/appxListingDetail?listingId=a0N3u00000OMemVEAT": dial tcp 13.109.215.252:443: socket: too many open files
ERRO[0003] url: https://appexchange.salesforce.com/appxConsultingListingDetail?listingId=a0N30000001gKHUEA2 fault: Get "https://appexchange.salesforce.com/appxConsultingListingDetail?listingId=a0N30000001gKHUEA2": dial tcp 13.109.215.252:443: socket: too many open files
...
ERRO[0003] https://www.heroku.com/flow#video-player-1 received 429
2021/01/31 23:59:40 getting https://www.heroku.com/flow#video-player-1: 429
ERRO[0003] https://www.heroku.com/flow#skip-link received 429
2021/01/31 23:59:40 getting https://www.heroku.com/flow#skip-link: 429
...
//!-output
*/
