package main

import "fmt"

func max(vals ...int) int {
	if vals == nil {
		return 0
	}
	x := vals[0]
	for _, val := range vals[1:] {
		if x < val {
			x = val
		}
	}
	return x
}

func min(vals ...int) int {
	if vals == nil {
		return 0
	}
	x := vals[0]
	for _, val := range vals[1:] {
		if x < val {
			continue
		}
		x = val
	}
	return x
}

func main() {
	fmt.Println("max()")
	fmt.Println(max(1, 3, 6, 2, 100, 9, 51, 1, 11))
	fmt.Println(max())
	fmt.Println("min()")
	fmt.Println(min(1, 3, 6, 2, -100, 9, -51, 1, 11))
	fmt.Println(min())
}
