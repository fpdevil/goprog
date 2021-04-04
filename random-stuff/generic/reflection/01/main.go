package main

import (
	"fmt"
	"os"
	"reflect"
)

type a struct {
	X int
	Y float64
	Z string
}

type b struct {
	F int
	G int64
	H string
	I float64
}

func main() {
	x := 100
	xRefl := reflect.ValueOf(&x).Elem()
	xType := xRefl.Type()
	fmt.Printf("Value of x is %#v Type is %s\n", xRefl, xType)

	A := a{100, 200.212, "Struct from A"}
	B := b{1, 2, "Struct from B", 102.8}
	var r reflect.Value

	args := os.Args
	if len(args) == 1 {
		r = reflect.ValueOf(&A).Elem()
	} else {
		r = reflect.ValueOf(&B).Elem()
	}

	iType := r.Type()
	fmt.Printf("Type of i: %s\n", iType)
	fmt.Printf("%s has %d fields\n", iType, r.NumField())

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("Field name: %s\n", iType.Field(i).Name)
		fmt.Printf("has type: %s ", r.Field(i).Type())
		fmt.Printf("and  value: %#v\n", r.Field(i).Interface())
	}
}
