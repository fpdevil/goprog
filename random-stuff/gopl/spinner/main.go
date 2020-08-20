package main

import (
	"fmt"
	"time"
)

const (
	msg = `
	45th Fibonacci Number
	`
	n = 45
)

func main() {
	fmt.Println(msg)

	go spinner(100 * time.Microsecond)
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// fib is a naive implementation of the Fibonacci number series
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
