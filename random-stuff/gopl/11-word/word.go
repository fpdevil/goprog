// Example from Chapter 11 on Testing
//!+

// Package word provides utilities for word games
package word

import "unicode"

// IsPalindrome function reports whether the provided word is a
// palindrome or not
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
