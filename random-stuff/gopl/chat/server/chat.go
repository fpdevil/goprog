// Chat server lets several users broadcast textual messages
// to each other
//
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// client is an outgoing message channel which is write-only
type client chan<- string

var (
	messages = make(chan string) // global channel for all incoming messages from all clients
	entering = make(chan client) // channel for arriving clients  (attach)
	leaving  = make(chan client) // channel for departing clients (disconnect)
)

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "TCP listen error %s\n", err.Error())
		return
	}

	// the broadcaster go routine handles 3 different kinds of messages
	go broadcaster()
	for {
		// accepting for each incoming connection from client
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "TCP accept error %s\n", err.Error())
			continue
		}
		// for each incoming network connection from client, a new
		// handleConn goroutine.
		go handleConn(conn)
	}
}

//!-main

//!+broadcaster

// broadcaster  function  listens  on  the global  entering  and  leaving
// channels  for  announcement  of  arriving and  departing  clients.  It
// selectively handles and responds to 3 different kinds of messages from
// clients.
func broadcaster() {
	// clients map keeps a record of the current set of connected clients
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// Broadcast the incoming message to all the connected
			// clients through their outgoing message channel
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			// This is when a new client joins or attaches
			clients[cli] = true
		case cli := <-leaving:
			// This is when  a client leaves or disconnects,  in which case
			// outgoing  message  channel  is  closed and  client  will  be
			// deleted from the clients map
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

//!+handleConn

// handleConn function handles the requests for each client. It creates a
// new outgoing message channel for each client and announces the arrival
// of the same to the broadcaster via the entering channel.
func handleConn(conn net.Conn) {
	ch := make(chan string)   // channel for all outgoing client messages
	go clientWriter(conn, ch) // write messages to clients connection

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
		fmt.Fprintf(conn, "%s\n", msg)
	}
}

//!-handleConn
