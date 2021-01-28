package main

import (
	"bytes"
	"fmt"
)

// comma function takes a string and inserts a comma every
// three places
// "12345" => "12,345"
func comma(s string) string {
	var buf bytes.Buffer
	i := len(s) % 3
	if i == 0 {
		i = 3
	}

	buf.WriteString(s[:i])
	for j := i; j < len(s); j += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[j : j+3])
	}
	return buf.String()
}

func main() {
	fmt.Println("vim-go")
	fmt.Printf("%v => %v\n", "12345", comma("12345"))
	fmt.Printf("%v => %v\n", "123456789", comma("123456789"))
}
