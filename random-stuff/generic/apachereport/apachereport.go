package main

// Test regular expression at https://regexr.com
// GET[ \t]+([^ \t\n]+[.]html?)

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"text/tabwriter"
)

var workers = runtime.NumCPU()

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <apache.log>\n", filepath.Base(os.Args[0]))
		return
	}
	lines := make(chan string, workers*4)
	results := make(chan map[string]int, workers)
	go readLines(os.Args[1], lines)
	getRegex := regexp.MustCompile(`GET[ \t]+([^ \t\n]+[.]html?)`)
	for i := 0; i < workers; i++ {
		go processLines(results, lines, getRegex)
	}
	totalPerPage := make(map[string]int)
	merge(results, totalPerPage)
	showResults(totalPerPage)
}

func readLines(filename string, lines chan<- string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("unable to open the file:", err)
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if line != "" {
			lines <- line
		}

		if err != nil {
			if err != io.EOF {
				log.Println("failed to finish reasding the file:", err)
			}
			break
		}
	}
	close(lines)
}

func processLines(results chan<- map[string]int, lines <-chan string, getRegex *regexp.Regexp) {
	countsPerPage := make(map[string]int)
	for line := range lines {
		if matches := getRegex.FindStringSubmatch(line); matches != nil {
			// first item of matches is the entire matching text and each
			// subsequent item corresponds to the text that matches a
			// parenthesized subexpression in the regex.
			countsPerPage[matches[1]]++
		}
	}
	results <- countsPerPage
}

func merge(results <-chan map[string]int, totalPerPage map[string]int) {
	for i := 0; i < workers; i++ {
		countsPerPage := <-results
		for page, count := range countsPerPage {
			totalPerPage[page] += count
		}
	}
}

func showResults(totalPerPage map[string]int) {
	// initialize tabwriter
	const format = "%4v\t%v\t\n"
	tw := new(tabwriter.Writer)
	tw.Init(os.Stdout, 0, 8, 2, ' ', 0)
	defer tw.Flush()
	fmt.Fprintf(tw, format, "Counts", "Page Link")
	fmt.Fprintf(tw, format, "------", "---------")
	for page, count := range totalPerPage {
		fmt.Fprintf(tw, format, count, page)
	}
}
