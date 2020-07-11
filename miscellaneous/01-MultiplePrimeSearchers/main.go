package main

import (
	"errors"
	"fmt"
	"log"
	"time"
	// log "github.com/sirupsen/logrus"
)

var (
	size      = 100000 // slice of capacity 100000
	searchers = 6      // 6 concurrent searches
)

func main() {
	// fmt.Println("/// 01 Multiple Prime Searcher ///")
	// fmt.Println()

	// primes(1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	data, err := generate()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < searchers; i++ {
		d := subslice(data, searchers, i)
		go primes(i+1, d)
	}
	// sleep for 5 seconds
	time.Sleep(5 * time.Second)
	// spinner(100 * time.Millisecond)
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

// primes handles the TODO 2 part
// TODO 2 - Write a Prime Search function
// essentially, the numbers in the input slice are checked whether
// its a prime or not and printed or skipped
func primes(id int, input []int) {
	var i int
	for _, v := range input {
		for i = 2; i < v; i++ {
			if v%i == 0 {
				break
			}
		}

		if i == v {
			// it is a prime now
			fmt.Println(v)
		}
	}
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
