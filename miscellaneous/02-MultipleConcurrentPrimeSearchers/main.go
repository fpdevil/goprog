package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
	// log "github.com/sirupsen/logrus"
)

const (
	size      = 100000 // slice of capacity
	searchers = 6
)

var (
	wg sync.WaitGroup
)

func main() {
	data, err := generate()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < searchers; i++ {
		d := subslice(data, searchers, i)
		wg.Add(1)
		go primes(i+1, d)
	}

	wg.Wait()
}

// generate data handles the part 1 of producing slice of size 100000
// TODO 1 - Write function which returns a slice of 100,000 integers.
func generate() (output []int, err error) {
	output = make([]int, size)
	if output == nil {
		err = errors.New("unable to allocate enough memory for storage")
		return
	}

	for i := 0; i < size; i++ {
		output[i] = i + 1
	}
	return
}

// primes function handles the TODO 2 part
// TODO 2 - Write a Prime Search function
// essentially, the numbers in the input slice are checked
// whether its a prime or not and printed or skipped
func primes(id int, input []int) {
	var i int
	for _, v := range input {
		for i = 2; i < v; i++ {
			if v%i == 0 {
				break
			}
		}

		if i == v {
			fmt.Println(v)
		}
	}
	// use the workgroup to handle concurrency
	wg.Done()
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
