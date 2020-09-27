package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type client chan<- string // an outgoing message channel for client identity

var (
	messages = make(chan string) // global channel for all incoming messages
	entering = make(chan client) // clients attach / enter
	leaving  = make(chan client) // when clients leave / disconnect
)

// broadcaster function listens on the global entering and leaving
// channels for announcements of arriving and departing clients.
func broadcaster() {
	// keep track of all the connected clients
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// broadcast incoming messages to all the connected
			// clients over outgoing message channels
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

// handleConn function creates a new outgoing message channel for its
// client and announces the arrival of this client to the broadcaster
// over the entering channel.
func handleConn(conn net.Conn) {
	ch := make(chan string) // channel for outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, "%v\n", msg) // ignore errors from n/w
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "listener error: %v\n", err)
		return
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "accept error: %v\n", err)
			continue
		}
		go handleConn(conn)
	}
}