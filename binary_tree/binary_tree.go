package binary_tree

import (
	"fmt"
	"strings"
)

type types interface {
	~int | ~string | ~float64
}

type node[T types] struct {
	value T
	left  *node[T]
	right *node[T]
}

type BinaryTree[T types] struct {
	Root  *node[T]
	Count int
}

func EmptyBinaryTree[T types]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

func NewBinaryTree[T types](value T) *BinaryTree[T] {
	return &BinaryTree[T]{
		Root: &node[T]{
			value: value,
		},
		Count: 1,
	}
}

// Insert inserts a node in the tree
//
// Time complexity: O(log n)
// The time complexity is O(log n) because we typically only traverse half the tree
func (bt *BinaryTree[T]) Insert(value T) {
	// Create a new node
	newNode := &node[T]{
		value: value,
	}

	// If the tree is empty, set the root to the new node
	if bt.Root == nil {
		bt.Root = newNode
		bt.Count = 1
		return
	}

	// Start at the root
	currentNode := bt.Root
	for {
		if value > currentNode.value {
			// if value is more, move right
			if currentNode.right == nil {
				// If node is nil, insert
				currentNode.right = newNode
				bt.Count++
				return
			}
			// Else keep going deeper
			currentNode = currentNode.right
		} else {
			// if the value is less than current, move left
			if currentNode.left == nil {
				// If node is nil, insert
				currentNode.left = newNode
				bt.Count++
				return
			}
			// Else keep going deeper
			currentNode = currentNode.left
		}
	}
}

// Clear removes all nodes from the tree
//
// Time complexity: O(1)
// Because we are just setting the root to nil, the time complexity is O(1)
func (bt *BinaryTree[T]) Clear() {
	bt.Root = nil
	bt.Count = 0
}

// ToArray returns the values of the tree as a slice, in order (left, middle, right)
//
// Time complexity: O(n)
// Because we have to visit every node, the time complexity is O(n)
func (bt *BinaryTree[T]) ToArray() []T {
	return bt.inOrder(bt.Root)
}

// InOrder returns the values of the tree in order (left, middle, right) as a string
//
// Time complexity: O(n)
// Because we have to visit every node, the time complexity is O(n)
func (bt *BinaryTree[T]) InOrder() string {
	// If the tree is empty, return an empty string
	if bt.Root == nil {
		return ""
	}

	// Get the values in order
	values := bt.inOrder(bt.Root)
	var output strings.Builder
	for _, value := range values {
		output.WriteString(fmt.Sprintf("%v ", value))
	}

	// Remove the last space
	return output.String()[0 : output.Len()-1]
}

func (bt *BinaryTree[T]) inOrder(node *node[T]) []T {
	// Stop case
	if node == nil {
		return nil
	}

	// Get left and right nodes
	left := bt.inOrder(node.left)
	right := bt.inOrder(node.right)
	value := node.value

	// Return left and right nodes with current node value
	return append(append(left, value), right...)
}

// PreOrder returns the values of the tree in pre-order (middle, left, right) as a string
//
// Time complexity: O(n)
// Because we have to visit every node, the time complexity is O(n)
func (bt *BinaryTree[T]) PreOrder() string {
	// If the tree is empty, return an empty string
	if bt.Root == nil {
		return ""
	}

	// Get the values in pre-order
	values := bt.preOrder(bt.Root)
	var output strings.Builder
	for _, value := range values {
		output.WriteString(fmt.Sprintf("%v ", value))
	}

	// Remove the last space
	return output.String()[0 : output.Len()-1]
}

func (bt *BinaryTree[T]) preOrder(node *node[T]) []T {
	if node == nil {
		return nil
	}

	// Get left and right nodes
	left := bt.preOrder(node.left)
	right := bt.preOrder(node.right)

	return append(append([]T{node.value}, left...), right...)
}

