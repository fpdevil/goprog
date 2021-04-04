// Fibonacci numbers computation using memoization
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fpdevil/goprog/random-stuff/generic/memoize"
)

var (
	// Fibonacci represents a variable of the type of memoized function
	Fibonacci memoize.MemoizedFunc
	// 45th Fibonacci number
	n = 45
)

// init is for initializing the Fibonacci variable with appropriate function
func init() {
	Fibonacci = memoize.Memoize(func(x int, xs ...int) interface{} {
		if x < 2 {
			return x
		}
		return Fibonacci(x-1).(int) + Fibonacci(x-2).(int)
	})
}

func main() {
	fmt.Println("* Computing Fibonacci numbers *")

	start := time.Now()

	// 1. memoized version
	go spinner(100 * time.Millisecond)
	start = time.Now()
	fibn := Fibonacci(n).(int)
	fmt.Printf("\r1. Memoized version Fibonacci(%d) = %d\n", n, fibn)
	fmt.Printf("* Time taken by Memoized function: %v\n", time.Since(start))
	fmt.Println()
	fmt.Println(">  now for some BIG Fibonacci numbers...")
	start = time.Now()
	fmt.Printf(">  a. Fibonacci(100): %v\n", Fibonacci(100).(int))
	fmt.Printf(">  b. Fibonacci(500): %v\n", Fibonacci(500).(int))
	fmt.Printf(">  c. Fibonacci(1000): %v\n", Fibonacci(1000).(int))
	fmt.Printf("* Time taken by the above Memoized version: %v\n", time.Since(start))

	// spin with a sleep timer of 100 milliseconds
	go spinner(100 * time.Millisecond)

	// pause for a second before the next invocation
	time.Sleep(1 * time.Second)

	// 2. recursive version
	fibn = fib(n)
	fmt.Printf("\r2. Recursive version Fibonacci(%d) = %d\n", n, fibn)
	fmt.Printf("* Time taken by Recursive function: %v\n", time.Since(start))
}

// fib is a naive recursive variant of Fibonacci
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

func took(start time.Time, value int, identifier string) {
	elapsed := time.Since(start)
	log.Printf("# %s | output: %v | took: %s", identifier, value, elapsed)
}
