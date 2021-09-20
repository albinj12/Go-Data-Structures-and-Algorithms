package main

import(
	"fmt"
)

var queueSize int

type queue struct{
	items []interface{}
	front int
	rear int
}

func main(){

	queueSize = 5

	var q = queue{
		items: make([]interface{}, 0, queueSize),
		front: -1,
		rear: -1,
	}

}

func (q *queue)enQueue(value interface{}){

	if q.rear == queueSize -1 {
		fmt.Println("Queue is full")
		return
	}
	
	if q.front == -1{
		q.items = append(q.items, value)
		q.front, q.rear = 0,0
	}else{
		q.items = append(q.items, value)
		q.rear++
	}
}

func (q *queue)deQueue(){

	if q.front == -1{
		fmt.Println("Queue is empty")
		return
	}

	if q.front == queueSize - 1{
		q.items = q.items[:0]
		q.front, q.rear = -1, -1
	}else{
		q.front++
		q.items = q.items[1:]
	}	
}

func (q *queue) isEmpty(){
	if q.front == -1{
		fmt.Println("Queue is empty")
	}else{
		fmt.Println("Queue is not empty")
	}
}

func (q *queue) isFull(){
	if q.rear == queueSize-1{
		fmt.Println("Queue is full")
	}else {
		fmt.Println("Queue is not full")
	}
}
