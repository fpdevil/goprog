package main

import (
	"fmt"
	"log"
	"time"
)

const (
	msg = `
	45th fibonacci Number
	`
	n = 45
)

var memo = map[uint64]uint64{
	0: 0,
	1: 1,
}

func main() {
	fmt.Println(msg)
	start := time.Now()

	go spinner(100 * time.Microsecond)
	fib1N := fib1(n)
	fmt.Printf("\rfibonacci(%d) = %d\n", n, fib1N)
	log.Printf("time taken for Fib1: %v\n", time.Since(start))

	start = time.Now()
	fib2N := fib2(n)
	fmt.Printf("\rfibonacci(%d) = %d\n", n, fib2N)
	log.Printf("time taken for Fib2: %v\n", time.Since(start))

	start = time.Now()
	fib3N := fib3(n)
	fmt.Printf("\rfibonacci(%d) = %d\n", n, fib3N)
	log.Printf("time taken for Fib3: %v\n", time.Since(start))
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
func fib1(x uint64) uint64 {
	if x < 2 {
		return x
	}
	return fib1(x-1) + fib1(x-2)
}

func fib2(x uint64) uint64 {
	if _, ok := memo[x]; !ok {
		memo[x] = fib2(x-1) + fib2(x-2)
	}
	return memo[x]
}

func fib3(x uint64) uint64 {
	if x == 0 {
		return x
	}

	var (
		i        uint64
		previous uint64 = 0
		next     uint64 = 1
	)
	for i = 1; i < x; i++ {
		previous, next = next, previous+next
	}
	return next
}
