// program to mimic the behaviour of unix du command
// to print the disk usage of files in a directory
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

var (
	verbose = flag.Bool("v", false, "show verbose progress messages")
	wg      sync.WaitGroup
	done    = make(chan struct{})
)

//!+cancelled
// cancelled function checks or polls the cancellation state
// at the instant it is called
func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

//!-cancelled

func main() {
	fmt.Println("** Disk Usage **")
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// this go routine will read from the standard input
	// and upon any input from the user, broadcasts the
	// cancellation by closing the done channel
	// cancel directory traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	// traverse through the file tree.
	fileSizes := make(chan int64)
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSizes)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	// print the results periodically for verbose mode
	// by generating a ticker every 500ms
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	// display the results
	var nfiles, nbytes int64

loop:
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.
			for range fileSizes {
				// perform nothing, just drain the fileSizes
			}
		case size, ok := <-fileSizes:
			if !ok {
				break loop // the fileSizes is closed, break out of loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	// for size := range fileSizes {
	// 	nfiles++
	// 	nbytes += size
	// }
	printDiskUsage(nfiles, nbytes)
}

// create a counting semaphore to prevent the error too many
// open files
var sem = make(chan struct{}, 20)

// dirents function returns the entries of a directory dir
// this version uses a counting semaphore to prevent opening
// too many files at once as the program creates thousands of
// gorotines at its peak
func dirents(dir string) []os.FileInfo {
	select {
	case sem <- struct{}{}: // acquiring the token
	case <-done:
		return nil // cancelled
	}
	defer func() {
		<-sem // releasing the token
	}()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

// walkDir function recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes
func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
