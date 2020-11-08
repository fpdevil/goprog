package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

// elements of the Periodic table
// https://en.wikipedia.org/wiki/Periodic_table
var elements = []string{
	"Nonmetals",
	"    Noble Gases",
	"        Helium",
	"        Neon",
	"        Argon",
	"        Krypton",
	"        Xenon",
	"    Reactive",
	"        Hydrogen",
	"        Carbon",
	"        Nitrogen",
	"        Oxygen",
	"        Bromine",
	"Metalloids",
	"    Silicon",
	"    Germanium",
	"    Boron",
	"    Tellurium",
	"Metals",
	"    Alkali Metals",
	"        Lithium",
	"        Sodium",
	"        Potassium",
	"        Rubedium",
	"    Alkaline Earth Metals",
	"        Berillium",
	"        Magnesium",
	"        Calcium",
	"        Strontium",
	"    Lanthanides",
	"        Lanthanum",
	"        Neodymium",
	"        Eurpium",
	"        Disprosium",
	"    Actanides",
	"        Actinium",
	"        Neptunium",
	"        Plutonium",
	"        Californium",
}

var original = []string{
	"Nonmetals",
	"    Hydrogen",
	"    Carbon",
	"    Nitrogen",
	"    Oxygen",
	"Inner Transitionals",
	"    Lanthanides",
	"        Europium",
	"        Cerium",
	"    Actinides",
	"        Uranium",
	"        Plutonium",
	"        Curium",
	"Alkali Metals",
	"    Lithium",
	"    Sodium",
	"    Potassium",
}

// Entry represents the structure of elements with key value pairs
// and the sub values represented as children
type Entry struct {
	key      string
	value    string
	children Entries
}

// Entries represents a list of type Entry
type Entries []Entry

// Helpers for trivially sorting the Entry types
// Len is the number of elements in the collection.
func (e Entries) Len() int {
	return len(e)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (e Entries) Less(i, j int) bool {
	return e[i].key < e[j].key
}

// Swap swaps the elements with indexes i and j.
func (e Entries) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func main() {
	fmt.Println("Elements of the Periodic Table")
	printElements(elements)
}

func printElements(slice []string) {
	const format = "%v\t%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer)
	// minwidth, tabwidth, padding, padchar, flags
	tw.Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "|", "---------------------------", "|", "---------------------------", "|")
	fmt.Fprintf(tw, format, "|", "         Original          ", "|", "           Sorted          ", "|")
	fmt.Fprintf(tw, format, "|", "---------------------------", "|", "---------------------------", "|")
	elems := SortIndentedStrings(slice)
	for e := range elems {
		fmt.Fprintf(tw, format, "|", slice[e], "|", elems[e], "|")
	}
	tw.Flush()
}

// SortIndentedStrings function  takes a string slice  with items at
// different levels of indent that are used to indicate parent-child
// relationships, sort the items case insensitively with child items
// sorted underneath their  parent items and so  on recursively. The
// indentation must  be either one or  more spaces or may  be one or
// more tabs.
func SortIndentedStrings(slice []string) []string {
	entries := fillEntries(slice)
	return sortedEntries(entries)
}

// addEntry function, based on the indent level of the key
// adds the key/value to the entries either directly as a child
// or as a child of another Entry
// if level = 0 it is a top level entry under *Entries
// if level > 0 it is added as a child under the preceding entry
func addEntry(level int, key, value string, entries *Entries) {
	if level == 0 {
		*entries = append(*entries, Entry{key, value, make(Entries, 0)})
	} else {
		xEntries := *entries
		lastEntry := &xEntries[xEntries.Len()-1]
		addEntry(level-1, key, value, &lastEntry.children)
	}
}

func sortedEntries(entries Entries) []string {
	var indentedSlice []string // indented strings are children of parents
	sort.Sort(entries)
	for _, e := range entries {
		fillIndentedStrings(e, &indentedSlice)
	}
	return indentedSlice
}

func fillIndentedStrings(entry Entry, indentedSlice *[]string) {
	*indentedSlice = append(*indentedSlice, entry.value)
	sort.Sort(entry.children)
	for _, child := range entry.children {
		fillIndentedStrings(child, indentedSlice)
	}
}

func getIndent(slice []string) (string, int) {
	for _, item := range slice {
		// check if the first element is a space ir a tab
		if len(item) > 0 && (item[0] == ' ' || item[0] == '\t') {
			whitespace := rune(item[0])
			for i, char := range item[1:] {
				if char != whitespace {
					i++
					return strings.Repeat(string(whitespace), i), i
				}
			}
		}
	}
	return "", 0
}

func fillEntries(slice []string) Entries {
	indent, indentSize := getIndent(slice)
	// fmt.Printf("* [%s] %d = %d *\n", indent, len(indent), indentSize)
	entries := make(Entries, 0)
	for _, item := range slice {
		i, level := 0, 0
		for strings.HasPrefix(item[i:], indent) {
			i += indentSize
			level++
		}
		key := strings.ToLower(strings.TrimSpace(item))
		addEntry(level, key, item, &entries)
	}
	return entries
}
