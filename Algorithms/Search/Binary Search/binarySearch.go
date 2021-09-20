package main

import (
	"fmt"
)

func main() {

	var pos = binarySearch([]int{1, 5, 10, 12, 25, 29, 30, 32}, 29)
	if pos != -1 {
		fmt.Println("Element found at position ", pos+1)
	} else {
		fmt.Println("Element is not in the list")
	}

}

func binarySearch(array []int, element int) int {
	left := 0
	right := len(array)

	fmt.Println("Entering into loop")
	for left <= right {
		middle := (left + right) / 2
		if element == array[middle] {
			return middle
		} else if element < array[middle] {
			right = middle - 1
		} else if element > array[middle] {
			left = middle + 1
		}
	}

	return -1
}
