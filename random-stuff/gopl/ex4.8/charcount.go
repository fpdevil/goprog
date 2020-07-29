package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	fmt.Println("* Unicode character count *")

	counts := make(map[rune]int)       // map for counting unicode characters
	categories := make(map[string]int) // map for counting unicode character categories
	var utflen [utf8.UTFMax + 1]int    // map for counting lengths of utf-8 encodings
	var invalid int                    // invalid utf-8 characters count

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			return
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utflen[n]++
		countUTF(r, categories)
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

	fmt.Print("\ncategory\tcount\n")
	for j, k := range categories {
		fmt.Printf("%v\t\t%v\n", j, k)
	}
}

func countUTF(r rune, categories map[string]int) {
	switch {
	case unicode.IsMark(r):
		categories["Mark"]++
	case unicode.IsDigit(r):
		categories["Digit"]++
	case unicode.IsLower(r):
		categories["Lower"]++
	case unicode.IsPrint(r):
		categories["Print"]++
	case unicode.IsPunct(r):
		categories["Punctuation"]++
	case unicode.IsSpace(r):
		categories["Space"]++
	case unicode.IsTitle(r):
		categories["Title"]++
	case unicode.IsUpper(r):
		categories["Title"]++
	case unicode.IsSymbol(r):
		categories["Symbol"]++
	case unicode.IsControl(r):
		categories["Control"]++
	case unicode.IsGraphic(r):
		categories["Graphic"]++
	}
}
