package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("NETCAT")
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // ignore errors
		fmt.Printf("finished\n")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done

	// defer conn.Close()
	// go mustCopy(os.Stdout, conn)
	// mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
}
