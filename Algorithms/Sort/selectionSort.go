package main

import "fmt"

func main() {
	var array = []int{29, 10, 14, 37, 14, 21, 49, 25, 11, 6, 8, 11, 6, 4, 3}
	fmt.Println(selectionSort(array))
}

func selectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		lowest := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[lowest] {
				lowest = j
			}
		}
		if lowest != i {
			array[i], array[lowest] = array[lowest], array[i]
		}
	}
	return array
}
