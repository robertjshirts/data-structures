package avl_tree

import (
	"fmt"
	"github.com/robertjshirts/data-structures/util"
	"testing"
)

func TestAVLTree_NewCreatesRoot(t *testing.T) {
	// Arrange
	expected := 10
	// Act
	avl := NewAVLTree[int](expected)
	// Assert
	util.SimpleAssert(t, expected, avl.Root.value)
}

func TestAVLTree_NewCorrectLength(t *testing.T) {
	// Arrange
	expected := 1
	// Act
	avl := NewAVLTree[int](10)
	// Assert
	util.SimpleAssert(t, expected, avl.Count)
}

func TestNewAVLTree(t *testing.T) {
	// Arrange
	left := 5
	right := 15
	root := 10
	// Act
	avl := NewAVLTree[int](root, left, right)
	// Assert
	util.SimpleAssert(t, root, avl.Root.value)
	util.SimpleAssert(t, left, avl.Root.left.value)
	util.SimpleAssert(t, right, avl.Root.right.value)
	util.SimpleAssert(t, avl.Count, 3)
}

func TestEmptyAVLTree(t *testing.T) {
	// Arrange
	// Act
	avl := EmptyAVLTree[int]()
	// Assert
	util.NilAssert(t, avl.Root)
	util.SimpleAssert(t, avl.Count, 0)
}

func TestAVLTree_Clear(t *testing.T) {
	// Arrange
	avl := NewAVLTree(10, 5, 15)
	// Act
	avl.Clear()
	// Assert
	util.NilAssert(t, avl.Root)
	util.SimpleAssert(t, avl.Count, 0)
}

func TestAVLTree_ClearOnEmptyTree(t *testing.T) {
	// Arrange
	avl := EmptyAVLTree[int]()
	// Act
	avl.Clear()
	// Assert
	util.NilAssert(t, avl.Root)
	util.SimpleAssert(t, avl.Count, 0)
}

