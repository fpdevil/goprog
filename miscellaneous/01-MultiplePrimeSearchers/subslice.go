package main

import (
	"math/rand"
	"time"
)

var (
	t = time.Now().UnixNano()
	s = rand.NewSource(t)
	r = rand.New(s)
)

// subSlice returns  a slice of the  'data' such that 'data'  can be
// considered broken up  into subsets for 'n'  subset. The parameter
// 'i' specifies which  0 based subset of the data  to return. A nil
// slice is return if 'i' or 'n' is negative, or if 'data' is nil
func subslice(data []int, n, i int) (output []int) {
	l := len(data)
	if l == 0 || i < 0 || n < 0 {
		return
	}

	// subsize of the maximum slice for n subsets
	maxSubset := l / n

	// check if the slice is of equal parts of multiples of n
	// increment the max subset size by 1 if it does not divide evenly
	if (l % n) != 0 {
		maxSubset++
	}

	offset := i * maxSubset
	if offset > l {
		return
	}

	actualSize := min(maxSubset, l-offset)
	// log.Infof("subSlice(%v, %v, %v): len: %v, max size, %v, offset: %v, actual size: %v", data, n, i, l, maxSubset, offset, actualSize)
	output = data[offset : offset+actualSize]
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
xdata := []int{1, 2, 3, 4, 5, 6, 7, 8}
x := 6
for i := 0; i < x; i++ {
	s := subslice(xdata, x, i)
	fmt.Printf("%#v\n", s)
	primes(i+1, s)
}
*/
