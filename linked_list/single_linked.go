package linked_list

import "fmt"

type node[T comparable] struct {
	Value T
	Next  *node[T]
}

type singleLinkedList[T comparable] struct {
	Head  *node[T]
	Count int
}

// Creates an empty list with no head
// Big-O is O(1) because the only operation we're doing is returning a struct
func EmptySingleLinkedList[T comparable]() singleLinkedList[T] {
	return singleLinkedList[T]{
		Head:  nil,
		Count: 0,
	}
}

// Creates a list with a first node
// Big-) is O(1) because there aren'T comparable for loops or data to iterate over
func NewSingleLinkedList[T comparable](firstItem T) singleLinkedList[T] {
	Head := node[T]{
		Value: firstItem,
		Next:  nil,
	}
	return singleLinkedList[T]{
		Head:  &Head,
		Count: 1,
	}
}

// Adds a value to the END list
// Big-O is O(n) becasue s.GetNode is O(n) and the rest are constant
func (s *singleLinkedList[T]) Add(item T) {
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

// Inserts a value at the specified index
// Big-O is O(n) because s.GetNode is O(n) and the rest of the expressions are constnant
func (s *singleLinkedList[T]) Insert(value T, index int) {
	// If insert at index 0, change head
	if index == 0 {
		s.Head = &node[T]{
			Value: value,
			Next:  s.Head,
		}
		return
	}

	newNode := node[T]{
		Value: value,
		Next:  nil,
	}

	// Pass one less than the index, we are trying to find the node
	// before our insert so we can change that node's
	currentNode := s.GetNode(index - 1)
	s.Count++

	// Assign current node;s pointer to our new node's pointer
	// So now the new node points to where the currentNode used to point
	newNode.Next = currentNode.Next

	// Assign currentNode's pointer to the new node we just inserted
	// Now currentNode points to newNode points to (the old) currentNode.Next
	currentNode.Next = &newNode
}

// Get the value at the specified index, and return it
// Big-O is O(n) becuase s.GetNode is O(n) and the rest are constant
func (s *singleLinkedList[T]) Get(index int) T {
	return s.GetNode(index).Value
}

// Return the value of the head and remove it from the list
// Big-O is O(1) because all the operations are constant, and there are no for loops or iterations
func (s *singleLinkedList[T]) Remove() T {
	if s.Head == nil {
		panic("No element")
	}
	head := *s.Head
	s.Head = head.Next
	s.Count--
	return head.Value
}

// Return the value at specified index and remove it from the list
// Big-O is O(n) because it calls s.GetNode which is O(n)
func (s *singleLinkedList[T]) RemoveAt(index int) T {
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

// Returns the value at the end of the list and removes it
// Big-O is O(n) because it calls s.RemoveAt which is O(n)
func (s *singleLinkedList[T]) RemoveLast() T {
	// Don't decrement bc RemoveAt does it
	return s.RemoveAt(s.Count - 1)
}

// Clears the list
// Big-O is O(1) because it's just setting the head to nil
func (s *singleLinkedList[T]) Clear() {
	s.Head = nil
	s.Count = 0
}

// Searches for the value in the list and returns the index
// Big-O is O(n) because it will have to iterate through the list
func (s *singleLinkedList[T]) Search(value T) int {
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

// Get the node at the specified index, and return it
// Big-O is O(n) because it will have to for loop through each
// element to possibly get the last one
func (s *singleLinkedList[T]) GetNode(index int) *node[T] {
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

func (s *singleLinkedList[T]) ToString() string {
	var result string
	currentNode := s.Head
	for currentNode != nil {
		result += fmt.Sprintf("%v ", currentNode.Value)
		currentNode = currentNode.Next
	}
	result += "\n"
	return result
}
