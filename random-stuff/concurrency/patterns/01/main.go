package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	t = time.Now().UnixNano() // unix timestamp for random seed
)

func main() {
	fmt.Println("/// Golang Concurrency Patterns ///")
	fmt.Println()

	// seed for random number generator
	rand.Seed(t)
	done := make(chan bool)
	d := numGen(done)
	for i := 0; i < 10; i++ {
		fmt.Printf("* %3v\n", <-d)
	}

	// send a signal to go routine to exit gracefully
	done <- true
}

// numGen function returns a channel on which random numbers are produced
func numGen(done chan bool) (out chan int) {
	out = make(chan int)

	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case out <- rand.Intn(200) + 1:
			}
		}
	}()
	return out
}
