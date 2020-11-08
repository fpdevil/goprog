package main

import (
	"flag"
	"fmt"
)

func main() {
	c := Config{}
	c.Setup()
	flag.Parse()
	fmt.Printf("%v\n", c.GetMessage())
}
