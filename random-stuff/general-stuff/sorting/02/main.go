package main

import (
	"fmt"
	"sort"
	"strings"
)

var (
	files = []string{"utils.go", "Makefile", "misc.go", "main.go", "Testing.conf", "metadata.xml"}
)

// FoldedStrings for a list of strings
//
type FoldedStrings []string

// SortFoldedStrings function performs case insensitive sort on list of strings
func SortFoldedStrings(slice []string) {
	sort.Sort(FoldedStrings(slice))
}

func (slice FoldedStrings) Len() int { return len(slice) }

func (slice FoldedStrings) Less(i, j int) bool {
	return strings.ToLower(slice[i]) < strings.ToLower(slice[j])
}

func (slice FoldedStrings) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	fmt.Println("* custom sorting 02 *")
	fmt.Printf("Unsorted: %v\n", files)

	SortFoldedStrings(files)

	fmt.Printf("Sorted: %v\n", files)
}
