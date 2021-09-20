package main

import "fmt"

type node struct {
	data interface{}
	prev *node
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func main() {

	l := linkedList{}

}

func (l *linkedList) display() {
	if l.head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	ptr := l.head
	for ptr.next != l.head {
		fmt.Println(ptr.data)
		ptr = ptr.next
	}
	fmt.Println(ptr.data)
}

func (l *linkedList) insertBeginning(value interface{}) {
	newNode := &node{
		data: value,
	}

	if l.head == nil {
		l.head = newNode
		newNode.next = l.head
		newNode.prev = l.head
		return
	}

	newNode.prev = l.head.prev
	newNode.next = l.head
	l.head.prev.next = newNode
	l.head.prev = newNode
	l.head = newNode
	l.length++
}

func (l *linkedList) insertEnd(value interface{}) {
	if l.head == nil {
		l.insertBeginning(value)
		return
	}

	newNode := &node{
		value,
		l.head.prev,
		l.head,
	}

	l.head.prev.next = newNode
	l.head.prev = newNode
}

func (l *linkedList) insertAt(value interface{}, pos int) {
	if pos < 0 || pos > l.length {
		return
	}

	if l.head == nil || pos == 1 {
		l.insertBeginning(value)
		return
	}

	ptr := l.head
	for i := 2; i < pos; i++ {
		ptr = ptr.next
	}

	newNode := &node{
		value,
		ptr,
		ptr.next,
	}

	ptr.next.prev = newNode
	ptr.next = newNode
	l.length++
}

func (l *linkedList) deleteBeginning() {
	if l.head == nil {
		fmt.Println("list is empty")
		return
	}

	l.head.next.prev = l.head.prev
	l.head.prev.next = l.head.next
	l.head = l.head.next
	l.length--
}

func (l *linkedList) deleteEnd() {
	if l.head == nil {
		fmt.Println("list is empty")
		return
	}

	l.head.prev = l.head.prev.prev
	l.head.prev.next = l.head
	l.length--
}

func (l *linkedList) deleteAt(pos int) {
	if l.head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	ptr := l.head
	for i := 2; i < pos; i++ {
		ptr = ptr.next
	}

	ptr.next = ptr.next.next
	ptr.next.prev = ptr
	l.length--
}
