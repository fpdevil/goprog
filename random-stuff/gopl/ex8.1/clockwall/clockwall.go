package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type clock struct {
	region, host, port string
}

const (
	usage = `
    usage %s region=host:port...
    `
)

func main() {
	args := os.Args
	// parse the supplied command line arguments and check for required format
	clocks, err := parseArgs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	for _, c := range clocks {
		// conn, err := net.Dial("tcp", c.host+":"+c.port)
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", c.host, c.port))
		if err != nil {
			fmt.Fprintf(os.Stderr, "connection error: %v\n", err)
			return
		}
		defer conn.Close()
		go c.mustCopy(os.Stdout, conn)
	}

	for {
		time.Sleep(1 * time.Second)
	}
}

func parseArgs(args []string) (clocks []*clock, err error) {
	if len(args) == 1 {
		err = fmt.Errorf(usage, filepath.Base(args[0]))
		return nil, err
	}
	for _, x := range args[1:] {
		params := strings.Split(x, "=")
		if len(params) != 2 {
			err = fmt.Errorf("bad argument: %v", x)
			return nil, err
		}
		hostinfo := strings.Split(params[1], ":")
		if len(hostinfo) != 2 {
			err = fmt.Errorf("invalid address: %v", hostinfo)
			return nil, err
		}
		c := &clock{params[0], hostinfo[0], hostinfo[1]}
		clocks = append(clocks, c)
	}
	return clocks, nil
}

func (c *clock) mustCopy(dst io.Writer, src io.Reader) {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		fmt.Fprintf(dst, "%-20s :: %-20s\n", c.region, scanner.Text())
	}
	fmt.Printf("finished %v\n", c.region)
	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "read error %s: %s\n", c.region, scanner.Err())
		return
	}
}
