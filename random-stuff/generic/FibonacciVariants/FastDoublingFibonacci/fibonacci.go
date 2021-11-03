package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"
)

func isEven(x int64) bool {
	// if number XOR 1 == number + 1 its Even
	return x^1 == x+1
}

func isOdd(x int64) bool {
	// if number XOR 1 == number - 1 its Odd
	return x^1 == x-1
}

func fastexp(b, e int64) int64 {
	if e == 0 {
		return 1
	} else if isOdd(e) {
		return b * fastexp(b, e-1)
	} else {
		return fastexp(b, e/2) * fastexp(b, e/2)
	}
}

// pow computes the power of a to n for floating point a and non-negative
// integer n iteratively using repeated squaring that uses the fact a^n =
// (a^(n/1))^2. The algorithm iterates over the bits of n, so it iterates
// log(n) times.
func pow(a, n int64) int64 {
	nbits := strconv.FormatInt(n, 2)
	res := int64(1)
	for _, b := range nbits {
		res *= res
		if string(b) == "1" {
			res *= a
		}
	}
	return res
}

//!+matmul

// matmul function returns the mulitplied results of 2 matrices
func matmul(m1, m2 [][]int64) [][]int64 {
	m := [][]int64{{0, 0}, {0, 0}}
	m[0][0] = m1[0][0]*m2[0][0] + m1[0][1]*m2[1][0]
	m[0][1] = m1[0][0]*m2[0][1] + m1[0][1]*m2[1][1]
	m[1][0] = m1[1][0]*m2[0][0] + m1[1][1]*m2[1][0]
	m[1][1] = m1[1][0]*m2[0][1] + m1[1][1]*m2[1][1]
	return m
}

//!-matmul

func matmulopt(m1 [][]int64) [][]int64 {
	// here m2 is {{1, 1}, {1, 0}}
	m := [][]int64{{0, 0}, {0, 0}}
	m[0][0] = m1[0][0] + m1[0][1]
	m[0][1] = m1[0][0]
	m[1][0] = m1[1][0] + m1[1][1]
	m[1][1] = m1[1][0]
	return m
}

func matpow(m [][]int64, n int64) [][]int64 {
	nbits := strconv.FormatInt(n, 2)
	res := [][]int64{{1, 0}, {0, 1}} // identity matrix
	for _, b := range nbits {
		res = matmul(m, res)
		if string(b) == "1" {
			res = matmulopt(res)
		}
	}
	return res
}

func fibonacci(n int64) *big.Int {
	if n < 0 {
		panic("Fibonacci of Negative numbers not allowed")
	}
	fst, _ := fibonacciDoubling(n)
	return fst
}

func fibonacciDoubling(n int64) (*big.Int, *big.Int) {
	if n == 0 {
		return big.NewInt(0), big.NewInt(1)
	}

	a, b := fibonacciDoubling(n >> 1)
	c := big.NewInt(0).Mul(a, (big.NewInt(0).Sub(big.NewInt(0).Mul(b, big.NewInt(2)), a))) // c = a(2b - a)
	d := big.NewInt(0).Add(big.NewInt(0).Mul(a, a), big.NewInt(0).Mul(b, b))               // d = a^2 + b^2

	if isOdd(n) {
		return d, big.NewInt(0).Add(c, d) // if n is odd then F(2n+1) i., (d, c + d)
	}

	return c, d // if n is even then F(2n)
}

//!+
func main() {
	defer trace("main")()
	var i int64
	fmt.Println(" Calculating large Fibonacci numbers")
	fmt.Println("-- From 0 to 100 multiples of 10 ---")
	for i = 0; i <= 100; i += 10 {
		fmt.Printf("* Fibonacci(%3v): %v\n", i, fibonacci(i))
	}

	fmt.Println()
	fmt.Println("-- From 100 to 1000 multiples of 100 ---")
	for i = 100; i <= 1000; i += 100 {
		fmt.Printf("* Fibonacci(%3v): %v\n", i, fibonacci(i))
	}

	fmt.Println()
	fmt.Println("---------------------------------------------")
	fmt.Println("* The 3,184th Fibonacci number *")
	fmt.Println("  This is an Apocalyptic number having 666 Digits")
	fmt.Println()
	fmt.Printf("%v\n", fibonacci(3184))
	fmt.Println()
	fmt.Println("---------------------------------------------")
	fmt.Println()
}

//!-

//!+trace

// trace functin provides runtime statistics of the execution of the
// method under which this is invoked using a defer
func trace(msg string) func() {
	start := time.Now()
	log.Printf("** [enter] %s **", msg)
	return func() {
		log.Printf("** [exit] %s took: %v **", msg, time.Since(start).String())
	}
}

//!-trace
