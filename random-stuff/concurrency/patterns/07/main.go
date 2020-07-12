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
	Fan-In Concurrency Pattern
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
	numbers := numGen(done)
	alphabets := alphaGen(done)

	// symbols := symGen(done)
	// mixed := fanIn(numbers, alphabets, symbols)

	mixed := fanIn(numbers, alphabets)
	show(mixed)

	time.Sleep(1 * time.Millisecond)
	done <- true
	done <- true
	wg.Wait()
}

// fanIn function handles clubbing of multiple inputs from various
// goroutines into a single channel input
func fanIn(numbers chan int, alphabets chan rune) (out chan string) {
	wg.Add(1)
	out = make(chan string)

	go func() {
		defer wg.Done()
		for {
			select {
			case n := <-numbers:
				out <- fmt.Sprintf("%v", n)
			case a := <-alphabets:
				out <- string(a)
			default:
				close(out)
				return
			}
		}
	}()
	return out
}

func show(in chan string) {
	wg.Add(1)
	go func() {
		for v := range in {
			fmt.Printf("%v,", v)
		}
		fmt.Println()
		wg.Done()
	}()
}

func alphaGen(done chan bool) (out chan rune) {
	wg.Add(1)
	out = make(chan rune)

	go func() {
		defer wg.Done()
		letters := []rune("abcdefghijklmnopqrstuvwxyz")
		for {
			select {
			case out <- letters[rand.Intn(len(letters))]:
			case <-done:
				close(out)
				return
			}
		}
	}()
	return
}

func symGen(done chan bool) (out chan rune) {
	wg.Add(1)
	out = make(chan rune)

	go func() {
		defer wg.Done()
		symbols := []rune("αβγδεζηθικλμνξοπρστυφχψω")
		for {
			select {
			case out <- symbols[rand.Intn(len(symbols))]:
			case <-done:
				close(out)
				return
			}
		}
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
