package main

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	t = time.Now().UnixNano() // unix timestamp for random seed
)

func main() {
	fmt.Println("/// SINK: Golang Concurrency Patterns ///")
	fmt.Println()

	// seed for random number generator
	rand.Seed(t)

	// sink - count numbers
	done := make(chan bool)
	d := numGen(done)
	counter(d)
	time.Sleep(1 * time.Second)

	// send a signal to go routine to exit gracefully
	done <- true
}

// counter function counts the numbers receieved on a channel
func counter(in chan int) {
	go func() {
		log.Info("Counter - starting to work...")
		start := time.Now()
		var count int
		for range in {
			count++
		}
		fmt.Printf("Counter processed %v items in %v\n", count, time.Since(start))
	}()
}

// numGen function returns a channel on which random numbers are produced
func numGen(done chan bool) (out chan int) {
	out = make(chan int)
	go func() {
		for {
			select {
			case <-done:
				close(out)
				return
			case out <- rand.Intn(200) + 1:
			}
		}
	}()
	return out
}
