package avl_tree

import (
	"fmt"
	"strings"
)

type AVLTree[T types] struct {
	Root  *node[T]
	Count int
}

func EmptyAVLTree[T types]() *AVLTree[T] {
	return &AVLTree[T]{
		Root:  nil,
		Count: 0,
	}
}

func NewAVLTree[T types](values ...T) *AVLTree[T] {
	avl := &AVLTree[T]{}
	for _, value := range values {
		avl.Insert(value)
	}
	return avl
}

// Insert inserts a node in the tree with the specified value, and balances the tree.
//
// Time complexity: O(log n) because node.insert() is O(log n)
func (avl *AVLTree[T]) Insert(value T) {
	avl.Root = avl.Root.insert(value)
	avl.Count++
}

// Remove removes a node in the tree with the specified value, then balances the tree.
//
// Time complexity: O(log n) because node.remove() is O(log n)
func (avl *AVLTree[T]) Remove(value T) {
	avl.Root = avl.Root.remove(value)
	avl.Count--
}

// Contains checks if the tree contains a node with the specified value.
//
// Time complexity: O(log n)
func (avl *AVLTree[T]) Contains(value T) bool {
	return avl.Root.contains(value)
}

// Clear clears the AVL tree and resets the Count
//
// Time complexity: O(1) because we're just changing pointers and values
func (avl *AVLTree[T]) Clear() {
	avl.Root = nil
	avl.Count = 0
}

// Height returns the height of the AVL tree
//
// Time complexity: O(1) because we're just getting the height from AVL.Root, or returning 0 if nil
func (avl *AVLTree[T]) Height() int {
	if avl.Root == nil {
		return 0
	}

	return avl.Root.height
}

// ToArray returns a breadth-first traversal of the tree as an array.
//
// Time complexity: O(n)
//
// Pseudo Code:
// 1. Make a multidimensional array as deep as the tree is tall
// 2. Call a recursive function that accepts a layer value (start at 0), and the array
// 3. In the recursive function, add the node's value to the correct row of the multidimensional array (gotten from the layer int)
// 4. Call the recursive function on the left and right nodes, but pass layer + 1 as the layer value.
// 5. After the recursive call finishes, add each of the arrays to one array from the top layer to the bottom
// 6. Return the array
func (avl *AVLTree[T]) ToArray() []T {
	if avl.Root == nil {
		return nil
	}

	// Get each layer
	layers := make([][]T, avl.Root.height)
	avl.Root.toArray(0, &layers)

	// Make the final array
	var array []T

	// Append each layer to the final array
	for i := 0; i < avl.Root.height; i++ {
		array = append(array, layers[i]...)
	}

	return array
}

// InOrder returns the values of the tree in order (left, middle, right) as a string
//
// Time Complexity: O(n)
// The time complexity is O(n) because we have to traverse every element in the tree
func (avl *AVLTree[T]) InOrder() string {
	// If the tree is empty, return an empty string
	if avl.Root == nil {
		return ""
	}

	values := avl.Root.inOrder()
	var output strings.Builder
	for _, value := range values {
		output.WriteString(fmt.Sprintf("%v ", value))
	}

	// Remove the last space
	return output.String()[0 : output.Len()-1]
}

// PostOrder returns the values of the tree in post-order (left, right, middle) as a string
//
// Time complexity: O(n)
// The time complexity is O(n) because we have to traverse every element of the tree
func (avl *AVLTree[T]) PostOrder() string {
	// If the tree is empty, return an empty string
	if avl.Root == nil {
		return ""
	}

	// Get the values in post-order
	values := avl.Root.postOrder()
	var output strings.Builder
	for _, value := range values {
		output.WriteString(fmt.Sprintf("%v ", value))
	}

	// Remove the last space
	return output.String()[0 : output.Len()-1]
}

// PreOrder returns the values of the tree in pre-order (middle, left, right) as a string
//
// Time Complexity: O(n)
// The time complexity is O(n) because we have to traverse every element of the tree
func (avl *AVLTree[T]) PreOrder() string {
	// If the tre is empty, return an empty string
	if avl.Root == nil {
		return ""
	}

	// Get the values in pre-order
	values := avl.Root.preOrder()
	var output strings.Builder
	for _, value := range values {
		output.WriteString(fmt.Sprintf("%v ", value))
	}

	// Remove the last space
	return output.String()[0 : output.Len()-1]
}
