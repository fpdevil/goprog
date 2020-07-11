package main

import "fmt"

const (
	msg = `
	constructing pipelines using channels; using channels
	for stage to receive, transform and emit the data
	*****************************************************

	`
)

func main() {
	fmt.Println(msg)

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 3)
	for v := range pipeline {
		fmt.Printf("%#v\n", v)
	}
}

// generator  function  takes  in  a  variadic  slice  of  integers,
// constructs a buffered channel of  integers with a length equal to
// the incoming integer  slice, starts a goroutine,  and returns the
// constructed channel
// it just converts a discrete set  of values into  a stream of data
// on a channel
func generator(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for _, v := range integers {
			select {
			case <-done:
				return
			case intStream <- v:
			}
		}
	}()
	return intStream
}

func multiply(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
	multiplyStream := make(chan int)
	go func() {
		defer close(multiplyStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case multiplyStream <- i * multiplier:
			}
		}
	}()
	return multiplyStream
}

func add(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
	addStream := make(chan int)
	go func() {
		defer close(addStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case addStream <- i + additive:
			}
		}
	}()
	return addStream
}
