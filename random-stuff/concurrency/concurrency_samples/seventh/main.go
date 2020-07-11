package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers = 10000
)

var (
	t            = time.Now().UnixNano()
	s            = rand.NewSource(t)
	r            = rand.New(s)
	wg           sync.WaitGroup
	counter      = 0
	counterMutex sync.Mutex
	m            map[int]int
)

func main() {
	fmt.Println("/// concurrent counter update go routines ///")
	fmt.Println()

	fmt.Println("--- running with mutexes ---")
	fmt.Printf("starting %v goroutines to increment 'counter'\n", numWorkers)
	start := time.Now()

	for i := 0; i < numWorkers; i++ {
		updateCounter()
	}
	wg.Wait()
	// this should be same as numWorkers
	fmt.Printf("counter = %v\n", counter)
	elapsed := time.Since(start)
	fmt.Printf("completed in %v\n", elapsed)
}

func updateCounter() {
	// this function is called concurrently to update 'counter'
	wg.Add(1)
	go func() {
		// prevent concurrent update or sync critical code section
		// wrap the counter with Lock and Unlock
		counterMutex.Lock()
		counter++
		counterMutex.Unlock()
		wg.Done()
	}()
}
