// Package shapes for drawing a set of related shapes
package shapes

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// A set of helper variables to hold functions all of which take
// an int and return an int.
var saneLength, saneRadius, saneSides func(int) int

// initialize the variables by assigning suitable anonymous functions
func init() {
	saneLength = makeBoundedIntFn(1, 4096)
	saneRadius = makeBoundedIntFn(1, 1024)
	saneSides = makeBoundedIntFn(3, 60)
}

// define the interfaces

// Filler interface defines methods to fill color of the type color.Color
// over an image of type draw.Image
type Filler interface {
	Fill() color.Color
	SetFill(fill color.Color)
}

// Drawer interface  defines a  method for drawing  over draw.Image  at a
// perticular x, y cartesian position
type Drawer interface {
	Draw(img draw.Image, x, y int) error
}

// Shaper interface defines  methods to get and set a  fill color of type
// color.Color and  a method for  drawing itself  over a draw.Image  at a
// particular position.
type Shaper interface {
	Filler // Fill(); Draw()
	Drawer // Draw()
}

// Radiuser interface defines method for getting and setting radius
type Radiuser interface {
	Radius() int
	SetRadius(radius int)
}

// CircularShaper interface is for a type of circle
type CircularShaper interface {
	Shaper   // Fill(); SetFill(); Draw()
	Radiuser // Radius(); SetRadius()
}

// Sideser interface defines methods for settings and getting sides of
// a regular polygon shape
type Sideser interface {
	Sides() int
	SetSides(sides int)
}

// Define Shapes

// shape is a simple  type which is defined as unexported,  so that it is
// only accessible within the shapes package. Essentially, no shape value
// can be created outside of shapes package
type shape struct {
	fill color.Color
}

// Circle type
type Circle struct {
	shape
	radius int
}

// RegularPolygon type represemts a multi sided figure
type RegularPolygon struct {
	*Circle
	sides int
}

// Option type
type Option struct {
	Fill   color.Color
	Radius int
}

// End of Shapes

// newShape function creates a new shape type with a default color filled
// with Black if none is provided. Use this for creating new shapes
func newShape(fill color.Color) shape {
	if fill == nil {
		fill = color.Black
	}
	return shape{fill}
}

// Fill function is the medium for satisfying the interfaces Filler,
// Drawer, Shaper
func (s shape) Fill() color.Color {
	return s.fill
}

// SetFill function sets an initial value for the fill type
func (s *shape) SetFill(fill color.Color) {
	if fill == nil {
		fill = color.Black
	}
	s.fill = fill
}

//+! NewCircle

// NewCircle function will create a new Circle type shape
func NewCircle(fill color.Color, radius int) *Circle {
	return &Circle{
		newShape(fill),
		saneRadius(radius),
	}
}

//!-

// Radius function is a way of circle type satisfying the Shaper
func (c *Circle) Radius() int {
	return c.radius
}

// SetRadius function is for satisfying Radiuser and Shaper
func (c *Circle) SetRadius(radius int) {
	c.radius = saneRadius(radius)
}

//!+ Draw

// Draw function satisfies the Shaper, Drawer interfaces
// The implementation of  this function to draw a circle  is based on the
// Midpoint Circle algorithm. The midpoint  circle drawing algorithm is a
// graphics  algorithm for  approximating  the pixels  needed  to draw  a
// circle given a  radius and a centre coordinate. It  is an extension to
// Bresenham's line algorithm. In this, we use eight-way symmetry so only
// ever calculate  the points for top  right eighth of a  circle and then
// use symmetry to get the rest of the points.
func (c *Circle) Draw(img draw.Image, x, y int) error {
	// The x, y coordinates might  fall outside of the image, which
	// can be  checked using  the checkBounds  function. It  is not
	// necessary to  check the radius  is in bounds  as newCircle()
	// takes care of it while creation
	if err := checkBounds(img, x, y); err != nil {
		return err
	}

	fill, radius := c.fill, c.radius

	x0, y0 := x, y
	d := roundUp(5/4.0) - (radius)
	x, y = 0, radius // start from point (0, r)

	for x < y {
		x++
		if d < 0 {
			d += 2*x + 1
		} else {
			y--
			d += 2*(x-y) + 1
		}
		img.Set(x0+x, y0+y, fill)
		img.Set(x0-x, y0+y, fill)
		img.Set(x0+x, y0-y, fill)
		img.Set(x0-x, y0-y, fill)
		img.Set(x0+x, y0+y, fill)
		img.Set(x0-x, y0+y, fill)
		img.Set(x0+x, y0-y, fill)
		img.Set(x0-x, y0-y, fill)
	}

	return nil
}

