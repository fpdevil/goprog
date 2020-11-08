package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

type Part struct {
	Id   int
	Name string
}

type LowerCase interface {
	LowerCase()
}

type UpperCase interface {
	UpperCase()
}

type LowerUpperCase interface {
	LowerCase
	UpperCase
}

type FixCaser interface {
	FixCase()
}

type ChangeCaser interface {
	LowerUpperCase
	FixCaser
}

type StringPair struct {
	first, second string
}

func main() {
	fmt.Println("* Testing Embedded Interfaces *")
	fmt.Println()

	// testing the StringPair implementations
	const size = 20

	jules := &StringPair{"Jules", "Verne"}
	alex := &StringPair{"Alexander", "Dumas"}
	canon := StringPair{"Sir Arthur Canon", "Doyle"}

	for _, reader := range []io.Reader{jules, alex, &canon} {
		raw, err := ToBytes(reader, size)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reader error: %v\n", err)
		}
		fmt.Fprintf(os.Stdout, "%q\n", raw)
	}

	tom := Part{9781495238956, "THE ADVENTURES OF TOM SAWYER"}
	mark := &StringPair{"Mark", "Twain"}
	fey := StringPair{"Richard P.", "Feynman"}

	tom.LowerCase()
	mark.FixCase()
	fey.FixCase()
	fmt.Println(tom)
	fmt.Println(mark)
	fmt.Println(fey)
}

// ToBytes function takes an io.Reader and a size limit and returns
// a []byte containing reader's data and an error if any
func ToBytes(reader io.Reader, size int) ([]byte, error) {
	data := make([]byte, size)
	n, err := reader.Read(data)
	if err != nil {
		return data, err
	}
	return data[:n], nil
}

// Exchange function swaps the two elements of StringPair
func (pair *StringPair) Exchange() {
	pair.first, pair.second = pair.second, pair.first
}

// Read fuction implementation to make StringPair satisfy the
// io.Reader interface
func (pair *StringPair) Read(p []byte) (n int, err error) {
	if pair.first == "" && pair.second == "" {
		return 0, io.EOF
	}
	if pair.first != "" {
		n = copy(p, pair.first)
		pair.first = pair.first[n:]
	}
	if n < len(p) && pair.second != "" {
		m := copy(p[n:], pair.second)
		pair.second = pair.second[m:]
		n += m
	}
	return n, nil
}

// fixCase function returns a copy of the string it is given in
// which every character has been lower cased, except for the very
// first character and the first character after each whitespace or
// hyphen character, which are upper cased
func fixCase(s string) string {
	var chars []rune
	upper := true
	for _, char := range s {
		if upper {
			char = unicode.ToUpper(char)
		} else {
			char = unicode.ToLower(char)
		}
		chars = append(chars, char)
		upper = unicode.IsSpace(char) || unicode.Is(unicode.Hyphen, char)
	}
	return string(chars)
}

// UpperCase function is an implementation for converting the provided
// string pair into upper case
func (pair *StringPair) UpperCase() {
	pair.first = strings.ToUpper(pair.first)
	pair.second = strings.ToUpper(pair.second)
}

// LowerCase function is an implementation for converting the provided
// string pair into lower case
func (pair *StringPair) LowerCase() {
	pair.first = strings.ToLower(pair.first)
	pair.second = strings.ToLower(pair.second)
}

// FixCase function is an implementation for converting the provided
// string pair into title case
func (pair *StringPair) FixCase() {
	pair.first = fixCase(pair.first)
	pair.second = fixCase(pair.second)
}

func (part *Part) FixCase() {
	part.Name = fixCase(part.Name)
}

func (part *Part) LowerCase() {
	part.Name = strings.ToLower(part.Name)
}

func (part *Part) UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}
