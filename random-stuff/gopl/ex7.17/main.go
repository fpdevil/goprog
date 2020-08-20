package main

// Token based XML Decoding

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "usage <prog> <url> [opt1, opt2...]\n")
		return
	}

	// dec := xml.NewDecoder(os.Stdin)
	body, err := fetch(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		return
	}

	dec := xml.NewDecoder(bytes.NewReader(body))
	var stack []string
	var attrs []map[string]string

	// tokenize the xml structure
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			return
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push the element
			attr := make(map[string]string)
			for _, a := range tok.Attr {
				attr[a.Name.Local] = a.Value
			}
			attrs = append(attrs, attr)
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop the element
			attrs = attrs[:len(attrs)-1]
		case xml.CharData:
			if containsAll(stringsToSlice(stack, attrs), args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// stringsToSlice converts the []map[string]string structure to
// a string slice like []string
// <div id="abc"> => []string{"div", "id=abc"]
func stringsToSlice(stack []string, attrs []map[string]string) []string {
	var output []string
	for i, stk := range stack {
		output = append(output, stk)
		for k, v := range attrs[i] {
			output = append(output, k+"+"+v)
		}
	}
	return output
}

// containsALl reports whether x contains the elements of y, in order
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

func fetch(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http error %s", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "read error %s", err)
		return nil, err
	}
	return body, nil
}
