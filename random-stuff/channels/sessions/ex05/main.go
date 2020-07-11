package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	chanCap = 10
)

var (
	t = time.Now().UnixNano()
	s = rand.NewSource(t)
	r = rand.New(s)
)

func main() {
	fmt.Println("/// passing channels to functions ///")
	fmt.Println()

	// passing channels to functions
	d := make(chan int, chanCap)
	if d == nil {
		fmt.Println("cannot allocate or initialize channels")
	}
	producer(d)
	consumer(d)

	fmt.Println()
	d = generator()
	consumer(d)

	fmt.Println()
	d = generator()
	d = counter(d)
	d = adder(d, 5)
	consumer(d)
}

// producer sends between 1 to cap(nums) random integers into the channel 'nums'
func producer(nums chan<- int) {
	n := r.Intn(cap(nums)) + 1
	for i := 0; i < n; i++ {
		nums <- r.Intn(200) // random number under 200
	}
	close(nums)
}

func consumer(nums <-chan int) {
	for v := range nums {
		fmt.Printf("consumer received: %#v\n", v)
	}
}

// generator creates a chan on which caller receives random numbers
func generator() (out chan int) {
	fmt.Println("--- generator of random integers ---")
	out = make(chan int, chanCap)
	for i := 0; i < cap(out); i++ {
		out <- r.Intn(200) // random number under 200
	}
	close(out)
	return
}

// counter takes a channel 'in' and copies the elemnt to
// 'out' channel, it writes the number of elements copied
func counter(in chan int) (out chan int) {
	fmt.Println("--- triggerring counter ---")
	var count int
	out = make(chan int, len(in))

	for v := range in {
		out <- v
		count++
	}
	close(out)
	fmt.Printf("counted %v elements\n", count)
	return
}

// adder takes a channel 'in' and adds the value 'c' to each
// element receieved from the channel, then write the result
// to 'out' channel
func adder(in chan int, c int) (out chan int) {
	fmt.Println("--- triggerring adder ---")
	out = make(chan int, len(in))
	for v := range in {
		out <- v + c
	}
	close(out)
	return
}
