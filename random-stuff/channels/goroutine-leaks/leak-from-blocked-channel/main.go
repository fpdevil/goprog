package main

import (
	"fmt"
	"time"
)

const (
	msg = `
	Demonstration of handling the memory leak in a goroutine
	from a blocked channel...
	********************************************************

`
)

func main() {
	fmt.Println(msg)

	done := make(chan interface{})
	// out goroutine should exit properly even when passed a
	// nil as an input value
	terminated := doWork(done, nil)

	// a goroutine spawned to cancel the one in doWork
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling the doWork() goroutine...")
		close(done)
	}()

	// this is the place where the foroutine spawned from
	// the doWork is joined with the main goroutine
	<-terminated
	fmt.Println("Done with the Work!")
}

// doWork function processes a string data as a read only channel
// the done channel passed as the first parameter acts as acancellation
// signal between the parent child channels
func doWork(done <-chan interface{}, strings <-chan string) <-chan interface{} {
	terminated := make(chan interface{})
	go func() {
		defer fmt.Println("func doWork() exited!")
		defer close(terminated)

		for {
			select {
			case <-done:
				return
			case s := <-strings:
				fmt.Printf("%s\n", s)
			}
		}
	}()
	return terminated
}
