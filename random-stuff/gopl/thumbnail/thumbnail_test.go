package thumbnail_test

import (
	"log"
	"os"
	"sync"

	"github.com/fpdevil/goprog/random-stuff/gopl/thumbnail"
)

//!+1
// makeThumbnails function makes thumbnails of specified files
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

//!-1

//!+2
// makeThumbnails2 incorrect version
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		// ignoring the errors
		go thumbnail.ImageFile(f)
	}
}

//!-2

//!+3
// makeThumbnails3 makes thumbnails of specified files in parallel.
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<-ch
	}
}

//!-3

//!+4
// makeThumbnails4 makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates
func makeThumbnails4(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}

//!-4
