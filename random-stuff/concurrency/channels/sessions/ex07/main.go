package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	t = time.Now().UnixNano()
	s = rand.NewSource(t)
	r = rand.New(s)

	wgProducers sync.WaitGroup
	wgConsumers sync.WaitGroup
)

func main() {
	fmt.Println("/// *multiple* concurrent data producers and *single* consumer ///")
	fmt.Println()

	// an unbuffered channel
	d := make(chan string)
	producer(1, d)
	producer(2, d)
	producer(3, d)
	go consumer(d)

	// wait for producers
	wgProducers.Wait()

	// since we have multiple producers, we will close the
	// channel only after the producers are finished
	close(d)

	// wait for consumers
	wgConsumers.Wait()

}

func producer(id int, out chan string) {
	wgProducers.Add(1)

	// now launch goroutine to produce data
	go func() {
		i := 1
		end := time.Now().Add(1000 * time.Millisecond)

		for time.Now().Before(end) {
			out <- fmt.Sprintf("Producer #%v - item: %v", id, i)
			i++
		}

		//-- we can't call close(out) anymore
		wgProducers.Done()
	}()
}

func consumer(in chan string) {
	wgConsumers.Add(1)
	count := 0
	for v := range in {
		count++
		fmt.Printf("Consumer received: %v\n", v)
	}

	if count == 0 {
		fmt.Printf("No data received...")
	}

	fmt.Printf("Processed %v items\n", count)
	wgConsumers.Done()
}
