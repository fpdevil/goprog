package main

import (
	"fmt"
	"sort"
)

// Palindrome type to represent a string or a byte slice
type Palindrome []byte

// Len is the number of elements in the collection.
func (p Palindrome) Len() int {
	return len(p)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p Palindrome) Less(i, j int) bool {
	return p[i] < p[j]
}

// Swap swaps the elements with indexes i and j.
func (p Palindrome) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func isEqual(i, j int, s sort.Interface) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}

//!+IsPalindrome

// IsPalindrome function checks whether the string value is a
// palindrome or not
func IsPalindrome(s sort.Interface) bool {
	x := s.Len() - 1
	for i, j := 0, x; i < j; i, j = i+1, j-1 {
		if !isEqual(i, j, s) {
			return false
		}
	}
	return true
}

//!-IsPalindrome

func main() {
	fmt.Printf("abracadabra is a palindrome: %v\n", IsPalindrome(Palindrome([]byte("abracadabra"))))
	fmt.Printf("abcdcba is a palindrome: %v\n", IsPalindrome(Palindrome([]byte("abcdcba"))))
	fmt.Printf("abcdecba is a palindrome: %v\n", IsPalindrome(Palindrome([]byte("abcdecba"))))
}
