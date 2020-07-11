package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("=== primes generation ===")
	fmt.Println()

	var limit *int = flag.Int("limit", 2, "upper limit till required where primes are needed")
	flag.Parse()

	primes := sieve(*limit)
	for {
		fmt.Println(<- primes)
	}
}

func generator() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ;i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in <-chan int, prime int) <-chan int {
	out := make(chan int)
	go func() {
		for {
			i := <- in
			if i != prime && i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve(limit int) chan int {
	out := make(chan int)
	go func() {
		ch := generator()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}
