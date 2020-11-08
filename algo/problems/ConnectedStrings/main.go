package main

// Find the length of connected cells of 1's in a matrix
// consisting of 0's and 1's

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("* Connected Cells *")
	fmt.Println()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	X, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	Y, _ := strconv.Atoi(scanner.Text())

	matrix := make([][]int, X)

	for i := 0; i < X; i++ {
		matrix[i] = make([]int, X)
		for j := 0; j < Y; j++ {
			scanner.Scan()
			v, _ := strconv.Atoi(scanner.Text())
			matrix[i][j] = v
		}
	}

	maximum := getMaxConnects(matrix, X, Y)
	fmt.Println(maximum)
}

// getConnects function traverses in all the 8 directions in a matrix
// and in each of those directions keep track of maximum region found
func getConnects(matrix [][]int, X, Y, r, c int) int {
	solution := 0
	if r < 0 || c < 0 || r >= X || c >= Y {
		solution = 0
	} else if matrix[r][c] == 1 {
		matrix[r][c] = 0
		solution = 1 +
			getConnects(matrix, X, Y, r-1, c) +
			getConnects(matrix, X, Y, r+1, c) +
			getConnects(matrix, X, Y, r, c-1) +
			getConnects(matrix, X, Y, r, c+1) +
			getConnects(matrix, X, Y, r-1, c-1) +
			getConnects(matrix, X, Y, r+1, c+1) +
			getConnects(matrix, X, Y, r-1, c+1) +
			getConnects(matrix, X, Y, r+1, c-1)
	}
	return solution
}

func getMaxConnects(matrix [][]int, X, Y int) int {
	maximum := 0
	for r := 0; r < X; r++ {
		for c := 0; c < Y; c++ {
			maximum = max(maximum, getConnects(matrix, X, Y, r, c))
		}
	}
	return maximum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
