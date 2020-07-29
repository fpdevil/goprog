package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("* demo of goruntime scheduling *")

	runtime.GOMAXPROCS(1)
	start := time.Now()

	wg.Add(3)

	n := 1000 // get primes between 1 and 1000
	fmt.Printf("spawning go routines for %d primes...", n)
	go printPrimes("A", n)
	go printPrimes("B", n)
	go printPrimes("C", n)

	fmt.Println("Waiting to finish...")
	wg.Wait()
	fmt.Printf("total time elapsed: %v\n", time.Since(start))
}

func printPrimes(id string, limit int) {
	defer wg.Done()
loop:
	for i := 2; i < limit; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue loop
			}
		}
		fmt.Printf("%s: %d\n", id, i)
	}
	fmt.Printf("Completed %v of %d primes\n", id, limit)
}
