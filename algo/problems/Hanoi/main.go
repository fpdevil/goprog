package main

import (
	"fmt"
	"os"
)

func main() {
	pegs := 2
	fmt.Printf("* Towers Of Hanoi with %d pegs *\n", pegs)
	TowersOfHanoi(pegs)
}

// TowersOfHanoi function invokes the hanoi function with the number of
// pegs submitted as an argument and using the the src, dst and aux as
// A, C & B
func TowersOfHanoi(pegs int) {
	hanoi(pegs, "A", "C", "B")
}

func hanoi(pegs int, src, dst, aux string) {
	if pegs == 1 {
		fmt.Fprintf(os.Stdout, "Move Peg %d from %s to %s\n", pegs, src, dst)
		return
	}

	// move the top n-1 pegs from src to aux via dst
	hanoi(pegs-1, src, aux, dst)

	// move the final peg from src to dst
	fmt.Fprintf(os.Stdout, "Move Peg %d from %s to %s\n", pegs, src, dst)

	// move the n-1 disks from aux to dst via src
	hanoi(pegs-1, aux, dst, src)
}
