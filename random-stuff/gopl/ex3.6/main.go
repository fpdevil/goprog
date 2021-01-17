// PLotting a Mandelbrot set with supersampling
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	iterations = 255              // maximum number of iterations
	contrast   = 75               // color contrast value
	output     = "mandelbrot.png" // output file name

	xmin = -2 // min x coordinate
	xmax = +2 // max x coordinate
	ymin = -2 // min y coordinate
	ymax = +2 // max y coordinate

	width  = 1024 // canvas width
	height = 1024 // canvas height

	// Create pixels for an image with twice the width and height
	// for more resolution
	twidth  = 2 * width
	theight = 2 * height
)

func main() {
	// calculate the amount of increment in each x and y intervals
	// for each pixel
	dx := (xmax - xmin) / float64(twidth)
	dy := (ymax - ymin) / float64(theight)

	// create a slice for holding the pixels to be distributed
	// over the image rectangle
	var pixels [twidth][theight]color.Color

	// get evenly spaced pixels or numbers on x over a specified
	// interval of width
	for px := 0; px < twidth; px++ {
		x := xmin + float64(px)*dx
		// Convert pixel coordinate x, y to complex number
		// Image point (px, py) represents complex value z
		for py := 0; py < theight; py++ {
			y := ymin + float64(py)*dy
			z := complex(x, y)

			pixels[px][py] = mandelbrot(z)
		}
	}

	// create the image canvas
	background := image.Rect(0, 0, width, height)
	img := image.NewRGBA(background)

	// now that we have all color pixels, we can apply supersampling
	// and set over the image canvas
	for px := 0; px < width; px++ {
		for py := 0; py < height; py++ {
			// scale factor of 2 for each pixel
			si, sj := 2*px, 2*py

			// now we can apply the average colors
			r1, g1, b1, a1 := pixels[si][sj].RGBA()
			r2, g2, b2, a2 := pixels[si+1][sj].RGBA()
			r3, g3, b3, a3 := pixels[si][sj+1].RGBA()
			r4, g4, b4, a4 := pixels[si+1][sj+1].RGBA()

			r := (r1 + r2 + r3 + r4) / 4
			g := (g1 + g2 + g3 + g4) / 4
			b := (b1 + b2 + b3 + b4) / 4
			a := (a1 + a2 + a3 + a4) / 4

			// now convert values from uint32 to uint8 by performing right shift operation
			// simply taking the 8 most significant bits of the 16 bit color
			// ur := r >> 24 & 0xff
			// ug := g >> 16 & 0xff
			// ub := b >> 8 & 0xff
			// ua := a >> 8 & 0xff
			ur := r >> 8
			ug := g >> 8
			ub := b >> 8
			ua := a >> 8

			c := color.RGBA{uint8(ur), uint8(ug), uint8(ub), uint8(ua)}

			img.Set(px, py, c)
		}
	}

	f, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to create file %s: %s", output, err.Error())
		return
	}

	if err := png.Encode(f, img); err != nil {
		fmt.Fprintf(os.Stderr, "unable to create png %s: %s", output, err.Error())
	}

	defer f.Close()
}

// step function calculats the amount of increment in each
// x and y intervals for each pixel
func steps(x1, x2, y1, y2, w, h int) (float64, float64) {
	dx := (x2 - x1) / w
	dy := (y2 - y1) / h
	return float64(dx), float64(dy)
}

func acosrgba(c complex128) color.Color {
	val := cmplx.Acos(c)
	b := uint8(real(val)*128) + 127
	r := uint8(imag(val)*128) + 127
	// return color.RGBA{r, uint8(math.Abs(float64(r - b))), b, 255}
	return color.YCbCr{192, b, r}
}

// getColors function returns a set of colors for plotting
// the mandelbrot. the list of colors are from below SO
// https://stackoverflow.com/questions/16500656/which-color-gradient-is-used-to-color-mandelbrot-in-wikipedia
func getColors() []color.RGBA {
	var palettes = []color.RGBA{
		color.RGBA{66, 30, 15, 255},
		color.RGBA{25, 7, 26, 255},
		color.RGBA{9, 1, 47, 255},
		color.RGBA{4, 4, 73, 255},
		color.RGBA{0, 7, 100, 255},
		color.RGBA{12, 44, 138, 255},
		color.RGBA{24, 82, 177, 255},
		color.RGBA{57, 125, 209, 255},
		color.RGBA{134, 181, 229, 255},
		color.RGBA{211, 236, 248, 255},
		color.RGBA{241, 233, 191, 255},
		color.RGBA{248, 201, 95, 255},
		color.RGBA{255, 170, 0, 255},
		color.RGBA{204, 128, 0, 255},
		color.RGBA{153, 87, 0, 255},
		color.RGBA{106, 52, 3, 255},
	}

	return palettes
}

func mandelbrot(c complex128) color.Color {
	var z complex128
	var p = getColors()
	var lp = len(p)

	for n := uint8(0); n < iterations; n++ {
		z = z*z + c
		// any sequence which generates a term greater than 2 units
		// from the origin will eventually escape the plot and will
		// be colored accordingly, else its just black
		if cmplx.Abs(z) > 2 {
			// r := 255 - contrast*n%255
			// g := 127 - contrast*n%127
			// b := 64 - contrast*n%64
			// return color.RGBA{r, g, b, 255}
			// return acosrgba(z)
			return p[n%uint8(lp)]
		}
	}

	return color.Black
}
