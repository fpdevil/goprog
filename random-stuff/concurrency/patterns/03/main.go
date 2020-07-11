package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	t = time.Now().UnixNano() // unix timestamp for random seed

	wg sync.WaitGroup
)

func main() {
	fmt.Println("/// SINK: Golang Concurrency Patterns (with WaitGroup) ///")
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
	wg.Wait()
}

// counter function counts the numbers receieved on a channel
func counter(in chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Info("Counter - starting work...")
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
	wg.Add(1)
	out = make(chan int)
	go func() {
		defer close(out)
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			case out <- rand.Intn(200):
			}
		}
	}()
	return out
}
