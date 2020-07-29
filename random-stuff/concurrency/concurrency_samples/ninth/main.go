package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("* create and schedule goroutines *")
	fmt.Println()

	runtime.GOMAXPROCS(1)
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("spawning 2 goroutines...")

	go func() {
		defer wg.Done()
		// print each alphabet sequence thrice
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		// print each alphabet sequence thrice
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	wg.Wait()
	fmt.Println("\nTerminating the program!")
	fmt.Printf("total time taken: %v\n", time.Since(start))
}
