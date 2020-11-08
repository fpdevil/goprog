package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

const (
	msg = `
	Convert numbers of integers into Roman Numerals
	`
)

type memoizeFunction func(int, ...int) interface{}

// RomanForDecimal is a variable that holds a value of type
// memoizeFunction for data caching
var RomanForDecimal memoizeFunction

func init() {
	decimals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	RomanForDecimal = Memoize(func(x int, xs ...int) interface{} {
		if x < 0 || x > 3999 {
			panic("RomanForDecimal() can only handle integers in the range [0, 3999]")
		}
		var buffer bytes.Buffer
		for i, decimal := range decimals {
			remainder := x / decimal
			// fmt.Printf("for i=%v, decimal=%v, x=%v\n", i, decimal, x)
			x %= decimal
			if remainder > 0 {
				buffer.WriteString(strings.Repeat(romans[i], remainder))
			}
		}
		return buffer.String()
	})
}

func main() {
	fmt.Println(msg)

	for _, x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 25, 30, 40, 50, 60, 69, 70, 80,
		90, 99, 100, 200, 300, 400, 500, 600, 666, 700, 800, 900,
		1000, 1009, 1444, 1666, 1945, 1997, 1999, 2000, 2008, 2010,
		2012, 2500, 3000, 3999} {
		fmt.Printf("%4d ==> %s\n", x, RomanForDecimal(x).(string))
	}
}

// Memoize function serves as a tool for memoizing the frequent
// calculations and storing them as a cache
func Memoize(function memoizeFunction) memoizeFunction {
	cache := make(map[string]interface{})
	return func(x int, xs ...int) interface{} {
		key := fmt.Sprint(x)
		for _, i := range xs {
			key += fmt.Sprintf(", %d", i)
		}
		if value, ok := cache[key]; ok {
			return value
		}
		value := function(x, xs...)
		cache[key] = value
		return value
	}
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
