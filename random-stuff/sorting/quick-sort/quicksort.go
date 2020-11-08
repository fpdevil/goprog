package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Quick Sort")
	var arr = []int{21, 11, 3, 99, 35, 54, 62, 49, 6, 1, 41}
	fmt.Print("pre sort: ")
	displayArr(arr)
	qsort(arr, 0, len(arr)-1)
	fmt.Print("final sort: ")
	displayArr(arr)
}

func partition(arr []int, start, end int) int {
	pivot := arr[start]
	i := start
	j := end

	for i < j {
		for arr[i] <= pivot && i < end {
			i++
		}
		for arr[j] > pivot && j > start {
			j--
		}
		if i < j {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	arr[start] = arr[j]
	arr[j] = pivot

	return j
}

func qsort(arr []int, start, end int) {
	if start < end {
		pivot := partition(arr, start, end)
		qsort(arr, start, pivot)
		qsort(arr, pivot+1, end)
	}
}

func displayArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(strconv.Itoa(arr[i]) + " ")
	}
	fmt.Println()
}
