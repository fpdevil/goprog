package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	msg = `
	Error Handling in concurrent programs
	Handling of errors in a sane way while making concurrent
	requests to multiple urls.
	========================================================
`
	n = 3 // number of errors to stop at

	timeout = time.Duration(5 * time.Second) // set custom http timeout
)

// Result struct for holding the error and response snapshots
type Result struct {
	Error    error
	URL      string
	Response *http.Response
}

func main() {
	fmt.Println(msg)
	fmt.Println()

	done := make(chan interface{})
	defer close(done)

	urls := []string{
		"http://httpbin.org",
		"https://www.rust-lang.org",
		"http://1.2.3.4",
		"a",
		"b",
		"c",
		"https://www.opera.com",
	}

	var errCount int
	for result := range getStatus(done, urls...) {
		if result.Error != nil {
			errCount++
			fmt.Printf("error from {%s}: %v\n", result.URL, result.Error)
			if errCount == n {
				fmt.Print("Too many errors... breaking the call!\n")
				break
			}
			continue
		}

		fmt.Printf("Response from {%s}: %#v\n", result.URL, result.Response.Status)
	}
}

// getStatus takes a receiver channel of type empty interface and a variadic
// list of  multiple  url's to make concurrent connections to.  It returns a
// a receiver channel of the type Result which can be handled by the caller
func getStatus(done <-chan interface{}, urls ...string) <-chan Result {
	results := make(chan Result)
	go func() {
		defer close(results)

		client := &http.Client{
			Timeout: timeout,
		}
		for _, url := range urls {
			var result Result
			response, err := client.Get(url)
			result = Result{
				Error:    err,
				URL:      url,
				Response: response,
			}

			select {
			case <-done:
				return
			case results <- result:
			}
		}
	}()
	return results
}
