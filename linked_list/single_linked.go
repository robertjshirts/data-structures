package linked_list

import "fmt"

type node[T comparable] struct {
	Value T
	Next  *node[T]
}

type SingleLinkedList[T comparable] struct {
	Head  *node[T]
	Count int
}

// EmptySingleLinkedList creates an empty list with no head
//
// # Returns a list with no head and a count of 0
//
// Big-O is O(1) because we're just instantiating a new node
func EmptySingleLinkedList[T comparable]() SingleLinkedList[T] {
	return SingleLinkedList[T]{
		Head:  nil,
		Count: 0,
	}
}

// NewSingleLinkedList creates a list with a first node
//
// # Returns a list with a single node with the value of firstItem
//
// Big-O is O(1) because we're just instantiating a new node
func NewSingleLinkedList[T comparable](firstItem T) SingleLinkedList[T] {
	Head := node[T]{
		Value: firstItem,
		Next:  nil,
	}
	return SingleLinkedList[T]{
		Head:  &Head,
		Count: 1,
	}
}

// Push adds an item to the start of the list
//
// Big-O is O(1) because it's just reassigning the head
func (s *SingleLinkedList[T]) Push(item T) {
	newNode := node[T]{
		Value: item,
		Next:  s.Head,
	}
	s.Head = &newNode
	s.Count++
}

// Add adds a value to the end of the list. Will reassign head if it's nil.
//
// Big-O is O(n) because it calls s.GetNode which is O(n) and the rest is just changing around pointers
func (s *SingleLinkedList[T]) Add(item T) {
	newNode := node[T]{
		Value: item,
		Next:  nil,
	}

	if s.Head == nil {
		s.Head = &newNode
		s.Count = 1
		return
	}

	lastNode := s.GetNode(s.Count - 1)
	lastNode.Next = &newNode
	s.Count++
}

// Insert inserts a value at the specified index
//
// # Panics if the index is out of bounds
//
// Big-O is O(n) because s.GetNode is O(n) and the rest is just reassigning pointers
func (s *SingleLinkedList[T]) Insert(value T, index int) {
	// If index is out of bounds, panic
	if index > s.Count || index < 0 {
		panic("Index out of bounds!")
	}

	// If insert at index 0, change head
	if index == 0 {
		s.Head = &node[T]{
			Value: value,
			Next:  s.Head,
		}
		return
	}

	newNode := &node[T]{
		Value: value,
		Next:  nil,
	}

	// Find previous node
	prevNode := s.GetNode(index - 1)
	// Assign newNode's pointer
	newNode.Next = prevNode.Next
	// Assign previousNode's pointer
	prevNode.Next = newNode

	s.Count++
}

// Get returns the value at the specified index
//
// # Returns the value at the specified index
//
// # Panics if the index is out of bounds
//
// Big-O is O(n) because s.GetNode is O(n) and the rest are constant
func (s *SingleLinkedList[T]) Get(index int) T {
	return s.GetNode(index).Value
}

// Remove removes the head node and returns its value
//
// # Returns the value of the head node
//
// # Panics if the list is empty
//
// Big-O is O(1) because it's just reassigning the head
func (s *SingleLinkedList[T]) Remove() T {
	if s.Head == nil {
		panic("No element")
	}
	head := *s.Head
	s.Head = head.Next
	s.Count--
	return head.Value
}

// RemoveAt removes the node at the specified index and returns its value
//
// # Returns the value of the node at the specified index
//
// Panics if the index is out of bounds.
//
// Big-O is O(n) because it calls s.GetNode which is O(n) and the rest is just reassigning pointers
func (s *SingleLinkedList[T]) RemoveAt(index int) T {
	// No need to panic explicitly because GetNode & Remove will do it
	if index == 0 {
		// If index is 0, just call Remove
		// Don't decrement because Remove does it
		return s.Remove()
	}

	beforeNode := s.GetNode(index - 1)
	removedNode := beforeNode.Next
	beforeNode.Next = removedNode.Next
	s.Count--

	return removedNode.Value
}

// RemoveLast removes the last node and returns its value
//
// # Returns the value of the last node
//
// # Panics if the list is empty
//
// Big-O is O(n) because it calls s.RemoveAt which is O(n)
func (s *SingleLinkedList[T]) RemoveLast() T {
	// Don't decrement bc RemoveAt does it
	if s.Count == 0 {
		panic("No elements in list")
	}
	return s.RemoveAt(s.Count - 1)
}

// Clear clears the list by setting the head to nil and count to 0
//
// Big-O is O(1) because it's just reassigning the head and count
func (s *SingleLinkedList[T]) Clear() {
	s.Head = nil
	s.Count = 0
}

// Search searches for a value in the list and returns its index
//
// # Returns the index of the value if found, -1 if not found
//
// Big-O is O(n) because it will have to for loop through each
func (s *SingleLinkedList[T]) Search(value T) int {
	currentNode := s.Head
	for i := 0; i < s.Count; i++ {
		if currentNode.Value == value {
			// If found, return the index
			return i
		}
		currentNode = currentNode.Next
	}

	// If not found, return -1
	return -1
}

// GetNode finds and returns the node at the specified index
//
// # Returns a pointer to the node at the specified index
//
// # Panics if the index is out of bounds
//
// Big-O is O(n) because it will have to iterate through the list
func (s *SingleLinkedList[T]) GetNode(index int) *node[T] {
	if index > (s.Count-1) || index < 0 {
		panic("Index out of bounds!")
	}

	// Find the node we want by iterating through until
	// we've done it index times
	currentNode := s.Head
	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}

	return currentNode
}

// ToString converts the list to a string
//
// Returns a string representation of the list in the following format: "1 2 3 4".
// Will return an empty string if the list is empty, not a nil pointer.
//
// Big-O is O(n) because it will have to iterate through the list
func (s *SingleLinkedList[T]) ToString() string {
	var result string
	currentNode := s.Head
	for currentNode != nil {
		result += fmt.Sprintf("%v ", currentNode.Value)
		currentNode = currentNode.Next
	}
	result += "\n"
	return result
}
