package main

import (
	"fmt"
	"time"
)

const msg = `
	Allowing a goroutine to make some progress on work while
	waiting for another goroutine to report the result(s)

`

func main() {
	fmt.Println(msg)

	start := time.Now()
	// create a done channel to take any data type
	done := make(chan interface{})
	go func() {
		time.Sleep(6 * time.Second) // sleep for 6 seconds
		close(done)
	}()

	workCounter := 0

loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		// simulate some work
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Performed %v cycles of CPU (work) prior to exiting\n", workCounter)
	fmt.Printf("Time elapsed: %v\n", time.Since(start))
}
