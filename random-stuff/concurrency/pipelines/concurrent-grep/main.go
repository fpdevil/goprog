package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	log "github.com/sirupsen/logrus"
)

// Result represents the structure of the response from running grep
type Result struct {
	// the name of the file
	filename string
	// current matching line
	line string
	// matching line number
	lineno int
}

// Job represents the current task of grepping and sending results
type Job struct {
	filename string
	results  chan<- Result
}

const (
	msg = `* Concurrent Grep over Files *

	`
)

var (
	workers runtime.NumCPU()
)

func main() {
	fmt.Println(msg)

	runtime.GOMAXPROCS(runtime.NumCPU())
	args := os.Args
	if len(args) < 3 || args[1] == "-h" || args[1] == "--help" {
		fmt.Printf("Usage: %s <regexp> <fiels...>\n", filepath.Base(args[0]))
		return
	}

	linerx, err := regexp.Compile(args[1])

	if err != nil {
		log.Printf("invalid regexp: %v\n", err)
		return
	}

	grep(linerx, cmdlineFiles(args[2:]))
}

func grep(line *regexp.Regexp, filenames []string) {
	jobs := make(chan Job, workers)
	results := make(chan Result, minimum(1000, len(filenames)))
	done := make(chan struct{}, workers)

}