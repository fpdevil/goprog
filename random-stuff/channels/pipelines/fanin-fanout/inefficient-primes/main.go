package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
We  will generate  a  stream of  random numbers  with  upper limit  of
50000000, converting the  stream into an integer stream  and then pass
them into  a prime  finder stage.  The prime  finder stage  defined as
findPrimes is  a naive  function which attempts  to divide  the number
provided in  the inputStream by  every number below that  input stream
number. If it is  unsuccessful, it will pass the value  on to the next
stage.  The algorithm  is certainly  very crude  and very  inefficient
taking a long time.

We will close the pipeline after m = 10 primes are found
*/

const (
	msg = `
	A Naive and an inefficient approach of generating prime numbers
	using the pipeline patterns
	***************************************************************
	`

	n = 50000000 // upper limit ot cap for the prime numbers
	m = 10       // number of primes to find
)

func main() {
	fmt.Println(msg)

	// random seed generation
	// t := time.Now().UTC().UnixNano()
	// rand.Seed(t)

	// random number generating function
	rand := func() interface{} {
		return rand.Intn(n) + 1
	}

	// create the predicate channel to mark closing
	done := make(chan interface{})
	defer close(done)

	// capture start time for benchmarking
	start := time.Now()

	// generate a random stream of numbers and converted to integer
	randFuns := repeatFn(done, rand)
	randIntStream := toInt(done, randFuns)
	fmt.Println("Generating Prime Numbers:")

	// generate the primes from randomStream now
	primeStream := findPrimes(done, randIntStream)

	// take m number of primes and print them using a range
	for prime := range take(done, primeStream, m) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Total time taken: %v\n", time.Since(start))
}

// repeatFn function is a stage which will run the given function
// repeatedly forever until stopped
func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	funcStream := make(chan interface{})

	go func() {
		defer close(funcStream)

		for {
			select {
			case <-done:
				return
			case funcStream <- fn():
			}
		}
	}()
	return funcStream
}

// take function is s stage which will take only the first num items off
// of it's incoming valueStream and then exits
func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})

	go func() {
		defer close(takeStream)

		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

// toInt function is a stage for applying the type assertion or
// type conversion to the valueStream pipelime
func toInt(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)

		for v := range valueStream {
			select {
			case <-done:
				return
			case intStream <- v.(int):
			}
		}
	}()
	return intStream
}

// findPrimes function is a very rudimentary way of calculating
// the prime numbers from the intStream values as range
func findPrimes(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
	primeStream := make(chan interface{})

	go func() {
		defer close(primeStream)

		for integer := range intStream {
			prime := true

			integer--
			for divisor := integer - 1; divisor > 1; divisor-- {
				if integer%divisor == 0 {
					prime = false
					break
				}
			}

			if prime {
				select {
				case <-done:
					return
				case primeStream <- integer:
				}
			}
		}
	}()
	return primeStream
}