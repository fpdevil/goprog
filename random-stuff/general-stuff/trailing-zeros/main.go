package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("usage %s <start> <end> <step>\n", filepath.Base(os.Args[0]))
		fmt.Printf("usage %s 5.0 100 5\n", filepath.Base(os.Args[0]))
		return
	}

	start, _ := strconv.ParseFloat(os.Args[1], 64)
	end, _ := strconv.ParseFloat(os.Args[2], 64)
	step, _ := strconv.ParseFloat(os.Args[3], 64)

	for num := start; num <= end; num += step {
		v, err := trailingZeros(num)
		if err != nil {
			fmt.Printf("error finding trailing zeros of factorial %v: %v\n", num, err)
			return
		}
		fmt.Printf("%3v! has %3v trailing zeros\n", num, v)
	}
}

func isPositiveInteger(val float64) bool {
	return !math.Signbit(val) && val == float64(int(val))
}

func log5(x float64) float64 {
	return math.Log(x) / math.Log(5.0)
}

func trailingZeros(num float64) (float64, error) {
	if isPositiveInteger(num) {
		k := int(math.Floor(log5(num)))
		zeros := 0.0
		for i := 1; i <= k+1; i++ {
			v := num / math.Pow(5, float64(i))
			// fmt.Printf("num: %v i: %v log(i): %v v: %v zeros: %v\n", num, i, log5(float64(i)), v, zeros)
			zeros += math.Floor(v)
		}
		return zeros, nil
	} else {
		return 0.0, fmt.Errorf("factorial for non-positive integer %v is not defined", num)
	}
}
