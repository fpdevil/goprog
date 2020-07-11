package main

import (
	"flag"
	"fmt"
)

// If the number  is not a multiple and the  next filter is not
// created,  then it is a  prime.  So  we print  it and  create
// another filter.

func main() {
	fmt.Println("=== Sieve Of Erathosthenes ===")
	fmt.Println()

	var limit *int = flag.Int("limit", 2, "upper limit till required where primes are needed")
	flag.Parse()
	sieve(*limit)
}

func generate(limit int, upstream chan<- int) {
	// first send 2 to the send only channel
	upstream <- 2
	// next filter out all the multiples of 2 (even numbers)
	// byt stepping 2 numbers at a time as we do not need the
	// evens anymore
	for i := 3; i <= limit; i += 2 {
		upstream <- i
	}
	close(upstream)
}

// The filter  function has  three arguments:  an input  channel, an
// output channel,  and a  prime number. It  copies values  from the
// input to the output, discarding anything divisible by the prime.
func filter(upstream <-chan int, downstream chan<- int, prime int) {
	// fmt.Printf("calling filter with %d... upstream %#v...\n", prime, upstream)
	for i := range upstream {
		if i != prime && i%prime != 0 {
			downstream <- i
		}
	}
	close(downstream)
}

func sieve(limit int) {
	primes := make([]int, 0)
	p1 := make(chan int)
	go generate(limit, p1)
	for p := 0; p <= limit; p++ {
		prime := <-p1
		if prime != 0 {
			primes = append(primes, prime)
		}
		p2 := make(chan int)
		go filter(p1, p2, prime)
		p1 = p2
	}
	fmt.Printf("primes generated: %v\n", primes)
}
