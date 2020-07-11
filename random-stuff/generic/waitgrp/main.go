package main

import (
	"fmt"
	"sync"
)

var (
	count      int
	lock       sync.Mutex
	arithmetic sync.WaitGroup
)

func main() {
	fmt.Println("/// mutexes & waitgroups ///")
	fmt.Println()

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic Completed...")
}

func increment() {
	lock.Lock()
	defer lock.Unlock()
	count++
	fmt.Printf("Incrementing: %d\n", count)
}

func decrement() {
	lock.Lock()
	defer lock.Unlock()
	count--
	fmt.Printf("Decrementing: %d\n", count)
}
