package binary_tree

import "testing"

func TestBTNewCreatesRoot(t *testing.T) {
	// Arrange
	expected := 10
	// Act
	bt := NewBinaryTree[int](expected)
	// Assert
	simpleAssert(t, expected, bt.Root.value)
}

func TestBTNewCorrectLength(t *testing.T) {
	// Arrange
	expected := 1
	// Act
	bt := NewBinaryTree[int](10)
	// Assert
	simpleAssert(t, expected, bt.Count)
}

func TestBTInsertAddsNodeLeft(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	expected := 5
	// Act
	bt.Insert(expected)
	// Assert
	simpleAssert(t, expected, bt.Root.left.value)
}

func TestBTInsertChangesCount(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	expected := 3
	// Act
	bt.Insert(5)
	bt.Insert(15)
	// Assert
	simpleAssert(t, expected, bt.Count)
}

func TestBTInsertAddsNodeRight(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	expected := 15
	// Act
	bt.Insert(expected)
	// Assert
	simpleAssert(t, expected, bt.Root.right.value)
}

func TestBTInsertAddsDuplicate(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	expected := 10
	// Act
	bt.Insert(expected)
	// Assert
	simpleAssert(t, expected, bt.Root.left.value)
}

func TestBTInOrder(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	bt.Insert(5)
	bt.Insert(15)
	bt.Insert(3)
	bt.Insert(7)
	bt.Insert(12)
	bt.Insert(17)
	expected := []int{3, 5, 7, 10, 12, 15, 17}
	// Act
	got := bt.InOrder()
	// Assert
	for i, v := range got {
		simpleAssert(t, expected[i], v)
	}
}

func TestBTInOrderEmpty(t *testing.T) {
	// Arrange
	bt := EmptyBinaryTree[int]()
	// Act
	got := bt.InOrder()
	// Assert
	simpleAssert(t, 0, len(got))
}

func simpleAssert[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
