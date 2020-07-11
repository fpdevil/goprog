package main

import "fmt"

func main() {
	fmt.Println("/// channels 02 ///")
	fmt.Println()

	var ch chan int // will be nil until initialized
	fmt.Printf("channel: %#v,  length: %d, capacity: %d\n", ch, len(ch), cap(ch))

	// sending data to a nil/un-initialized channel will fail
	// ch <- 1

	// receiving from a nil or un-initialized channel will fail
	// <- ch

	// make or initialize an unbuffered channel
	// ch = make(chan int)
	// fmt.Printf("channel: %#v,  length: %d, capacity: %d\n", ch, len(ch), cap(ch))

	// sending data to a  unbuffered channel without recevier will fail
	// ch <- 3
	// fmt.Printf("channel: %#v,  length: %d, capacity: %d\n", ch, len(ch), cap(ch))

	// for receiving the data from channel
	// fmt.Println(<- ch)
	// fmt.Printf("channel: %#v,  length: %d, capacity: %d\n", ch, len(ch), cap(ch))

	// making buffered channel
	ch = make(chan int, 2)
	fmt.Printf("channel: %#v,  length: %d, capacity: %d\n", ch, len(ch), cap(ch))

	// send
	ch <- 3
	ch <- 4
	fmt.Printf("channel: %#v,  length: %d, capacity: %d\n", ch, len(ch), cap(ch))
	fmt.Printf("received first value from %#v <- %v\n", ch, <-ch)
	fmt.Printf("received second value from %#v <- %v\n", ch, <-ch)

	// receive value from channel
	ch <- 99 // first send
	v := <- ch
	fmt.Printf("channel: %#v,  length: %d, capacity: %d\n", ch, len(ch), cap(ch))
	fmt.Printf("received value from %#v <- %v\n", ch, v)

	// Read-Onlky & Write-Only channels
	chs := make(chan string, 5)

	// Sender or Write Only channels
	// initializing by assigning to defined channel chs
	var out chan<- string
	out = chs
	out <- "Pride"
	out <- "Prejudice"
	fmt.Printf("send only channel: %#v,  length: %d, capacity: %d\n", out, len(out), cap(out))

	// receive only channel
	var in <-chan string
	in = chs
	fmt.Println(<- in)
	fmt.Println(<- in)
}
