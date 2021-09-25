package main

import "fmt"

func main() {
	fmt.Println(mergeSort([]int{29, 10, 14, 37, 14, 21, 49, 25, 11, 6, 8, 11, 6, 4, 3}))
}

func mergeSort(array []int) []int {
	if len(array) == 1 {
		return array
	}

	middle := len(array) / 2

	leftArray := mergeSort(array[:middle])
	rightArray := mergeSort(array[middle:])

	return merge(leftArray, rightArray)
}

func merge(leftArray, rightArray []int) []int {
	sortedArray := make([]int, 0)
	var i, j int
	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] < rightArray[j] {
			sortedArray = append(sortedArray, leftArray[i])
			i++
		} else {
			sortedArray = append(sortedArray, rightArray[j])
			j++
		}

	}
	sortedArray = append(sortedArray, leftArray[i:]...)
	sortedArray = append(sortedArray, rightArray[j:]...)
	return sortedArray
}
