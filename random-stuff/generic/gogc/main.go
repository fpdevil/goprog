package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("/// Go Collection Stats ///")
	fmt.Println()

	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		// create some v.big memory slices
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation Failure!")
		}
	}

	printStats(mem)
}

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc", mem.HeapAlloc)
	fmt.Println("mem.NumGC", mem.NextGC)
	fmt.Println("------------------------")
}
