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

func (bc *ByteCounter) Write(p []byte) (int, error) {
	bc.writer.Write(p)
	bc.written += int64(len(p))
	return len(p), nil
}

// CountingWriter function as proposed in exercise
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := ByteCounter{w, 0}
	return &cw, &cw.written
}

func main() {
	writer, written := CountingWriter(os.Stdout)
	fmt.Fprintf(writer, "Testing the CountingWriter\n")
	fmt.Printf("Count: %v\n", *written)
}
