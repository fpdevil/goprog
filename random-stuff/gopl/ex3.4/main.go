// Compute the SVG rendering of a 3 dimensional function
package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
)

const (
	cells         = 100                 // number of grid cells (100 X 100)
	width, height = 600, 400            // size of SVG canvas in pixels
	xyrange       = 30.0                // amount of spread for range of axes from -xy to +xy
	xyscale       = width / 2 / xyrange // pixels per x or y unit to decide boundary within x & y
	zscale        = height / 2          // pixels per z unit with arbitrary height 0.5 *height
	angle         = math.Pi / 6         // angle of x and y axes = 30º
)

var (
	sin30 = math.Sin(angle)
	cos30 = math.Cos(angle)

	red  = "#ff0000"
	blue = "#0000ff"
)

func main() {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		svg(w)
	}

	http.HandleFunc("/", f)
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		fmt.Fprintf(os.Stderr, "error starting server")
		return
	}

}

func svg(w io.Writer) {
	sp := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: #87CEFA; fill: #ffffff; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n",
		width, height)
	_, _ = w.Write([]byte(sp))

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			// check for finiteness of the values
			if isNonFinite(ax) || isNonFinite(ay) || isNonFinite(bx) || isNonFinite(by) ||
				isNonFinite(cx) || isNonFinite(cy) || isNonFinite(dx) || isNonFinite(dy) {
				continue
			}

			sp = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)

			_, _ = w.Write([]byte(sp))
		}
	}

	sp = fmt.Sprintf("</svg>\n")
	_, _ = w.Write([]byte(sp))
}

// corner function translates the (i, j) in 2D grid cells to
// 2D isometric projection
func corner(i, j int) (float64, float64) {
	// find the point (x, y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// compute the surface height z
	z := f(x, y)

	// now project the coordinates (x, y, z) isometrically onto
	// 2D SVG canvas over scaled coordinates (sx, sy)

	// if z is added, the function will be printed upside down
	// sx := width/2 + (x-y)*cos30*xyscale
	// sy := height/2 + (x+y)*sin30*xyscale + z*zscale

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

// f is the surface function along z over x and y
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // this is the distance from (0, 0)
	if r == 0 {
		return math.Inf(0)
	}

	return math.Sin(r) / r
}

// check if the value if Non-Finite
func isNonFinite(f float64) bool {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return true
	}

	return false
}
