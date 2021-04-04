package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// client struct represents the structure of the client message
// model with name and an outgoing messaging channel
type client struct {
	name    string
	xclient chan<- string
}

var (
	messages = make(chan string) // global channel for all incoming messages from all clients
	entering = make(chan client) // channel for arriving clients  (attach)
	leaving  = make(chan client) // channel for departing clients (disconnect)
)

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "tcp listen error: %s\n", err.Error())
		return
	}

	// the broadcaster go routine handles 3 different kinds of messages
	go broadcaster()

	for {
		// accepting for each incoming connection from client
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "accept connection error : %s\n", err.Error())
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
	// clients here keeps a track of all the connected clients
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// Broadcast the incoming message to all the connected
			// clients through their outgoing message channel
			for cli := range clients {
				cli.xclient <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			var names []string // list of all existing clients
			for c := range clients {
				names = append(names, c.name)
			}
			cli.xclient <- fmt.Sprintf("%d clients are online %v\n", len(names), strings.Join(names, ", "))
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.xclient)
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
	messages <- who + " has arrived!"
	entering <- client{name: who, xclient: ch}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- "Message from " + who + ": " + input.Text()
	}

	leaving <- client{name: who, xclient: ch}
	messages <- who + " has left the group!"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, "%v\n", msg)
	}
}

//!-handleConn
