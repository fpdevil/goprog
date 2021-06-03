package main

import (
	"fmt"
	"strings"
	"unsafe"
)

/**
Count the number of  set bits (1s) in an integer or  more specifically, given an
integer, Count the number of set bits(1s) in the integer.

The number of  set bits in an  integer is also called Hamming  Weight. There are
quite a few ways to this.

Using Lookup Table

	if n is even: number of set bits in n is number of set bits in n/2
	if n is odd: number of set bits in n is number of set bit in n/2 + 1
				(as in case of odd number last bit is a set)


	For  example, if  we want  to  generate number  20(binary 10100)  from
	number 10(01010) then we have to left shift number 10 by 1. We can see
	number of set bits in 10 and 20  is same except for the fact that bits
	in 20 are left  shifted by 1 position compared to  number 10. So, from
	here we can conclude  that number of set bits in the  number n is same
	as that of number of set bit in n/2 (if n is even).

	In case of odd numbers, like 21(10101) all bits will be same as number
	20 except  for the  last bit, which  will be  set to 1  in case  of 21
	resulting in an extra one set bit for odd number.


	* More generic formula
	BitsSetTable256[i] = (i & 1) +  BitsSetTable256[i / 2];

	where BitsetTable256 is table we are  building for bit count. For base
	case  we can  set BitsetTable256[0]  =  0; rest  of the  table can  be
	computed using above formula in bottom up approach.
*/

// pc[i] is the population count of i
// this would be the bit set table
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = byte(i&1) + pc[i/2]
	}
}

//!+setBitCount1
// Step 1: Keep a counter to track the number of set bits.
// Step 2: If (N&1) is 1, i.e our rightmost bit is set, increment the counter.
// Step 3: After we have checked the rightmost bit, right-shift our number so
//		   that the second last bit can take the place of last bit.
// Step 4: Keep doing Steps 2-3 until our number reduces to 0.
// Step 5: Report the counter.
func setBitCount1(n int64) uint {
	var count uint
	for i := 0; i < 8*int(unsafe.Sizeof(n)); i++ {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}
	return count
}

//!-setBitCount1

//!+setBitCount2
// setBitCount2 function uses Brian Kernighan’s Algorithm
// Step 1: Keep a counter to track the number of set bits.
// Step 2: Loop until our N is not equals to 0.
// Step 3: Update our N, N = N & (N-1) and also update the counter.
// Step 4: After our N becomes 0. Report the counter.
func setBitCount2(n int64) uint {
	var count uint
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}

//!-setBitCount2

//!+PopCount
// PopCount function returns the population count (number of set bits)
// of the input value
func PopCount(x uint64) int {
	var b byte
	for i := 0; i < 8; i++ {
		b += pc[byte(x>>(i*8))]
	}

	return int(b)
}

//!-PopCount

func main() {
	nums := []int{-1, 128, 14, 15, 77, 22}
	fmt.Println(strings.Repeat("-", 36))
	for _, n := range nums {
		fmt.Printf("* PopCount(%v) = %v\n", n, PopCount(uint64(n)))
	}
	fmt.Println(strings.Repeat("-", 36))

	for _, n := range nums {
		fmt.Printf("* 1 - setBitCount1(%v) = %v\n", n, setBitCount1(int64(n)))
		fmt.Printf("* 2 - setBitCount2(%v) = %v\n", n, setBitCount2(int64(n)))
	}

}
