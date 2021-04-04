package main

/*

Lissajous  figure or  Bowditch  curve, is  the graph  of  a system  of
parametric equations which describe complex harmonic motion.
These figures are created through a superposition of two perpendicular
sinusoidal waves. One  mirror oscillates in the X  direction while the
other one oscillates in the Y  direction. The laser reflected from the
mirrors traces  patterns which depend  on the relative  frequencies of
the sounds. The below 2 equations represent Lissajous figures.

x = Asin(pt)
y = Bsin(qt + φ)

where  the constants  A  & B  represent  the amplitudes  along  X &  Y
directions The values A and B  determine the scale of the figure, with
the entire graph being contained with in a box of dimension 2A by 2B.

p & q represents the angular frequencies along X & Y directions.
	1 ≤ p ≤ q
	if n = q/p (where n is irrational, the equations may be written as)

	x = Asin(t)
	y = Bsin(nt + φ) ; 0 ≤ φ ≤ π/2p

t is a parameter representing time passed in seconds
	0 ≤ t ≤ 2π

φ is the phase  shift which is inserted to account  for the change in
phases of the two vibrations.
	0 ≤ φ ≤ π/2

@ref: https://mathcurve.com/courbes2d.gb/lissajous/lissajous.shtml
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
	"path/filepath"
	"strings"
	"time"
)

// create a palette which is an array of colors for artboard
// ref: https://www.rapidtables.com/web/color/RGB_Color.html
var palette = []color.Color{color.White, color.Black, color.RGBA{0x99, 0x00, 0x4c, 0xff}}

// select the appropriate color index to pick the colors of the
// figure and use it in SetColorIndex to use same from palette
const (
	whiteIndex = 0 // this is the first color White in palette
	blackIndex = 1 // this is the next color Black in palette
	colorIndex = 2 // this is the last color in palette
)

func main() {
	fmt.Println("Lissajous Figure Generation")
	fmt.Printf("%s\n", strings.Repeat("-", 36))

	args := os.Args
	if len(args) == 1 || args[1] == "-h" || args[1] == "--help" {
		fmt.Fprintf(os.Stdout, "usage %s -option=<web|file>\n", filepath.Base(os.Args[0]))
		fmt.Println(" - if option=web go to browser and check http://localhost:8000")
		fmt.Println(" - if option=file open the output file lissajous.gif in a browser")
		return
	}

	var options string
	if strings.Contains(args[1], "=") {
		options = strings.Split(args[1], "=")[1]
	} else {
		fmt.Fprintf(os.Stderr, "options should be \"=\" delimited string. check input %s\n", args[1])
		fmt.Fprintf(os.Stderr, "run %[1]s -h or %[1]s --help\n", filepath.Base(args[0]))
		fmt.Println("program exited!")
		return
	}

	// random seed generation
	t := time.Now().UTC().UnixNano()
	rand.Seed(t)

	if options == "web" {
		fmt.Println("open browser and check http://localhost:8000")
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}

		http.HandleFunc("/", handler)
		if err := http.ListenAndServe("localhost:8000", nil); err != nil {
			fmt.Fprintf(os.Stderr, "error starting server")
			return
		}
	}

	// since options is now file, dump gif to a file
	outputfile := "lissajous.gif"
	out, err := os.Create(outputfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to create file %s", err.Error())
		return
	}
	defer out.Close()

	lissajous(out)
}

func lissajous(out io.Writer) {
	const (
		res     = 0.001 // angular resolution or step value
		size    = 100   // image canvas coverage size [-size .. +size] 2A x 2B
		cycles  = 5     // number of complete x oscillator revolutions
		nframes = 64    // number of animated frames
		delay   = 8     // delay which maps to 80ms between frames
	)

	// relative frequency of the y oscillator
	freq := rand.Float64() * 3.0

	anim := gif.GIF{}
	anim.LoopCount = nframes

	phase := 0.0 // provide a startup phase difference Φ

	// the outer loop runs for 64 iterations each producing a single
	// frame of animation; creates a 201 X 201 image with 2 color palette
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // return a rectangle (0, 0) to (2x+1, 2y+1)
		img := image.NewPaletted(rect, palette)      // get new Paletted image
		// inner loop runs 2 Oscillators, X oscillator as just a Sine function
		// and Y oscillator a Sine with frequency relative to X oscillator as a
		// random number between 0 and 3, phase relative to X oscillator
		// n = p/q = 1 & A = B
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
