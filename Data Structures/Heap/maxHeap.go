package main

import "fmt"

type maxHeap struct {
	data []int
}

func main() {
	m := maxHeap{}
	// m.insert(10)
	// m.insert(20)
	// m.insert(15)
	// m.insert(30)
	// m.insert(40)
	// m.display()
	// m.extract()
	// m.display()
	// m.insert(5)
	// m.display()
	// m.extract()
	// m.display()
}

func (m *maxHeap) insert(value int) {
	m.data = append(m.data, value)
	m.HeapifyUp(len(m.data) - 1)
}

func (m *maxHeap) extract() {
	m.data[0] = m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	m.HeapifyDown(0)
}

func (m *maxHeap) HeapifyUp(index int) {
	parentElementIndex := (index - 1) / 2
	if index < 0 {
		return
	}

	if m.data[index] > m.data[parentElementIndex] {
		m.data[index], m.data[parentElementIndex] = m.data[parentElementIndex], m.data[index]
		m.HeapifyUp(parentElementIndex)
	}
}

func (m *maxHeap) HeapifyDown(index int) {
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2

	greater := index
	if leftChildIndex <= len(m.data) -1 && m.data[leftChildIndex] > m.data[index] {
 		greater = leftChildIndex
	}

	if rightChildIndex <= len(m.data) -1 && m.data[rightChildIndex] > m.data[greater]{
		greater = rightChildIndex
	}

	if greater != index{
		m.data[greater], m.data[index] = m.data[index], m.data[greater]
		m.HeapifyDown(greater)
	}
}

func (m *maxHeap) display() {
	fmt.Println(m.data)
}