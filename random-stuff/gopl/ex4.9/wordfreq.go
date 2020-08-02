package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freq := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		freq[in.Text()]++
	}

	if in.Err() != nil {
		fmt.Fprintf(os.Stderr, "%v", in.Err())
		return
	}

	for w, count := range freq {
		fmt.Printf("%-30v %d\n", w, count)
	}
}
