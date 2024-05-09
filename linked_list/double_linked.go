package linked_list

import "fmt"

type doubleNode[T comparable] struct {
	Value T
	Next  *doubleNode[T]
	Prev  *doubleNode[T]
}

type DoubleLinkedList[T comparable] struct {
	Head  *doubleNode[T]
	Tail  *doubleNode[T]
	Count int
}

// EmptyDoubleLinkedList creates an empty list with no head
//
// Returns a list with no head or tail and a count of 0.
//
// Big-O is O(1) because we're just instantiating a new node
func EmptyDoubleLinkedList[T comparable]() DoubleLinkedList[T] {
	return DoubleLinkedList[T]{
		Head:  nil,
		Tail:  nil,
		Count: 0,
	}
}

// NewDoubleLinkedList creates a list with a first node
//
// Returns a list with a single node with the value of firstItem.
//
// Big-O is O(1) because we're just instantiating a new node
func NewDoubleLinkedList[T comparable](firstItem T) DoubleLinkedList[T] {
	Node := doubleNode[T]{
		Value: firstItem,
		Next:  nil,
		Prev:  nil,
	}
	return DoubleLinkedList[T]{
		Head:  &Node,
		Tail:  &Node,
		Count: 1,
	}
}

// Add adds a value to the end of a list. Will reassign head and tail if they're nil.
//
// Big-O is O(1) because we just change the tail pointers around
func (s *DoubleLinkedList[T]) Add(item T) {
	newNode := &doubleNode[T]{
		Value: item,
		Next:  nil,
		Prev:  nil,
	}

	if s.Count == 0 {
		s.Head = newNode
		s.Tail = newNode
		s.Count = 1
		return
	}

	newNode.Prev = s.Tail
	s.Tail.Next = newNode
	s.Tail = newNode
	s.Count++
}

// Insert inserts a value at the specified index
//
// Panics if the index is out of bounds.
//
// Big-O is O(n) because s.GetNode is O(n) and the rest of the expressions are constant
func (s *DoubleLinkedList[T]) Insert(value T, index int) {
	// If index out of bounds, panic
	// Don't use s.Count -1 bc we could insert at the tail
	if index < 0 || index > s.Count {
		panic("Index out of bounds")
	}

	// If index is tail, call Add function
	if index == s.Count {
		s.Add(value)
		return
	}

	newNode := &doubleNode[T]{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}

	// If index is head, change accordingly
	if index == 0 {
		newNode.Next = s.Head
		s.Head.Prev = newNode
		s.Head = newNode
		s.Count++
		return
	}

	// Find previous node
	prevNode := s.GetNode(index - 1)
	// Assign newNode's pointers
	newNode.Next = prevNode.Next
	newNode.Prev = prevNode
	// Reassign surrounding node's pointers
	prevNode.Next.Prev = newNode
	prevNode.Next = newNode

	s.Count++
}

// Get returns the value at the specified index
//
// Returns the value at the specified index.
//
// Panics if the index is out of bounds.
//
// Big-O is O(n) because s.GetNode is O(n)
func (s *DoubleLinkedList[T]) Get(index int) T {
	return s.GetNode(index).Value
}

// Remove removes the head node returns its value
//
// Returns the value of the head node.
//
// Panics if the list is empty.
//
// Big-O is O(1) because we're just reassigning the head node
func (s *DoubleLinkedList[T]) Remove() T {
	if s.Count == 0 {
		panic("No elements in list")
	}

	head := s.Head
	// If count is 1, we have to change the tail as well
	if s.Count == 1 {
		s.Head = nil
		s.Tail = nil
		s.Count = 0
		return head.Value
	}

	// Reassign pointers
	s.Head = s.Head.Next
	s.Head.Prev = nil

	s.Count--
	return head.Value
}

// RemoveAt removes a node at the specified index and returns its value
//
// Returns the value of the removed node.
//
// Panics if the index is out of bounds.
//
// Big-O is O(n) because it calls s.GetNode which is O(n)
func (s *DoubleLinkedList[T]) RemoveAt(index int) T {
	// If index is 0 just call remove, don't decrement bc its alr handled
	// No need to panic here because Remove & RemoveLast will panic if index is out of bounds
	if index == 0 {
		return s.Remove()
	} else if index == s.Count-1 {
		return s.RemoveLast()
	}

	// No need to panic here because GetNode will panic if index is out of bounds
	beforeNode := s.GetNode(index - 1)
	removedNode := beforeNode.Next
	// If we have a node after the one we're removing, change its prev pointer
	if removedNode.Next != nil {
		removedNode.Next.Prev = beforeNode
	}
	beforeNode.Next = removedNode.Next
	s.Count--

	return removedNode.Value
}

// RemoveLast removes the tail node and returns its value
//
// Returns the value of the tail node.
//
// Panics if the list is empty.
//
// Big-O is O(1) because we are just reassigning pointers
func (s *DoubleLinkedList[T]) RemoveLast() T {
	if s.Count == 0 {
		panic("No elements in list")
	}

	tail := s.Tail
	// If count is 1, we have to change the head as well
	if s.Count == 1 {
		s.Head = nil
		s.Tail = nil
		s.Count = 0
		return tail.Value
	}

	// Reassign pointers
	s.Tail = s.Tail.Prev
	s.Tail.Next = nil

	s.Count--
	return tail.Value
}

// Clear removes all nodes from the list
//
// Big-O is O(1) because we're just reassigning pointers
func (s *DoubleLinkedList[T]) Clear() {
	s.Head = nil
	s.Tail = nil
	s.Count = 0
}

// Search searches for a value in the list and returns its index
//
// Returns the index of the value if found, -1 otherwise.
//
// Big-O is O(n) because there is a for loop that iterates over the list
func (s *DoubleLinkedList[T]) Search(value T) int {
	currentNode := s.Head
	for i := 0; i < s.Count; i++ {
		if currentNode.Value == value {
			return i
		}
		currentNode = currentNode.Next
	}
	return -1
}

// GetNode finds and returns the node at the specified index
//
// Returns a pointer to the node at the specified index.
//
// Panics if the index is out of bounds.
//
// Big-O is O(n) because there is a for loop that iterates over the list
func (s *DoubleLinkedList[T]) GetNode(index int) *doubleNode[T] {
	if index < 0 || index >= s.Count {
		panic("Index out of bounds")
	}

	node := s.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node
}

// ToString converts the list to a string
//
// Returns a string representation of the list in the following format: "1 2 3 4".
// Will return an empty string if the list is empty, not a nil pointer.
//
// Big-O is O(n) because it will have to iterate through the list
func (s *DoubleLinkedList[T]) ToString() string {
	// Add this here so we can use the substring at the bottom w/o panicking
	if s.Head == nil {
		return ""
	}

	result := ""
	currentNode := s.Head
	for currentNode != nil {
		result += fmt.Sprintf("%v ", currentNode.Value)
		currentNode = currentNode.Next
	}
	// Return a substring to remove the trailing space
	return result[:len(result)-1]
}
