package main

import "fmt"

type node struct {
	data int
	next *node
}

type queue struct {
	front *node
	rear  *node
}

func main() {
	q := queue{}
}

func (q *queue) display() {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return
	}

	ptr := q.front

	for ptr != q.rear {
		fmt.Println(ptr.data)
		ptr = ptr.next
	}
	fmt.Println(ptr.data)
}

func (q *queue) enQueue(value int) {
	newNode := &node{
		value,
		nil,
	}

	if q.front == nil {
		q.front, q.rear = newNode, newNode
		return
	}

	q.rear.next = newNode
	q.rear = newNode

}

func (q *queue) deQueue() {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return
	}
	if q.front == q.rear {
		q.front, q.rear = nil, nil
		return
	}

	q.front = q.front.next
}
