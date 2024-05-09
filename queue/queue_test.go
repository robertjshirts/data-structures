package queue

import "testing"

func TestNewQueue(t *testing.T) {
	// Arrange
	// Act
	queue := NewQueue[int]()
	// Assert
	if queue == nil {
		t.Errorf("Expected queue to not be nil")
	}
}

func TestQueue_EnqueueAddsValue(t *testing.T) {
	// Arrange
	queue := NewQueue[int]()
	want := 1
	// Act
	queue.Enqueue(want)
	got := queue.list.Head.Value
	// Assert
	simpleAssert(t, got, want)
}

func TestQueue_EnqueueAddsSecondValue(t *testing.T) {
	// Arrange
	queue := NewQueue[string]()
	want := "second added"
	// Act
	queue.Enqueue("first added")
	queue.Enqueue(want)
	// Get the last value
	got := queue.list.Tail.Value
	// Assert
	simpleAssert(t, got, want)
}

func TestQueue_DequeueRemovesFirstValue(t *testing.T) {
	// Arrange
	queue := NewQueue[int]()
	want := 1
	queue.Enqueue(want)
	// Act
	got := queue.Dequeue()
	// Assert
	simpleAssert(t, *got, want)
}

func TestQueue_DequeueRemovesSecondValue(t *testing.T) {
	// Arrange
	queue := NewQueue[int]()
	want := 2
	queue.Enqueue(1)
	queue.Enqueue(want)
	queue.Enqueue(3)
	queue.Enqueue(4)
	// Act
	// Dequeue the first value
	queue.Dequeue()
	got := *queue.Dequeue()
	// Assert
	simpleAssert(t, got, want)
}

func TestQueue_DequeueReturnsNilOnNoValues(t *testing.T) {
	// Arrange
	queue := NewQueue[int]()
	// Act
	got := queue.Dequeue()
	// Assert
	nilAssert(t, got)
}

func TestQueue_PeekReturnsFirstValue(t *testing.T) {
	// Arrange
	queue := NewQueue[int]()
	want := 1
	queue.Enqueue(want)
	// Act
	got := queue.Peek()
	// Assert
	simpleAssert(t, *got, want)
}

func TestQueue_PeekReturnsSecondValue(t *testing.T) {
	// Arrange
	queue := NewQueue[int]()
	want := 2
	queue.Enqueue(1)
	queue.Enqueue(want)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)
	// Act
	// Dequeue the first value
	queue.Dequeue()
	got := *queue.Peek()
	// Assert
	simpleAssert(t, got, want)
}

func TestQueue_PeekReturnsNilOnNoValues(t *testing.T) {
	// Arrange
	queue := NewQueue[int]()
	// Act
	got := queue.Peek()
	// Assert
	nilAssert(t, got)
}

func TestQueue_GetReturnsValueAtFirstIndex(t *testing.T) {
	// Arrange
	queue := NewQueue[int]()
	want := 1
	queue.Enqueue(want)
	// Act
	got := queue.Get(0)
	// Assert
	simpleAssert(t, *got, want)
}

func TestQueue_GetReturnsValueAtThirdIndex(t *testing.T) {
	// Arrange
	queue := NewQueue[int]()
	want := 2
	queue.Enqueue(0)
	queue.Enqueue(1)
	queue.Enqueue(want)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)
	// Act
	// Dequeue the first value
	got := *queue.Get(2)
	// Assert
	simpleAssert(t, got, want)
}

func TestQueue_GetReturnsNilOnOutOfBoundsIndex(t *testing.T) {
	// Arrange
	queue := NewQueue[int]()
	queue.Enqueue(0)
	queue.Enqueue(1)
	// Act
	got := queue.Get(2)
	// Assert
	nilAssert(t, got)
}

func TestQueue_GetReturnsNilOnNegativeIndex(t *testing.T) {
	// Arrange
	queue := NewQueue[int]()
	queue.Enqueue(0)
	queue.Enqueue(1)
	// Act
	got := queue.Get(-1)
	// Assert
	nilAssert(t, got)
}

func TestQueue_GetReturnsNilOnEmptyQueue(t *testing.T) {
	// Arrange
	queue := NewQueue[int]()
	// Act
	got := queue.Get(0)
	// Assert
	nilAssert(t, got)
}

func TestQueue_ContainsReturnsTrue(t *testing.T) {
	// Arrange
	queue := NewQueue[int]()
	queue.Enqueue(0)
	queue.Enqueue(1)
	// Act
	got := queue.Contains(1)
	// Assert
	if !got {
		t.Errorf("Expected true, got false")
	}
}

func TestQueue_ContainsReturnsFalse(t *testing.T) {
	// Arrange
	queue := NewQueue[int]()
	queue.Enqueue(0)
	queue.Enqueue(1)
	// Act
	got := queue.Contains(2)
	// Assert
	if got {
		t.Errorf("Expected false, got true")
	}
}

func TestQueue_ContainsReturnsFalseOnEmptyQueue(t *testing.T) {
	// Arrange
	queue := NewQueue[int]()
	// Act
	got := queue.Contains(0)
	// Assert
	if got {
		t.Errorf("Expected false, got true")
	}
}

func nilAssert[T comparable](t *testing.T, got *T) {
	if got != nil {
		t.Errorf("Expected nil, got %v", *got)
	}
}

func simpleAssert[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
