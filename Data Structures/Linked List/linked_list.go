package main

import (
	"fmt"
)

type node struct{
	data interface{}
	next *node
}

type linkedlist struct{
	head *node
	length int
}

func main(){

	l := linkedlist{}
	
}

func (l *linkedlist) display(){
	ptr := l.head
	for ptr != nil{
		fmt.Println(ptr.data)
		ptr = ptr.next
	}
}

func (l *linkedlist) insertBeginning(value interface{}){
	newNode := &node{
		data: value,
		next: l.head,
	}
	l.head = newNode
	l.length++
}


func (l *linkedlist) insertEnd(value interface{}){
	if l.head == nil{
		l.insertBeginning(value)
		return
	}

	newNode := &node{
		data: value,
		next: nil,
	}

	ptr := l.head
	for ptr.next != nil{
		ptr = ptr.next
	}

	ptr.next = newNode
	l.length++
}


func (l *linkedlist) insertAt(value interface{}, position int){

	if position < 1  || position > l.length{
		return
	}

	if l.length == 0 || position == 1{
		l.insertBeginning(value)
		return
	}

	newNode := &node{
		data: value,
	}

	ptr := l.head
	for i:= 2; i<position; i++{
		ptr = ptr.next
	}
	newNode.next = ptr.next
	ptr.next = newNode

	l.length++
}

func (l *linkedlist) deleteBeginning(){
	if l.head == nil{
		fmt.Println("Linked list is empty")
		return
	}

	l.head = l.head.next
	l.length--
}


func (l *linkedlist) deleteEnd(){
	if l.head == nil{
		fmt.Println("Linked list is empty")
		return
	}

	ptr := l.head

	for ptr.next.next != nil{
		ptr = ptr.next
	}

	ptr.next = nil

	l.length--
}


func (l *linkedlist) deleteAt(position int){
	if position < 0 || position > l.length {
		return
	}

	if position == 1{
		l.deleteBeginning()
		return
	}

	ptr := l.head

	for i:=2; i < position; i++{
		ptr = ptr.next
	}

	ptr.next = ptr.next.next
	l.length--
}