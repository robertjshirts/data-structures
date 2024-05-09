package stack

import "github.com/robertjshirts/data-structures/linked_list"

type Stack[T comparable] struct {
	list *linked_list.SingleLinkedList[T]
}

func NewStack[T comparable]() *Stack[T] {
	list := linked_list.EmptySingleLinkedList[T]()
	return &Stack[T]{
		list: &list,
	}
}

// Push adds a value to the top of the stack
//
// # Arguments
// value: The value to add to the stack
//
// Time complexity: O(1)
// Because list.Push is O(1), the time complexity is O(1)
func (s *Stack[T]) Push(value T) {
	s.list.Push(value)
}

// Pop removes the top value from the stack and returns it as a pointer
//
// # Returns the top value of the stack, as a pointer. Returns nil if the stack is empty
//
// Time complexity: O(1)
// Because we're just removing the first value in the list, the time complexity is O(1)
// Also s.list.Remove() is O(1) because we're just removing the first value in the list
func (s *Stack[T]) Pop() *T {
	if s.list.Head == nil {
		return nil
	}

	value := s.list.Remove()
	return &value
}

// Peek returns the top value of the stack, as a pointer
//
// # Returns the top value of the stack, as a pointer. Returns nil if the stack is empty
//
// Time complexity: O(1)
// Because we're just getting the first value in the list, the time complexity is O(1)
func (s *Stack[T]) Peek() *T {
	if s.list.Head == nil {
		return nil
	}

	return &s.list.Head.Value
}

// Get returns the value at the given index
//
// # Returns the value at the given index, as a pointer. Returns nil if the index is out of bounds
//
// Time complexity: O(n)
// Because it calls list.Get(), which has a complexity of O(n)
func (s *Stack[T]) Get(index int) *T {
	if index < 0 || index >= s.list.Count {
		return nil
	}

	value := s.list.Get(index)
	return &value
}

// Contains checks if the stack contains a value
//
// # Returns true if the value is in the stack, false otherwise
//
// Time complexity: O(n)
// Because we have to check every value in the stack, the time complexity is O(n)
//
// Pseudo code:
// if the stack is empty, return false
// if list.count == 0
//
//	return false
//
// return s.List.Search(value) != -1
func (s *Stack[T]) Contains(value T) bool {
	if s.list.Count == 0 {
		return false
	}
	return s.list.Search(value) != -1
}
