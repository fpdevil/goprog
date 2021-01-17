// All incoming requests are forwarded to a channel, which processes them
// one by one. When the channel is  done processing a request, it sends a
// message to  the original caller saying  that it is ready  to process a
// new one. So,  the capacity of the buffer of  the channel restricts the
// number of simultaneous requests that it can keep.
package main

import "fmt"

func main() {
	numbers := make(chan int, 5)
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	for i := 0; i < counter+5; i++ {
		select {
		case num := <-numbers:
			fmt.Println(num)
		default:
			fmt.Println("Nothing more to be done!")
			break
		}
	}
}
