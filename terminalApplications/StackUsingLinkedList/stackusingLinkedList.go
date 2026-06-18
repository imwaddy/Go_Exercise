package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	head *Node
}

func main() {

	Stack := Stack{}
	Stack.Push()
	Stack.Push()
	Stack.Push()
	Stack.Push()
	Stack.DisplayElements()

}

func (stk *Stack) Push() {
	fmt.Println("Enter element ")
	var a int
	fmt.Scanf("%d", &a)

	node := &Node{data: a, next: nil}

	if stk.head == nil {
		stk.head = node
		return
	}

	node.next = stk.head
	stk.head = node

}

// DisplayElements - Displays elements in linked list
func (lst *Stack) DisplayElements() {

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
