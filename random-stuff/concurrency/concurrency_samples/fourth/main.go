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
	fmt.Println("/// copying sync.WaitGroup value ///")
	fmt.Println()

	for i := 0; i < 4; i++ {
		producer(i, wg)
	}
	// because of the copy being passed to producer, the wg
	// value is still 0 and hence the execution never happens
	wg.Wait()

	// introduce sleep to make it execute
	time.Sleep(1 * time.Second)
}

func producer(id int, wg sync.WaitGroup) {
	wg.Add(1)
	go func() {
		n := r.Intn(1000) + 1
		d := time.Duration(n) * time.Millisecond
		time.Sleep(d)
		fmt.Printf("* producer #%v ran for %v\n", id, d)
		wg.Done()
	}()
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
