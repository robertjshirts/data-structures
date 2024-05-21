package linked_list

import (
	"github.com/robertjshirts/data-structures/util"
	"testing"
)

func TestNewSingleLinkedList(t *testing.T) {
	// Arrange
	expected := "the first value"
	// Act
	list := NewSingleLinkedList(expected)
	// Assert
	util.SimpleAssert(t, list.Head.Value, expected)
	util.NilAssert(t, list.Head.Next)
}

func TestEmptySingleLinkedList(t *testing.T) {
	// Arrange
	// Act
	list := EmptySingleLinkedList[string]()
	// Assert
	util.NilAssert(t, list.Head)
	util.SimpleAssert(t, list.Count, 0)
}

func TestSingleLinkedList_Add(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList("first value")
	expected := "expected value"
	// Act
	list.Add(expected)
	actual := list.Head.Next.Value
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestSingleLinkedList_AddChangesCount(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList("first value")
	expected := 2
	// Act
	list.Add("second value")
	actual := list.Count
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestSingleLinkedList_AddMultipleNodes(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList("first value")
	list.Add("second value")
	expected := "third value"
	// Act
	list.Add(expected)
	actual := list.Head.Next.Next.Value
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestSingleLinkedList_GetNode(t *testing.T) {
	// Arrange
	value := "value"
	list := NewSingleLinkedList(value)
	// Act
	actual := list.GetNode(0)
	// Assert
	util.SimpleAssert(t, actual, list.Head)
}

func TestSingleLinkedList_GetNodePanicsOnInvalidIndex(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList("value")
	defer func() {
		// Assert
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Act
	list.GetNode(1)
}

func TestSingleLinkedList_Insert(t *testing.T) {
	// Arrange
	expected := "second value"
	index := 1
	list := NewSingleLinkedList("first value")
	list.Add("third value")
	// Act
	list.Insert(expected, index)
	actual := list.Get(index)
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestSingleLinkedList_InsertAtZero(t *testing.T) {
	// Arrange
	index := 0
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(3)
	// Act
	list.Insert(0, index)
	actual := list.Get(index)
	// Assert
	util.SimpleAssert(t, actual, 0)
}

func TestSingleLinkedList_InsertPanicsOnInvalidIndex(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList(1)
	defer func() {
		// Assert
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Act
	list.Insert(2, 3)
}

func TestSingleLinkedList_InsertChangesCount(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(3)
	expected := 4
	// Act
	list.Insert(4, 3)
	actual := list.Count
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestSingleLinkedList_Get(t *testing.T) {
	// Arrange
	expected := "second value"
	list := NewSingleLinkedList("first value")
	list.Add(expected)
	// Act
	actual := list.Get(1)
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestSingleLinkedList_GetPanicsOnInvalidIndex(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList(1)
	defer func() {
		// Assert
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Act
	list.Get(1)
}

func TestSingleLinkedList_RemoveReturnsCorrectValue(t *testing.T) {
	// Arrange
	expected := 1
	list := NewSingleLinkedList(expected)
	list.Add(2)
	list.Add(3)
	// Act
	actual := list.Remove()
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestSingleLinkedList_RemoveChangesCount(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(3)
	expected := 2
	// Act
	list.Remove()
	actual := list.Count
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestSingleLinkedList_RemoveAtReturnsCorrectValue(t *testing.T) {
	// Arrange
	expected := 2
	list := NewSingleLinkedList(1)
	list.Add(expected)
	list.Add(3)
	// Act
	actual := list.RemoveAt(1)
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestSingleLinkedList_RemoveAtChangesCount(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(3)
	expected := 2
	// Act
	list.RemoveAt(1)
	actual := list.Count
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestSingleLinkedList_RemoveLastReturnsCorrectValue(t *testing.T) {
	// Arrange
	expected := 3
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(expected)
	// Act
	actual := list.RemoveLast()
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestSingleLinkedList_RemoveLastChangesCount(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(3)
	expected := 2
	// Act
	list.RemoveLast()
	actual := list.Count
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestSingleLinkedList_ClearChangesCount(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(3)
	expected := 0
	// Act
	list.Clear()
	actual := list.Count
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestSingleLinkedList_SearchReturnsCorrectIndex(t *testing.T) {
	// Arrange
	search := 2
	expected := 1
	list := NewSingleLinkedList(1)
	list.Add(search)
	list.Add(3)
	// Act
	actual := list.Search(search)
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestSingleLinkedList_SearchReturnsMinusOneIfNotFound(t *testing.T) {
	// Arrange
	search := 4
	expected := -1
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(3)
	// Act
	actual := list.Search(search)
	// Assert
	util.SimpleAssert(t, actual, expected)
}
