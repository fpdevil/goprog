package main

import "fmt"

// ByteCounter for counting bytes
type ByteCounter int

// Write method satisfies the io.Writer interface
func (bc *ByteCounter) Write(p []byte) (int, error) {
	*bc += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	fmt.Println("-- ByteCounter --")
	var c ByteCounter
	c.Write([]byte("Hello!"))
	fmt.Println(c)

	c = 0
	var name = "Nutshell"
	fmt.Fprintf(&c, "hello, %s!", name)
	fmt.Println(c)
}
