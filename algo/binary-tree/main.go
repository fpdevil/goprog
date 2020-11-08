package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	msg = `
	/* Binary Tree */
	`
)

var seed = time.Now().UnixNano()

// Tree represents an int binary tree
type Tree struct {
	Left  *Tree
	Node  int
	Right *Tree
}

func main() {
	fmt.Println(msg)
	tree := create(10)
	fmt.Printf("Root of Tree: %v\n", tree.Node)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -5)
	tree = insert(tree, -3)
	traverse(tree)
	fmt.Println()
	fmt.Printf("Root of Tree: %v\n", tree.Node)
	tree = insert(tree, 9)
	traverse(tree)
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(seed)
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Node {
		return t
	}
	if v < t.Node {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Printf(" <%v> ", t.Node)
	traverse(t.Right)
}
