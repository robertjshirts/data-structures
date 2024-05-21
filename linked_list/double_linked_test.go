package linked_list

import (
	"github.com/robertjshirts/data-structures/util"
	"testing"
)

// DLL stands for DoubleLinkedList

func TestEmptyDoubleLinkedList(t *testing.T) {
	// Arrange
	// Act
	list := EmptyDoubleLinkedList[int]()
	// Assert
	util.SimpleAssert(t, list.Count, 0)
	util.NilAssert(t, list.Head)
}

func TestNewDoubleLinkedList(t *testing.T) {
	// Arrange
	expected := 0
	// Act
	list := NewDoubleLinkedList(expected)
	// Assert
	util.SimpleAssert(t, list.Head.Value, expected)
}

func TestNewDoubleLinkedListHasCorrectCount(t *testing.T) {
	// Arrange
	expected := 5
	// Act
	list := NewDoubleLinkedList[int](expected)
	// Assert
	util.SimpleAssert(t, list.Count, 1)
}

func TestDoubleLinkedList_AddAddsNode(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	expected := 1
	// Act
	list.Add(expected)
	// Assert
	util.SimpleAssert(t, list.Tail.Value, expected)
}

func TestDoubleLinkedList_AddChangesCount(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	expected := 2
	// Act
	list.Add(1)
	// Assert
	util.SimpleAssert(t, list.Count, expected)

}

func TestDoubleLinkedList_AddOnEmptyLinkedList(t *testing.T) {
	// Arrange
	list := EmptyDoubleLinkedList[int]()
	expected := 1
	// Act
	list.Add(expected)
	// Assert
	util.SimpleAssert(t, list.Head.Value, expected)
}

func TestDoubleLinkedList_InsertInsertsValue(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(2)
	expected := 1
	index := 1
	// Act
	list.Insert(expected, index)
	// Assert
	util.SimpleAssert(t, list.Head.Next.Value, expected)
}

func TestDoubleLinkedList_InsertChangesCount(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(2)
	expected := 3
	// Act
	list.Insert(1, 1)
	// Assert
	util.SimpleAssert(t, list.Count, expected)
}

func TestDoubleLinkedList_InsertChangesHead(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(1)
	list.Add(2)
	expected := 0
	index := 0
	// Act
	list.Insert(expected, index)
	// Assert
	util.SimpleAssert(t, list.Head.Value, expected)
}
func TestDoubleLinkedList_InsertChangesTail(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	expected := 2
	index := 2
	// Act
	list.Insert(expected, index)
	// Assert
	util.SimpleAssert(t, list.Tail.Value, expected)
}

