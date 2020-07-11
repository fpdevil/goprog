package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	t = time.Now().UnixNano()
	s = rand.NewSource(t)
	r = rand.New(s)

	wg sync.WaitGroup
)

func main() {
	fmt.Println("/// sync.WaitGroup pitfalls - incorrect initialization ///")
	fmt.Println()

	// sync.WaitGroup pitfalls - incorrect initialization
	// ---
	// case 1
	// here we add Wait without the Add so the work is never completed
	// as the go routine never got enough time to complete
	// go producer1(1)
	// wg.Wait()

	// ---
	// case 2
	// here a mismatch in the number of go routines in Add() and the actual
	// number of go routines spawned will cause a dedlock
	// wg.Add(2)
	// go producer1(1)
	// wg.Wait()

	// ---
	// running correctly with 4
	n := 4

	wg.Add(n)
	for i := 0; i < n; i++ {
		go producer1(i)
	}
	wg.Wait()

	// adding another here after Wait will cause deadlock
	// also just running the function directly also will fail as it
	// has a Done defined in it which tries to reduce the number of
	// workgroup count
	// producer1(1)

	fmt.Println()
	fmt.Println("/// encapsulating the waitgroup in closure ///")
	fmt.Println()
	for i := 0; i < n; i++ {
		producer2(i)
	}
	wg.Wait()

	fmt.Println()
	fmt.Println("/// launching go routines with anonymous functions ///")
	fmt.Println()
	launchWorkers(4)
	wg.Wait()
}

// lanchWorkers creates 'n' goroutines using ananymous function
func launchWorkers(c int) {
	for i := 0; i < c; i++ {
		wg.Add(1)
		id := i
		go func() {
			n := r.Intn(1000) + 1
			d := time.Duration(n) * time.Millisecond
			time.Sleep(d)
			fmt.Printf("* worker #%v ran for %v\n", id, d)
			wg.Done()
		}()
	}
}

func producer2(id int) {
	wg.Add(1)
	go func() {
		n := r.Intn(1000) + 1
		d := time.Duration(n) * time.Millisecond
		time.Sleep(d)
		fmt.Printf("* producer2 #%v ran for %v\n", id, d)
		wg.Done()
	}()
}

func producer1(id int) {
	n := r.Intn(1000) + 1
	// n := (r.Int() % 1000) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("* producer1 #%v ran for %v\n", id, d)
	wg.Done()
}
