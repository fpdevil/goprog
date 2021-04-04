package main

import "fmt"

func main() {
	fmt.Println("*Hofstadter Sequence*")
	females := make([]int, 20)
	males := make([]int, len(females))
	for n := range females {
		females[n] = hofstadterFemale(n)
		males[n] = hofstadterMale(n)
	}
	fmt.Printf("Hofstadter Males: %#v\n", males)
	fmt.Printf("Hofstadter Females: %#v\n", females)
}

func hofstadterFemale(n int) int {
	if n <= 0 {
		return 1
	}
	return n - hofstadterMale(hofstadterFemale(n-1))
}

func hofstadterMale(n int) int {
	if n <= 0 {
		return 0
	}
	return n - hofstadterFemale(hofstadterMale(n-1))
}
