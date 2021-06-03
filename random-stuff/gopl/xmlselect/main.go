// Package xmlselect is for decoding the XML. It prints the
// text of selected elements of an XML documet.
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//!+
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage: go run %s [<url1> <url2>...]\n", filepath.Base(os.Args[0]))
		return
	}

	sb, err := fetch(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error from fetch: %s", err.Error())
		return
	}

	in := bytes.NewReader(sb)
	dec := xml.NewDecoder(in)
	var stack []string // keep a stack of element names
	for {
		token, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %#v\n", err)
			return
		}

		switch token := token.(type) {
		case xml.StartElement:
			stack = append(stack, token.Name.Local) // push it to stack
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop the item out of stack
		case xml.CharData:
			if containsAll(stack, args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), token)
			}
		}
	}
}

//!-

//!+containsAll

// containsAll function reports whether x contains the elements of
// y, all in order or not
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

//!-containsAll

//#+!fetch
// fetch function will loop over a list of http url links provided
// as argument and makes a `GET` call to each to fetch the contetnt
// and store them into a byte slice to return the same

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error calling url %s: %s", url, err.Error())
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading the response: %s", err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	// aggregate all the response bytes from all url's
	return body, nil
}

//!-fetch
