// Author: Mayur Wadekar
// This is Terminal application for Stack DS.

package main

import (
	"fmt"
	"os"
	"time"
)

// Stack type struct
type Stack struct {
	items []int
	top   int
}

// main function
func main() {
	s := Stack{}
	choice := 0
	fmt.Println("===================================")
	fmt.Println("This is simple stack program")
	fmt.Println("===================================")
	// loop
	for {
		fmt.Println("==================Menu=================")
		fmt.Println("1.Initialize Stack")
		fmt.Println("2.Push")
		fmt.Println("3.Pop")
		fmt.Println("4.Peek(top)")
		fmt.Println("5.Print Stack")
		fmt.Println("6.Exit")
		fmt.Println("===================================")
		fmt.Printf("Enter your choice=")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			size := 0
			fmt.Printf("Enter stack size=")
			fmt.Scanf("%d", &size)
			flag := s.Init(size)
			if flag {
				fmt.Println("Stack Initialized with size ", size)
				time.Sleep(time.Millisecond * 1000)
			}
		case 2:
			IsInitialized := s.IsInitialized()
			if !IsInitialized {
				isfull := s.IsFull()
				if !isfull {
					element := 0
					fmt.Printf("Enter stack element=")
					fmt.Scanf("%d", &element)
					s.Push(element)
					fmt.Println("Element pushed into stack")
					// time.Sleep(time.Millisecond * 1000)
				} else {
					fmt.Println("Stack is full")
					// time.Sleep(time.Millisecond * 1000)
				}
			} else {
				fmt.Println("Initialize stack first")
			}
			time.Sleep(time.Millisecond * 1000)
		case 3:
			IsInitialized := s.IsInitialized()
			if !IsInitialized {
				isempty := s.IsEmpty()
				if !isempty {
					s.Pop()
					fmt.Println("Element poped from stack")
				} else {
					fmt.Println("Stack is already empty")
				}
			} else {
				fmt.Println("Initialize stack first")
			}
			time.Sleep(time.Millisecond * 1000)
		case 4:
			fmt.Println("Top pointer points on ", s.Peek(), " location")
			time.Sleep(time.Millisecond * 1000)
		case 5:
			fmt.Println("Please wait....Printing elements")
			time.Sleep(time.Millisecond * 1000)
			s.Print()
		case 6:
			os.Exit(0)
		default:
			fmt.Println("Please enter valid choice from below")
		}
	}
}

// Init - Stack initialization
func (s *Stack) Init(size int) bool {
	if cap(s.items) > 0 {
		fmt.Println("Stack already initialized. Please choose other options from below")
		time.Sleep(time.Millisecond * 1000)
		return false
	}
	s.items = make([]int, size)
	s.top = -1
	return true
}

// IsInitialized - checks stack initialized or not
func (s *Stack) IsInitialized() bool {
	if cap(s.items) == 0 {
		return true
	}
	return false
}

// IsFull - checks if stack is full
func (s *Stack) IsFull() bool {
	if (cap(s.items) - 1) == s.top {
		return true
	}
	return false
}

// IsEmpty - checks if stack is empty
func (s *Stack) IsEmpty() bool {
	if -1 == s.top {
		return true
	}
	return false
}

// Push - pushes element into stack
func (s *Stack) Push(element int) {
	s.top++
	if s.top == -1 {
		s.items[0] = element
	} else {
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
	s.items[s.top] = 0
	s.top--
}

// Peek - gives top element
func (s *Stack) Peek() int {
	return s.top
}
