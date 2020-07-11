package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	// rnaomd value with seeding
	t = time.Now().UnixNano()
	s = rand.NewSource(t)
	r = rand.New(s)

	// create a WaitGroup and add the total number of goroutines that
	// are going to be spawned
	wg sync.WaitGroup
)

func main() {
	fmt.Println("/// Idealistic Way: Waiting & Synchronization ///")

	n := 4 //  number of go routines spawned
	fmt.Printf("spawning %d go routines...\n", n)
	fmt.Println()

	start := time.Now()

	// execute n number of concurrent transactions
	wg.Add(n)

	for i := 0; i < n; i++ {
		go producer(i, &wg)
	}

	// ---
	// now we will wait for groutines to complete work
	// This will block the execution of the main function until the
	// waitGroup's count is back down to zero.
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("non-idealistic way: took %v seconds to complete\n", elapsed)
}

func producer(id int, wg *sync.WaitGroup) {
	// call wg.Done() method using defer as it's the easiest way thats guaranteed
	// to be called during every exit
	defer wg.Done()
	n := (r.Intn(1000) + 1)
	// cast n to time.Duration and multiply with Millisecond
	d := time.Duration(n) * time.Millisecond
	// fmt.Printf("sleeping for %d seconds\n", d)
	time.Sleep(d)
	fmt.Printf("* producer #%v ran for %v\n", id, d)
}
