package main

import (
	"fmt"
)

// Stack type struct
type Stack struct {
	items []int
	top   int
}

// main function
func main() {
	s := Stack{}
	s.Init(5) // Initialize stack with size 5

	// Push elements into the stack
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	s.Push(50)

	// Print stack elements
	fmt.Println("Stack Elements:")
	s.Print()

	// Pop element from the stack
	s.Pop()

	// Print stack elements after pop
	fmt.Println("\nStack Elements after Pop:")
	s.Print()

	// Print the top element
	fmt.Println("\nTop pointer points on", s.Peek(), "location")
}

// Init - Stack initialization
func (s *Stack) Init(size int) {
	s.items = make([]int, size)
	s.top = -1
}

// Push - pushes element into stack
func (s *Stack) Push(element int) {
	s.top++
	if s.top < len(s.items) {
		s.items[s.top] = element
	}
}

// Print - prints element from stack
func (s *Stack) Print() {
	for i, element := range s.items {
		fmt.Println("Number=", i, "Element=", element)
	}
}

// Pop - pop element from stack
func (s *Stack) Pop() {
	if s.top >= 0 {
		s.items[s.top] = 0
		s.top--
	}
}

// Peek - gives top element
func (s *Stack) Peek() int {
	return s.top
}
