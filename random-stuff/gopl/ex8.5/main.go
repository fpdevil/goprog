// Exercise 8.5 Fractals - Mandelbrot set plotting
package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	xmin   = -2   // min x coordinate
	xmax   = +2   // max x coordinate
	ymin   = -2   // min y coordinate
	ymax   = +2   // max y coordinate
	width  = 1024 // canvas width
	height = 1024 // canvas height

	iterations = 255              // maximum number of iterations
	contrast   = 75               // color contrast value
	output     = "mandelbrot.png" // output file name
)

func main() {
	for num := 1; num <= runtime.NumCPU(); num++ {
		img := image.NewRGBA(image.Rect(0, 0, width, height))

		runtime.GOMAXPROCS(num)
		wg := sync.WaitGroup{}
		start := time.Now()

		yranges := make([][]int, num)
		for i := range yranges {
			yranges[i] = make([]int, 0)
		}
		// Divide height by number of goroutine.
		for y := 0; y < height; y++ {
			i := y % num
			yranges[i] = append(yranges[i], y)
		}
		for i := 0; i < num; i++ {
			wg.Add(1)
			go func(yrange []int) {
				for _, py := range yrange {
					y := float64(py)/height*(ymax-ymin) + ymin
					for px := 0; px < width; px++ {
						x := float64(px)/width*(xmax-xmin) + xmin
						z := complex(x, y)
						// Image point (px, py) represents complex value z.
						img.Set(px, py, mandelbrot(z))
					}
				}
				wg.Done()
			}(yranges[i])
		}
		wg.Wait()
		// png.Encode(os.Stdout, img) // NOTE: ignoring errors
		// for testing in go playground comment below and enable
		// the function displayImage(img)
		f, err := os.Create(output)
		if err != nil {
			fmt.Printf("error occurred creating image %v", err.Error())
			return
		}

		defer f.Close()

		if err := png.Encode(f, img); err != nil {
			fmt.Printf("error encoding the image %v", err.Error())
			return
		}

		log.Printf("\nMAX CPU NUM %d\nUSE CPU NUM %d\nCalculated time : %v\n", runtime.NumCPU(), num, time.Since(start))
	}
}

// linspace function is for generating n evenly spaced points in the interval [a, b]
// diff = (b - a) / (n - 1)
// for i in 1 to n:
//     a + i * diff
func linspace(start, end float64, num int) []float64 {
	result := make([]float64, num)
	step := (end - start) / float64(num-1)
	for i := range result {
		result[i] = start + float64(i)*step
	}
	return result
}

func mandelbrot(c complex128) color.Color {
	var z complex128
	for n := uint8(0); n < iterations; n++ {
		z = z*z + c
		// any sequence which generates a term greater than 2 units
		// from the origin will eventually escape the plot and will
		// be colored accordingly, else its just black
		if cmplx.Abs(z) > 2 {
			r := 255 - contrast*n%255
			g := 127 - contrast*n%127
			b := 64 - contrast*n%64
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.Black
}

// displayImage renders an image to the playground's console by
// base64-encoding the encoded image and printing it to stdout
// with the prefix "IMAGE:".
func displayImage(img image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		panic(err)
	}
	fmt.Println("IMAGE:" + base64.StdEncoding.EncodeToString(buf.Bytes()))
}
