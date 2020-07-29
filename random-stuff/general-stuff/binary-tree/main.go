package main

import (
	"fmt"
	"io"
	"os"
)

/*
Wikipedia definition:
A tree  is a  data structure made  up of nodes  or vertices  and edges
without having any cycle. The tree with no nodes is called the null or
empty tree.  A tree  that is  not empty  consists of  a root  node and
potentially many levels of additional nodes that form a hierarchy.
*/

const (
	msg = ` * Binary Tree Implementation *

 `
)

// Node represents a single node of a tree
type Node struct {
	left  *Node // left child node
	right *Node // right child node
	value int   // value at the node
}

// Tree represents a Binary Tree structure
type Tree struct {
	root *Node
	size int
}

// NewTree creates a new Binary Tree
func NewTree() *Tree {
	tree := new(Tree)
	tree.size = 0
	return tree
}

// Size function returns the size of tree
func (tree *Tree) Size() int {
	return tree.size
}

// Root function returns a node at the root
func (tree *Tree) Root() *Node {
	return tree.root
}

// Insert function inserts a value into the Binary Tree
func (tree *Tree) Insert(value int) {
	if tree.root == nil {
		tree.root = &Node{nil, nil, value}
	}
	tree.size++
	tree.root.insert(&Node{nil, nil, value})
}

func (root *Node) insert(newNode *Node) {
	if newNode.value > root.value {
		if root.right == nil {
			root.right = newNode
		} else {
			root.right.insert(newNode)
		}
	} else {
		if root.left == nil {
			root.left = newNode
		} else {
			root.left.insert(newNode)
		}
	}
}

// Min function finds the miniumum value of the tree keeping
// the fact that always left part of the tree contains lesser
// valued elements
func (root *Node) Min() int {
	if root.left == nil {
		return root.value
	}
	return root.left.Min()
}

func print(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.value)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func stringify(node *Node, size int) {
	if node != nil {
		f := ""
		for i := 0; i < size; i++ {
			f += "	"
		}
		f += "*>{"
		size++
		stringify(node.left, size)
		fmt.Printf("%s%d}\n", f, node.value)
		stringify(node.right, size)
	}
}

func main() {
	fmt.Println(msg)

	e := &Node{nil, nil, 99}
	fmt.Fprintf(os.Stdout, "an empty node: %#v\n", e)

	data := []int{2, 5, 7, -1, -1, 5, 5}

	t := NewTree()
	for _, v := range data {
		t.Insert(v)
	}

	fmt.Printf("size: %d\n", t.Size())
	print(os.Stdout, t.root, 0, 'C')
}
