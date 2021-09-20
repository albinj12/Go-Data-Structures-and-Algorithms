package main

import "fmt"

type node struct {
	data int
	next *node
}

type stack struct {
	head *node
}

func main() {
	s := stack{}
}

func (s *stack) display() {
	if s.head == nil {
		fmt.Println("Stack is empty")
		return
	}

	ptr := s.head
	for ptr != nil {
		fmt.Println(ptr.data)
		ptr = ptr.next
	}
}

func (s *stack) push(value int) {
	newNode := &node{
		value,
		s.head,
	}

	s.head = newNode
}

func (s *stack) pop() {
	if s.head == nil {
		fmt.Println("Stack Underflow")
		return
	}

	s.head = s.head.next
}
