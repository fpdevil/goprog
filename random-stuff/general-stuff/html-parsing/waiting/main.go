package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	timeout = 1 * time.Minute
)

func main() {
	fmt.Println("Wait & Reconnect...")
	fmt.Println()

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: wait url\n")
		return
	}

	url := os.Args[1]
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "site is down: %v\n", url)
		return
	}
}

// WaitForServer function attemps to contact the URL of the server
// It then tries for a minute using the exponential back-off and
// will report an error in case all attempts have failed
func WaitForServer(url string) error {
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			// Happy case
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // this is the exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
