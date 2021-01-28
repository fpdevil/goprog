package main

import (
	"fmt"
	"io"
	"os"
)

// ByteCounter stores the io writer and number of bytes written
type ByteCounter struct {
	writer  io.Writer
	written int64
}

func (bc *ByteCounter) Write(p []byte) (nb int, err error) {
	nb, err = bc.writer.Write(p)
	bc.written += int64(nb)
	return
}

// CountingWriter function as proposed in exercise
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &ByteCounter{}
	cw.writer = w
	cw.written = 0
	return cw, &cw.written
}

func main() {
	writer, written := CountingWriter(os.Stdout)
	fmt.Fprintf(writer, "Testing the CountingWriter\n")
	fmt.Printf("Count: %v\n", *written)
}
