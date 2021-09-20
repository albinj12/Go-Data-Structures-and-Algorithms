package main

import "fmt"

func main() {
	var array = []int{1, 5, 10, 12, 25, 29, 30, 32}
	var pos = binarySearch(array, 9, 0, len(array))
	if pos != -1 {
		fmt.Println("Element found at position ", pos+1)
	} else {
		fmt.Println("Element is not in the list")
	}
}

func binarySearch(array []int, element, left, right int) int {

	middle := (left + right) / 2
	if left > right {
		return -1
	}

	if element == array[middle] {
		return middle
	} else if element < array[middle] {
		return binarySearch(array, element, left, middle-1)
	} else {
		return binarySearch(array, element, middle+1, right)
	}
}
