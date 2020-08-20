package main

import (
	"bytes"

	"golang.org/x/net/html"

	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func elemMap(r io.Reader) (map[string]int, error) {
	emap := make(map[string]int, 0)
	var err error
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		name, _ := z.TagName()
		if len(name) > 0 {
			emap[string(name)]++
		}
	}
	if err != io.EOF {
		return emap, err
	}
	return emap, err
}

func fetch(urls []string) ([]byte, error) {
	var b []byte
	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error %v from %s\n", err, url)
			return nil, err
		}
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error %v\n", err)
		}
		b = append([]byte(nil), body...)
	}
	return b, nil
}

func main() {
	body, err := fetch(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	emap, err := elemMap(bytes.NewReader(body))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	for t, count := range emap {
		fmt.Printf("%-5d %-10s\n", count, t)
	}
}