// PostOrder returns the values of the tree in post-order (left, right, middle) as a string
//
// Time complexity: O(n)
// The time complexity is O(n) because we have to traverse every element of the tree
func (bt *BinaryTree[T]) PostOrder() string {
	// If the tree is empty, return an empty string
	if bt.Root == nil {
		return ""
	}

	// Get the values in post-order
	values := bt.postOrder(bt.Root)
	var output strings.Builder
	for _, value := range values {
		output.WriteString(fmt.Sprintf("%v ", value))
	}

	// Remove the last space
	return output.String()[0 : output.Len()-1]
}

func (bt *BinaryTree[T]) postOrder(node *node[T]) []T {
	if node == nil {
		return nil
	}

	// Get left and right nodes
	left := bt.postOrder(node.left)
	right := bt.postOrder(node.right)

	return append(append(left, right...), node.value)
}

// Remove removes a node
//
// Time complexity: O(log n)
// The time complexity is O(log n) because we typically only traverse half the tree
// But in the worst case, it can be O(n)
// Psuedo code:
// If 0 children - set parent left/right pointer to null
// if 1 child - set parent left/right pointer to child
// if 2 children
//
//	If deleting left/right find largest/smallest (respectively)
//	replace deleted node with that value
//	call remove on left/right with value (recurse)
func (bt *BinaryTree[T]) Remove(value T) {
	bt.Root = bt.remove(bt.Root, value)
}

// Revised Psuedo code:
// This is a recursive function that returns the new reference to the node
// If a node with no children is getting removed, it will return a nil reference
// If a node with one child is getting removed, it will return the reference to the child
// If a node with two children is getting removed, it will find the smallest value in the right subtree, replace the current node with that value, and call remove on the right subtree with that value
//
// if node is nil, return nil
// if value = node.value
//
//	if node has no children, return nil
//	if node has one child, return that child
//	if node has two children
//		find the smallest value in right subtree
//		replace current node with that value
//		call remove on right subtree with that value
//
// Move closer to the value
// if value > node.value, call remove on right subtree
// if value < node.value, call remove on left subtree
func (bt *BinaryTree[T]) remove(node *node[T], value T) *node[T] {
	// Stop case
	if node == nil {
		return nil
	}

	if value == node.value {
		// If node has no children
		if node.left == nil && node.right == nil {
			return nil
		}

		// If node has one child
		if node.left == nil {
			return node.right
		}

		if node.right == nil {
			return node.left
		}

		// If node has two children
		smallestValue := bt.findSmallest(node.right)
		node.value = smallestValue
		node.right = bt.remove(node.right, smallestValue)
		return node
	}

	// If value is greater than current node, go right
	if value > node.value {
		node.right = bt.remove(node.right, value)
		return node
	}

	// If value is less than current node, go left
	node.left = bt.remove(node.left, value)
	return node
}

// findSmallest finds the smallest value in the tree
//
// Time complexity: O(log n)
// The time complexity is O(log n) because we typically only traverse half the tree
func (bt *BinaryTree[T]) findSmallest(node *node[T]) T {
	if node.left == nil {
		return node.value
	}
	return bt.findSmallest(node.left)
}

// Contains checks if a value is in the tree
//
// # Returns true if the value is in the tree, false otherwise
//
// Time complexity: O(log n)
// Time complexity is O(log n) because we typically only traverse half the tree
func (bt *BinaryTree[T]) Contains(value T) bool {
	return bt.contains(bt.Root, value)
}

func (bt *BinaryTree[T]) contains(node *node[T], value T) bool {
	// Stop case
	if node == nil {
		return false
	}

	if value == node.value {
		return true
	}

	if value > node.value {
		return bt.contains(node.right, value)
	}

	return bt.contains(node.left, value)
}

// Height returns the height of the tree
//
// # Returns the height of the tree
//
// Time complexity: O(n)
// The time complexity is O(n) because we have to visit every node
func (bt *BinaryTree[T]) Height() int {
	return bt.height(bt.Root)
}

func (bt *BinaryTree[T]) height(node *node[T]) int {
	if node == nil {
		return 0
	}

	left := bt.height(node.left)
	right := bt.height(node.right)

	if left > right {
		return left + 1
	}
	return right + 1
}
