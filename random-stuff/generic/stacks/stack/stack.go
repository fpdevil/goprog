package stack

import "errors"

// Stack is a custom type built using the interface slice, so that
// it can accept any type of values.
type Stack []interface{}

// Len function returns the lenght of a slice
func (stack Stack) Len() int {
	return len(stack)
}

// Cap method returns the capacity of the stack
func (stack Stack) Cap() int {
	return cap(stack)
}

// IsEmpty method checks if the stack is empty or not
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

// Push method inserts an element at the top of the stack
func (stack *Stack) Push(i interface{}) {
	*stack = append(*stack, i)
}

// Top method returns an item at the top of the stack
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("could not run Top() over empty stack")
	}
	return stack[len(stack)-1], nil
}

// Pop method pops out an element from the top of stack
func (stack *Stack) Pop() (interface{}, error) {
	s := *stack
	if len(s) == 0 {
		return nil, errors.New("could not run Pop() over empty stack")
	}
	x := s[len(s)-1]
	*stack = s[:len(s)-1]
	return x, nil
}
