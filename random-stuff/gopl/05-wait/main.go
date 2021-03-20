package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

//!+WaitForServer

// WaitForServer function attempts  to contact the server  of remote URL,
// tries for a minute using the exponential back-off. It returns an error
// if all the attempts fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err != nil {
			return nil
		}
		log.Printf("server not responding (%v); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // an exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

//!-WaitForServer

//!+exec

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage go run %s <url>\n", filepath.Base(os.Args[0]))
		return
	}

	url := os.Args[1]
	if err := WaitForServer(url); err != nil {
		log.Fatalf("site is down: %s\n", err.Error())
		// fmt.Fprintf(os.Stderr, "site is down: %s\n", err.Error())
		// return
	}
}

//!-exec
