package main

import "fmt"

func main() {
	fmt.Println("Towers of Hanoi")
	moveTower(3, "A", "B", "C")
}

func moveDisk(src, dst string) {
	fmt.Println("moving disk form", src, "to", dst)
}

func moveTower(height int, src, dst, aux string) {
	if height >= 1 {
		moveTower(height-1, src, aux, dst)
		moveDisk(src, dst)
		moveTower(height-1, aux, dst, src)
	}
}
