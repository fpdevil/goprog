package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	msg = `
	Generating random numbers...
	`
)

func main() {
	fmt.Println(msg)
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: go run %s.go <startup integer>", filepath.Base(os.Args[0]))
		return
	}

	t := time.Now().UTC().UnixNano()
	rand.Seed(t)
	createInt := make(chan int)
	done := make(chan bool)

	i, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("creating %d random numbers...\n", i)
	go gen(0, 2*i, createInt, done)

	for j := 0; j < i; j++ {
		fmt.Printf("%d ", <-createInt)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("done...")
	done <- true
}

func gen(min, max int, createInt chan int, done chan bool) {
	for {
		select {
		case createInt <- rand.Intn(max-min) + min:
		case <-done:
			close(done)
			return
		case <-time.After(2 * time.Second):
			fmt.Printf("\ntime.After(2)!\n")
		}
	}
}
