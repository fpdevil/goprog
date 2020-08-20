package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

// StringReader is type containing a string whihc should
// satisfy the Read method from io.Reader
type StringReader struct {
	str string
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
	n = copy(p, []byte(sr.str))
	sr.str = sr.str[n:]
	fmt.Println(sr)
	if n == 0 {
		err = io.EOF
	}
	return
}

// NewReader is a custom implementation analogous to strings.NewReader
func NewReader(s string) io.Reader {
	fmt.Println(s)
	return &StringReader{s}
}

func main() {
	shtml := "<html><body><h1><p>Testing</p></h1></body></html>"
	doc, err := html.Parse(NewReader(shtml))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}
	fmt.Println(doc)
}