func TestDoubleLinkedList_InsertPanicsOnInvalidIndex(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	expected := 2
	index := 3
	// Act
	defer func() {
		// Assert
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	list.Insert(expected, index)
}

func TestDoubleLinkedList_Get(t *testing.T) {
	// Arrange
	expected := 1
	list := NewDoubleLinkedList(0)
	list.Add(expected)
	// Act
	actual := list.Get(1)
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestDoubleLinkedList_GetStartOfList(t *testing.T) {
	// Arrange
	expected := 0
	index := 0
	list := NewDoubleLinkedList(expected)
	list.Add(1)
	list.Add(2)
	// Act
	actual := list.Get(index)
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestDoubleLinkedList_GetEndOfList(t *testing.T) {
	// Arrange
	expected := 2
	index := 2
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(expected)
	// Act
	actual := list.Get(index)
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestDoubleLinkedList_GetPanicsOnInvalidIndex(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	index := 3
	// Act
	defer func() {
		// Assert
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	list.Get(index)
}

func TestDoubleLinkedList_GetPanicsOnNegativeIndex(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	index := -1
	// Act
	defer func() {
		// Assert
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	list.Get(index)
}

func TestDoubleLinkedList_RemoveReturnsCorrectValue(t *testing.T) {
	// Arrange
	expected := 0
	list := NewDoubleLinkedList(expected)
	list.Add(1)
	list.Add(2)
	// Act
	actual := list.Remove()
	// Assert
	util.SimpleAssert(t, actual, expected)
}

// Renaming functions

func TestDoubleLinkedList_RemoveChangesHead(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	expected := 1
	list.Add(expected)
	list.Add(2)
	// Act
	list.Remove()
	// Assert
	util.SimpleAssert(t, list.Head.Value, expected)
}

func TestDoubleLinkedList_RemoveChangesCount(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	// Act
	list.Remove()
	// Assert
	util.SimpleAssert(t, list.Count, 2)
}

func TestDoubleLinkedList_RemovePanicsOnEmptyList(t *testing.T) {
	// Arrange
	list := EmptyDoubleLinkedList[int]()
	// Act
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	list.Remove()
}

func TestDoubleLinkedList_RemoveAtReturnsCorrectValue(t *testing.T) {
	// Arrange
	expected := 1
	index := 1
	list := NewDoubleLinkedList(0)
	list.Add(expected)
	list.Add(2)
	// Act
	actual := list.RemoveAt(index)
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestDoubleLinkedList_RemoveAtChangesCount(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	index := 1
	// Act
	list.RemoveAt(index)
	// Assert
	util.SimpleAssert(t, list.Count, 2)
}

func TestDoubleLinkedList_RemoveAtPanicsOnInvalidIndex(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	index := 3
	// Act
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	list.RemoveAt(index)
}

func TestDoubleLinkedList_RemoveAtPanicsOnNegativeIndex(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	index := -1
	// Act
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	list.RemoveAt(index)
}

func TestDoubleLinkedList_RemoveLastReturnsCorrectValue(t *testing.T) {
	// Arrange
	expected := 2
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(expected)
	// Act
	actual := list.RemoveLast()
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestDoubleLinkedList_RemoveLastChangesTail(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	expected := 1
	list.Add(expected)
	list.Add(2)
	// Act
	list.RemoveLast()
	// Assert
	util.SimpleAssert(t, list.Tail.Value, expected)
}

func TestDoubleLinkedList_RemoveLastChangesCount(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	// Act
	list.RemoveLast()
	// Assert
	util.SimpleAssert(t, list.Count, 2)
}

func TestDoubleLinkedList_RemoveLastPanicsOnEmptyList(t *testing.T) {
	// Arrange
	list := EmptyDoubleLinkedList[int]()
	// Act
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	list.RemoveLast()
}

func TestDoubleLinkedList_ClearChangesCount(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	// Act
	list.Clear()
	// Assert
	util.SimpleAssert(t, list.Count, 0)
}

func TestDoubleLinkedList_ClearChangesHead(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	// Act
	list.Clear()
	// Assert
	util.NilAssert(t, list.Head)
}

func TestDoubleLinkedList_ClearChangesTail(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	// Act
	list.Clear()
	// Assert
	util.NilAssert(t, list.Tail)
}

func TestDoubleLinkedList_ClearClearsListOnEmptyList(t *testing.T) {
	// Arrange
	list := EmptyDoubleLinkedList[int]()
	// Act
	list.Clear()
	// Assert
	util.SimpleAssert(t, list.Count, 0)
	util.NilAssert(t, list.Head)
	util.NilAssert(t, list.Tail)
}

func TestDoubleLinkedList_SearchReturnsCorrectIndex(t *testing.T) {
	// Arrange
	value := 1
	expected := 1
	list := NewDoubleLinkedList(0)
	list.Add(value)
	list.Add(2)
	list.Add(3)
	// Act
	actual := list.Search(value)
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestDoubleLinkedList_SearchReturnsMinusOneIfNotFound(t *testing.T) {
	// Arrange
	value := 4
	expected := -1
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	// Act
	actual := list.Search(value)
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestDoubleLinkedList_SearchReturnsMinusOneOnEmptyList(t *testing.T) {
	// Arrange
	value := 1
	expected := -1
	list := EmptyDoubleLinkedList[int]()
	// Act
	actual := list.Search(value)
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestDoubleLinkedList_GetNodeReturnsCorrectNode(t *testing.T) {
	// Arrange
	expected := 1
	list := NewDoubleLinkedList(0)
	list.Add(expected)
	list.Add(2)
	list.Add(3)
	// Act
	actual := list.GetNode(1)
	// Assert
	util.SimpleAssert(t, actual.Value, expected)
}

func TestDoubleLinkedList_GetNodePanicsOnInvalidIndex(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	index := 3
	// Act
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	list.GetNode(index)
}

func TestDoubleLinkedList_GetNodePanicsOnNegativeIndex(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	index := -1
	// Act
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	list.GetNode(index)
}

func TestDoubleLinkedList_ToStringWorks(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	expected := "0 1 2 3"
	// Act
	actual := list.ToString()
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestDoubleLinkedList_ToStringWorksOnSingleElementList(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	expected := "0"
	// Act
	actual := list.ToString()
	// Assert
	util.SimpleAssert(t, actual, expected)
}

func TestDoubleLinkedList_ToStringWorksOnEmptyList(t *testing.T) {
	// Arrange
	list := EmptyDoubleLinkedList[int]()
	expected := ""
	// Act
	actual := list.ToString()
	// Assert
	util.SimpleAssert(t, actual, expected)
}
