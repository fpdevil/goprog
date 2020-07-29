package main

/*

Lissajous figure or Bowditch curve, is the graph of a system of parametric equations
which describe complex harmonic motion.

x = Asin(at + δ),
y = Bsin(bt)

where A & B are frquencies along X & Y directions
t is a parameter representing time passed in seconds
δ is inserted to account for the change in phases of the two vibrations


Exercise 1.6:

Modify the Lissajous  program to produce images in  multiple colors by
adding more values to palette and then displaying them by changing the
third argument of SetColorIndex in some interesting way.

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

	log "github.com/sirupsen/logrus"
)

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
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}

	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas coverage size [-size .. +size]
		cycles  = 5     // number of complete x oscillator revolutions
		nframes = 64    // number of animated frames
		delay   = 8     // delay which maps to 10ms between frames
	)

	// create a palette for colors
	var palette = make([]color.Color, 0, nframes)
	palette = append(palette, color.RGBA{0, 0, 0, 255})
	for i := 0; i < nframes; i++ {
		ratio := float64(i) / float64(nframes)
		clr := color.RGBA{uint8(200*ratio + 55), uint8(200*ratio + 55), uint8(200*ratio + 55), 255}
		palette = append(palette, clr)
	}

	// rellative frequency of the y oscillator
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	phase += 0.0
	// the outer loop runs for 64 iterations each producing a single
	// frame of animation; creates a 201 X 201 image with 2 color palette
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			// inner loop runs 2 Oscillators, X oscillator as just a Sine function
			// and Y oscillator a Sine with frequency relative to X oscillator as a
			// random number between 0 and 3, phase relative to X oscillator
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
