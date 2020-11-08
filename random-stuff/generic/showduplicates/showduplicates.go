package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"sync"
)

const (
	msg = `
	* Finding Duplicate files *
	`
	usage = `
	usage: %s <path>
	`
	maxGoroutines      = 100       // maximum spawned
	maxSizeOfSmallFile = 1024 * 32 // 32Kb
)

type pathsInfo struct {
	size  int64    // file size
	paths []string // file path
}

type fileInfo struct {
	sha1 []byte // SHA1 checksum of file
	size int64  // file size
	path string // file path
}

func main() {
	fmt.Println(msg)
	if len(os.Args) == 1 || os.Args[1] == "--h" || os.Args[1] == "--help" {
		fmt.Printf(usage, filepath.Base(os.Args[0]))
		return
	}

	infoChan := make(chan fileInfo, maxGoroutines*2)
	go findDuplicates(infoChan, os.Args[1])
	pathData := mergeResults(infoChan)
	renderResults(pathData)
}

func findDuplicates(infoChan chan fileInfo, dirname string) {
	waiter := &sync.WaitGroup{}
	filepath.Walk(dirname, makeWalkFunc(infoChan, waiter))
	waiter.Wait()
	close(infoChan)
}

// makeWalkFunc function creates and returns an anonymous function
// of type filepath.WaitFunc that would be called for every file
// and directory that the filepath.Walk() function encounters.
// check if a link or a regular file (nonzero size) with the below
// - 	info.Mode()&os.ModeType == 0
// once the computation of the SHA1 is complete for a file, the
// results will sent over to the channel infoChan
func makeWalkFunc(infoChan chan fileInfo, waiter *sync.WaitGroup) func(string, os.FileInfo, error) error {
	return func(path string, fi os.FileInfo, err error) error {
		if err == nil && fi.Size() > 0 && (fi.Mode()&os.ModeType == 0) {
			if fi.Size() < maxSizeOfSmallFile || runtime.NumGoroutine() > maxGoroutines {
				// file size is small < 32Kb, calculate SHA1
				processFile(path, fi, infoChan, nil)
			} else {
				// file size exceeds 32Kb or the number of current goroutines
				// exceed the maximum, run processFile to get SHA1 as a separate
				// goroutine to run asynchronously
				waiter.Add(1)
				// pass anonymous function to the goroutine to call the
				// sync.WaitGroup.Done()
				go processFile(path, fi, infoChan, func() { waiter.Done() })
			}
		}
		return nil // ignore all errors
	}
}

func processFile(filename string, info os.FileInfo, infoChan chan fileInfo, done func()) {
	if done != nil {
		defer done()
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Println("error:", err)
		return
	}

	defer file.Close()

	// Open a new SHA1 hash interface to write to
	hash := sha1.New()

	// Copy the file in the hash interface and check for any error
	size, err := io.Copy(hash, file)
	if size != info.Size() || err != nil {
		if err != nil {
			log.Println("error:", nil)
		} else {
			log.Println("error: unable to read the whole file:", filename)
		}
		return
	}
	// Sum method returns the 20-byte SHA1 hash value
	infoChan <- fileInfo{hash.Sum(nil), info.Size(), filename}
}

func mergeResults(infoChan <-chan fileInfo) map[string]*pathsInfo {
	pathData := make(map[string]*pathsInfo)
	// create a format string with 16 zero-padded hexadecimal digits
	// to represent file's size and enough hex digits representing SHA1
	format := fmt.Sprintf("%%016X:%%%dX", sha1.Size*2) // == "%016X:%40X"
	for info := range infoChan {
		key := fmt.Sprintf(format, info.size, info.sha1)
		value, ok := pathData[key]
		if !ok {
			// create and add a suitable value with the given file size
			// and with an empty slice of paths
			value = &pathsInfo{size: info.size}
			pathData[key] = value
		}
		// duplicate files will have more than one paths
		value.paths = append(value.paths, info.path)
	}
	return pathData
}

func renderResults(pathData map[string]*pathsInfo) {
	keys := make([]string, 0, len(pathData))
	for key := range pathData {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := pathData[key]
		if len(value.paths) > 1 {
			fmt.Printf("%d duplicate files (%s bytes):\n", len(value.paths), delimiter(value.size, ","))
			sort.Strings(value.paths)
			for _, name := range value.paths {
				fmt.Printf("\t%s\n", name)
			}
		}
	}
}

// delimiter function returns a string representating the
// whole number with comma grouping for himan readability
func delimiter(x int64, delim string) string {
	value := fmt.Sprint(x)
	for i := len(value) - 3; i > 0; i -= 3 {
		value = value[:i] + delim + value[i:]
	}
	return value
}
