package main

import (
	"fmt"
	"net"
	"os"
)

const (
	msg = `
* GO Simple FTP Server *
`
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
	fmt.Println(msg)

	service := "0.0.0.0:1225"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// transfer buffer of 512 bytes
	var trfrBuf [512]byte
	for {
		n, err := conn.Read(trfrBuf[0:])
		if err != nil {
			conn.Close()
			return
		}

		s := string(trfrBuf[0:n])
		// now decode the client request
		if s[0:2] == CD {
			chdir(conn, s[3:])
		} else if s[0:3] == DIR {
			dirlist(conn)
		} else if s[0:3] == PWD {
			pwd(conn)
		}
	}
}

func chdir(conn net.Conn, s string) {
	if os.Chdir(s) == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("ERROR"))
	}
}

func pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(""))
		return
	}
	conn.Write([]byte(s))
}

func dirlist(conn net.Conn) {
	// a blank line will be sent upon termination
	defer conn.Write([]byte("\r\n"))

	dir, err := os.Open(".")
	if err != nil {
		return
	}

	files, err := dir.Readdirnames(-1)
	if err != nil {
		return
	}

	for _, f := range files {
		conn.Write([]byte(f + "\r\n"))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred: %v\n", err.Error())
		return
	}
}
