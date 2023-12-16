package main

import "fmt"

// Node represents a node in the singly linked list.
type Node struct {
	Value int   // The value of the node.
	Next  *Node // Pointer to the next node in the list.
}

// SinglyLinkedList represents a singly linked list.
type SinglyLinkedList struct {
	Head *Node // Pointer to the first node in the list.
}

func main() {
	// Create a new singly linked list.
	s := &SinglyLinkedList{}

	// Insert elements at the end of the list.
	s.InsertAtTheEnd(1)
	s.InsertAtTheEnd(2)
	s.InsertAtTheEnd(3)

	// Insert an element at the beginning of the list.
	s.InsertAtTheBeginning(4)

	// Find the number of nodes in the list.
	s.FindNoOfNodes()

	// Insert elements at specific positions in the list.
	s.InsertAtPosition(6, 2)
	s.InsertAtPosition(8, 3)

	// Print the elements in the list.
	s.Print()

	// Search for the position of a specific value in the list.
	s.SearchForPosition(6)
}

// InsertAtTheEnd inserts a new node with the given value at the end of the list.
func (sll *SinglyLinkedList) InsertAtTheEnd(v int) {
	node := sll.Head
	attachNode := &Node{Value: v}
	if node == nil {
		sll.Head = attachNode
		return
	}
	for node.Next != nil {
		node = node.Next
	}
	node.Next = attachNode
}

// InsertAtTheBeginning inserts a new node with the given value at the beginning of the list.
func (sll *SinglyLinkedList) InsertAtTheBeginning(v int) {
	node := sll.Head
	attachNode := &Node{Value: v, Next: node}
	sll.Head = attachNode
}

// InsertAtPosition inserts a new node with the given value at the specified position in the list.
func (sll *SinglyLinkedList) InsertAtPosition(v, pos int) {
	if sll.Head == nil {
		fmt.Errorf("ERROR: 0 NODES IN LINKED LIST")
		return
	}

	pos -= 1
	node := sll.Head
	for pos != 0 {
		node = node.Next
		pos--
	}
	temp := node.Next
	attachNode := &Node{Value: v}
	node.Next = attachNode
	attachNode.Next = temp
}

// Print prints the elements of the list.
func (sll *SinglyLinkedList) Print() {
	node := sll.Head
	for node != nil {
		fmt.Print(node.Value)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
		node = node.Next
	}
	fmt.Println()
}

// FindNoOfNodes prints the total number of nodes in the list.
func (sll *SinglyLinkedList) FindNoOfNodes() {
	var count int
	node := sll.Head
	for node != nil {
		count++
		node = node.Next
	}
	fmt.Println("Total nodes:", count)
}

// SearchForPosition searches for the position of a node with a specific value in the list.
func (sll *SinglyLinkedList) SearchForPosition(v int) {
	var count int
	node := sll.Head
	for node != nil {
		if node.Value == v {
			fmt.Println("Value found at position ", count)
			return
		}
		count++
		node = node.Next
	}
	fmt.Println("Value not found")
}
