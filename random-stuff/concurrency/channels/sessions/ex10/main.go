package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("/// multiple channels ///")
	fmt.Println()

	fmt.Println("get a notification from a channel with notifyAfter()")
	fmt.Printf("Message 1 at: %v\n", time.Now())
	alarm := notifyAfter(1 * time.Second)
	<-alarm
	fmt.Printf("Message 2 at: %v\n", time.Now())

	fmt.Println()
	fmt.Println("get a notification from a channel with time.After()")
	fmt.Printf("Message 1 at: %v\n", time.Now())
	<-time.After(1 * time.Second)
	fmt.Printf("Message 1 at: %v\n", time.Now())

	fmt.Println()
	fmt.Println("selecting from multiple channels")
	var ch1, ch2 chan string
	select {
	case <-ch1:
		log.Info("Received Msg *FROM* ch1")
	case ch2 <- "hello":
		log.Info("Sent Msg *TO* ch2")
	default:
		log.Info("No communication from ch1 or ch2")
	}

	fmt.Println()
	fmt.Println("Random BIT Stream Generation:")
	bits := genRandomBits(10)
	for v := range bits {
		fmt.Print(v)
	}
	fmt.Println()
}

func genRandomBits(l int) (out chan int8) {
	out = make(chan int8, l)

	defer close(out)

	for {
		select {
		case out <- 0:
		case out <- 1:
		default:
			// channel is full
			return
		}
	}
}

func notifyAfter(delay time.Duration) (out chan time.Time) {
	out = make(chan time.Time)
	go func() {
		time.Sleep(delay)
		out <- time.Now()
		close(out)
	}()
	return out
}
