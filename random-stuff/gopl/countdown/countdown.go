package main

import (
	"fmt"
	"os"
	"time"
)

// countdown for a rocket launcher
func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Starting countdown...Press return to abort launch")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("Launch aborted.")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift Off!")
}
