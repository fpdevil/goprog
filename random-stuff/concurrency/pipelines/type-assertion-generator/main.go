package main

import "fmt"

const (
	msg = `
	Type Assertion Generator. This generator performs the
	type assertion when placed as a stage when we need to
	deal with specific types
	*****************************************************
	`
)

func main() {
	fmt.Println(msg)

	done := make(chan interface{})
	defer close(done)

	var message string
	for token := range toString(done, take(done, repeat(done, "I ", "am ", "doing "), 10)) {
		message += token
	}
	fmt.Printf("message: %s...", message)
	fmt.Println()
}

// toString function for handling the toString pipelime stage
func toString(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)

		for v := range valueStream {
			select {
			case <-done:
				return
			case stringStream <- v.(string):
			}
		}
	}()
	return stringStream
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
