// A  Mandelbrot set  intersects with  X-Axis at  (-2.5, +0.25)  and with
// Y-Axis at (-1, +1) on a Cartesian plane. The scaled values of x become
// real part  of the mandelbrot  equation's c and  the scaled value  of y
// becomes the imaginary part of c.
//
// For representing the  image of mandelbrot, the  complex coordinates of
// the mandelbrot set need to be mapped somehow on to the 2D pixel image.
//
// ref: https://lodev.org/cgtutor/index.html
//

package main

import (
	"errors"
	"fmt"
)

// Conf struct defines the boundaries of the plot with the
// complex coordinates
type Conf struct {
	width  int
	height int
	minRe  float64
	maxRe  float64
	minIm  float64
	maxIm  float64
}

func (c Conf) ToReal(x int) (float64, error) {
	if x >= c.width || x < 0 {
		return 0, errors.New("X out of bounds")
	}
	size := ((c.maxRe - c.minRe) / float64(c.width-1))
	return c.minRe + float64(x)*size, nil
}

func main() {
	fmt.Println("vim-go")
}
