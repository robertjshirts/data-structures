package linked_list

import (
	"testing"
)

// DLL stands for DoubleLinkedList

func TestEmptyDLLWorksCorrectly(t *testing.T) {
	// Arrange
	// Act
	list := EmptyDoubleLinkedList[int]()
	// Assert
	if list.Count != 0 {
		simpleAssert(t, "Count", 0, list.Count)
	}
	if list.Head != nil {
		simpleAssert(t, "Head", nil, list.Head)
	}
	if list.Tail != nil {
		simpleAssert(t, "Tail", nil, list.Tail)
	}
}

func TestNewDLLWorksCorrectly(t *testing.T) {
	// Arrange
	expected := 5
	// Act
	list := NewDoubleLinkedList[int](expected)
	// Assert
	if list.Count != 1 {
		simpleAssert(t, "Count", 1, list.Count)
	}
	if list.Head.Value != expected {
		simpleAssert(t, "Head.Value", expected, list.Head.Value)
	}
	if list.Tail.Value != expected {
		simpleAssert(t, "Tail.Value", expected, list.Tail.Value)
	}
}

func TestDLLAdd(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	expected := 1
	// Act
	list.Add(expected)
	actual := list.Head.Next.Value
	// Assert
	if list.Head.Next.Value != list.Tail.Value {
		simpleAssert(t, "Head.Next.Value", list.Tail.Value, list.Head.Next.Value)
	}
	if expected != actual {
		simpleAssert(t, "Head.Next.Value", expected, actual)
	}
}

func TestDLLAddWorksOnEmptyList(t *testing.T) {
	// Arrange
	list := EmptyDoubleLinkedList[int]()
	expected := 1
	// Act
	list.Add(expected)
	actual := list.Head.Value
	// Assert
	if expected != actual {
		simpleAssert(t, "Head.Value", expected, actual)
	}
}

func TestDLLInsert(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(2)
	expected := 1
	index := 1
	// Act
	list.Insert(expected, index)
	actual := list.Head.Next.Value
	// Assert
	if expected != actual {
		simpleAssert(t, "Head.Next.Value", expected, actual)
	}
}

func TestDLLInsertWorksOnEmptyList(t *testing.T) {
	// Arrange
	list := EmptyDoubleLinkedList[int]()
	expected := 0
	index := 0
	// Act
	list.Insert(expected, index)
	actual := list.Head.Value
	// Assert
	if expected != actual {
		simpleAssert(t, "Head.Value", expected, actual)
	}
}

func TestDLLInsertWorksOnEndOfList(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	expected := 2
	index := 2
	// Act
	list.Insert(expected, index)
	actual := list.Tail.Value
	// Assert
	if expected != actual {
		simpleAssert(t, "Tail.Value", expected, actual)
	}
}

func TestDLLInsertPanicsOnInvalidIndex(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	expected := 2
	index := 3
	// Act
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	list.Insert(expected, index)
}

func TestDLLGet(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	expected := 1
	index := 1
	list.Add(expected)
	list.Add(2)
	// Act
	actual := list.Get(index)
	// Assert
	if expected != actual {
		simpleAssert(t, "Get", expected, actual)
	}
}

func TestDLLGetStartOfList(t *testing.T) {
	// Arrange
	expected := 0
	index := 0
	list := NewDoubleLinkedList(expected)
	list.Add(1)
	list.Add(2)
	// Act
	actual := list.Get(index)
	// Assert
	if expected != actual || actual != list.Head.Value {
		simpleAssert(t, "Get", expected, actual)
	}
}

func TestDLLGetEndOfList(t *testing.T) {
	// Arrange
	expected := 2
	index := 2
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(expected)
	// Act
	actual := list.Get(index)
	// Assert
	if expected != actual || actual != list.Tail.Value {
		simpleAssert(t, "Get", expected, actual)
	}
}

func TestDLLGetPanicsOnInvalidIndex(t *testing.T) {
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
	list.Get(index)
}

func TestDLLRemoveReturnsCorrectValue(t *testing.T) {
	// Arrange
	expected := 0
	list := NewDoubleLinkedList(expected)
	list.Add(1)
	list.Add(2)
	// Act
	actual := list.Remove()
	// Assert
	if expected != actual {
		simpleAssert(t, "Remove", expected, actual)
	}
}

