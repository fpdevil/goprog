package main

import (
	"fmt"

	"github.com/fpdevil/goprog/random-stuff/generic/FuzzyBool/api"
)

func main() {
	// var x *api.FuzzyBool // this is same as api.New(0)
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
	fmt.Printf("%-25s %v %3v %3v %3v %3v\n", "Original :", a, b, c, d, e)
	fmt.Printf("%-25s %v %3v %3v %3v %3v\n", "NOT      :", a.Not(), b.Not(), c.Not(), d.Not(), e.Not())
	fmt.Printf("%-25s %v %3v %3v %3v %3v\n", "NOT NOT  :", a.Not().Not(), b.Not().Not(), c.Not().Not(), d.Not().Not(), e.Not().Not())

	fmt.Printf("%-25s %v\n", "0.And(.25)    :", a.And(b))
	fmt.Printf("%-25s %v\n", "0.25.And(.50) :", b.And(c))
	fmt.Printf("%-25s %v\n", "0.50.And(.75) :", c.And(d))
	fmt.Printf("%-25s %v\n", "0.75.And(1.0) :", d.And(e))
	fmt.Printf("%-25s %v\n", "0.Or(.25)     :", a.Or(b))
	fmt.Printf("%-25s %v\n", "0.25.Or(.50)  :", b.Or(c))
	fmt.Printf("%-25s %v\n", "0.50.Or(.75)  :", c.Or(d))
	fmt.Printf("%-25s %v\n", "0.75.Or(1.0)  :", d.Or(e))

	fmt.Printf("%-25s %v\n", "a < c :", a.Less(c))
	fmt.Printf("%-25s %v\n", "a = c :", a.Equal(c))
	fmt.Printf("%-25s %v\n", "a > c :", c.Less(a))

	fmt.Printf("%-25s %v %v %v %v %v\n", "Boolean values:", a.Bool(), b.Bool(), c.Bool(), d.Bool(), e.Bool())
	fmt.Printf("%-25s %v %v %v %v %v\n", "Floating values:", a.Float(), b.Float(), c.Float(), d.Float(), e.Float())
}
