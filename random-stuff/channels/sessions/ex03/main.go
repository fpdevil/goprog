package main

import "fmt"

func main() {
	fmt.Println("/// Working with Closed Channels ///")
	fmt.Println()

	// create a string Channel
	sc := make(chan string, 2)
	sc <- "Pride"

	// close the Channel
	close(sc)
	fmt.Printf("sc: %#v, length: %v, capacity: %v\n", sc, len(sc), cap(sc))

	// sending data to a closed channel will fail
	// sc <- "Prejudice"

	// receiving value from a closed channel
	s := <-sc
	fmt.Printf("sc 1st value received: %v, len: %v, capacity: %v\n", s, len(sc), cap(sc))
	// now length is 0, but we can still try to receive values from a closed channel
	s = <-sc
	fmt.Printf("sc 2nd value received: %v, len: %v, capacity: %v\n", s, len(sc), cap(sc))

	fmt.Println()
	// testing values sent before closure
	sc = make(chan string, 10)
	sc <- "Pride"
	sc <- "Prejudice"
	close(sc)

	// receive values using the `ok` to know if channel is open or closed
	var ok bool
	s, ok = <-sc
	fmt.Printf("sc 1st value received: %v, ok: %v\n", s, ok)
	s, ok = <-sc
	fmt.Printf("sc 2nd value received: %v, ok: %v\n", s, ok)
	s, ok = <-sc
	fmt.Printf("sc 3rd value received: %v, ok: %v\n", s, ok)
}
