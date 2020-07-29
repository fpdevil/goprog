package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("another variant of uniq")
	fmt.Println()

	count := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// returns bytes, err
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "uniq error: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(b), "\n") {
			count[line]++
		}
	}

	for line, i := range count {
		if i > 1 {
			fmt.Printf("%d\t%v\n", i, line)
		}
	}
}
