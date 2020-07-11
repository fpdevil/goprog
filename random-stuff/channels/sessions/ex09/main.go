package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("/// select-case ///")
	fmt.Println()

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(2 * time.Second)
		}
	}()
	go func() {
		for {
			c3 <- "from 3"
			time.Sleep(3 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case msg3 := <-c3:
				fmt.Println(msg3)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
