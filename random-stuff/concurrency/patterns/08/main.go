package main

import (
	"fmt"
	"time"
)

//!+orfun

// orfun function  takes in  a variadic  slice of  channels and
// returns a  single channel.  The function enables  to combine
// any number of  channels together into a  single channel that
// will  close as  soon as  any  of it  component channels  are
// closed or written to.
func orfun(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		// if variadic slice is empty we return a nil channel
		return nil
	case 1:
		// if contains only a single element, just return it
		return channels[0]
	}

	orDone := make(chan interface{})
	// create a go routine that will wait for messages on our
	// channels without blocking
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			// every recursive call will have atleast 2 channels
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-orfun(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

//!-

func main() {
	fmt.Println("OR Channel")
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = orfun

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v\n", time.Since(start))
}
