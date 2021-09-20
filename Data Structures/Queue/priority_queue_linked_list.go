// Ascending order priority queue
package main

import "fmt"

type node struct {
	data     int
	priority int
	next     *node
}

type pqueue struct {
	front *node
}

func main() {
	p := pqueue{}
}

func (p *pqueue) display() {
	if p.front == nil {
		fmt.Println("Prioriy Queue is empty")
		return
	}

	ptr := p.front
	fmt.Println("Data, Priority")
	for ptr != nil {
		fmt.Println(ptr.data, ptr.priority)
		ptr = ptr.next
	}
}

func (p *pqueue) enQueue(value, prio int) {
	newNode := &node{
		value,
		prio,
		nil,
	}

	if p.front == nil || newNode.priority < p.front.priority {
		newNode.next = p.front
		p.front = newNode
		return
	}

	ptr := p.front
	for ptr.next != nil && newNode.priority > ptr.next.priority {
		ptr = ptr.next
	}

	newNode.next = ptr.next
	ptr.next = newNode
}

func (q *pqueue) deQueue() {
	if q.front == nil {
		fmt.Println("Prioriy Queue is empty")
		return
	}

	q.front = q.front.next
}
