package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	ch := make(chan struct{})
	in := bufio.NewScanner(c)
	go func() {
		for {
			if in.Scan() {
				ch <- struct{}{}
			} else {
				close(ch)
				return
			}
		}
	}()
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				c.Close()
				return
			}
			go echo(c, in.Text(), 1*time.Second)
		case <-time.After(10 * time.Second):
			c.Close()
			return
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}
		go handleConn(conn)
	}
}
