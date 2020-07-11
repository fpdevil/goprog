package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	msg = `
	Demonstration of handling the memory leak in a goroutine
	from a blocked channel while perform a WRITE...
	********************************************************

`

	m = 100 // generate random number from 0 to m inclusive
	n = 10  // generate n such random numbers
)

func main() {
	fmt.Println(msg)

	t := time.Now().UTC().UnixNano()
	rand.Seed(t)

	done := make(chan interface{})
	fmt.Printf("generating %d random numbers each ranging from 0 to %d\n", n, m+1)
	randStream := genRandomStream(done)

	for i := 0; i <= n; i++ {
		fmt.Printf("%2d -> %3d\n", i, <-randStream)
	}
	close(done)

	// simulate some work being done...
	time.Sleep(1 * time.Second)
	fmt.Println("Done with generation of Random numbers!")
}

// gerRandomStream function takes a decider channel done
// which has a read-only chahnel of type empty interface
// and it returns a read-only integer channel consisting
// of a random stream of integers
func genRandomStream(done <-chan interface{}) <-chan int {
	randStream := make(chan int)
	go func() {
		defer fmt.Println("closing func genRandomStream()!")
		defer close(randStream)

		for {
			select {
			case <-done:
				return
			case randStream <- rand.Intn(m) + 1:
			}
		}
	}()
	return randStream
}
