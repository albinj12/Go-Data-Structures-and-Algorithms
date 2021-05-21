package main

import (
	"fmt"
)

var stackSize = 5 // size of the struct

type Stack struct{
		items []interface{}
		top int
	}

func (s *Stack) Push(value interface{}){
	if s.top == stackSize -1{
		fmt.Println("Stack Overflow, cannot insert more values")
		return
	}

	s.items = append(s.items, value)
	s.top++
	fmt.Printf("Inserted value %v\n", value)
	return
}

func (s *Stack) Pop(){
	if s.top == -1{
		fmt.Println("Stack Underflow, No values to remove")
		return
	}

	s.items = s.items[:s.top]
	s.top--
	return
}

func (s *Stack) Peek(){
	if s.top == -1 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Printf("Removed value %v\n", s.items[s.top])
	fmt.Println(s.items[s.top])
}

func (s *Stack) isEmpty(){
	if s.top == -1{
		fmt.Println("Stack is empty")
		return
	}

	fmt.Println("Stack is not empty")
}

func main(){
	
	var s = Stack{
		items : make([]interface{}, 0, 10),
		top : -1,
	}

}

