package queue

import "github.com/robertjshirts/data-structures/linked_list"

type Queue[T comparable] struct {
	list *linked_list.DoubleLinkedList[T]
}

func NewQueue[T comparable]() *Queue[T] {
	list := linked_list.EmptyDoubleLinkedList[T]()
	return &Queue[T]{
		list: &list,
	}
}

// Enqueue adds a value to the end of the queue
//
// # Arguments
// value: The value to add to the queue
//
// Time complexity: O(1)
// Because list.Push is O(1), the time complexity is O(1)
func (q *Queue[T]) Enqueue(value T) {
	q.list.Add(value)
}

// Dequeue removes the first value from the queue and returns it as a pointer
//
// # Returns the first value of the queue, as a pointer. Returns nil if the queue is empty
//
// Time complexity: O(1)
// Because we're just removing the first value in the list, the time complexity is O(1)
//
// Pseudo code:
// if the queue is empty
//
//	return nil
//
// value = list.Remove()
// return value
func (q *Queue[T]) Dequeue() *T {
	if q.list.Head == nil {
		return nil
	}

	value := q.list.Remove()
	return &value
}

// Peek returns the first value of the queue, as a pointer
//
// # Returns the first value of the queue, as a pointer. Returns nil if the queue is empty
//
// Time complexity: O(1)
// Because we're just getting the first value in the list, the time complexity is O(1)
func (q *Queue[T]) Peek() *T {
	if q.list.Head == nil {
		return nil
	}

	return &q.list.Head.Value
}

// Get returns the value at the given index
//
// # Returns the value at the given index, as a pointer. Returns nil if the index is out of bounds
//
// Time complexity: O(n), because q.list.Get is O(n)
func (q *Queue[T]) Get(index int) *T {
	if index < 0 || index >= q.list.Count {
		return nil
	}

	value := q.list.Get(index)
	return &value
}

// Contains returns true if the queue contains the given value, false otherwise
//
// # Returns true if the value is in the queue, false otherwise
//
// Time complexity: O(n), because q.list.Search is O(n)
func (q *Queue[T]) Contains(value T) bool {
	return q.list.Search(value) != -1
}
