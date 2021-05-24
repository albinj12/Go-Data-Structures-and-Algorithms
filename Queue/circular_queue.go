package main

import (
	"fmt"
)

type cqueue struct{
	items []interface{}
	front int
	rear int
}

var queueSize int

func main(){

	queueSize = 5

	cq := cqueue{
		items: make([]interface{}, queueSize),
		front: -1,
		rear: -1,
	}
}

func (cq *cqueue) enQueue(value interface{}){
	
	if cq.front == 0 && cq.rear == queueSize-1 || cq.front == cq.rear+1{
		fmt.Println("Queue is full, insertion not possible")
		return
	}

	if cq.front == -1{
		cq.front, cq.rear = 0, 0
		cq.items[cq.rear] = value
		return
	}

	cq.rear = (cq.rear+1) % queueSize
	cq.items[cq.rear] = value
}

func (cq *cqueue) deQueue(){
	if cq.front == -1{
		fmt.Println("Queue is empty, deletion not possible")
		return
	}

	cq.items[cq.front] = nil
	if cq.front == cq.rear{
		cq.front, cq.rear = -1, -1
	}else{
		cq.front = (cq.front+1) % queueSize
	}
	
}

func (cq *cqueue) isFull() bool{
	if cq.front == 0 && cq.rear == queueSize-1 || cq.front == cq.rear +1{
		fmt.Println("Queue is full")
		return true
	}
	return false
}

func (cq *cqueue) isEmpty() bool{
	if cq.front == -1 {
		fmt.Println("Queue is empty")
		return true
	}
	return false
}

func (cq *cqueue) display(){
	if cq.isEmpty() {
		return
	}
	if cq.front > cq.rear{
		for i := 0; i <= cq.rear; i++ {
			fmt.Printf("%v", cq.items[i])
		}
		for i:= cq.front; i < queueSize; i++{
			fmt.Printf("%v", cq.items[i])
		}
		fmt.Println()
		return
	}else{
		for i := cq.front; i <= cq.rear; i++ {
			fmt.Printf("%v", cq.items[i])
		}
		fmt.Println()
	}
}