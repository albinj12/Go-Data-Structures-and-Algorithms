package main

type maxHeap struct {
	data []int
}

func main() {
	m := maxHeap{}
}

func (m *maxHeap) insert(value int) {
	m.data = append(m.data, value)
	m.maxHeapifyUp(len(m.data) - 1)
}

func (m *maxHeap) extract() {
	m.data[0] = m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	m.maxHeapifyDown(0)
}

func (m *maxHeap) maxHeapifyUp(index int) {
	parentElementIndex := (index - 1) / 2
	if index < 0 {
		return
	}

	if m.data[index] > m.data[parentElementIndex] {
		m.data[index], m.data[parentElementIndex] = m.data[parentElementIndex], m.data[index]
		m.maxHeapifyUp(parentElementIndex)
	}
}

func (m *maxHeap) maxHeapifyDown(index int) {
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2
	if leftChildIndex >= len(m.data)-1 {
		return
	}

	if leftChildIndex == len(m.data)-1 {
		m.data[index], m.data[leftChildIndex] = m.data[leftChildIndex], m.data[index]

	} else if m.data[leftChildIndex] > m.data[rightChildIndex] {
		m.data[index], m.data[leftChildIndex] = m.data[leftChildIndex], m.data[index]
		m.maxHeapifyDown(leftChildIndex)
	} else {
		m.data[index], m.data[rightChildIndex] = m.data[rightChildIndex], m.data[index]
		m.maxHeapifyDown(rightChildIndex)
	}
}
