package main

import (
	"fmt"
	"os"
	"time"
)

// countdown for a rocket launcher
func main() {
	//!+abort
	abort := make(chan struct{}, 1)
	// a go routine to read a single byte from standard input
	// and if succeeds, sends a value on a channel abort
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	//!-abort

	// fmt.Println("Starting countdown...Press return to abort launch")
	// select {
	// case <-time.After(10 * time.Second):
	// 	// do nothing here
	// case <-abort:
	// 	fmt.Println("Launch aborted!")
	// 	return
	// }

	// tick := time.Tick(1 * time.Second)
	// for countdown := 10; countdown > 0; countdown-- {
	// 	fmt.Println(countdown)
	// 	<-tick
	// }

	fmt.Println("Starting countdown...Press return to abort launch")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// do nothing
		case <-abort:
			fmt.Println("Launch aborted")
			return
		}
	}

	launch()
}

// lanunch function is invoked for actual rocket launch
func launch() {
	fmt.Println("Lift Off!")
}
