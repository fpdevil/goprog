package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("/// Sieve Of Erathostenes ///")

	var wg sync.WaitGroup
	p0 := make(chan int)
	p1 := make(chan int)

	limit := 25
	go generate(limit, p0)

	wg.Add(limit)
	go sieve(p0, p1, &wg)
	wg.Wait()

	printp(p1)
	fmt.Println()
}

func generate(limit int, upstream chan int) {
	for i := 2; i <= limit; i++ {
		upstream <- i
	}

	close(upstream)
}

func sieve(upstream chan int, downstream chan int, wg *sync.WaitGroup) {
	var i int
	for v := range upstream {
		for i := 2; i < v; i++ {
			if v%i == 0 {
				break
			}
		}
		if i == v {
			fmt.Println(i)
			downstream <- i
		}
	}
	defer wg.Done()
	close(downstream)
}

func printp(upstream chan int) {
	for v := range upstream {
		fmt.Printf("%+v\n", v)
	}
}
