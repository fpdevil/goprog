package main

import (
	"bytes"
	"fmt"
	"strings"
)

// helper function takes a string and inserts a comma every
// three places
// "12345" => "12,345"
func helper(s string, buf *bytes.Buffer) {
	i := len(s) % 3
	if i == 0 {
		i = 3
	}

	buf.WriteString(s[:i])
	for j := i; j < len(s); j += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[j : j+3])
	}
}

func comma(s string) string {
	// separate out the sign of the number if available
	var numsign string
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		numsign = s[:1]
		s = s[1:]
	}

	var buf bytes.Buffer
	buf.WriteString(numsign)

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		helper(s[:dot], &buf)
		buf.WriteString(".")
		helper(s[dot+1:], &buf)
	} else {
		helper(s, &buf)
	}

	return buf.String()
}

func main() {
	fmt.Println("vim-go")
	fmt.Printf("%v => %v\n", "123.45", comma("123.45"))
	fmt.Printf("%v => %v\n", "-10241.678", comma("-10241.678"))
	fmt.Printf("%v => %v\n", "1234567.891", comma("1234567.891"))
}
