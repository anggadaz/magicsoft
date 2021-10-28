package main

import (
	"fmt"
)

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

type QueueFix struct {
	size int
	data []interface{}
}

func New(size int) *QueueFix {
	return &QueueFix{
		size: size,
	}
}

// IsEmpty : checks whether the queue is empty
func (q *QueueFix) IsEmpty() bool {
	return len(q.data) == 0
}

// Peek : returns the next element in the queue
func (q *QueueFix) Peek() (interface{}, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	return q.data[0], nil
}

// Queue : adds an element onto the queue and returns an pointer to the current queue
func (q *QueueFix) Push(n interface{}) *QueueFix {
	if q.Len() < q.size {
		q.data = append(q.data, n)
	} else {
		q.Pop()
		q.Push(n)
	}
	return q
}

// Dequeue : removes the next element from the queue and returns its value
//func (q *QueueFix) Pop() (interface{}, error) {
func (q *QueueFix) Pop() interface{} {
	if len(q.data) == 0 {
		//return 0, fmt.Errorf("Queue is empty")
		return 0
	}
	element := q.data[0]
	q.data = q.data[1:]
	//return element, nil
	return element
}

func (q *QueueFix) Len() int {
	return len(q.data)
}

func (q *QueueFix) Keys() []interface{} {
	return q.data
}

func (q *QueueFix) Contains(key interface{}) bool {
	cont := false
	for i := 0; i < q.Len(); i++ {
		if q.data[i] == key {
			cont = true
		}
	}
	return cont
}

func main() {
	queue := New(5)
	result, _ := queue.Push(1).Push(2).Push(3).Peek()
	fmt.Println(result)
	fmt.Println(queue.IsEmpty())
	// result, _ = queue.Pop()
	// fmt.Println(result)
	// result, _ = queue.Pop()
	// fmt.Println(result)
	queue.Pop()
	fmt.Println(queue.IsEmpty())
	_, err := queue.Peek()
	fmt.Println(err)
	fmt.Println(queue.Keys())
	fmt.Println(queue.Len())
	fmt.Println(queue.Contains(2))
	fmt.Println(queue.Contains(9))
}
