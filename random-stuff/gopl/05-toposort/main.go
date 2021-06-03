package main

import (
	"fmt"
	"sort"
)

// prereqs maps computer courses to their prerequisites.
// this information forms an acyclic graph
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {
		"ordinary differential equations",
		"partial differential equations",
		"linear algebra",
	},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming", "discrete mathematics using a computer"},
	"formal languages":      {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool) // keep track of visited nodes
	var visitAll func(items []string)

	// define an anonymous function with DFS logic
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
