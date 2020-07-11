package main

import "fmt"

var d = []int{6, 5, 3, 1, 8, 7, 2, 4}

func main() {
	fmt.Println("/// Bubble Sort ///")
	fmt.Printf("initial: %v\n", d)

	if len(d) <= 1 {
		fmt.Printf("sorted: %v\n", d)
		return
	}

	sort(d)
	fmt.Printf("sorted: %v\n", d)
}

func swap(i, j int, digits []int) {
	temp := digits[i]
	digits[i] = digits[j]
	digits[j] = temp
}

func sort(digits []int) {
	for i := 0; i < len(digits)-1; i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[i] > digits[j] {
				swap(i, j, digits)
			}
		}
	}
}