func TestAVLTree_Height(t *testing.T) {
	// Arrange
	avl := NewAVLTree(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	// Act
	util.SimpleAssert(t, avl.Height(), 4)
}

func TestAVLTree_HeightOnEmptyTree(t *testing.T) {
	// Arrange
	avl := EmptyAVLTree[int]()
	// Act
	avl.Clear()
	// Assert
	util.NilAssert(t, avl.Root)
	util.SimpleAssert(t, avl.Count, 0)
}

func TestAVLTree_InsertAddsNode(t *testing.T) {
	// Arrange
	avl := NewAVLTree[int](10)
	expected := 5
	// Act
	avl.Insert(expected)
	// Assert
	util.SimpleAssert(t, expected, avl.Root.left.value)
}

func TestAVLTree_InsertChangesCount(t *testing.T) {
	// Arrange
	avl := NewAVLTree[int](10)
	expected := 3
	// Act
	avl.Insert(5)
	avl.Insert(15)
	// Assert
	util.SimpleAssert(t, expected, avl.Count)
}

func TestAVLTree_InsertAddsDuplicate(t *testing.T) {
	// Arrange
	avl := NewAVLTree[int](10)
	expected := 10
	// Act
	avl.Insert(expected)
	// Assert
	util.SimpleAssert(t, avl.Root.right.value, expected)
}

func TestAVLTree_InsertRotatesRight(t *testing.T) {
	// Arrange
	expectedRight := 7
	expectedLeft := 3
	expectedRoot := 5
	avl := NewAVLTree[int](expectedRight)
	// Act
	avl.Insert(expectedRoot)
	avl.Insert(expectedLeft)
	// Assert
	util.SimpleAssert(t, avl.Root.value, expectedRoot)
	util.SimpleAssert(t, avl.Root.right.value, expectedRight)
	util.SimpleAssert(t, avl.Root.left.value, expectedLeft)
}

func TestAVLTree_InsertRotatesLeft(t *testing.T) {
	// Arrange
	expectedRight := 7
	expectedLeft := 3
	expectedRoot := 5
	avl := NewAVLTree[int](expectedLeft)
	// Act
	avl.Insert(expectedRoot)
	avl.Insert(expectedRight)
	// Assert
	util.SimpleAssert(t, avl.Root.value, expectedRoot)
	util.SimpleAssert(t, avl.Root.right.value, expectedRight)
	util.SimpleAssert(t, avl.Root.left.value, expectedLeft)
}

func TestAvlTree_InsertRotatesLeftRight(t *testing.T) {
	// Arrange
	expectedRight := 5
	expectedLeft := 3
	expectedRoot := 4
	avl := NewAVLTree[int](expectedRight)
	// Act
	avl.Insert(expectedRoot)
	avl.Insert(expectedLeft)
	// Assert
	util.SimpleAssert(t, avl.Root.value, expectedRoot)
	util.SimpleAssert(t, avl.Root.right.value, expectedRight)
	util.SimpleAssert(t, avl.Root.left.value, expectedLeft)
}

func TestAVLTree_InsertRotatesRightLeft(t *testing.T) {
	// Arrange
	expectedRight := 5
	expectedLeft := 3
	expectedRoot := 4
	avl := NewAVLTree[int](expectedRight)
	// Act
	avl.Insert(expectedRoot)
	avl.Insert(expectedLeft)
	// Assert
	util.SimpleAssert(t, avl.Root.value, expectedRoot)
	util.SimpleAssert(t, avl.Root.right.value, expectedRight)
	util.SimpleAssert(t, avl.Root.left.value, expectedLeft)
}

func TestAVLTree_RemoveRotatesRight(t *testing.T) {
	// Arrange
	// Insert tree nodes all at once bc we aren't testing Insert
	expectedRoot := 15
	expectedRight := 17
	expectedLeft := 10
	valueToRemove := 5
	avl := NewAVLTree[int](expectedLeft, valueToRemove, expectedRoot, expectedRight)
	// Act
	avl.Remove(valueToRemove)
	// Assert
	util.SimpleAssert(t, avl.Root.value, expectedRoot)
	util.SimpleAssert(t, avl.Root.right.value, expectedRight)
	util.SimpleAssert(t, avl.Root.left.value, expectedLeft)
}

func TestAVLTree_RemoveRotatesLeft(t *testing.T) {
	// Arrange
	// Insert all tree nodes at once
	expectedRoot := 5
	expectedRight := 10
	expectedLeft := 3
	valueToRemove := 15
	avl := NewAVLTree[int](expectedRight, valueToRemove, expectedRoot, expectedLeft)
	// Act
	avl.Remove(valueToRemove)
	// Assert
	util.SimpleAssert(t, avl.Root.value, expectedRoot)
	util.SimpleAssert(t, avl.Root.right.value, expectedRight)
	util.SimpleAssert(t, avl.Root.left.value, expectedLeft)
}

func TestAVLTree_RemoveRotatesLeftRight(t *testing.T) {
	// Arrange
	// Insert tree nodes all at once
	expectedRoot := 13
	expectedRight := 15
	expectedLeft := 10
	valueToRemove := 5
	avl := NewAVLTree[int](expectedLeft, valueToRemove, expectedRight, expectedRoot)
	// Act
	avl.Remove(valueToRemove)
	// Assert
	util.SimpleAssert(t, avl.Root.value, expectedRoot)
	util.SimpleAssert(t, avl.Root.right.value, expectedRight)
	util.SimpleAssert(t, avl.Root.left.value, expectedLeft)
}

func TestAVLTree_RemoveRotatesRightLeft(t *testing.T) {
	// Arrange
	// Insert tree nodes all at once
	expectedRoot := 7
	expectedRight := 10
	expectedLeft := 5
	valueToRemove := 15
	avl := NewAVLTree(expectedRight, expectedLeft, valueToRemove, expectedRoot)
	// Act
	avl.Remove(valueToRemove)
	// Assert
	util.SimpleAssert(t, avl.Root.value, expectedRoot)
	util.SimpleAssert(t, avl.Root.right.value, expectedRight)
	util.SimpleAssert(t, avl.Root.left.value, expectedLeft)
}

func TestAVLTree_ContainsReturnsTrue(t *testing.T) {
	// Arrange
	avl := NewAVLTree(10, 15, 5, 3, 7, 17)
	value := 13
	avl.Insert(value)
	want := true
	// Act
	got := avl.Contains(value)
	// Assert
	util.SimpleAssert(t, got, want)
}

func TestAVLTree_ContainsReturnsFalse(t *testing.T) {
	// Arrange
	avl := NewAVLTree(10, 15, 5, 3, 7, 17)
	value := 13
	want := false
	// Act
	got := avl.Contains(value)
	// Assert
	util.SimpleAssert(t, got, want)
}

func TestAVLTree_ContainsReturnsFalseOnEmptyTree(t *testing.T) {
	// Arrange
	avl := EmptyAVLTree[int]()
	want := false
	// Act
	got := avl.Contains(0)
	// Assert
	util.SimpleAssert(t, got, want)
}

func TestAVLTree_ToArrayReturnsCorrectArray2(t *testing.T) {
	// Arrange
	avl := NewAVLTree(10, 15, 5, 3, 6)
	expectedArray := []int{10, 5, 15, 3, 6}
	// Act
	actualArray := avl.ToArray()
	fmt.Printf("%v\n", actualArray)
	// Assert
	for i := 0; i < len(expectedArray); i++ {
		util.SimpleAssert(t, actualArray[i], expectedArray[i])
	}
}

func TestAVLTree_ToArrayReturnsCorrectArray(t *testing.T) {
	// Arrange
	avl := NewAVLTree(10, 5, 15, 3, 7, 13, 17)
	expectedArray := []int{10, 5, 15, 3, 7, 13, 17}
	// Act
	actualArray := avl.ToArray()
	// Assert
	for i := 0; i < len(expectedArray); i++ {
		util.SimpleAssert(t, actualArray[i], expectedArray[i])
	}
}

func TestAVLTree_ToArrayReturnsNilOnEmptyTree(t *testing.T) {
	// Arrange
	avl := EmptyAVLTree[int]()
	// Act
	actualArray := avl.ToArray()
	// Assert
	util.SimpleAssert(t, len(actualArray), 0)
}

func TestAVLTree_PreOrder(t *testing.T) {
	// Arrange
	avl := NewAVLTree(1, 2, 3, 4, 5, 6, 7)
	want := "4 2 1 3 6 5 7"
	// Act
	got := avl.PreOrder()
	// Assert
	util.SimpleAssert(t, got, want)
}

func TestAVLTree_PreOrderReturnsEmptyStringOnEmptyTree(t *testing.T) {
	// Arrange
	avl := EmptyAVLTree[int]()
	want := ""
	// Act
	got := avl.PreOrder()
	// Assert
	util.SimpleAssert(t, got, want)
}

func TestAVLTree_PostOrder(t *testing.T) {
	// Arrange
	avl := NewAVLTree(1, 2, 3, 4, 5, 6, 7)
	want := "1 3 2 5 7 6 4"
	// Act
	got := avl.PostOrder()
	// Assert
	util.SimpleAssert(t, got, want)
}

func TestAVLTree_PostOrderReturnsEmptyStringOnEmptyTree(t *testing.T) {
	// Arrange
	avl := EmptyAVLTree[int]()
	want := ""
	// Act
	got := avl.PostOrder()
	// Assert
	util.SimpleAssert(t, got, want)
}

func TestAVLTree_InOrder(t *testing.T) {
	// Arrange
	avl := NewAVLTree(1, 2, 3, 4, 5, 6, 7)
	want := "1 2 3 4 5 6 7"
	// Act
	got := avl.InOrder()
	// Assert
	util.SimpleAssert(t, got, want)
}

func TestAVLTree_InOrderReturnsEmptyStringOnEmptyTree(t *testing.T) {
	// Arrange
	avl := EmptyAVLTree[int]()
	want := ""
	// Act
	got := avl.InOrder()
	// Assert
	util.SimpleAssert(t, got, want)
}
