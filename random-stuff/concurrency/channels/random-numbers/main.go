package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	n = 100 //random numbers from 0 to n
	m = 10  // generate m number of random numbers
)

func main() {
	fmt.Println("/// Random Number Stream ///")
	fmt.Println()

	// create a random seed
	t := time.Now().UTC().UnixNano()
	rand.Seed(t)

	done := make(chan interface{})
	randStream := newRandIntStream(done)
	fmt.Printf("generating %v random numbers each between 0 to %v\n", m, n)
	for i := 1; i <= m; i++ {
		fmt.Printf("%2d: %2d\n", i, <-randStream)
	}
	close(done)

	// now simulate doing some work
	time.Sleep(500 * time.Millisecond)
}

//!+ newRandIntStream

// The function newRandIntStream takes a receiver or read only channel
// that acts like a stopping predicate and returns another read only
// channel with random number stream from which random numbers must
// be consumed by another party. In essence, this is a Producer
func newRandIntStream(done <-chan interface{}) <-chan int {
	randStream := make(chan int)
	go func() {
		defer fmt.Println("exiting the closure newRandIntStream...")
		defer close(randStream)

		for {
			select {
			case <-done:
				// producer closure is indicated by done
				return
			case randStream <- rand.Intn(n) + 1:
			}
		}
	}()
	return randStream
}

//!-
