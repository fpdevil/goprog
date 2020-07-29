package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("* Ping-Pong *")

	rand.Seed(time.Now().UTC().UnixNano())
	c := make(chan int)
	wg.Add(2)
	go run("Ping", c)
	go run("Pong", c)

	c <- 1

	wg.Wait()
}

func run(id string, count chan int) {
	defer wg.Done()
	for {
		c, ok := <-count
		if !ok {
			fmt.Fprintf(os.Stdout, "%s completed\n", id)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("%s\n", id)
			close(count)
			return
		}
		fmt.Printf("%s %d\n", id, c)
		c++
		count <- c
	}
}
