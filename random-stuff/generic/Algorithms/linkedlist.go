// Package linkedlist defined the linkedlist datastructure that
// contains a series of structures called nodes. Each node consists
// of a record that contains 2 parts.
// The first part contains the data and the second part contains
// a pointer to a node.
// Basic operations over a linked list are
// 1. Traverse list
// 2. Insert item in list
// 3. Delete item from list
package linkedlist

import (
	"fmt"
	"os"
)

// Node contains data as well as a pointer to the next node.
type Node struct {
	Data interface{} // the data within linked list
	Next *Node       // pointer to the next node
}

// LinkedList contains a pointer to the first node in the list and
// its size.
type LinkedList struct {
	head *Node // memory address of first node in linked list
	size int   // length of the linked list
}

//!+Show

// Show method is for printing the list data with additional information
func (ll *LinkedList) Show() error {
	if ll.head == nil {
		return fmt.Errorf("empty list")
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%v =>", current.Data)
		current = current.Next
	}
	fmt.Println()
	return nil
}

//!-Show

//!+Length

// Length method takes a linked list and returns the count of number of
// nodes within the list. Time Complexity: O(n)
func (ll *LinkedList) Length() int {
	// return ll.size
	var size int
	current := ll.head
	for current != nil {
		size++
		current = current.Next
	}
	return size
}

//!-Lenght

//!+InsertAtStart

// InsertAtStart inserts an item at the beginning of linked list
// Time Complexity: O(1), Space Complexity: O(1)
func (ll *LinkedList) InsertAtStart(data interface{}) {
	node := &Node{
		Data: data,
	}
	if ll.head == nil {
		ll.head = node
	} else {
		node.Next = ll.head
		ll.head = node
	}
	ll.size++
}

//!-InsertAtStart

//!+InsertAtEnd

// InsertAtEnd inserts an item at the end of linked list
// Time Complexity: O(n), Space Complexity: O(1)
func (ll *LinkedList) InsertAtEnd(data interface{}) {
	node := &Node{
		Data: data,
	}
	if ll.head == nil {
		ll.head = node
	} else {
		current := ll.head
		for current.Next != nil {
			// get the last node of list
			current = current.Next
		}
		current.Next = node
	}
}

//!-InsertAtEnd

//!+InsertAtPos

// InsertAtPos inserts a node at a given position
// Time Complexity: O(n), Space Complexity: O(1)
func (ll *LinkedList) InsertAtPos(data interface{}, position int) {
	if position < 1 || position > ll.size+1 {
		fmt.Fprintf(os.Stderr, "invalid position: index out of bounds")
		return
	}
	newNode := &Node{
		Data: nil,
	}
	var prev, current *Node
	prev = nil
	current = ll.head
	for position > 1 {
		prev = current
		current = current.Next
		position = position - 1
	}

	if prev != nil {
		prev.Next = newNode
		newNode.Next = current
	} else {
		newNode.Next = current
		ll.head = newNode
	}
	ll.size++
}

//!-InsertAtPos

//!+DeleteFirst

// DeleteFirst method deletes first node or the node at the beginning
// of the linked list
// Time Complexity: O(n), Space Complexity: O(1)
func (ll *LinkedList) DeleteFirst() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("delete at first: empty list")
	}
	data := ll.head.Data
	ll.head = ll.head.Next
	ll.size--
	return data, nil
}

//!-DeleteFirst

//!+DeleteLast

// DeleteLast method deletes a node at the end of the linked list
// Time Complexity: O(n), Space Complexity: O(1)
func (ll *LinkedList) DeleteLast() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("delete at last: empty list")
	}
	var prev *Node
	current := ll.head
	for current.Next != nil {
		prev = current
		current = current.Next
	}
	if prev != nil {
		prev.Next = nil
	} else {
		ll.head = nil
	}
	ll.size--
	return current.Data, nil
}

//!-DeleteLast

//!+DeleteAtPos

// DeleteAtPos deletes a node at a apecific position
func (ll *LinkedList) DeleteAtPos(position int) (interface{}, error) {
	if position < 1 || position > ll.size+1 {
		return nil, fmt.Errorf("insert at position: index out of bounds")
	}

	var prev, current *Node
	prev = nil
	current = ll.head
	pos := 0

	if position == 1 {
		ll.head = ll.head.Next
	} else {
		for pos != position-1 {
			pos++
			prev = current
			current = current.Next
		}
		if current != nil {
			prev.Next = current.Next
		}
	}
	ll.size--
	return current.Data, nil
}

//!-DeleteAtPos

var (
	root = new(Node) // first element of the list
)

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{v, nil}
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
		t = &Node{v, nil}
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