//!-

//!+ String

// String function ensures that the circle type if now satisfying the Stringer
func (c *Circle) String() string {
	str := "circle(fill: %v, radius: %d)"
	return fmt.Sprintf(str, c.fill, c.radius)
}

//!-

func checkBounds(img image.Image, x, y int) error {
	if !image.Rect(x, y, x, y).In(img.Bounds()) {
		return fmt.Errorf("%s(): point (%d, %d) is outside the image", caller(1), x, y)
	}
	return nil
}

// roundUp rounds the floating point value to int
func roundUp(val float64) int {
	if val > 0 {
		return int(val + 1.0)
	}
	return int(val)
}

//!+ NewRegularPolygon

// NewRegularPolygon function is the way of creating a regulat polygon
func NewRegularPolygon(fill color.Color, radius, sides int) *RegularPolygon {
	// NewCircle() function may be called to perform any bounds check on
	return &RegularPolygon{
		NewCircle(fill, radius),
		saneSides(sides),
	}
}

//!-

// Sides function is a getter for number of sides of polygon
func (p *RegularPolygon) Sides() int {
	return p.sides
}

// SetSides function is a setter for number of sides of a polygon
func (p *RegularPolygon) SetSides(sides int) {
	p.sides = saneSides(sides)
}

// Draw function here is the one for drawing a Regular polygon
func (p *RegularPolygon) Draw(img draw.Image, x, y int) error {
	// the x, y bounds will be checked for len(points) = sides + 1
	if err := checkBounds(img, x, y); err != nil {
		return err
	}

	points := getPoints(x, y, p.sides, float64(p.Radius()))
	// now draw lines between the apexes of polygon
	for i := 0; i < p.sides; i++ {
		drawLine(img, points[i], points[i+1], p.Fill())
	}
	return nil
}

func drawLine(img draw.Image, start, end image.Point, fill color.Color) {
	x0, x1 := start.X, end.X
	y0, y1 := start.Y, end.Y

	minX, maxX := minmax(x0, x1)
	minY, maxY := minmax(y0, y1)

	dx := float64(x1 - x0)
	dy := float64(y1 - y0)

	if maxX-minX > maxY-minY {
		d := 1
		if x0 > x1 {
			d = -1
		}
		for x := 0; x != x1-x0+d; x += d {
			y := int(float64(x) * dy / dx)
			img.Set(x0+x, y0+y, fill)
		}
	} else {
		d := 1
		if y0 > y1 {
			d = -1
		}
		for y := 0; y != y1-y0+d; y += d {
			x := int(float64(y) * dx / dy)
			img.Set(x0+x, y0+y, fill)
		}
	}
}

func minmax(p, q int) (int, int) {
	if p < q {
		return p, q
	}
	return q, p
}

func (p *RegularPolygon) String() string {
	s := "polygon(fill: %v, radius: %d, sides: %d)"
	return fmt.Sprintf(s, p.Fill(), p.Radius(), p.sides)
}

