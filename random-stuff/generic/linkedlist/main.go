package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

var (
	root = new(Node) // first element of the list
)

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{ v, nil }
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Printf("Node %v already exists\n", v)
		return -1
	}

	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -1
	}

	return addNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}

	fmt.Println()
}

func lookup(t *Node, v int) bool {
	if root == nil {
		t = &Node{ v, nil }
		root = t
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookup(t.Next, v)
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func main() {
	fmt.Println("/// Linked Lists ///")
	fmt.Println()

	fmt.Printf("root: %#v\n", root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, -1)
	traverse(root)
	addNode(root, 11)
	addNode(root, 8)
	addNode(root, 5)
	addNode(root, 16)
	addNode(root, 0)
	addNode(root, 5)
	traverse(root)
	addNode(root, 50)
	traverse(root)

	if lookup(root, 50) {
		fmt.Printf("Node exists!\n")
	} else {
		fmt.Println("No Node exists!")
	}

	if lookup(root, -50) {
		fmt.Printf("Node exists!\n")
	} else {
		fmt.Println("No Node exists!")
	}
}

