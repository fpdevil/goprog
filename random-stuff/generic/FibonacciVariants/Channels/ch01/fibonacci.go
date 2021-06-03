package main

import (
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/fpdevil/goprog/random-stuff/generic/helper"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage go run %s <limit>", filepath.Base(os.Args[0]))
		return
	}
	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("invalid input %v %s\n", limit, err)
		return
	}

	defer helper.Trace("main")()
	start := time.Now()
	ch := make(chan *big.Int, 2) // buffered channel of capacity 2
	defer close(ch)
	ch <- big.NewInt(0)
	ch <- big.NewInt(1)

	fmt.Printf("Fibonacci sequence from 0 to %v\n", limit)
	for z := 0; z <= limit; z++ {
		n := <-ch
		fmt.Printf("* Fibonacci(%2v): %10v\n", z, n)
		ch <- big.NewInt(0).Add(n, <-ch) // add next value in pipeline to previous
		ch <- n
	}
	fmt.Println()
	fmt.Printf("Total time elapsed: %v\n", time.Since(start))
}
