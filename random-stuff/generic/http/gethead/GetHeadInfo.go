package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("* HTTP HEAD *")
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s host:port\n", filepath.Base(args[0]))
		return
	}

	service := args[1]

	// tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	// checkErr(err)
	// fmt.Printf("checking http HTTP against %s\n", tcpAddr)
	// conn, err := net.DialTCP("tcp", nil, tcpAddr)

	conn, err := net.Dial("tcp", service)
	checkErr(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErr(err)

	result, err := readAll(conn)
	checkErr(err)

	fmt.Println(string(result))
}

func readAll(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte

	for {
		nbytes, err := conn.Read(buf[0:])
		result.Write(buf[0:nbytes])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal error: %s\n", err.Error())
		return
	}
}
