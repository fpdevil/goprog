package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers = 45
)

var (
	t            = time.Now().UnixNano()
	s            = rand.NewSource(t)
	r            = rand.New(s)
	wg           sync.WaitGroup
	counter      = 0
	counterMutex sync.Mutex
	// keep track of each count in a map
	m map[int]int
)

func main() {
	fmt.Println("/// concurrent counter update go routines ///")
	fmt.Println()

	fmt.Println("--- running with mutexes tracking with a map ---")
	fmt.Printf("starting %v goroutines to increment 'counter'\n", numWorkers)

	start := time.Now()
	// intialize the counter map
	m = make(map[int]int)

	for i := 0; i < numWorkers; i++ {
		updateCounter(i)
	}
	wg.Wait()

	// this should print the same as numWorkers
	fmt.Printf("counter = %v\n", counter)
	fmt.Printf("map m: len = %v, values = %#v\n", len(m), m)

	elapsed := time.Since(start)
	fmt.Printf("completed in %v\n", elapsed)
}

func updateCounter(id int) {
	// this function is called concurrently to update 'counter'
	wg.Add(1)
	go func() {
		// prevent concurrent update or sync critical code section
		// wrap the counter with Lock and Unlock
		counterMutex.Lock()
		counter++
		// track each count with a map
		m[id] = m[id] + 1
		counterMutex.Unlock()
		wg.Done()
	}()
}
