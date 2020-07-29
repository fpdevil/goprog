package main

import (
	"bufio"
	"fmt"
	"os"
)

// Found is the matching string
type Found string

func main() {
	args := os.Args
	count := make(map[string]int)       // map of words to their count
	countf := make(map[string][]string) // map of words to list of file names
	files := args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, count, countf)
	} else {
		for _, f := range files {
			// Open each file and get a pointer to the file
			o, err := os.Open(f)
			if err != nil {
				fmt.Fprintf(os.Stdout, "uniq: %v\n", err)
				continue
			}

			// run function to count the words and files
			countLines(o, count, countf)
			o.Close()
		}
	}

	for line, i := range count {
		if i > 1 {
			fmt.Printf("%d\t%v\t%s\n", i, countf[line], line)
		}
	}
}

func (in Found) isIn(str []string) bool {
	for _, s := range str {
		if in == Found(s) {
			return true
		}
	}
	return false
}

func countLines(f *os.File, count map[string]int, countf map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		count[line]++
		if !Found(f.Name()).isIn(countf[line]) {
			countf[line] = append(countf[line], f.Name())
		}
	}
}
