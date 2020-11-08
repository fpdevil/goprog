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

// DirInfo struct stores root name and size
type DirInfo struct {
	id   int   // identifier for root directory
	size int64 // size of the root directory
}

var (
	wg      sync.WaitGroup
	verbose = flag.Bool("v", false, "show verbose progress messages")
)

func main() {
	fmt.Println("** Disk Usage")
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	dirstats := make(chan *DirInfo)

	// traverse through the directories
	for i, root := range roots {
		wg.Add(1)
		go walkDir(root, i, &wg, dirstats)
	}
	go func() {
		wg.Wait()
		defer close(dirstats)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	// track number of files and number of bytes
	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))

loop:
	for {
		select {
		case ds, ok := <-dirstats:
			if !ok {
				break loop
			}
			nfiles[ds.id]++
			nbytes[ds.id] += ds.size
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes)
		}
	}
	printDiskUsage(roots, nfiles, nbytes)
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for i, root := range roots {
		fmt.Printf("%-20s => %.1f GB for %-7dfiles\n", root, float64(nbytes[i])/1e9, nfiles[i])
	}
}

// sem is a counting semaphore for limiting concurrency in dirents.
// we will limit the parallelism using a buffered channel with capacity
// of 20; sending to channel acquires the token
//        receiving from channel releases the token

var sem = make(chan struct{}, 20)

// dirents function reads the directory named by dir and
// returns a list of directory entries
func dirents(dir string) []os.FileInfo {
	sem <- struct{}{} // acquire the token
	defer func() {
		<-sem // release token
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
func walkDir(dir string, id int, wg *sync.WaitGroup, dirstats chan<- *DirInfo) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, id, wg, dirstats)
		} else {
			dirstats <- &DirInfo{id: id, size: entry.Size()}
		}
	}
}
