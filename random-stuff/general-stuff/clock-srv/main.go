package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const (
	msg = `
	A TCP based clock server to fetch and display time periodically
	***************************************************************

	`
)

func main() {
	fmt.Println(msg)

	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %v", err)
		return
	}
	for {
		con, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go connectionHandler(con)
	}
}

func connectionHandler(con net.Conn) {
	defer con.Close()
	for {
		_, err := io.WriteString(con, time.Now().Format("15:04:05 -0700\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
