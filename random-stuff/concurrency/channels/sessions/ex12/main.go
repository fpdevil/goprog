package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("/// Channel Selection: Wait timeout ///")
	fmt.Println()

	d := producer()
	consumer(d)
}

func producer() (out chan string) {
	out = make(chan string)
	go func() {
		count := 1
		for i := 0; i < 10; i++ {
			out <- fmt.Sprintf("Producer sent message %v\n", count)
			count++
		}

		// specify a sleep time in between
		time.Sleep(3 * time.Millisecond)

		for i := 0; i < 10; i++ {
			out <- fmt.Sprintf("Producer sent message %v\n", count)
			count++
		}

		close(out)
	}()
	return
}

func consumer(in chan string) {
	for {
		// provide a suitable timeout for consuming the messages
		// alarm := time.After(2 * time.Millisecond)
		alarm := time.After(4 * time.Millisecond)
		select {
		case m, ok := <-in:
			if !ok {
				fmt.Println("no more data from `in`")
				return
			}
			fmt.Printf("Consumer received - %v", m)
		case <-alarm:
			fmt.Println("Timedout waiting for the data")
			return
		}
	}
}
