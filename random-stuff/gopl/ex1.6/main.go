package main

/*

Lissajous figure or Bowditch curve, is the graph of a system of parametric equations
which describe complex harmonic motion.

x = Asin(at + δ),
y = Bsin(bt)

where A & B are frquencies along X & Y directions
t is a parameter representing time passed in seconds
δ is inserted to account for the change in phases of the two vibrations

*/

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// RGBColor struct for generating random colors
type RGBColor struct {
	Red   int
	Blue  int
	Green int
	Alpha int
}

func main() {
	fmt.Println("Lissajous Figure Generation")
	fmt.Println()

	// random seed generation
	t := time.Now().UTC().UnixNano()
	rand.Seed(t)

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}

		http.HandleFunc("/", handler)
		if err := http.ListenAndServe("localhost:8000", nil); err != nil {
			fmt.Fprintf(os.Stderr, "error starting server")
			return
		}
	}

	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas coverage size [-size .. +size]
		cycles  = 5     // number of complete x oscillator revolutions
		nframes = 64    // number of animated frames
		delay   = 8     // delay which maps to 80ms between frames
	)

	// create a palette for random colors
	palette := make([]color.Color, 0, nframes)
	palette = append(palette, color.RGBA{0, 0, 0, 255})
	for c := 0; c < nframes; c++ {
		p := color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
		palette = append(palette, p)
	}

	// rellative frequency of the y oscillator
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // provide a startup phase difference

	// the outer loop runs for 64 iterations each producing a single
	// frame of animation; creates a 201 X 201 image with 2 color palette
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // return a rectangle (0, 0) to (2x+1, 2y+1)
		img := image.NewPaletted(rect, palette)      // get new Paletted image
		// inner loop runs 2 Oscillators, X oscillator as just a Sine function
		// and Y oscillator a Sine with frequency relative to X oscillator as a
		// random number between 0 and 3, phase relative to X oscillator
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8((i%(len(palette)-1))+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

// getHex parses a provided decimal integer and returns its Hex equivalent
func getHex(d int) string {
	hex := fmt.Sprintf("%x", d)
	if len(hex) == 1 {
		// prepend with 0 to represent as Hex string
		hex = "0" + hex
	}
	return hex
}

// GetRandomRGBColor generates a random RGB Color struct
func GetRandomRGBColor() RGBColor {
	t := time.Now().UTC().UnixNano()
	rand.Seed(t)
	R := rand.Intn(255)
	G := rand.Intn(255)
	B := rand.Intn(255)

	return RGBColor{R, G, B, 255}
}

// GetRandomHexColor a random color as Hex string
func GetRandomHexColor() string {
	color := GetRandomRGBColor()
	hex := "#" + getHex(color.Red) + getHex(color.Green) + getHex(color.Blue) + getHex(color.Alpha)
	return hex
}
