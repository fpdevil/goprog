package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	msg = `
	Repeat Function Generator The generator would repeatedly
	call a function until it's called to stop
	********************************************************
	`

	m = 10    // generate 10 random numbers
	n = 10000 // random number range between 0 to n
)

func main() {
	fmt.Println(msg)

	done := make(chan interface{})
	defer close(done)

	// random seed generation
	t := time.Now().UnixNano()
	rand.Seed(t)

	// a function for generating the random number 0 < number < n
	r := func() interface{} {
		return rand.Intn(n)
	}

	for num := range take(done, repeatFn(done, r), m) {
		fmt.Printf("%4d\n", num)
	}
}

// repeatFn function is a stage which runs the function repeatedly
// forever until stopped
func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)

		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}


// take function is a stage which grabs the first num elements
func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)

		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- (<-valueStream):
			}
		}
	}()
	return takeStream
}
