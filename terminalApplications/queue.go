// Author: Mayur Wadekar
// This is Terminal application for Queue DS.

package main

import (
	"fmt"
	"os"
	"time"
)

// Queue type struct
type Queue struct {
	items []int
	Front int
	Rear  int
}

// main function
func main() {
	s := Queue{}
	choice := 0
	fmt.Println("===================================")
	fmt.Println("This is simple stack program")
	fmt.Println("===================================")
	// loop
	for {
		fmt.Println("==================Menu=================")
		fmt.Println("1.Initialize Queue")
		fmt.Println("2.Enqueue")
		fmt.Println("3.Dequeue")
		fmt.Println("4.Front")
		fmt.Println("5.Rear")
		fmt.Println("6.Print Queue Element")
		fmt.Println("7.Exit")
		fmt.Println("===================================")
		fmt.Printf("Enter your choice=")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			size := 0
			fmt.Printf("Enter queue size=")
			fmt.Scanf("%d", &size)
			flag := s.Init(size)
			if flag {
				fmt.Println("Queue Initialized with size ", size)
				time.Sleep(time.Millisecond * 1000)
			}
		case 2:
			IsInitialized := s.IsInitialized()
			if !IsInitialized {
				isfull := s.IsFull()
				if !isfull {
					element := 0
					fmt.Printf("Enter queue element=")
					fmt.Scanf("%d", &element)
					s.Enqueue(element)
					fmt.Println("Element queued into queue")
					// time.Sleep(time.Millisecond * 1000)
				} else {
					fmt.Println("Queue is full")
					// time.Sleep(time.Millisecond * 1000)
				}
			} else {
				fmt.Println("Initialize queue first")
			}
			// time.Sleep(time.Millisecond * 1000)
		case 3:
			IsInitialized := s.IsInitialized()
			if !IsInitialized {
				isempty := s.IsEmpty()
				if !isempty {
					s.Dequeue()
					fmt.Println("Element removed from queue")
				} else {
					fmt.Println("Queue is already empty")
				}
			} else {
				fmt.Println("Initialize queue first")
			}
		// 	time.Sleep(time.Millisecond * 1000)
		case 4:
			fmt.Println("Front pointer points on ", s.FrontPointer(), " location")
			time.Sleep(time.Millisecond * 1000)
		case 5:
			fmt.Println("Rear pointer points on ", s.RearPointer(), " location")
			time.Sleep(time.Millisecond * 1000)
		case 6:
			fmt.Println("Please wait....Printing elements")
			time.Sleep(time.Millisecond * 1000)
			s.Print()
		case 7:
			os.Exit(0)
		default:
			fmt.Println("Please enter valid choice from below")
		}
	}
}

// Init - Queue initialization
func (s *Queue) Init(size int) bool {
	if cap(s.items) > 0 {
		fmt.Println("Queue already initialized. Please choose other options from below")
		time.Sleep(time.Millisecond * 1000)
		return false
	}
	s.items = make([]int, size)
	s.Front = -1
	s.Rear = -1
	return true
}

// IsInitialized - checks queue initialized or not
func (s *Queue) IsInitialized() bool {
	if cap(s.items) == 0 {
		return true
	}
	return false
}

// IsFull - checks if stack is full
func (s *Queue) IsFull() bool {
	if (cap(s.items) - 1) == s.Rear {
		return true
	}
	return false
}

// IsEmpty - checks if stack is empty
func (s *Queue) IsEmpty() bool {
	if -1 == s.Front || -1 == s.Rear {
		return true
	}
	return false
}

// Enqueue - pushes element into stack
func (s *Queue) Enqueue(element int) {
	if s.Front == -1 {
		s.items[0] = element
		s.Front++
		s.Rear++
	} else {
		s.Rear++
		s.items[s.Rear] = element
	}
}

// Dequeue - removes element from queue
func (s *Queue) Dequeue() {
	if s.Rear == 0 {
		s.items[s.Front] = 0
		s.Rear--
		s.Front--
	} else {
		x := make([]int, cap(s.items))
		copy(x, s.items[s.Front+1:])
		s.items = x
		s.Rear--
	}
}

// Print - prints element from queue
func (s *Queue) Print() {
	for i, element := range s.items {
		fmt.Println("Number=", i, "Element=", element)
	}
}

// FrontPointer - gives Front pointer
func (s *Queue) FrontPointer() int {
	return s.Front
}

// RearPointer - gives Rear pointer
func (s *Queue) RearPointer() int {
	return s.Rear
}
