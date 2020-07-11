package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("/// No Channels are ready, but need to perform work ///")
	fmt.Println()

	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("Matched `default` after %v\n\n", time.Since(start))
	}
}
