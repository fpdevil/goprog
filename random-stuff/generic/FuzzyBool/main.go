package main

import (
	"fmt"

	"github.com/fpdevil/goprog/random-stuff/generic/FuzzyBool/api"
)

func main() {
	a, _ := api.New(0)
	b, _ := api.New(.25)
	c, _ := api.New(0.5)
	d, _ := api.New(0.75)
	e := d.Copy()
	if err := e.Set(1); err != nil {
		fmt.Println(err)
	}

	processResults(a, b, c, d, e)
}

func processResults(a, b, c, d, e *api.FuzzyBool) {
	fmt.Printf("Original values			: %v %v %v %v %v\n", a, b, c, d, e)
	fmt.Printf("NOT values			: %v %v %v %v %v\n", a.Not(), b.Not(), c.Not(), d.Not(), e.Not())
	fmt.Printf("NOT NOT values			: %v %v %v %v %v\n", a.Not().Not(), b.Not().Not(), c.Not().Not(), d.Not().Not(), e.Not().Not())

	fmt.Printf("0.And(.25)			: %v\n", a.And(b))
	fmt.Printf("0.25.And(.50)			: %v\n", b.And(c))
	fmt.Printf("0.50.And(.75)			: %v\n", c.And(d))
	fmt.Printf("0.75.And(1.0)			: %v\n", d.And(e))
	fmt.Printf("0.Or(.25)			: %v\n", a.Or(b))
	fmt.Printf("0.25.Or(.50)			: %v\n", b.Or(c))
	fmt.Printf("0.50.Or(.75)			: %v\n", c.Or(d))
	fmt.Printf("0.75.Or(1.0)			: %v\n", d.Or(e))

	fmt.Printf("a < c				: %v\n", a.Less(c))
	fmt.Printf("a = c				: %v\n", a.Equal(c))
	fmt.Printf("a > c				: %v\n", c.Less(a))

	fmt.Printf("Boolean valeus			: %v %v %v %v %v\n", a.Bool(), b.Bool(), c.Bool(), d.Bool(), e.Bool())
	fmt.Printf("Floating values			: %v %v %v %v %v\n", a.Float(), b.Float(), c.Float(), d.Float(), e.Float())
}
