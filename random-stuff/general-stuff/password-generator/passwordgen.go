package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("* generating random numbers upto 94 *")
	var LENGTH int64 = 8
	args := os.Args

	switch len(args) {
	case 2:
		LENGTH, _ = strconv.ParseInt(args[1], 10, 64)
		if LENGTH <= 0 {
			LENGTH = 8
		}
	default:
		fmt.Println("using default values!")
	}

	pass, err := genPassword(LENGTH)
	if err != nil {
		fmt.Printf("error %s\n", err.Error)
		return
	}
	fmt.Printf("password: %s\n", pass[0:LENGTH])
}

func genBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func genPassword(s int64) (string, error) {
	b, err := genBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
