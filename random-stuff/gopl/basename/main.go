// basename removes the directory components and a .suffix
// e.g., a => a
//		 a.go => a
//		 a/b/c.go => c
//		 a/b.c.go => b.c
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// comma inserts commas in a non-negative decimal integer string
// abcde => ab,cde
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// intsToString function converts the int list to csv string
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(intsToString([]int{1, 2, 3, 4, 5}))
	fmt.Printf("%v => %v\n", "12345", comma("12345"))
	fmt.Printf("%v => %v\n", "123456789", comma("123456789"))
}
