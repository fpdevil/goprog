package main

import (
	"fmt"
	"time"
)

// A Memoized version of the Fibonacci series

// memoizedFunc is a function type which takes atleast one int
// and returns an empty interface{}
type memoizedFunc func(int, ...int) interface{}

// Fibonacci variable stores a function of the type of memoizedFunc
var Fibonacci memoizedFunc

// 45th Fibonacci number
var n = 45

func init() {
	// Memoize will memorizes any function that takes atleast one int
	// and which returns an interface.
	Fibonacci = Memoize(func(x int, xs ...int) interface{} {
		if x < 2 {
			return x
		}
		// using unchecked type assertion to convert the returned
		// interfaces to their underlying ints
		return Fibonacci(x-1).(int) + Fibonacci(x-2).(int)
	})
}

func main() {
	fmt.Println("Computing Fibonacci numbers...")
	start := time.Now()
	// sping with a sleep timer of 100 milliseconds
	go spinner(100 * time.Millisecond)
	fibn := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibn)
	fmt.Printf("* time taken for Recursive function: %v\n", time.Since(start))

	fmt.Println()
	start = time.Now()
	fmt.Printf("Fibonacci(%d) = %d\n", n, Fibonacci(n).(int))
	fmt.Printf("* time taken for Memoized function: %v\n", time.Since(start))
}

// Memoize function takes memoizedFunc as an argument and returns a function
// with the same signature
func Memoize(function memoizedFunc) memoizedFunc {
	cache := make(map[string]interface{})
	return func(x int, xs ...int) interface{} {
		key := fmt.Sprint(x)
		for _, i := range xs {
			key += fmt.Sprintf(",%d", i)
		}

		if value, found := cache[key]; found {
			return value
		}

		value := function(x, xs...)
		cache[key] = value
		return value
	}
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
