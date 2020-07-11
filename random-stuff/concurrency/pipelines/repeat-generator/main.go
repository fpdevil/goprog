package main

import "fmt"

const (
	msg = `
	The repeat generator function will repeat the values passed
	to it infinitely until it's  stopped.  The generic pipeline
	stage that will be used in combination with this is 'take'
	Here, take will only grab the first num elements and exists
	***********************************************************
	`
)

func main() {
	fmt.Println(msg)

	done := make(chan interface{})
	defer close(done)

	for num := range take(done, repeat(done, 1, 2, 3), 10) {
		fmt.Printf("%v ", num)
	}
	fmt.Println()
}

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)

		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- (<-valueStream):
			}
		}
	}()
	return takeStream
}