func TestDLLRemoveChangesHeadAndCount(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	expected := 1
	list.Add(expected)
	list.Add(2)
	// Act
	list.Remove()
	// Assert
	if list.Head.Value != expected {
		simpleAssert(t, "Head.Value", expected, list.Head.Value)
	}
	if list.Count != 2 {
		simpleAssert(t, "Count", 2, list.Count)
	}
}

func TestDLLRemovePanicsOnEmptyList(t *testing.T) {
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

func TestDLLRemoveAtReturnsCorrectValue(t *testing.T) {
	// Arrange
	expected := 1
	index := 1
	list := NewDoubleLinkedList(0)
	list.Add(expected)
	list.Add(2)
	// Act
	actual := list.RemoveAt(index)
	// Assert
	if expected != actual {
		simpleAssert(t, "RemoveAt", expected, actual)
	}
}

func TestDLLRemoveAtChangesCount(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	index := 1
	// Act
	list.RemoveAt(index)
	// Assert
	if list.Count != 2 {
		simpleAssert(t, "Count", 2, list.Count)
	}
}

func TestDLLRemoveAtPanicsOnInvalidIndex(t *testing.T) {
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

func TestDLLRemoveLastReturnsCorrectValue(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	expected := 2
	list.Add(expected)
	// Act
	actual := list.RemoveLast()
	// Assert
	if expected != actual {
		simpleAssert(t, "RemoveLast", expected, list.Tail.Value)
	}
}

func TestDLLRemoveLastChangesTailAndCount(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	expected := 1
	list.Add(expected)
	list.Add(2)
	// Act
	list.RemoveLast()
	// Assert
	if list.Tail.Value != expected {
		simpleAssert(t, "Tail.Value", expected, list.Tail.Value)
	}
	if list.Count != 2 {
		simpleAssert(t, "Count", 2, list.Count)
	}
}

func TestDLLRemoveLastPanicsOnEmptyList(t *testing.T) {
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

func TestDLLClearClearsList(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	// Act
	list.Clear()
	// Assert
	if list.Count != 0 {
		simpleAssert(t, "Count", 0, list.Count)
	}
	if list.Head != nil {
		simpleAssert(t, "Head", nil, list.Head)
	}
	if list.Tail != nil {
		simpleAssert(t, "Tail", nil, list.Tail)
	}
}

func TestDLLClearClearsListOnEmptyList(t *testing.T) {
	// Arrange
	list := EmptyDoubleLinkedList[int]()
	// Act
	list.Clear()
	// Assert
	if list.Count != 0 {
		simpleAssert(t, "Count", 0, list.Count)
	}
	if list.Head != nil {
		simpleAssert(t, "Head", nil, list.Head)
	}
	if list.Tail != nil {
		simpleAssert(t, "Tail", nil, list.Tail)
	}
}

func TestDLLSearchReturnsCorrectIndex(t *testing.T) {
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
	if expected != actual {
		simpleAssert(t, "Search", expected, actual)
	}
}

func TestDLLReturnsMinusOneIfNotInList(t *testing.T) {
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
	if expected != actual {
		simpleAssert(t, "Search", expected, actual)
	}
}

func TestDLLSearchReturnsMinusOneOnEmptyList(t *testing.T) {
	// Arrange
	value := 1
	expected := -1
	list := EmptyDoubleLinkedList[int]()
	// Act
	actual := list.Search(value)
	// Assert
	if expected != actual {
		simpleAssert(t, "Search", expected, actual)
	}
}

func TestDLLToString(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList(0)
	list.Add(1)
	list.Add(2)
	expected := "0 1 2"
	// Act
	actual := list.ToString()
	// Assert
	if expected != actual {
		simpleAssert(t, "ToString", expected, actual)
	}
}

func TestDLLToStringWorksOnEmptyList(t *testing.T) {
	// Arrange
	list := EmptyDoubleLinkedList[int]()
	expected := ""
	// Act
	actual := list.ToString()
	// Assert
	if expected != actual {
		simpleAssert(t, "ToString", expected, actual)
	}
}
