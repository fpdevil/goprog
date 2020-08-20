package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var port = flag.Int("port", 8000, "listen port")

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05 Mon Jan 2 2006\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		fmt.Fprintf(os.Stderr, "clock error: %v\n", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("clock error: %v\n", err)
			continue
		}
		go handleConn(conn)
	}
}
