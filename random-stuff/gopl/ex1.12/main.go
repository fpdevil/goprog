package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
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
	fmt.Println("open browser at http://localhost:8000/?cycles=<number>")
	fmt.Printf("%s\n", strings.Repeat("-", 36))

	// random seed generation
	t := time.Now().UTC().UnixNano()
	rand.Seed(t)

	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	lissajous(w)
	// }

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		fmt.Fprintf(os.Stderr, "error starting server")
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	// for k, v := range r.Form {
	// 	log.Println(w, "query parameters %q = %q\n", k, v)
	// }

	var ncycles int
	if c := r.Form.Get("cycles"); c != "" {
		log.Printf("using number of cycles: %v\n", c)
		ncycles, _ = strconv.Atoi(c)
	}

	lissajous(w, float64(ncycles))
}

func lissajous(out io.Writer, ncycles float64) {
	const (
		res     = 0.001 // angular resolution or step value
		size    = 100   // image canvas coverage size [-size .. +size] 2A x 2B
		cycles  = 5     // number of complete x oscillator revolutions
		nframes = 64    // number of animated frames
		delay   = 8     // delay which maps to 80ms between frames
	)

	if ncycles == 0 {
		ncycles = cycles
	}

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
		for t := 0.0; t < ncycles*2*math.Pi; t += res {
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
