package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("* Parse IP *")
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <ip address> <ones> <bits>\n", filepath.Base(args[0]))
		return
	}

	ip := args[1]

	addr := net.ParseIP(ip)
	if addr == nil {
		fmt.Println("invalid ip address.")
		return
	}

	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println("Address: ", addr.String())
	fmt.Println("Default IP Mask Length: ", bits)
	fmt.Println("Leading ones count: ", ones)
	fmt.Println("Mask is (hex): ", mask.String())
	fmt.Println("Network: ", network.String())
}
