package main

import (
	"fmt"
)

func main() {
	fmt.Println("/// multiple channels ///")
	fmt.Println()
	fmt.Println("Random BIT Stream Generation:")
	const total = 10000
	bits := genRandomBits(total)
	m := make(map[int8]int)
	for v := range bits {
		m[v]++
	}

	for k, v := range m {
		f := (float64(v) / total) * 100
		fmt.Printf("%v occurred %.2f%% of the time\n", k, f)
	}
	fmt.Println()
}

func genRandomBits(l int) (out chan int8) {
	out = make(chan int8) // unbuffered channel

	go func() {
		for i := 0; i < l; i++ {
			// sending 1 twice will give 66.66% chance for 1
			// and 33.33% chance for 0 as 2/3rd is 1 snd 1/3rd
			// is sent as 0
			select {
			case out <- 0:
			case out <- 1:
			case out <- 1:
			}
		}
		close(out)
	}()

	return out
}
