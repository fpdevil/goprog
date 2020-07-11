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
	fmt.Println("/// Processor: Concurrency Patterns Counting & Runnning a Filter ///")
	fmt.Println()

	// seed for random number generator
	rand.Seed(t)
	// sink - count numbers
	done := make(chan bool)
	d := numGen(done)

	// pass through proxy to get number count
	d = proxyCounter(d)

	// filter for even numbers
	d = filter(d, func(x int) bool {
		if x%2 == 0 {
			return true
		}
		return false
	})

	counter(d)
	time.Sleep(1 * time.Second)

	// send a signal to go routine to exit gracefully
	done <- true
	wg.Wait()
}

// proxyCounter function counts the numbers received on a channel
// but it just passes or proxied them untouched
func proxyCounter(in chan int) (out chan int) {
	wg.Add(1)
	out = make(chan int)

	go func() {
		defer wg.Done()
		defer close(out)
		log.Info("Counter {Proxy} - starting work...")
		start := time.Now()
		var count int
		for v := range in {
			count++
			out <- v
		}
		fmt.Printf("Counter {Proxy} processed %v items in %v\n", count, time.Since(start))
	}()
	return out
}

// counter function counts the numbers receieved on a channel
func counter(in chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Info("Counter - starting wwork...")
		start := time.Now()
		var count int
		for range in {
			count++
		}
		fmt.Printf("Counter processed %v items in %v\n", count, time.Since(start))
	}()
}

// filter function keeps the numbers which match a specific predicate
// defined by comp; if predicate  = 'nil' no filtering
func filter(in <-chan int, predicate func(int) bool) (out chan int) {
	wg.Add(1)
	out = make(chan int)

	if predicate == nil {
		predicate = func(int) bool {
			return true
		}
	}

	go func() {
		defer wg.Done()
		defer close(out)
		for v := range in {
			if predicate(v) {
				out <- v
			}
		}
	}()
	return
}

// numGen function returns a channel on which random numbers are produced
func numGen(done <-chan bool) (out chan int) {
	wg.Add(1)
	out = make(chan int)
	go func() {
		defer wg.Done()
		defer close(out)
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
