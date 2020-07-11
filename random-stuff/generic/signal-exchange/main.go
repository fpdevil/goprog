package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("/// Using sync Cond for event exchange between goroutines ///")
	fmt.Println()

	// Initialize a new Cond.
	// sync.NewCond takes its argument a type which satisfies the sync.Locker
	// interface which allows Cond type to facilitate coordination with other
	// goroutines in a concurrent safe way.
	c := sync.NewCond(&sync.Mutex{})
	// create a zero length slice of capcity 10
	queue := make([]interface{}, 0, 10)

	dequeue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		// point the head of slice to second item; this way we are sort of
		// doing qequeuing operation
		queue = queue[1:]
		fmt.Println("Removed from the Q")
		c.L.Unlock()
		// keep goroutine waiting on the condition that something has occurred
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to Q")
		queue = append(queue, struct{}{})
		go dequeue(1 * time.Second)
		c.L.Unlock()
	}
}
