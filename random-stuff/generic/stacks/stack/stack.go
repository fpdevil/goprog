package stack

import "errors"

// Stack is a custom type of interface slice
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
func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

// Top method returns an item at the top of the stack
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("unable to Top() from an empty stack")
	}
	return stack[len(stack)-1], nil
}

// Pop method pops out an element from te top of stack
func (stack *Stack) Pop() (interface{}, error) {
	s := *stack
	if len(s) == 0 {
		return nil, errors.New("unable to Pop() from an empty stack")
	}
	top := s[len(s)-1]
	*stack = s[:len(s)-1]
	return top, nil
}
