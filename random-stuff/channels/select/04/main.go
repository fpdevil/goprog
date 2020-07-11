package main

import (
	"fmt"
	"time"
)

const msg = `
	Allowing a goroutine to make progress on work while waiting
	for another goroutine to report the result
`

func main() {
	fmt.Println(msg)
	fmt.Println()

	done := make(chan interface{})
	go func() {
		time.Sleep(6 * time.Second)
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

	fmt.Printf("Achieved %v cycles of work before signalling to stop\n", workCounter)
}
