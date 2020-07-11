package main

/*
Conways  Game  of  Life   is  a  simulation  of  underpopulation,
overpopulation, and reproduction. The simulation is played out on
a two-dimensional grid of cells.  As such, this challenge focuses
on slices.

Each cell has  eight adjacent cells in  the horizontal, vertical,
and diagonal  directions. In each  generation, cells live  or die
based on the number of living neighbors.

The rules for Conway’s Game of Life:

1. A live cell with less than two live neighbors dies.

2. A live cell  with two or three live neighbors  lives on to the
   next  generation.

3. A  live  cell with  more  than three  live neighbors dies.

4. A dead  cell with exactly three live neighbors  becomes a live
   cell.

Rule 1 represents `death  by under-population`; Rule 2 represents
`sustainable life`; Rule 3 represents `death by over-population`,
and  Rule 4  represents  `birth` or  `reproduction`. The  initial
state  of the  game  is  the `seed`  and  all  cells are  updated
simultaneously. Time steps are sometimes called `generations`.
*/

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

// Universe is a two-dimensional field of cells.
type Universe [][]bool

// NewUniverse would return a blank universe
func NewUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

// Seed method will randomly set approximately 25% of the
// cells to alive (true)
func (u Universe) Seed() {
	for i := 0; i < (width*height)/4; i++ {
		u.Place(rand.Intn(width), rand.Intn(height), true)
	}
}

// Place the state of the appropriate cell under context
func (u Universe) Place(x, y int, z bool) {
	u[y][x] = z
}

// Alive method checks whether the specified cell is alive.
// If the coordinates are outside of the universe, then they wrap
// around using the modulo (%) operator.
// So the grid size of 80 X 15 cannot be crossed
func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

// Neighbors method counts the number of live neighbors for a
// given cell, from 0 to 8.
func (u Universe) Neighbors(x, y int) int {
	// we will count the adjacent cells excluding the current
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) && u.Alive(x+i, y+j) {
				count++
			}
		}
	}
	return count
}

// Next method determines whether a cell has two, three,
// or more neighbors, It returns the state of the specified
// cell at the next step
func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 2 || n == 3 && u.Alive(x, y)
}

// String would return the universe as a string
func (u Universe) String() string {
	var b byte
	buf := make([]byte, 0, (width+1)*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b = ' '
			if u[y][x] {
				b = '*'
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}

// Step updates the state of the next universe (b) from
// the current universe (a).
func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Place(x, y, a.Next(x, y))
		}
	}
}

// Show clears the screen and displays the universe.
func (u Universe) Show() {
	fmt.Print("\x0c", u.String())
}

func main() {
	fmt.Println("---- Conways Game Of Life ----")

	a, b := NewUniverse(), NewUniverse()
	a.Seed()

	for i := 0; i < 300; i++ {
		// Clear all the characters on the screen
		// screen.Clear()
		a.Show()
		time.Sleep(time.Second / 30)
		a, b = b, a
	}
}
