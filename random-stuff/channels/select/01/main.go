package main

import "fmt"

// case -1
// multiple channels have something to read at the same time

func main() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var total = 1000
	var c1Count, c2Count int

	for i := total; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	f1 := (float64(c1Count) / float64(total)) * 100
	f2 := (float64(c2Count) / float64(total)) * 100
	fmt.Printf("Probability of c1: %.4v%%\n", f1)
	fmt.Printf("Probability of c2: %.4v%%\n", f2)
}
