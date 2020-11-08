/*
Echo Server
*/

package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func main() {
	// bind to a port 8080 over all interfaces
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:8080")
	for {
		// wait for a connection and create net.Conn over connection established.
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// go echo(conn)
		go echo(conn)
	}
}

// updEcho function is a handler which simply echoes the received data
func updEcho(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	log.Printf("Read %d bytes: %s", len(s), s)

	log.Println("Writing data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}
	writer.Flush()
}

// echo function is a handler which simply echoes the received data
func echo(conn net.Conn) {
	defer conn.Close()

	// create a buffer to store the received data
	b := make([]byte, 512)
	for {
		// receive data through conn.Read into a buffer.
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s", size, string(b))

		// send data through conn.Write.
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}
