package main

/*
How many go routines can be reasonably spawned for an application
- Need to handle concurency and data corruption
  For those needs we have the mutexes from sync
*/

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	// number of workers/go routines that can be spawned
	numWorkers = 10000
)

var (
	t  = time.Now().UnixNano()
	s  = rand.NewSource(t)
	r  = rand.New(s)
	wg sync.WaitGroup
)

func main() {
	fmt.Println("/// 10k go routines ///")
	fmt.Printf("creating %v goroutines\n", numWorkers)
	fmt.Println()
	start := time.Now()

	for i := 0; i < numWorkers; i++ {
		producer(i)
	}
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("completed in %v\n", elapsed)
}

func producer(id int) {
	wg.Add(1)
	go func() {
		fmt.Printf("producer #%v - finished\n", id)
		n := r.Intn(5000)
		d := time.Duration(n) * time.Nanosecond
		time.Sleep(d)
		wg.Done()
	}()
}
