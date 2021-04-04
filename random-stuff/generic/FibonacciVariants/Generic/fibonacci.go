package main

import "fmt"

var (
	// memo will be used as a cache with initial values of 0, 1
	memo = map[int64]int64{0: 0, 1: 1}
)

// FibResult struct stores the values of previous and next values of Fibonacci
type FibResult struct {
	previous int64
	next     int64
}

// fib1 is calculated using a local cache consisting of a map from int64
// to int64 stored in the memo variable
func fib1(n int64) int64 {
	val, ok := memo[n]
	if !ok {
		memo[n] = fib1(n-1) + fib1(n-2)
		return memo[n]
	}
	return val
}

// fib2 function calculates the Fibonacci numbers by swapping the values and
// assigning previous values there by preventing creation of temporary values
func fib2(n int64) int64 {
	var i, prev, next int64
	if n == 0 {
		return 1
	}
	prev, next = 0, 1
	for i = 1; i < n; i++ {
		prev, next = next, prev+next
	}
	return next
}

/**
 * Function: fib3 uses channels for calculating the Fibonacci series
 * that essentially exploits the go's concurrency model for faster computation
 *
 * @param ch: input channel with integers
 * @param n: input integer value
 */
func fib3(ch chan<- int64, n int64) {
	var (
		x int64 = 0
		y int64 = 1
	)

	for n >= 0 {
		ch <- x
		n--
		x, y = y, x+y
	}

	close(ch)
}

/**
 * [FibResult] Method: fib4 uses the variable swapping technique with
 * the results formulates as a struct
 *
 * @return: [int64] returns a sequence of Fibonacci values
 */
func (fs *FibResult) fib4() int64 {
	result := fs.previous
	fs.previous, fs.next = fs.next, fs.previous+fs.next
	return result
}

func main() {
	N := 90
	ch := make(chan int64)

	fmt.Println("Fibonacci using caching...")
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci1(%2v): %25v\n", i*10, fib1(int64(i*10)))
	}

	fmt.Println()
	fmt.Println("Fibonacci using variable swapping...")
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci2(%2v): %25v\n", i*10, fib2(int64(i*10)))
	}

	fmt.Println()
	fmt.Println("Fibonacci sequence using Channels...")
	go fib3(ch, int64(N))
	for i := 0; i <= N; i++ {
		if i%10 == 0 {
			fmt.Printf("Fibonacci3(%2v): %25v\n", i, <-ch)
		} else {
			<-ch
		}
	}

	fmt.Println()
	fmt.Println("Fibonacci sequence using struct type...")
	fs := &FibResult{0, 1}
	for i := 0; i <= N; i++ {
		if i%10 == 0 {
			fmt.Printf("Fibonacci4(%2v): %25v\n", i, fs.fib4())
		} else {
			fs.fib4()
		}
	}
}
