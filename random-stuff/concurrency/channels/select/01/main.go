package main

import (
	"fmt"
	"time"
)

// case -1
// multiple channels being ready simultaneously have
// something to read at the same time

func main() {
	start := time.Now()

	var (
		total = 1000
		c1Count, c2Count int
	)

	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)


	for i := total; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	elapsed := time.Since(start)

	// get the percentage count of messages to channel in select
	f1 := (float64(c1Count) / float64(total)) * 100
	f2 := (float64(c2Count) / float64(total)) * 100

	fmt.Printf("* Probability of c1: %.4v%%\n", f1)
	fmt.Printf("* Probability of c2: %.4v%%\n", f2)

	fmt.Println()
	fmt.Printf("Total time taken: %v\n", elapsed)
}
