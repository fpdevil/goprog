// Package roman_numerals provides a conversion of integer to roman
// numbers and it uses memoization or repeated calculations
package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/fpdevil/goprog/random-stuff/generic/memoize"
)

const (
	msg = `
	* Convert Integers into Roman Numerals *
	`
)

// RomanForDecimal is a variable that holds a value of type
// MemoizedFunct for data caching
var RomanForDecimal memoize.MemoizedFunc

// init initializes the RomanForDecimal with an appropriate version of the function
// and also defines and initializes certain values like decimals and romans
// The largest number that can be represented in roman notation is 3,999 (MMMCMXCIX)
func init() {
	decimals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	decimalsToRomans := func(x int, xs ...int) interface{} {
		if x < 0 || x > 3999 {
			panic("RomanForDecimal: will only handle numbers in the range [0, 3999]")
		}
		var buffer bytes.Buffer
		for i, decimal := range decimals {
			// get quotient and remainder for x
			quotient := x / decimal
			if quotient > 0 {
				buffer.WriteString(strings.Repeat(romans[i], quotient))
			}
			x %= decimal
		}
		return buffer.String()
	}

	RomanForDecimal = memoize.Memoize(decimalsToRomans)
}

func main() {
	fmt.Println(msg)
	for _, x := range []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18,
		19, 20, 25, 30, 40, 50, 60, 69,
		70, 80, 90, 99, 100, 200, 300,
		400, 500, 600, 666, 700, 800,
		900, 1000, 1009, 1444, 1666, 1945,
		1997, 1999, 2000, 2008, 2010, 2012,
		2020, 2021, 2500, 3000, 3999,
	} {
		fmt.Printf("%4d = %s\n", x, RomanForDecimal(x).(string))
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
