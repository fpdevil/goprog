package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("-- Sieve Of Erathosthenes --")
	if len(os.Args) != 2 {
		fmt.Println("provide an upper limit for primes")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid arguments provided \n")
		return
	}

	// create the primes slice with capacity = n
	primes := make([]int, 0, n)

	// populate the primes slice with values from 2 to n
	// we will first insert 2 and there after filter out all the
	// even numbers or multiples of 2
	primes = append(primes, 2)
	for i := 3; i <= n; i+=2 {
		primes = append(primes, i)
	}

	for i, v := range primes {
		if v != 0 {
			sieve(v, primes[i:])
		}
	}

	for _, v := range primes {
		if v != 0 {
			fmt.Println(v)
		}
	}
}

func sieve(factor int, primes []int) {
	for i, v := range primes {
		if v != 0 && factor != v {
			if v%factor == 0 {
				primes[i] = 0
			}
		}
	}
}
