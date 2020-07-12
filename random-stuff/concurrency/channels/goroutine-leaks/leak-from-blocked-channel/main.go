package main

import (
	"fmt"
	"strings"
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

	// sc := insert("I am feeling lucky")
	// terminated := perform(done, sc)
	terminated := perform(done, nil)

	// a goroutine spawned to cancel the one in perform
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling the perform() goroutine...")
		close(done)
	}()

	// this is the place where the foroutine spawned from
	// the perform is joined with the main goroutine
	fmt.Printf("%v\n", <-terminated)
	fmt.Println("Done with the Work!")
}

// perform function processes a string data as a read only channel
// the done channel passed as the first parameter acts as acancellation
// signal between the parent child channels
func perform(done <-chan interface{}, strings <-chan string) <-chan interface{} {
	terminated := make(chan interface{})
	go func() {
		defer fmt.Println("func perform() exited!")
		defer close(terminated)

		for {
			select {
			case <-done:
				return
			case s := <-strings:
				fmt.Printf("%#v\n", s)
			}
		}
	}()
	return terminated
}

func insert(in string) chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, v := range strings.Fields(in) {
			out <- v
		}
	}()
	return out
}
