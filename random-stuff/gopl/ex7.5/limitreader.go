package main

import (
	"fmt"
	"io"
	"strings"
)

// A LimitIOReader reads from R but limits the amount of
// data returned to just N bytes. Each call to Read
// updates N to reflect the new amount remaining.
// Read returns EOF when N <= 0 or when the underlying R returns EOF.
type LimitIOReader struct {
	R io.Reader
	N int64
}

func (l *LimitIOReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

// LimitReader accepts an io Reader r and a number of bytes n
// and returns another reader reading from r but stops with
// EOF condition after n bytes.
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitIOReader{r, n}
}

func main() {
	s := `<!DOCTYPE html>
<html>
	<head>
		<title>A test page</title>
	</head>
<body>
	<h1>
		<p>Testing</p>
	</h1>
</body>
</html>
	`
	r := LimitIOReader{
		R: strings.NewReader(s),
		N: 81,
	}
	buffer := make([]byte, 1024)
	n, err := r.Read(buffer)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}
	buffer = buffer[:n]
	fmt.Printf("reading first %d bytes\n", n)
	fmt.Printf("buffer data: %s\n", string(buffer))
}
