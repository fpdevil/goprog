package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	fmt.Println("* Parsing IPv4 Addresses *")
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s logfil\n", filepath.Base(args[0]))
		return
	}

	for _, filename := range args[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s: %s\n", filename, err)
			return
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading the file %s\n", err)
				break
			}

			ip := findIP(line)
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}
			fmt.Printf("ip: %v\n", ip)
		}
	}
}

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}
