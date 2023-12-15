package main

import "fmt"

// TreeNode represents a node in the binary search tree.
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

// BST represents a binary search tree.
type BST struct {
	Root *TreeNode
}

// main function to demonstrate the binary search tree operations.
func main() {
	// Initialize a new binary search tree
	binarySearchTree := &BST{}

	// Add values to the tree
	binarySearchTree.Insert(100)
	binarySearchTree.Insert(90)
	binarySearchTree.Insert(110)
	binarySearchTree.Insert(101)
	binarySearchTree.Insert(92)

	// Print the tree values
	binarySearchTree.Print()

	// Search for values in the tree
	binarySearchTree.Search(101)
	binarySearchTree.Search(10)
}

// Insert adds a new value to the binary search tree.
func (b *BST) Insert(value int) {
	b.Root = insertNode(b.Root, value)
}

// insertNode is a helper function to recursively insert a value into the binary search tree.
func insertNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value}
	}

	if value < node.Value {
		node.Left = insertNode(node.Left, value)
	} else if value > node.Value {
		node.Right = insertNode(node.Right, value)
	}

	return node
}

// Print displays the values of the binary search tree in-order.
func (b *BST) Print() {
	printInOrder(b.Root)
	fmt.Println()
}

// printInOrder is a helper function to recursively print the values of the tree in-order.
func printInOrder(node *TreeNode) {
	if node != nil {
		printInOrder(node.Left)
		fmt.Print(node.Value, " ")
		printInOrder(node.Right)
	}
}

// Search checks if a value exists in the binary search tree.
func (b *BST) Search(value int) {
	found := searchNode(b.Root, value)
	fmt.Println(found)
}

// searchNode is a helper function to recursively search for a value in the binary search tree.
func searchNode(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}

	if node.Value == value {
		return true
	} else if value < node.Value {
		return searchNode(node.Left, value)
	} else if value > node.Value {
		return searchNode(node.Right, value)
	}

	return false
}
