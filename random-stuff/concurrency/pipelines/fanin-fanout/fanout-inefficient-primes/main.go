package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/*
FanOut Version of the Primes

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
***************************************************************
 Applying  the  Fan-Out  pattern to  the  inefficient  primes
 version, speeding up the  prime numbers generation using the
 pipeline patterns.
 generating %d primes with upper limit of %v
***************************************************************
`

	n = 50000000 // upper limit ot cap for the prime numbers
	m = 10       // number of primes to find
)

func main() {
	defer trace("main")()
	fmt.Printf(msg, m, n)

	// random seed generation
	// t := time.Now().UTC().UnixNano()
	// rand.Seed(t)

	// random number generating function
	rand := func() interface{} {
		return rand.Intn(n)
	}

	// create the predicate channel to mark closing
	done := make(chan interface{})
	defer close(done)

	// generate a random stream of numbers and converted to integer
	randFuns := repeatFn(done, rand)
	randIntStream := toInt(done, randFuns)

	// Fan-Out
	// generate the primes from randomStream now
	// primeStream := findPrimes(done, randIntStream)
	// we will FanOut the stage in pipeline by starting or spawning
	// multiple versions of the primeStream as below
	numFinders := runtime.NumCPU()
	fmt.Printf("spawning %d goroutines for primes\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)
	fmt.Println("Generating Prime Numbers:")
	for i := 0; i < numFinders; i++ {
		finders[i] = findPrimes(done, randIntStream)
	}

	// take m number of primes and print them using a range
	for prime := range take(done, fanIn(done, finders...), m) {
		fmt.Printf("\t%d\n", prime)
	}
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

// fanIn function implements the fan-in pattern
// the fanin pattern is a method of multiplexing or joining
// together multiple streams of data into a single stream
func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})

	// multiplex function when passed with a channel will read
	// from the channel and pass the value read from it onto the
	// multiplexedStream channel
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	// select from all the channels
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// wait for all channels to complete
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("* [enter] %s *", msg)
	return func() {
		log.Printf("* [exit] time spent in %s: %s *", msg, time.Since(start))
	}
}
