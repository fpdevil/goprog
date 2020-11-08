package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("NETCAT")
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "usage %v <localport>\n", filepath.Base(args[0]))
		return
	}

	// conn, err := net.Dial("tcp", "localhost:8000")
	// port := strconv.Atoi(args[1])

	port := args[1]
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	done := make(chan struct{})

	// read data from the connection and write it to the standard
	// output until an end-of-file condition or an error occurs
	go func() {
		io.Copy(os.Stdout, conn) // ignore errors
		fmt.Printf("done\n")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	// this closes both halves (Write + Read) of the network connection
	defer conn.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		if err == io.EOF {
			fmt.Fprintf(os.Stderr, "%v\n", err.Error())
			return
		}
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		return
	}
}
