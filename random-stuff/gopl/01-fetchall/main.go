package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Fetching URL's concurrently...")

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Printf("%v\n", <-ch)
	}

	fmt.Printf("Total time elapsed: %.2fs\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("error reading from url (%s): %v", url, err)
		return
	}
	elapsed := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", elapsed, nbytes, url)
}
