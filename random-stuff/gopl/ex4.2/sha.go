package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

/*
Write a program  that prints the SHA256 hash of  its standard input by
default but supports a command-line flag to print the SHA284 or SHA512
hash instead.
*/

var in = flag.Int("sha", 256, "hash width (384 / 512)")

func main() {
	flag.Parse()
	var shasum func(b []byte) []byte

	switch *in {
	case 256:
		shasum = func(b []byte) []byte {
			fmt.Println("processing SHA256")
			h := sha256.Sum256(b)
			return h[:]
		}
	case 384:
		shasum = func(b []byte) []byte {
			fmt.Println("processing SHA384")
			h := sha512.Sum384(b)
			return h[:]
		}
	case 512:
		shasum = func(b []byte) []byte {
			fmt.Println("processing SHA512")
			h := sha512.Sum512(b)
			return h[:]
		}
	default:
		log.Errorln("invalid sha value specified")
	}

	// handle input data as byres
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	fmt.Fprintf(os.Stdout, "%x\n", shasum(b))
}
