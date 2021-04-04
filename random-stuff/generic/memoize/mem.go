// Package memoize helps in caching the function results for helping with
// faster results of repeated computation using memoization.
package memoize

import (
	"fmt"
)

// MemoizedFunc is a type for defining the memoized functions which takes
// a series of integers and returns any type
type MemoizedFunc func(int, ...int) interface{}

//!+ Memoize

// Memoize  function takes  a function  of the  type of  MemoizedFunc and
// returns  the same  type. The  returned function  remembers the  return
// values(s) of the function call. It memoizes any function that takes at
// least one int and that returns  an interface type. It cannot work with
// pointers as they only provide addreeses which cannot be cached
func Memoize(function MemoizedFunc) MemoizedFunc {
	// cache is a map to cache precomputed results with string keys and
	// interface{} values
	cache := make(map[string]interface{})
	// create closure of the type MemoizedFunc
	f := func(x int, xs ...int) interface{} {
		// key here is a comma delimited string of all int arguments joined
		key := fmt.Sprint(x) // convert the first value to a string
		for _, i := range xs {
			key += fmt.Sprintf(", %d", i)
		}

		if value, found := cache[key]; found {
			return value
		}
		value := function(x, xs...)
		cache[key] = value
		return value
	}
	return f
}

//!- Memoize
