package main

import "fmt"

type minHeap struct{
	data []int
}

func main(){
	m := minHeap{}
	// m.insert(10)
	// m.insert(20)
	// m.insert(15)
	// m.insert(30)
	// m.insert(40)
	// m.display()
	// m.extract()
	// m.display()
	// m.extract()
	// m.display()
	// m.insert(5)
	// m.display()
	// m.extract()
	// m.display()
}

func (m *minHeap) insert(value int){
	m.data = append(m.data, value)
	m.heapifyUp(len(m.data)-1)
}

func (m *minHeap) extract() {
	m.data[0] = m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	m.heapifyDown(0)
}


func (m *minHeap) heapifyUp(index int){
	parentElementIndex := (index-1)/2

	if index < 0 {
		return
	}

	if m.data[index] < m.data[parentElementIndex]{
		m.data[index], m.data[parentElementIndex] = m.data[parentElementIndex], m.data[index]
		m.heapifyUp(parentElementIndex)
	}
}

func (m *minHeap) heapifyDown(index int) {
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2

	smaller := index
	if leftChildIndex <= len(m.data) -1 && m.data[leftChildIndex] < m.data[index] {
 		smaller = leftChildIndex
	}

	if rightChildIndex <= len(m.data) -1 && m.data[rightChildIndex] < m.data[smaller]{
		smaller = rightChildIndex
	}

	if smaller != index{
		m.data[smaller], m.data[index] = m.data[index], m.data[smaller]
		m.heapifyDown(smaller)
	}
}

func (m *minHeap) display() {
	fmt.Println(m.data)
}