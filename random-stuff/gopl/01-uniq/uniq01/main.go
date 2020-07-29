package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	count := make(map[string]int)
	files := args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, count)
	} else {
		for _, f := range files {
			o, err := os.Open(f)
			if err != nil {
				fmt.Fprintf(os.Stdout, "uniq: %v\n", err)
				continue
			}

			countLines(o, count)
			o.Close()
		}
	}

	for line, i := range count {
		if i > 1 {
			fmt.Printf("%d\t%s\n", i, line)
		}
	}
}

func countLines(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}
}
