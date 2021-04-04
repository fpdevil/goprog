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
	Generating %d random numbers...
	`
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: go run %s.go <startup integer>", filepath.Base(os.Args[0]))
		return
	}

	t := time.Now().UTC().UnixNano()
	rand.Seed(t)
	createInt := make(chan int)
	done := make(chan bool)

	i, _ := strconv.Atoi(os.Args[1])
	fmt.Printf(msg, i)
	go gen(0, 2*i, createInt, done)

	for j := 0; j < i; j++ {
		fmt.Printf("%d ", <-createInt)
	}

	time.Sleep(250 * time.Millisecond)
	fmt.Print("\n	done... \n")
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
