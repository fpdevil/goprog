package main

import (
	"fmt"
	"time"
)

// case - 2
// There are never any channels which become ready
//
// channels are said to be ready if they are populated or closed channels
// in case of reads; and channels which are not at capacity in the case of
// writes

func main() {
	fmt.Println("/// no channels are ready ///")
	fmt.Println()

	var c <-chan int
	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Printf("Timed out\n")
	}
}
