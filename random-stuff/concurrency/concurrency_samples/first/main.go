package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	t = time.Now().UnixNano()
	s = rand.NewSource(t)
	r = rand.New(s)
)

func main() {
	fmt.Println("/// Non-Idealistic Way: Waiting & Synchronization ///")

	n := 4 // spawning 4 goroutines
	fmt.Printf("spawning %d go routines...\n", n)
	fmt.Println()

	start := time.Now()
	for i := 0; i < n; i++ {
		go producer(i)
	}

	// wait for more that n Milliseconds to get the work finished
	time.Sleep(1 * time.Second)
	elapsed := time.Since(start)
	fmt.Printf("non-idealistic way: took %v seconds to complete\n", elapsed)
}

func producer(id int) {
	n := (r.Intn(1000) + 1)
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("* producer #%v ran for %v\n", id, d)
}
