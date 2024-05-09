package stack

import "testing"

func TestNewStack(t *testing.T) {
	stack := NewStack[int]()
	if stack == nil {
		t.Errorf("Expected stack to not be nil")
	}
}

func TestStack_PushPushesValue(t *testing.T) {
	// Arrange
	stack := NewStack[int]()
	want := 1
	// Act
	stack.Push(want)
	got := stack.list.Get(0)
	// Assert
	simpleAssert(t, got, want)
}

func TestStack_PushPushesSecondValue(t *testing.T) {
	// Arrange
	stack := NewStack[int]()
	want := 1
	// Act
	stack.Push(0)
	stack.Push(want)
	// Get 0 bc it was pushed second, so it's at the top of the stack
	got := stack.list.Get(0)
	// Assert
	simpleAssert(t, got, want)
}

func TestStack_PeekReturnsTopValue(t *testing.T) {
	// Arrange
	stack := NewStack[int]()
	want := 1
	stack.Push(0)
	stack.Push(want)
	// Act
	got := *stack.Peek()
	// Assert
	simpleAssert(t, got, want)
}

func TestStack_PeekReturnsNilOnNoValues(t *testing.T) {
	// Arrange
	stack := NewStack[int]()
	// Act
	got := stack.Peek()
	// Assert
	nilAssert(t, got)
}

func TestStack_GetReturnsValueAtTop(t *testing.T) {
	// Arrange
	stack := NewStack[int]()
	want := 1
	index := 0 // At index 0 bc we push what we want AFTER the other value
	stack.Push(0)
	stack.Push(want)
	// Act
	got := *stack.Get(index)
	// Assert
	simpleAssert(t, got, want)
}

func TestStack_GetReturnsValue(t *testing.T) {
	// Arrange
	stack := NewStack[int]()
	stack.Push(0)
	want := 1
	index := 1
	stack.Push(want)
	stack.Push(2)
	// Act
	got := *stack.Get(index)
	// Assert
	simpleAssert(t, got, want)
}

func TestStack_GetReturnsNilOnOutOfBoundsIndex(t *testing.T) {
	// Arrange
	stack := NewStack[int]()
	stack.Push(0)
	stack.Push(1)
	// Act
	got := stack.Get(2)
	// Assert
	nilAssert(t, got)
}

func nilAssert[T any](t *testing.T, got *T) {
	if got != nil {
		t.Errorf("Got %v, wanted nil", got)
	}
}

func simpleAssert[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
