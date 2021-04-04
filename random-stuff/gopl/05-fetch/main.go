package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

//!+fetch

// fetch function downloads the URL and returns the name and
// length of the local file
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)
	// now close the file, but prefer error from Copy if any
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	return
}

//!-fetch

//!+exec
func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "usage go run %s [<url1>, <url2>...]\n", path.Base(os.Args[0]))
	}
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err.Error())
			continue
		}

		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
	}
}

//!-exec
