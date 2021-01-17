package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Println("URL Fetching...")
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "url fetch: %v\n", err)
			return
		}

		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "url fetch: reading %s: %v\n", url, err)
			return
		}
		fmt.Println(string(body))
	}
}
