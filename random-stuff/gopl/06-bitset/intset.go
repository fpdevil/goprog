package main

import (
	"bytes"
	"fmt"
)

// IntSet structure represents a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64 // words is an unsigned integer slice
}

//!+Has

// Has reports whether the set contains the non-negative value i.
// The set would contain a value i if the i-th bit is a set.
// @param i: a non-negative integer value
// @return: [bool] check if set contains i
func (s *IntSet) Has(i int) bool {
	// each word will have 64 bits, so for locating the bit for x, we can
	// use the quotient x/64 as its word index and the remainder x%64 as
	// the bit index within that word.
	word, bit := i/64, uint(i%64)
	predicate := word < len(s.words)
	return predicate && s.words[word]&(1<<bit) != 0
}

//!-Has

//!+Add

// Add method adds the non-negative value i to the set, which is nothing
// but setting nth bit at a value a using => a | (1 << n)
// @param i: a non negative integer value
// @returns: [bool]
func (s *IntSet) Add(i int) {
	// each word will have 64 bits, so for locating the bit for x, we can
	// use the quotient x/64 as its word index and the remainder x%64 as
	// the bit index within that word.
	word, bit := i/64, uint(i%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

//!-Add

//!+UnionOf

// UnionOf method sets s to the union of s and t
// @param s: a non-negative integer set
// @param t: a non-negative integer set
func (s *IntSet) UnionOf(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			// use bitwise OR to compute union 64 elements at a time
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-UnionOf

//!+String

// String returns the set as a string representation for printing
// in the form of "{1, 2, 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		// skip zero values
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-String

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(145)
	x.Add(5)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(5)
	y.Add(87)
	fmt.Println(y.String())

	x.UnionOf(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(5), x.Has(120))
}
