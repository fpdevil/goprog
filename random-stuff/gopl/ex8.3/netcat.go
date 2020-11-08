package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("* NETCAT *")
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage %s <local port>\n", filepath.Base(os.Args[0]))
		return
	}

	port := os.Args[1]
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %s\n", err.Error())
		return
	}

	TCPConn, ok := conn.(*net.TCPConn)
	if !ok {
		fmt.Fprintf(os.Stderr, "no TCP network connection\n")
		return
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	TCPConn.CloseWrite()
	<-done // wait for background goroutine to finish
	TCPConn.CloseRead()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
