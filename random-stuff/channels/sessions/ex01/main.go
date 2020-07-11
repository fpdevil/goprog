package main

import (
	"fmt"
	"strings"
)

var (
	datax = []string{
		"the bad guns of navarone",
		"theory of relativity",
		"apples and bees",
		"the bad boys club",
		"to be or not to be",
	}

	datay = []string{"a", "b", "b", "c", "d", "d", "d", "e"}
)

func main() {
	fmt.Println("--- first channel setup")
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(datax, c0)
	go filterGopher(c0, c1)
	printGopher(c1)
	fmt.Println()

	fmt.Println("--- second channel setup")
	d0 := make(chan string)
	d1 := make(chan string)
	go sourceGopher(datay, d0)
	go removeDuplicates(d0, d1)
	printGopher(d1)
	fmt.Println()

	fmt.Println("--- third channel setup")
	e0 := make(chan string)
	e1 := make(chan string)
	go sourceGopher(datax, e0)
	go splitWords(e0, e1)
	printGopher(e1)
	fmt.Println()

	fmt.Println("--- fourth channel setup")
	sieve()

}

// sourceGopher defined a send only channel to which we pump data
func sourceGopher(data []string, downstream chan<- string) {
	for _, v := range data {
		downstream <- v
	}
	close(downstream)
}

// filterGopher filters data of strings containing `bad`
func filterGopher(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Printf("%+v\n", v)
	}
}

func removeDuplicates(upstream, downstream chan string) {
	previous := ""
	for v := range upstream {
		if v != previous {
			downstream <- v
			previous = v
		}
	}
	close(downstream)
}

func splitWords(upstream, downstream chan string) {
	for v := range upstream {
		for _, word := range strings.Fields(v) {
			downstream <- word
		}
	}
	close(downstream)
}

func generate(limit int, upstream chan int) {
	for i := 2; i <= limit; i++ {
		upstream <- i
	}

	close(upstream)
}

func filter(upstream, downstream chan int, digit int) {
	for v := range upstream {
		// fmt.Println(digit, v)
		if v != digit && v%digit != 0 {
			downstream <- v
		}
	}
	close(downstream)
}

func sieve() {
	n := 20
	p0 := make(chan int)
	go generate(n, p0)
	for i := 0; i <= n; i++ {
		digit := <-p0
		if digit != 0 {
			println(digit)
		}
		p1 := make(chan int)
		go filter(p0, p1, digit)
		p0 = p1
	}
}

func printp(upstream chan int) {
	for v := range upstream {
		fmt.Printf("%+v\n", v)
	}
}
