package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

// Clock is a TCP server that periodically writes the current time

func main() {
	p := "8000"
	if len(os.Args) == 2 {
		p = os.Args[1]
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", p))
	if err != nil {
		log.Printf("failure %s", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			// may fail for any reason eg. connection abort
			log.Printf("accept error %s", err.Error())
			continue
		}
		// handle connections sequentially 1 at a time
		// handleConn(conn)
		// making server concurrent by defining as a goroutine
		go handleConn(conn)
	}
}

//!+handleConn

// handleConn function handles one complete client connection
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // client is disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

//!-
