package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// WordCounter for counting words
type WordCounter struct {
	words int
}

// LineCounter for counting lines
type LineCounter struct {
	lines int
}

func (wc *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	wc.words += count
	return len(p), nil
}

func (lc *LineCounter) Write(p []byte) (int, error) {
	var count int
	for _, b := range p {
		if b == '\n' {
			count++
		}
	}
	lc.lines += count
	return len(p), nil
}

func main() {
	s := "⇒ the quick brown fox jumps over the lazy dog"
	var wc WordCounter
	fmt.Fprintf(&wc, s)
	fmt.Printf("Number of words: %#v\n", wc)

	lc := &LineCounter{}
	b := []byte(`
	One
	Two
	Three
	Four
	`)
	lc.Write(b)
	fmt.Printf("Line Count: %#v\n", lc)
}
