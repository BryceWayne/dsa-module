package main

import "dsa/trees"

func main() {
	// Create a new binary tree
	tree := trees.NewBinaryTree()

	// Insert elements into the tree
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)

	// Perform in-order traversal of the tree
	tree.InOrderTraversal() // Outputs: 5, 10, 15

}
