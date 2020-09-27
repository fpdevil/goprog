package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	msg = `
* Go FTP Client *
`
	userDir  = "dir"
	userCd   = "cd"
	userPwd  = "pwd"
	userQuit = "quit"
)

// constants transmitted across the network
const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
	fmt.Println(msg)
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %v host", os.Args[0])
		return
	}

	host := os.Args[1]

	conn, err := net.Dial("tcp", host+":1225")
	checkErr(err)

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		// remove the trailing whitespaces
		line = strings.TrimRight(line, " \t\r\n")
		if err != nil {
			break
		}
		strs := strings.SplitN(line, " ", 2)
		switch strs[0] {
		case userDir:
			dirRequest(conn)
		case userCd:
			if len(strs) != 2 {
				fmt.Println("cd <dir>")
				continue
			}
			fmt.Println("CD \"", strs[1], "\"")
			cdRequest(conn, strs[1])
		case userPwd:
			pwdRequest(conn)
		case userQuit:
			conn.Close()
			return
		default:
			fmt.Println("Unknown command!")
		}
	}
}

func dirRequest(conn net.Conn) {
	conn.Write([]byte(DIR + " "))

	// transfer buffer of 512 bytes
	var trfbuf [512]byte
	result := bytes.NewBuffer(nil)

	for {
		n, _ := conn.Read(trfbuf[0:])
		result.Write(trfbuf[0:n])
		length := result.Len()
		contents := result.Bytes()
		if string(contents[length-4:]) == "\r\n\r\n" {
			fmt.Println(string(contents[0 : length-4]))
			return
		}
	}
}

func cdRequest(conn net.Conn, dir string) {
	conn.Write([]byte(CD + " " + dir))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	if s != "OK" {
		fmt.Println("Failed to change dir")
	}
}

func pwdRequest(conn net.Conn) {
	conn.Write([]byte(PWD))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	fmt.Println("Current dir\"" + s + "\"")
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred: %v\n", err.Error())
		return
	}
}
