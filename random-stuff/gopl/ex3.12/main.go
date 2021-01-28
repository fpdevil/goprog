package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage %s <string1> <string1>\n", filepath.Base(os.Args[0]))
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]
	fmt.Printf("strings %s and %s are anagrams: %t\n", s1, s2, isAnagram(s1, s2))
}

func fillMap(s string) map[rune]int {
	fill := make(map[rune]int)
	for _, v := range s {
		fill[v]++
	}

	// fmt.Printf("fill map: %v\n", fill)
	return fill
}

func areEqual(x, y map[rune]int) bool {
	for k, v := range x {
		if y[k] != v {
			return false
		}
	}

	for k, v := range y {
		if x[k] != v {
			return false
		}
	}

	return true
}

func isAnagram(s1, s2 string) bool {
	m1 := fillMap(s1)
	m2 := fillMap(s2)
	return areEqual(m1, m2)
}
