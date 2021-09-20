package main

import "fmt"

func main() {

	var array = []int{29, 10, 14, 37, 14, 21, 49, 25, 11, 6, 8, 11, 6, 4, 3}
	fmt.Println(insertionSort(array))

}

func insertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		selectedElement := array[i]
		j := i - 1
		for j >= 0 && array[j] > selectedElement {
			array[j+1] = array[j]
			j--
		}

		array[j+1] = selectedElement
	}

	return array
}
