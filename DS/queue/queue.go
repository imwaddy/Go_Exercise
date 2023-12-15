package main

import (
	"fmt"
)

// Queue type struct
type Queue struct {
	items []int
	Front int
	Rear  int
}

// main function
func main() {
	q := Queue{}
	q.Init(5) // Initialize queue with size 5

	// Enqueue elements into the queue
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)

	// Print queue elements
	fmt.Println("Queue Elements:")
	q.Print()

	// Dequeue element from the queue
	q.Dequeue()

	// Print queue elements after dequeue
	fmt.Println("\nQueue Elements after Dequeue:")
	q.Print()

	// Print Front and Rear pointers
	fmt.Println("\nFront pointer points on", q.FrontPointer(), "location")
	fmt.Println("Rear pointer points on", q.RearPointer(), "location")
}

// Init - Queue initialization
func (q *Queue) Init(size int) {
	q.items = make([]int, size)
	q.Front = -1
	q.Rear = -1
}

// Enqueue - pushes element into queue
func (q *Queue) Enqueue(element int) {
	if q.Front == -1 {
		q.items[0] = element
		q.Front++
		q.Rear++
	} else {
		q.Rear++
		q.items[q.Rear] = element
	}
}

// Dequeue - removes element from queue
func (q *Queue) Dequeue() {
	if q.Rear == 0 {
		q.items[q.Front] = 0
		q.Rear--
		q.Front--
	} else {
		x := make([]int, cap(q.items))
		copy(x, q.items[q.Front+1:])
		q.items = x
		q.Rear--
	}
}

// Print - prints element from queue
func (q *Queue) Print() {
	for i, element := range q.items {
		fmt.Println("Number=", i, "Element=", element)
	}
}

// FrontPointer - gives Front pointer
func (q *Queue) FrontPointer() int {
	return q.Front
}

// RearPointer - gives Rear pointer
func (q *Queue) RearPointer() int {
	return q.Rear
}
