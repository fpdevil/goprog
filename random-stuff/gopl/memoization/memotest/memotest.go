// Package memotest provides some common functions
// for testing variants of the memo package
package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

// HTTPGetBody holds a reference to the function httpGetBody
var HTTPGetBody = httpGetBody

// M interface has a single method Get which needs to be satisfied
type M interface {
	Get(key string) (interface{}, error)
}

//!+httpGetBody

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//!-httpGetBody

//!+fillURLs
func fillURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

//!-fillURLs

//!+Sequential

func Sequential(t *testing.T, m M) {
	for url := range fillURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(url)
			continue
		}
		fmt.Printf("* %-30s, %-5.15s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

//!-Sequential

//!+Concurrent

func Concurrent(t *testing.T, m M) {
	var wg sync.WaitGroup
	for url := range fillURLs() {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("* %-30s, %-5.15s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	wg.Wait()
}

//!-Concurrent
