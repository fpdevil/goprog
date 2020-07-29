package main

import (
	"fmt"
	"time"
)

const (
	n = 45 // 45 th fibonacci number
)

func main() {
	fmt.Println("== naive fibonacci of 45th number ==")
	start := time.Now()
	// sping with a sleep timer of 100 milliseconds
	go spinner(100 * time.Millisecond)
	fibn := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibn)
	fmt.Printf("total time taken: %v\n", time.Since(start))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
