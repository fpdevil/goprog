package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	msg = `
	Concurrent Directory Traversal
	Report disk usage of the specified directories from command line
	much similar to how posix based du command works
	****************************************************************
	`
)

var (
	// to display verbose for progress
	verbose = flag.Bool("v", false, "display verbose progress messages")

	// number of files and number of bytes of each
	nfiles, nbytes int64

	// introducing concurrency
	wg sync.WaitGroup
)

func main() {
	fmt.Println(msg)

	// handle the command line arguments to take initial directories
	flag.Parse()
	roots := flag.Args()
	// in case if the input arguments are `0` we take the current directory
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// now traverse through the file tree and receive over the channel
	fileSizes := make(chan int64)

	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSizes)
	}

	wg.Wait()

	defer close(fileSizes)

	// --- for reading from the channel and getting the file info
	// get progress of the program statistics occassionally in verbose mode
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	// now read from te channel and print the results
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			showDiskUsage(nfiles, nbytes)
		}
	}

	// finally show the totals...
	showDiskUsage(nfiles, nbytes)

	// for size := range fileSizes {
	// 	nfiles++
	// 	nbytes += size
	// }
	// showDiskUsage(nfiles, nbytes)
}

// dirEntry function returns all the os level file/dir statistics and entries
// of the supplied directory dir sorted by filename.
func dirEntry(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %#v\n", err)
		return nil
	}
	return entries
}

// walkDir function wlks through the file tree rooted at the dir recursively
// and sends the size of each file found to fileSizes channel.
func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	for _, entry := range dirEntry(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func showDiskUsage(nfiles, nbytes int64) {
	gb := float64(nbytes) / 1e9
	fmt.Printf("* %d files %.2f GB\n", nfiles, gb)
}
