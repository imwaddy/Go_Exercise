// Author: Mayur Wadekar
// This is Terminal application for Linked List DS.

package main

import (
	"fmt"
	"os"
)

// Node holds data and pointer to next node
type Node struct {
	data int
	next *Node
}

// Linked list struct contains head in it
type List struct {
	head *Node
}

// main function
func main() {

	lst := &List{}
	choice := 0
	fmt.Println("===================================")
	fmt.Println("This is simple Linked List program")
	fmt.Println("===================================")
	// loop
	for {
		fmt.Println("==================Menu=================")
		fmt.Println("1. Add Element")
		fmt.Println("2. Delete Element")
		fmt.Println("3. Print Linked List Element")
		fmt.Println("4. Exit")
		fmt.Println("===================================")
		fmt.Printf("Enter your choice=")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			var element int
			fmt.Printf("Enter queue element=")
			fmt.Scanf("%d", &element)
			lst.InsertElement(element)
			fmt.Printf("Element inserted successfully")
		case 2:
			var element int
			fmt.Printf("Enter element which you want to delete=")
			fmt.Scanf("%d", &element)
			lst.DeleteElement(element)
			fmt.Printf("Element deleted successfully")
		case 3:
			lst.DisplayElements()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Please enter valid choice from below")
		}
	}
}

// InsertElement - Insert element in linked list
func (lst *List) InsertElement(elem int) {

	// prepare node
	node := &Node{data: elem, next: nil}

	// If head nil then assign node and return
	if lst.head == nil {
		lst.head = node
		return
	}

	// If head non-nil then iterate through
	head := lst.head
	for head.next != nil {
		head = head.next
	}
	// If list is nit then assign the node
	head.next = node

}

// DisplayElements - Displays elements in linked list
func (lst *List) DisplayElements() {

	node := lst.head

	if node == nil {
		fmt.Println("Something wrong, Wait you have not added any nodes in your list...!")
		return
	}

	fmt.Print("Linked Elements are: ")
	for node != nil {
		fmt.Print(node.data, "-> ")
		node = node.next
	}
	fmt.Println()

}

// DeleteElement - Delete element in linked list
func (lst *List) DeleteElement(elem int) {

	// If empty then return
	if lst.head == nil {
		fmt.Print("Please insert elements so you can delete :) ")
		return
	}

	// If first element in list then point it to next node
	if lst.head.data == elem {
		lst.head = lst.head.next
		return
	}

	node := lst.head
	for node.next != nil {
		if node.next.data == elem {
			node.next = node.next.next
			return
		}
		node = node.next
	}

	fmt.Println("There is no elements in linked list..")
}
