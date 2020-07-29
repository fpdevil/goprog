package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Exercise 1.9: Modify to also print http status code

func main() {
	fmt.Println("URL fetching with io.Copy")

	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "url fetch: %v\n", err)
			return
		}

		fmt.Printf("HTTP Status for %s: %v\n", url, res.StatusCode)
		_, err = io.Copy(os.Stdout, res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "url %s fetch: read %v\n", url, err)
			return
		}
	}
}
