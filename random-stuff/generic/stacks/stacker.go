package main

import (
	"fmt"
	"strings"

	"github.com/fpdevil/goprog/random-stuff/generic/stacks/stack"
)

func main() {
	fmt.Println("** STACK Data Structure **")
	var hayStack stack.Stack
	hayStack.Push("Hay")
	hayStack.Push("The Stack!")
	hayStack.Push(-99)
	hayStack.Push([]string{"mickey", "mouse", "donald", "duck"})
	hayStack.Push(102.80)
	for {
		item, err := hayStack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}

	var pStack stack.Stack
	pStack.Push("Rustic Art")
	pStack.Push("Rendezvous")
	pStack.Push(99.999999999999)
	pStack.Push(007)
	x, _ := pStack.Top()
	fmt.Println(x)
	pStack.Push(-6e-4)
	pStack.Push("Baker")
	pStack.Push(-3)
	pStack.Push("Cake")
	pStack.Push("Dancer")
	x, _ = pStack.Top()
	fmt.Println(x)
	pStack.Push(11.7)
	fmt.Println("stack is empty", pStack.IsEmpty())
	fmt.Printf("Len() == %d  Cap == %d\n", pStack.Len(), pStack.Cap())
	difference := pStack.Cap() - pStack.Len()
	for i := 0; i < difference; i++ {
		pStack.Push(strings.Repeat("*", difference-i))
	}
	fmt.Printf("Len() == %d  Cap == %d\n", pStack.Len(), pStack.Cap())
	for pStack.Len() > 0 {
		x, _ = pStack.Pop()
		fmt.Printf("%T %v\n", x, x)
	}
	fmt.Println("stack is empty", pStack.IsEmpty())
	x, err := pStack.Pop()
	fmt.Println(x, err)
	x, err = pStack.Top()
	fmt.Println(x, err)
}
