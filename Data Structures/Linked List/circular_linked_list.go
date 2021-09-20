package main

import "fmt"

type node struct {
	data interface{}
	next *node
}

type linkedlist struct {
	head   *node
	length int
}

func main() {

	l := linkedlist{}

}

func (l *linkedlist) display() {
	if l.head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	ptr := l.head
	for ptr.next != l.head {
		fmt.Println(ptr.data, ptr.next)
		ptr = ptr.next
	}
	fmt.Println(ptr.data, ptr.next)
}

func (l *linkedlist) insertBeginning(value interface{}) {
	newNode := &node{
		data: value,
	}
	if l.head == nil {
		l.head = newNode
		newNode.next = l.head
		return
	}
	ptr := l.head
	for ptr.next != l.head {
		ptr = ptr.next
	}
	ptr.next = newNode
	newNode.next = l.head
	l.head = newNode
	l.length++
}

func (l *linkedlist) insertEnd(value interface{}) {
	if l.head == nil {
		l.insertBeginning(value)
		return
	}

	ptr := l.head
	for ptr.next != l.head {
		ptr = ptr.next
	}

	ptr.next = &node{
		value,
		l.head,
	}
	l.length++
}

func (l *linkedlist) insertAt(value interface{}, pos int) {
	if pos < 1 || pos > l.length {
		return
	}

	if l.length == 0 || pos == 1 {
		l.insertBeginning(value)
		return
	}

	ptr := l.head
	for i := 2; i < pos; i++ {
		ptr = ptr.next
	}

	newNode := &node{
		value,
		ptr.next,
	}

	ptr.next = newNode
	l.length++
}

func (l *linkedlist) deleteBeginning() {
	if l.head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	if l.head == l.head.next {
		l.head = nil
		return
	}

	ptr := l.head
	for ptr.next != l.head {
		ptr = ptr.next
	}
	ptr.next = l.head.next
	l.head = l.head.next
	l.length--

}

func (l *linkedlist) deleteEnd() {
	if l.head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	if l.head == l.head.next {
		l.head = nil
		return
	}

	ptr := l.head
	for ptr.next.next != l.head {
		ptr = ptr.next
	}

	ptr.next = l.head
	l.length--
}

func (l *linkedlist) deleteAt(pos int) {
	if pos < 1 || pos > l.length {
		return
	}

	if pos == 1 {
		l.deleteBeginning()
		return
	}

	if l.head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	ptr := l.head
	for i := 2; i < pos; i++ {
		ptr = ptr.next
	}
	ptr.next = ptr.next.next
	l.length--

}
