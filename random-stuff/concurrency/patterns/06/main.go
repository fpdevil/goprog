package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	msg = `
	FanOut Concurrency Pattern
	Handling of multiple processes in the pipeline
	stage to handle cpu intensive or computationally
	intensive tasks.

	`
)

var (
	t  = time.Now().UTC().UnixNano()
	wg sync.WaitGroup
)

func main() {
	fmt.Println(msg)

	rand.Seed(t)
	done := make(chan bool)
	nums := numGen(done)
	// counter(nums)

	c1, c2, c3 := fanOut(nums)
	counter(c1)
	counter(c2)
	counter(c3)

	time.Sleep(250 * time.Millisecond)
	done <- true
	wg.Wait()
}

// fanOut function splits the stream of input ints into
// multiple (3) separate streams
func fanOut(in <-chan int) (out1, out2, out3 chan int) {
	wg.Add(1)
	out1 = make(chan int)
	out2 = make(chan int)
	out3 = make(chan int)

	go func() {
		defer wg.Done()
		for v := range in {
			select {
			case out1 <- v:
			case out2 <- v:
			case out3 <- v:
			}
		}
		defer close(out1)
		defer close(out2)
		defer close(out3)
	}()
	return
}

// counter function is a stage just to count the total number
// of messages goung through the channel at that moment
func counter(in <-chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Info("Counter - initiating the work...")
		start := time.Now()
		var count int
		for range in {
			count++
		}
		fmt.Printf("Counter processed %v items in %v\n", count, time.Since(start))
	}()
}

// numGen function takes a boolean read only channel and
// returns a stream of random numbers in the output channel
func numGen(done <-chan bool) (out chan int) {
	out = make(chan int)
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(out)
		for {
			select {
			case <-done:
				return
			case out <- rand.Intn(200) + 1:
			}
		}
	}()
	return out
}
