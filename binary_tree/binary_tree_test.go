package binary_tree

import (
	"github.com/robertjshirts/data-structures/util"
	"testing"
)

func TestBTNewCreatesRoot(t *testing.T) {
	// Arrange
	expected := 10
	// Act
	bt := NewBinaryTree[int](expected)
	// Assert
	util.SimpleAssert(t, expected, bt.Root.value)
}

func TestBTNewCorrectLength(t *testing.T) {
	// Arrange
	expected := 1
	// Act
	bt := NewBinaryTree[int](10)
	// Assert
	util.SimpleAssert(t, expected, bt.Count)
}

func TestBTInsertAddsNodeLeft(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	expected := 5
	// Act
	bt.Insert(expected)
	// Assert
	util.SimpleAssert(t, expected, bt.Root.left.value)
}

func TestBTInsertChangesCount(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	expected := 3
	// Act
	bt.Insert(5)
	bt.Insert(15)
	// Assert
	util.SimpleAssert(t, expected, bt.Count)
}

func TestBTInsertAddsNodeRight(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	expected := 15
	// Act
	bt.Insert(expected)
	// Assert
	util.SimpleAssert(t, expected, bt.Root.right.value)
}

func TestBTInsertAddsDuplicate(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	expected := 10
	// Act
	bt.Insert(expected)
	// Assert
	util.SimpleAssert(t, expected, bt.Root.left.value)
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
	want := "3 5 7 10 12 15 17"
	// Act
	got := bt.InOrder()
	// Assert
	util.SimpleAssert(t, got, want)
}

func TestBTInOrderEmpty(t *testing.T) {
	// Arrange
	bt := EmptyBinaryTree[int]()
	want := ""
	// Act
	got := bt.InOrder()
	// Assert
	util.SimpleAssert(t, got, want)
}

func TestBTToArrayReturnsArray(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	bt.Insert(5)
	bt.Insert(15)
	bt.Insert(3)
	bt.Insert(7)
	bt.Insert(12)
	bt.Insert(17)
	want := []int{3, 5, 7, 10, 12, 15, 17}
	// Act
	got := bt.ToArray()
	// Assert
	for i, v := range got {
		util.SimpleAssert(t, want[i], v)
	}
}

func TestBTToArrayEmpty(t *testing.T) {
	// Arrange
	bt := EmptyBinaryTree[int]()
	// Act
	got := bt.ToArray()
	// Assert
	if len(got) != 0 {
		t.Errorf("Expected empty array, got %v", got)
	}
}

func TestBTPreOrder(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	bt.Insert(5)
	bt.Insert(15)
	bt.Insert(3)
	bt.Insert(7)
	bt.Insert(12)
	bt.Insert(17)
	want := "10 5 3 7 15 12 17"
	// Act
	got := bt.PreOrder()
	// Assert
	util.SimpleAssert(t, got, want)
}

func TestBTPreOrderEmpty(t *testing.T) {
	// Arrange
	bt := EmptyBinaryTree[int]()
	want := ""
	// Act
	got := bt.PreOrder()
	// Assert
	util.SimpleAssert(t, got, want)
}

func TestBTPostOrder(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	bt.Insert(5)
	bt.Insert(15)
	bt.Insert(3)
	bt.Insert(7)
	bt.Insert(12)
	bt.Insert(17)
	want := "3 7 5 12 17 15 10"
	// Act
	got := bt.PostOrder()
	// Assert
	util.SimpleAssert(t, got, want)
}

func TestBTPostOrderEmpty(t *testing.T) {
	// Arrange
	bt := EmptyBinaryTree[int]()
	want := ""
	// Act
	got := bt.PostOrder()
	// Assert
	util.SimpleAssert(t, got, want)
}

func TestBinaryTree_Height(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	bt.Insert(5)
	bt.Insert(15)
	bt.Insert(3)
	bt.Insert(7)
	bt.Insert(6)
	bt.Insert(12)
	bt.Insert(17)
	expected := 4
	// Act
	got := bt.Height()
	// Assert
	util.SimpleAssert(t, expected, got)
}

func TestBinaryTree_HeightEmptyTree(t *testing.T) {
	// Arrange
	bt := EmptyBinaryTree[int]()
	expected := 0
	// Act
	got := bt.Height()
	// Assert
	util.SimpleAssert(t, expected, got)
}

func TestBinaryTree_HeightOneNode(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	expected := 1
	// Act
	got := bt.Height()
	// Assert
	util.SimpleAssert(t, expected, got)
}

func TestBinaryTree_HeightTwoNodes(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	bt.Insert(5)
	expected := 2
	// Act
	got := bt.Height()
	// Assert
	util.SimpleAssert(t, expected, got)
}

func TestBinaryTree_HeightThreeNodes(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	bt.Insert(5)
	bt.Insert(15)
	expected := 2
	// Act
	got := bt.Height()
	// Assert
	util.SimpleAssert(t, expected, got)
}

func TestBinaryTree_HeightFourNodes(t *testing.T) {
	// Arrange
	bt := NewBinaryTree[int](10)
	bt.Insert(5)
	bt.Insert(15)
	bt.Insert(3)
	expected := 3
	// Act
	got := bt.Height()
	// Assert
	util.SimpleAssert(t, expected, got)
}
