package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	// CLOSEA to check whether the first channel can be closed
	CLOSEA = false
	// DATA for storing the data
	DATA = make(map[int]bool)
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, " ")
		_, ok := DATA[x]
		if ok {
			CLOSEA = true
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	var sum int
	for x := range in {
		sum += x
	}
	fmt.Printf("sum of random numbers is %d\n", sum)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("input two integer parameters.")
		return
	}

	x1, _ := strconv.Atoi(os.Args[1])
	x2, _ := strconv.Atoi(os.Args[2])

	if x1 > x2 {
		fmt.Printf("first param (%d) should be smaller than the second (%d)\n", x1, x2)
		return
	}

	t := time.Now().UTC().UnixNano()
	rand.Seed(t)
	A := make(chan int)
	B := make(chan int)

	go first(x1, x2, A)
	go second(B, A)
	third(B)
}
