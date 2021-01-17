// Fractals - Mandelbrot set plotting
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin   = -2   // min x coordinate
	xmax   = +2   // max x coordinate
	ymin   = -2   // min y coordinate
	ymax   = +2   // max y coordinate
	width  = 1024 // canvas width
	height = 1024 // canvas height
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// get evenly spaced pixels or numbers on y over a specified
	// interval of height
	for py := 0; py < height; py++ {
		y := ymin + float64(py)/height*(ymax-ymin)
		// get evenly spaced pixels or numbers on x over a specified
		// interval of width
		for px := 0; px < width; px++ {
			x := xmin + float64(px)/width*(xmax-xmin)
			// get a complex representation of x, y
			// Image point (px, py) represents complex value z
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(c complex128) color.Color {
	const (
		iterations = 200
		contrast   = 15
	)

	var z complex128
	for n := uint8(0); n < iterations; n++ {
		z = z*z + c
		if cmplx.Abs(z) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
