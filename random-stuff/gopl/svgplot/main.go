// Compute the SVG rendering of a 3 dimensional function
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 400            // size of SVG canvas in pixels
	cells         = 100                 // number of grid cells (100 X 100)
	xyrange       = 30.0                // amount of spread for range of axes from -xy to +xy
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height / 2          // pixels per z unit with arbitrary height 0.5 *height
	angle         = math.Pi / 6         // angle of x and y axes = 30º
)

const (
	starthtml = `<!DOCTYPE html>
	<html>
	<head>
		<title>%s</title>
	</head>
	<body>
	`
	endhtml = `</body>
	</html>
	`
)

var (
	sin30 = math.Sin(angle)
	cos30 = math.Cos(angle)
)

func main() {
	fmt.Printf(starthtml, "Polygon Mesh")
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: blue; fill: #ffffff; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n",
		width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
	fmt.Printf(endhtml)
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
	// sx := width/2 + (x-y)*cos30*xyscale
	// sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

// f is the surface function along z over x and y
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}
