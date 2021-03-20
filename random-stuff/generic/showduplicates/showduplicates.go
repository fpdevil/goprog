// package showduplicates filters  all the files which have  the same MD5
// checksum and lists them as duplicate files
//
// The SHA-1 secure Hash algorithm produces a 20-byte value for any given
// chunk of data, such as a file. Files which are identical will have the
// same  SHA-1  values  and  different  files  will  almost  always  have
// different SHA-1 values.

package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"sync"
)

const (
	msg                = `*usage: %s <path>*` // usage/error message
	maxSizeOfSmallFile = 32 * 1024            // 32Kb cap for small files
	maxGoroutines      = 100                  // maximum number of go routines
)

// fileInfo struct is for summarizing the metadata information of each file
// if 2 files SHA1 and sizes are same then they are both equal
type fileInfo struct {
	sha1 []byte
	size int64
	path string
}

// pathsInfo struct is for storing the details of each duplicate file and
// it holds the file size and all files paths and filenames
type pathsInfo struct {
	size  int64
	paths []string
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // using all the available cores
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf(msg, filepath.Base(os.Args[0]))
		return
	}

	wg := &sync.WaitGroup{}
	infoChan := make(chan fileInfo, maxGoroutines*2)
	go findDuplicates(infoChan, os.Args[1], wg)
	pathData := mergeResults(infoChan)
	outputResults(pathData)
}

//!+ findDuplicates

// findDuplicates function  calls the  filepath.WalkDir to walk  over the
// directory(s) tree  (starting from the  dirname) and for each  file and
// directory it calls the fs.WalkDirFunc Func function passed as a second
// argument
// The function is internal and is not exported
func findDuplicates(infoChan chan fileInfo, dirname string, waiter *sync.WaitGroup) {
	filepath.WalkDir(dirname, walkDirFunc(infoChan, waiter))
	waiter.Wait() // this blocks untill all the work is done
	defer close(infoChan)
}

//!-

//!+ walkDirFunc

// walkDirFunc  function  creates  and   returns  an  anonymous  function
// complying  to func(string,  fs.DirEntry, error)  that would  return an
// error.  The anonymous  function would  be  called for  every file  and
// directory that the filepath.WalkDir() function encounters
func walkDirFunc(infoChan chan fileInfo, waiter *sync.WaitGroup) func(string, fs.DirEntry, error) error {
	// create and return the callback function that will be used in WalkDir
	callback := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		info, err := d.Info() // fetch the file info of file/dir
		// process only the regular files of non-zero size
		// fs.ModeType  provides a  bitmask  value which  has bits  set
		// for directories,  symbolic links,  named pipes,  sockets and
		// devices; if none  of these are set then it's  a regular file
		// which we can process
		if err == nil && info.Size() > 0 && getType(info.Mode()) == 'f' {
			if info.Size() < maxSizeOfSmallFile || runtime.NumGoroutine() > maxGoroutines {
				// file size is less than 32Kb, calculate  SHA-1 directly
				processFile(path, info, infoChan, nil)
			} else {
				// if files are bigger than 32Kb delegate computation to
				// separate goroutines
				waiter.Add(1)
				go processFile(path, info, infoChan, func() { waiter.Done() })
			}
		}
		return nil
	}
	return callback
}

//!-

//!+ processFile

// processFile function  processes each  file at  the specified  path and
// sends the results to the input channel of type fileInfo
func processFile(filename string, info fs.FileInfo, infoChan chan fileInfo, done func()) {
	// if done is not nil, the function is invooked as a goroutine
	if done != nil {
		defer done()
	}

	hash := sha1.New()
	file, err := os.Open(filename)
	if err != nil {
		log.Println("error:", err)
		return
	}

	defer file.Close()
	size, err := io.Copy(hash, file)
	if size != info.Size() || err != nil {
		if err != nil {
			log.Println("error:", err)
		} else {
			log.Println("error: failed to read the whole file:", filename)
		}
		return
	}

	infoChan <- fileInfo{hash.Sum([]byte("")), info.Size(), filename}
}

//!-

//!+ mergeResults

// mergeResults function  takes the  fileInfo channel  and returns  a map
// that stores duplicate files
func mergeResults(infoChan <-chan fileInfo) map[string]*pathsInfo {
	pathData := make(map[string]*pathsInfo)
	format := fmt.Sprintf("%%016X:%%%dX", sha1.Size*2)
	for info := range infoChan {
		key := fmt.Sprintf(format, info.size, info.sha1)
		value, ok := pathData[key]
		if !ok {
			value = &pathsInfo{size: info.size}
			pathData[key] = value
		}
		value.paths = append(value.paths, info.path)
	}
	return pathData
}

//!-

//!+ outputResults

func outputResults(pathData map[string]*pathsInfo) {
	keys := make([]string, 0, len(pathData))
	for key := range pathData {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := pathData[key]
		if len(value.paths) > 1 {
			fmt.Printf("%d duplicate files [%s bytes]:\n", len(value.paths), commas(value.size))
			sort.Strings(value.paths)
			for _, name := range value.paths {
				fmt.Printf("\t%s\n", name)
			}
		}
	}
}

//!-

//!+ commas

// commas function returns a string representing the whole number grouped
// by commas
func commas(x int64) string {
	value := fmt.Sprint(x)
	for i := len(value) - 3; i > 0; i -= 3 {
		value = value[:i] + "," + value[i:]
	}
	return value
}

//!-

//!+ getType

// getType function takes the file's mode and permission bits which is
// represented by the type fs.FileMode and returns the corresponding type
func getType(mode fs.FileMode) byte {
	switch {
	case mode.IsDir():
		return 'd' // is directory
	case mode.IsRegular():
		return 'f' // is regular file
	case mode&fs.ModeSymlink != 0:
		return 'l' // is a symbolic link
	case mode&fs.ModeSocket != 0:
		return 's' // is a posix socket
	case mode&fs.ModeNamedPipe != 0:
		return 'p' // is a posix pipe
	case mode&fs.ModeDevice != 0:
		// is a device file
		if mode&fs.ModeCharDevice != 0 {
			return 'C'
		}
		return 'D'
	default:
		// is a directory
		return 'd'
	}
}

//!-
