package main

import (
	"fmt"
)

type node struct{
	data interface{}
	prev *node
	next *node
}

type linkedList struct{
	head *node
	length int
}

func main(){

	l := linkedList{}

}

func (l *linkedList) display(){
	ptr := l.head
	for ptr != nil{
		// fmt.Printf("%v, %v, %v\n", ptr.prev, ptr.data, ptr.next)
		fmt.Println(ptr.data)
		ptr = ptr.next
	}
}

func (l *linkedList) insertBeginning(value interface{}){
	newNode := &node{
		data: value,
		next: l.head,
		prev: nil,
	}

	if l.head != nil{
		l.head.prev = newNode
	}

	l.head = newNode
	l.length++
}

func (l *linkedList) insertEnd(value interface{}){
	if l.head == nil{
		l.insertBeginning(value)
		return
	}

	ptr := l.head
	for ptr.next != nil{
		ptr = ptr.next
	}

	newNode := &node{
		data: value,
		next: nil,
		prev: ptr,
	}

	ptr.next = newNode
	l.length++

}

func (l *linkedList) insertAt(value interface{}, position int){
	if position < 1 || position > l.length{
		return
	}

	if l.head == nil || position == 1{
		l.insertBeginning(value)
		return
	}

	ptr := l.head

	for i:=2; i<position; i++{
		ptr = ptr.next
	}

	newNode := &node{
		data: value,
		next: ptr.next,
		prev: ptr,
	}

	ptr.next.prev = newNode
	ptr.next = newNode
	l.length++
}

func (l *linkedList) deleteBeginning(){
	if l.head == nil {
		fmt.Println("list is empty")
		return
	}

	l.head.prev = nil
	l.head = l.head.next
	l.length--
}

func (l *linkedList) deleteEnd(){
	if l.head == nil{
		fmt.Println("list is empty")
		return
	}

	ptr := l.head
	for ptr.next != nil{
		ptr = ptr.next
	}

	ptr.prev.next = nil
	l.length--

}

func (l *linkedList) deleteAt(position int){
	if position < 0 || position > l.length{
		return
	}

	if position == 1{
		l.deleteBeginning()
		return
	}

	ptr := l.head

	for i:=2; i<position; i++{
		ptr = ptr.next
	}


	ptr.next = ptr.next.next
	if (position < l.length){
		ptr.next.prev = ptr
	}

	l.length--

}