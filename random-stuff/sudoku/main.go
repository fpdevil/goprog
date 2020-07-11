package main

import (
	"errors"
	"fmt"
	"strings"
)

// define size here
const (
	rows, cols = 9, 9
	empty      = 0
)

// Block is a single square of the sudoku grid
type Block struct {
	digit int8
	start bool
}

// Errors
var (
	ErrBounds     = errors.New("out of bounds")
	ErrDigit      = errors.New("invalid digit")
	ErrInRow      = errors.New("digit already present in the row")
	ErrInCol      = errors.New("digit already present in the column")
	ErrFixedDigit = errors.New("cannot change starting fixed digits")
)

// Grid is a Sudoku Grid
type Grid [rows][cols]Block

// SudokuError is a slice of error to satisfy the error
// interface with a method to join multiple errors together
// to return as a single string
type SudokuError []error

// Error returns one or more errors separated by commas.
// satisfy the error interface implementing Error()
func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error()) // convert errors to string
	}
	return strings.Join(s, ", ")
}

// NewSudoku is the constructor function for preparing thhe
// sudoku puzzle initially
func NewSudoku(digits [rows][cols]int8) *Grid {
	var g Grid
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			d := digits[row][col]
			if d != empty {
				g[row][col].digit = d
				g[row][col].start = true
			}
		}
	}
	return &g
}

// isInBounds function ensures that the row and column
// stays within the boundary of sudoku grid and does
// not cross the bounds
func isInBounds(row, col int) bool {
	if row < 0 || row >= rows {
		return false
	}

	if col < 0 || col >= cols {
		return false
	}
	return true
}

// isValidDigit function checks and accepts a valid digit
func isValidDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

// isFixedDigit checks if the current digit under context
// is the starting digit during sudoko creationg
func (g *Grid) isFixedDigit(row, col int) bool {
	return g[row][col].start
}

// isInRow function checks if a given digit is already
// present on the specified row or not
func (g *Grid) isInRow(row int, digit int8) bool {
	for c := 0; c < cols; c++ {
		if g[row][c].digit == digit {
			return true
		}
	}
	return false
}

// isInCol function checks if a given digit is already
// present on the specified column or not
func (g *Grid) isInCol(col int, digit int8) bool {
	for r := 0; r < rows; r++ {
		if g[r][col].digit == digit {
			return true
		}
	}
	return false
}

// Set method validates and fixes the input digit into
// the sudoku grid at specified row, col
func (g *Grid) Set(row, col int, digit int8) error {
	// combine multiple errors
	var errs SudokuError

	// switch {
	// case !isInBounds(row, col):
	// 	errs = append(errs, ErrBounds)
	// case !isValidDigit(digit):
	// 	errs = append(errs, ErrDigit)
	// case g.isFixedDigit(row, col):
	// 	errs = append(errs, ErrFixedDigit)
	// case g.isInRow(row, digit):
	// 	errs = append(errs, ErrInRow)
	// case g.isInCol(col, digit):
	// 	errs = append(errs, ErrInCol)
	// }
	fmt.Println(isInBounds(row, col))
	fmt.Println(isValidDigit(digit))
	fmt.Println(g.isFixedDigit(row, col))
	fmt.Println(g.isInRow(row, digit))

	if !isInBounds(row, col) {
		errs = append(errs, ErrBounds)
	}

	if !isValidDigit(digit) {
		errs = append(errs, ErrDigit)
	}

	if g.isFixedDigit(row, col) {
		errs = append(errs, ErrFixedDigit)
	}
	if g.isInRow(row, digit) {
		errs = append(errs, ErrInRow)
	}
	if g.isInCol(col, digit) {
		errs = append(errs, ErrInCol)
	}

	// check the errors slice
	if len(errs) != 0 {
		return errs
	}

	g[row][col].digit = digit
	return nil
}

// Clear removes a digit from the grid
func (g *Grid) Clear(row, col int) error {
	var errs SudokuError

	if !isInBounds(row, col) {
		errs = append(errs, ErrBounds)
	}

	if g.isFixedDigit(row, col) {
		errs = append(errs, ErrFixedDigit)
	}

	// check the errors slice
	if len(errs) != 0 {
		return errs
	}

	g[row][col].digit = empty
	return nil
}

// describe is a helper for printing the dynamic value and
// dynamic type during type assertion
func describe(i interface{}) {
	fmt.Printf("Dynamic Type: %T\nDynamic Value: %[1]v\n", i)
}

func main() {
	fmt.Println("==== SUDOKU ====")
	fmt.Println()

	// var g Grid
	// err := g.Set(10, 0, 15)

	s := NewSudoku([rows][cols]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	err := s.Set(1, 1, 4)
	// err := s.Set(10, 0, 15)

	// using Type assertion, we convert the error interface
	// type to its underlying concrete type SudokuError
	if err != nil {
		describe(err) // a debugger
		fmt.Println()
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf(">>> %v\n", e)
			}
		}
		return
	}
}
