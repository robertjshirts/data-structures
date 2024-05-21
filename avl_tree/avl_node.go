package avl_tree

type types interface {
	~int | ~string | ~float64
}

type node[T types] struct {
	value  T
	left   *node[T]
	right  *node[T]
	height int
}

// insert inserts a node in the tree, recursively. returns the new node to be assigned to the parent node
//
// Time complexity: O(log n)
// The time complexity is O(log n) because we typically only traverse half the tree when inserting, and balancing is all O(1).
func (n *node[T]) insert(value T) *node[T] {
	//Stop case
	if n == nil {
		return &node[T]{value: value, height: 1}
	}

	// Traverse further down
	if value < n.value {
		n.left = n.left.insert(value)
	} else {
		n.right = n.right.insert(value)
	}

	// ON THE WAY BACK UP
	// Balance node and return it.
	return n.balanceNode()
}

// remove removes a node from the tree, recursively. Returns the new node to be assigned for the parent node
func (n *node[T]) remove(value T) *node[T] {
	if n == nil {
		return nil
	}

	if value == n.value {
		// if node has no children
		if n.left == nil && n.right == nil {
			return nil
		}

		if n.left == nil {
			return n.right
		}

		if n.right == nil {
			return n.left
		}

		smallest := findSmallest(n.right)
		// Reassign current node to the smallest from the right subtree
		n.value = smallest.value
		// Remove the copied value from the subtree by recursively calling the algo
		n.right = n.right.remove(smallest.value)
	}

	// If value is greater than the current node, recursively call right
	if value > n.value {
		n.right = n.right.remove(value)
		return n.balanceNode()
	}

	n.left = n.left.remove(value)
	return n.balanceNode()
}

// balanceNode checks the balance factor of the subtree from itself, and rotates as necessary
//
// Time complexity: O(1) because we are just checking the height field and changing pointers around.
func (n *node[T]) balanceNode() *node[T] {
	// Calc balance
	balance := height(n.left) - height(n.right)

	// balance > 1 means left is heavier
	// balance < -1 means right is heavier

	// Left Left case
	if balance > 1 && height(n.left.left) >= height(n.left.right) {
		return n.rotateRight()
	}

	// Left Right case
	if balance > 1 && height(n.left.left) < height(n.left.right) {
		n.left = n.left.rotateLeft()
		return n.rotateRight()
	}

	// Right Right case
	if balance < -1 && height(n.right.right) >= height(n.right.left) {
		return n.rotateLeft()
	}

	// Right Left case
	if balance < -1 && height(n.right.right) < height(n.right.left) {
		n.right = n.right.rotateRight()
		return n.rotateLeft()
	}

	// Calc height
	n.height = max(height(n.left), height(n.right)) + 1
	return n
}

// contains checks if a value is in the tree, recursively
//
// Time complexity: O(log n)
// The time complexity is O(log n) because we typically only traverse half the tree
func (n *node[T]) contains(value T) bool {
	if n == nil {
		return false
	}

	if value == n.value {
		return true
	}

	if value > n.value {
		return n.right.contains(value)
	}

	return n.left.contains(value)
}

// toArray returns a breadth-first traversal of the tree as an array
//
// Time complexity: O(n) because we have to visit every node
func (n *node[T]) toArray(layer int, array *[][]T) {
	// Stop case
	if n == nil {
		return
	}

	// Add to correct layer's array
	(*array)[layer] = append((*array)[layer], n.value)

	// Do the same, incrementing the layer
	n.left.toArray(layer+1, array)
	n.right.toArray(layer+1, array)
}

// findSmallest finds the smallest node in a given subtree, based on the provided node.
//
// Big-O: O(log n) because we only have to traverse half the tree.
func findSmallest[T types](n *node[T]) *node[T] {
	if n.right == nil {
		return n
	}

	return findSmallest(n.right)
}

// inOrder recursively retrieves every value in the tree, and returns it in order. (left, root, right)
func (n *node[T]) inOrder() []T {
	if n == nil {
		return nil
	}

	// Get left and right nodes
	left := n.left.inOrder()
	right := n.right.inOrder()

	return append(append(left, n.value), right...)
}

// preOrder recursively retrieves every value in the tree, and returns it in preorder. (root, left, right)
func (n *node[T]) preOrder() []T {
	if n == nil {
		return nil
	}

	// Get left and right nodes
	left := n.left.preOrder()
	right := n.right.preOrder()

	return append(append([]T{n.value}, left...), right...)
}

// postOrder recursively retrieves every value in the tree, and returns it in postorder. (left, right, root)
func (n *node[T]) postOrder() []T {
	if n == nil {
		return nil
	}

	// Get left and right nodes
	left := n.left.postOrder()
	right := n.right.postOrder()

	return append(append(left, right...), n.value)
}

// height is a utility function that just returns 0 for height if the node is nil, else returns the node's height.
func height[T types](n *node[T]) int {
	if n == nil {
		return 0
	}

	return n.height
}

func (n *node[T]) rotateLeft() *node[T] {
	pivot := n.right
	n.right = pivot.left
	pivot.left = n
	// Update heights
	n.height = max(height(n.right), height(n.left)) + 1
	pivot.height = max(height(pivot.right), height(pivot.left)) + 1
	return pivot
}

func (n *node[T]) rotateRight() *node[T] {
	pivot := n.left
	n.left = pivot.right
	pivot.right = n
	// Update heights
	n.height = max(height(n.right), height(n.left)) + 1
	pivot.height = max(height(pivot.right), height(pivot.left)) + 1
	return pivot
}