// getPoints function takes the initial x, y coordinates and number of sides
// as well as radius of the shape and returns the x, y coordinates of all the
// remaining sides of the polygon by sweeping a full 360ª
func getPoints(x, y, sides int, radius float64) []image.Point {
	// number of points = number of sides + 1
	points := make([]image.Point, sides+1)
	fullSweep := 2 * math.Pi
	x0, y0 := float64(x), float64(y)
	for i := 0; i < sides; i++ {
		𝛉 := float64((fullSweep / float64(sides)) * float64(i))
		x1 := x0 + (radius * math.Sin(𝛉))
		y1 := y0 + (radius * math.Cos(𝛉))
		points[i] = image.Pt(int(x1), int(y1))
	}

	// close the shape by assigning starting point to ending point
	points[sides] = points[0]
	return points
}

//!+ makeBoundedIntFn

// makeBoundedIntFn function  returns a function  which takes a  value x,
// returns  x if  it is  in between  minimum and  maximum (inclusive)  or
// returns the closest bounding value
func makeBoundedIntFn(minimum, maximum int) func(int) int {
	return func(x int) int {
		result := x
		switch {
		case x < minimum:
			result = minimum
		case x > maximum:
			result = maximum
		}
		if result != x {
			// we log the name of the caller of the functio here
			log.Printf("%s(): swapped %d with %d\n", caller(1), x, result)
		}
		return result
	}
}

//!-

func caller(steps int) string {
	who := "?"
	if pc, _, _, ok := runtime.Caller(steps + 1); ok {
		who = filepath.Base(runtime.FuncForPC(pc).Name())
	}
	return who
}

//!+ New

// New is a Factory function: This  helps in the situations when we would
// like to create  shapes whose shape was determined  during runtime. for
// example, using  a shape name.  The factory function will  return shape
// type values where the specific type of the value returned depends upon
// an argument.
func New(shape string, option Option) (Shaper, error) {
	sidesForShape := map[string]int{
		"triangle": 3,
		"square":   4,
		"pentagon": 5,
		"hexagon":  6,
		"heptagon": 7,
		"octagon":  8,
		"enneagon": 9,
		"nonagon":  9,
		"decagon":  10,
	}

	if sides, found := sidesForShape[shape]; found {
		return NewRegularPolygon(option.Fill, option.Radius, sides), nil
	}

	if shape != "circle" {
		return nil, fmt.Errorf("shapes.New(): invalid shape %s", shape)
	}

	return NewCircle(option.Fill, option.Radius), nil
}

//!-

//!+ FilledImage

// FilledImage function generates an image of given size filled uniformly
// with the provided color value
func FilledImage(width, height int, fill color.Color) draw.Image {
	// treat a nil color as Black implicitly
	if fill == nil {
		fill = color.Black
	}
	// ensure that both the dimensions are sensible
	width = saneLength(width)
	height = saneLength(height)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// source image to copy from
	// To fill a rectangle with a solid color, use an image.Uniform
	// source.  In  Go,  a  nil  mask image  is  equivalent  to  an
	// infinitely sized, fully opaque mask image.
	sp := &image.Uniform{fill}

	draw.Draw(img, img.Bounds(), sp, image.Point{0, 0}, draw.Src)

	return img
}

//!-

//!+ DrawShapes

// DrawShapes  function  takes a  draw.Image  to  draw upon,  a  position
// coordinate and zero or mode shapes to draw over the image.
func DrawShapes(img draw.Image, x, y int, shapes ...Drawer) error {
	for _, shape := range shapes {
		if err := shape.Draw(img, x, y); err != nil {
			return err
		}

		// for showing better in screenshots make it thick
		if err := shape.Draw(img, x+1, y); err != nil {
			return err
		}

		if err := shape.Draw(img, x, y+1); err != nil {
			return err
		}
	}
	return nil
}

//!-

//!+ SaveImage

// SaveImage  function attempts  to  save an  image  which satisfies  the
// image.Image interface to a given file name.
func SaveImage(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	switch strings.ToLower(filepath.Ext(filename)) {
	case ".jpg", ".jpeg":
		return jpeg.Encode(file, img, nil)
	case ".png":
		return png.Encode(file, img)
	}
	return fmt.Errorf("shapes.SaveImage(): '%s' has an unrecognized extension", filename)
}

//!-
