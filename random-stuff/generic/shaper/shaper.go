package main

import (
	"fmt"

	"github.com/fpdevil/goprog/random-stuff/generic/shaper/shapes"
)

func sanityCheck(name string, shape shapes.Shaper) {

}

//!+ showShapeDetails

// showShapeDetails function is useful getting the details of any
// of kind of shape.
func showShapeDetails(shape shapes.Shaper) {
	fmt.Print("fill: ", shape.Fill(), " ") // all shapes have a fill color
	if shape, ok := shape.(shapes.CircularShaper); ok {
		fmt.Print("radius: ", shape.Radius(), " ")
	}
	if shape, ok := shape.(shapes.Radiuser); ok {
		fmt.Print("radius: ", shape.Radius(), " ")
	}
	if shape, ok := shape.(shapes.Sideser); ok {
		fmt.Print("sides: ", shape.Sides(), " ")
	}
	fmt.Println()
}

//!-
