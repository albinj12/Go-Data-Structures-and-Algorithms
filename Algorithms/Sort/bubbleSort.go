package main

import (
	"fmt"
)

func main() {
	var array = []int{29, 10, 14, 37, 14, 21, 49, 25, 11, 6, 8, 11, 6, 4, 3}
	fmt.Println(bubbleSort(array))
}

func bubbleSort(array []int) []int {
	var swap = true
	var i int

	for swap {
		swap = false
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swap = true
			}
		}
		i++
	}

	return array
}
