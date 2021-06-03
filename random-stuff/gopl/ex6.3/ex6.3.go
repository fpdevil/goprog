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

//!+UnionWith

// UnionWith method sets s to the union of s and t
// @param s: a non-negative integer set
// @param t: a non-negative integer set
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			// use bitwise OR to compute union 64 elements at a time
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-UnionWith

//!+IntersectWith

// IntersectWith method performs a set intersection of 2 non-negative
// integer sets
// @param s: a non-negative integer set
// @param t: a non-negative integer set
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-IntersectWith

//!+DifferenceWith

// DifferenceWith method returns a set difference between 2 sets
// @param s: a non-negative integer set
// @param t: a non-negative integer set
func (s *IntSet) DifferenceWith(t IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-DifferenceWith

//!+SymmetricDifference

// SymmetricDifference method returns difference of 2 sets containing
// the elements present in one set or the other but not both
// @param s: a non-negative integer set
// @param t: a non-negative integer set
func (s *IntSet) SymmetricDifference(t IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-SymmetricDifference

//!+setBitCount
// setBitCount function uses Brian Kernighan’s Algorithm
// Step 1: Keep a counter to track the number of set bits.
// Step 2: Loop until our N is not equals to 0.
// Step 3: Update our N, N = N & (N-1) and also update the counter.
// Step 4: After our N becomes 0. Report the counter.
func setBitCount(n uint64) int {
	var count int
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}

//!-setBitCount

//!+Len

// Len method returns the number of elements
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += setBitCount(word)
	}
	return count
}

//!-Len

//!+Remove

// Remove method is fr removing an element x from the set
// Toggle the Kth Bit in X => X NOT (1 << K - 1)
// Turn Off the Kth bit in X by => X AND (X NOT (1 << K - 1))
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word > len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

//!-Remove

//!+Clear

// Clear method removes all the elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

//!-Clear

//!+Copy

// Copy method returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	x := &IntSet{}
	x.words = make([]uint64, len(s.words))
	copy(x.words, s.words)
	return x
}

//!-Copy

//!+AddAll

// AddAll method takes a list of values and addes them to the set
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

//!-AddAll

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

	fmt.Println("Length of x", x.Len())
	fmt.Println("Length of y", y.Len())

	x.Add(101)
	fmt.Println("Length of x", x.Len())
	fmt.Println(x.String())

	x.Remove(9)
	fmt.Println("Length of x", x.Len())
	fmt.Println(x.String())

	z := x.Copy()
	fmt.Println("z:", z.String())
	fmt.Println("length of z:", z.Len())

	y.Clear()
	fmt.Println("y:", y.String())
	fmt.Println("length of y:", y.Len())
	y.AddAll(1, 2, 3)
	fmt.Println("y:", y.String())
	fmt.Println("length of y:", y.Len())
}
