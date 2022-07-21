// We are implement a range sum query here

package main

import "fmt"

type segmentTree struct {
	data []int
}

func main() {
	array := []int{1,2,3,4,5,6,7}
	st := segmentTree{
		data: make([]int, 4*len(array)),
	}
	st.build(array, 0, len(array)-1, 0)
	fmt.Println(st.data)
	fmt.Println(st.query(0, len(array)-1, 0, 1, 5))
	st.update(0, len(array)-1, 0, 2, 10)
	fmt.Println(st.data)
}


// array represents the original array from which we want to build segment tree
// startIndex and endIndex represent start and end of segment
// index represent current position of segment array where we need to insert value
func (st *segmentTree) build (array []int, startIndex, endIndex, index int){

	if startIndex == endIndex{
		st.data[index] = array[startIndex]
		return
	}
	mid := (startIndex+endIndex)/2
	st.build(array, startIndex, mid, 2*index+1)
	st.build(array, mid+1, endIndex, 2*index+2)
	st.data[index] = st.data[2*index+1] + st.data[2*index+2]
}


// nodeStartIndex and nodeEndIndex represent the start and end index of range of current node
// index represent the current postion in segment array
// startIndex and endIndex represent the start and end index for range query
func (st *segmentTree) query(nodeStartIndex, nodeEndIndex, index, startIndex, endIndex int) int {

	fmt.Println(nodeStartIndex, " ", nodeEndIndex, " ", index, " ", startIndex, " ", endIndex)

	if (startIndex <= nodeStartIndex && endIndex >= nodeEndIndex) {
		return st.data[index]
	}

	if (startIndex > nodeEndIndex || endIndex < nodeStartIndex) {
		return 0
	}

	mid := (nodeStartIndex+nodeEndIndex)/2
	sumLeft := st.query(nodeStartIndex, mid, 2*index+1, startIndex, endIndex)
	sumRight := st.query(mid+1, nodeEndIndex, 2*index+2, startIndex, endIndex)
	return sumLeft+sumRight
}


// startIndex and endIndex represent start and end index of segment
// nodeIndex represent the present index of node
// index represent the position to update
// value represent the new value
func (st *segmentTree) update(startIndex, endIndex, nodeIndex, index, value int){

	if startIndex == endIndex{
		st.data[nodeIndex] = value
		return
	}

	mid := (startIndex+endIndex)/2
	if index <= mid{
		st.update(startIndex, mid, 2*nodeIndex+1, index, value)
	}else{
		st.update(mid+1, endIndex, 2*nodeIndex+2, index, value)
	}

	st.data[nodeIndex] = st.data[2*nodeIndex+1] + st.data[2*nodeIndex+2]
}