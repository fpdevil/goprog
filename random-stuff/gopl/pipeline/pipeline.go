// Demo of pipelines using 3 goroutines and 2 channels
// 3 goroutines are
// - counter
// - squarer
// - printer
// 2 channels are
//		-- naturals and
//		-- squares
// -----------------------------------------------------
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

// counter function takes a write only channel out and sends the
// values from 0 to 100
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

// squares function takes a write only channel out and a read only
// channel in and passes values from in to out
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

// printer function takes a read only channel in and prints the
// values form the channel
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
