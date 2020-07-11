package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	t = time.Now().UnixNano()
	s = rand.NewSource(t)
	r = rand.New(s)
)

func main() {
	fmt.Println("/// non-concurrent data producer and consumer ///")
	fmt.Println()

	d := make(chan string, 5)
	producer(1, d)
	consumer(d)
}

func producer(id int, out chan string) {
	num := r.Intn(cap(out))
	for i := 0; i < num; i++ {
		out <- fmt.Sprintf("Producer #%v => item: %v", id, i+1)
	}
	close(out)
}

func consumer(in chan string) {
	count := 0
	for v := range in {
		count++
		fmt.Printf("Consumer received %v\n", v)
	}

	if count == 0 {
		fmt.Printf("No data received")
		return
	}

	fmt.Printf("Processed %v items\n", count)
}
