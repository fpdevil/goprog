package main

import (
	"fmt"
	"strings"
)

const (
	msg = `
	/* Hash Table */
	`
	// SIZE represents number of buckets of hash table
	SIZE = 15
)

// Node of the hash table which is a linked list
type Node struct {
	Value int
	Next  *Node
}

// HashTable represents a hash table with a table of specific size
type HashTable struct {
	Table map[int]*Node
	Size  int
}

func main() {
	fmt.Println(msg)
	fmt.Printf("%s\n", strings.Repeat("-", 60))
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Printf("number of spaces: %d\n", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	traverse(hash)

	fmt.Printf("%s\n", strings.Repeat("-", 60))
	fmt.Println("performing lookup...")
	for i := 10; i < 125; i++ {
		if !lookup(hash, i) {
			fmt.Printf("%d not present in the hash table\n", i)
		}
	}
}

func hashFunction(i, size int) int {
	return (i % size)
}

func insert(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := Node{
		Value: value,
		Next:  hash.Table[index],
	}
	hash.Table[index] = &element
	return index
}

func traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%3d => ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

func lookup(hash *HashTable, value int) bool {
	index := hashFunction(value, hash.Size)
	var present bool
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Value == value {
				present = true
			}
			t = t.Next
		}
	}
	return present
}
