package main

import "fmt"

var (
	ch chan int
)

func main() {
	fmt.Println("/// iterating over channel values ///")
	fmt.Println()

	fillChan(5, 5)
	for i := 0; i < cap(ch); i++ {
		fmt.Println(<-ch)
	}
}

func fillChan(c, l int) {
	ch = make(chan int, c)
	for i := 0; i < l; i++ {
		ch <- i
	}
}
