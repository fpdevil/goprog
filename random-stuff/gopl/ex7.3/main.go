package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

// binary tree
type tree struct {
	value       int
	left, right *tree
}

// Sort function does in place sorting
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func (t *tree) String() string {
	values := make([]int, 0)
	values = appendValues(values, t)
	if len(values) == 0 {
		return "[]"
	}

	b := &bytes.Buffer{}
	fmt.Fprintf(b, "[%d", values[0])
	for _, v := range values[1:] {
		fmt.Fprintf(b, " %d", v)
	}
	fmt.Fprintf(b, "]")
	return b.String()
}

func main() {
	data := make([]int, 20)
	for i := range data {
		data[i] = rand.Int() % 20
	}
	var root *tree
	for _, v := range data {
		root = add(root, v)
	}
	fmt.Println(root.String())
}
