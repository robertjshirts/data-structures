package binary_tree

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
// Because the tree is balanced, the time complexity is O(log n)
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

// InOrder returns the values of the tree in order
//
// # Returns a slice of the tree in order
//
// Time complexity: O(n)
// Because we have to visit every node, the time complexity is O(n)
func (bt *BinaryTree[T]) InOrder() []T {
	return bt.inOrder(bt.Root)
}

func (bt *BinaryTree[T]) inOrder(node *node[T]) []T {
	// Stop case
	if node == nil {
		return nil
	}

	// Get left and right nodes
	left := bt.inOrder(node.left)
	right := bt.inOrder(node.right)

	// Return left and right nodes with current node value
	return append(append(left, node.value), right...)
}

// Remove removes a node
//
// Psuedo code:
// If 0 children - set parent left/right pointer to null
// if 1 child - set parent left/right pointer to child
// if 2 children
//   if deleling left/right find largest/smallest (respectively)
//   replace deleted node with that value
//   call remove on left/right with value (recurse)
