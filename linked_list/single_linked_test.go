package linked_list

import "testing"

func TestNewListAndHeadValue(t *testing.T) {
	// Arrange
	expected := "the first value"
	// Act
	list := NewSingleLinkedList(expected)
	// Assert
	if list.Head.Value != expected {
		simpleAssert(t, "list.Head.Value", expected, list.Head.Value)
	}
	if list.Head.Next != nil {
		simpleAssert(t, "list.Head.Next", "nil", "something else lol")
	}
}

func TestEmptyList(t *testing.T) {
	// Arrange
	// Act
	list := EmptySingleLinkedList[string]()
	// Assert
	if list.Head != nil {
		simpleAssert(t, "list.Head", "nil", "something else lol")
	}
	if list.Count != 0 {
		simpleAssert(t, "list.Count", 0, list.Count)
	}
}

func TestListAdd(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList("first value")
	expected := "expected value"
	// Act
	list.Add(expected)
	actual := list.Head.Next.Value
	// Assert
	if expected != actual {
		simpleAssert(t, "list.Head.Next.Value", expected, actual)
	}
}

func TestListAddCount(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList("first value")
	expected := 2
	// Act
	list.Add("second value")
	actual := list.Count
	// Assert
	if expected != actual {
		simpleAssert(t, "list.Count", expected, actual)
	}
}

func TestListAddMultiple(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList("first value")
	list.Add("second value")
	list.Add("third value")
	expected := "third value"
	// Act
	actual := list.Head.Next.Next.Value
	// Assert
	if expected != actual {
		simpleAssert(t, "list.Head.Next.Next.Value", expected, actual)
	}
}

func TestGetNode(t *testing.T) {
	// Arrange
	value := "value"
	list := NewSingleLinkedList(value)
	// Act
	actual := list.GetNode(0)
	// Assert
	if actual != list.Head {
		simpleAssert(t, "list.GetNode(0)", "list.Head", "something else lmao")
	}
}

func TestGetNodeOutOfBoundsPanics(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList("value")
	defer func() {
		if r := recover(); r == nil {
			simpleAssert(t, "list.GetNode(1)", "panic", "no panic")
		}
	}()
	// Act
	list.GetNode(1)
	// Should panic, so we have to recover with the previously deffered function
}

func TestListInsert(t *testing.T) {
	// Arrange
	expected := "second value"
	index := 1
	list := NewSingleLinkedList("first value")
	list.Add("third value")
	// Act
	list.Insert(expected, index)
	actual := list.Get(index)
	// Assert
	if expected != actual {
		simpleAssert(t, "list.Insert(1)", expected, actual)
	}
}

func TestListInsertAtZero(t *testing.T) {
	// Arrange
	index := 0
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(3)
	// Act
	list.Insert(0, index)
	actual := list.Get(index)
	// Assert
	if actual != 0 {
		simpleAssert(t, "list.Insert(0)", 0, actual)
	}
}

func TestListInsertOutOfBoundsPanics(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList(1)
	defer func() {
		if r := recover(); r == nil {
			simpleAssert(t, "list.Insert(X, 1)", "panic", "no panic")
		}
	}()
	// Act
	// Insert val 2 at index 3, far out of bounds
	list.Insert(2, 3)
	// Should panic, so we have to recover with the previously deffered function
}

func TestListInsertCount(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(3)
	expected := 4
	// Act
	list.Insert(4, 3)
	actual := list.Count
	// Assert
	if expected != actual {
		simpleAssert(t, "list.Count", expected, actual)
	}
}

func TestListGet(t *testing.T) {
	// Arrange
	expected := "second value"
	list := NewSingleLinkedList("first value")
	list.Add(expected)
	// Act
	actual := list.Get(1)
	// Assert
	if expected != actual {
		simpleAssert(t, "list.Get(1)", expected, actual)
	}
}

func TestListGetOutOfBoundsPanics(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList(1)
	defer func() {
		if r := recover(); r == nil {
			simpleAssert(t, "list.Get(1)", "panic", "no panic")
		}
	}()
	// Act
	list.Get(1)
	// Should panic, so we have to recover with the previously deffered function
}

func TestListRemoveReturnsCorrectValue(t *testing.T) {
	// Arrange
	expected := 1
	list := NewSingleLinkedList(expected)
	list.Add(2)
	list.Add(3)
	// Act
	actual := list.Remove()
	// Assert
	if expected != actual {
		simpleAssert(t, "list.Remove()", expected, actual)
	}
}

func TestListRemoveRemovesNode(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(3)
	expected := 2
	// Act
	list.Remove()
	actual := list.Count
	// Assert
	if expected != actual {
		simpleAssert(t, "list.Count", expected, actual)
	}
}

func TestListRemoveAtReturnsCorrectValue(t *testing.T) {
	// Arrange
	expected := 2
	list := NewSingleLinkedList(1)
	list.Add(expected)
	list.Add(3)
	// Act
	actual := list.RemoveAt(1)
	// Assert
	if expected != actual {
		simpleAssert(t, "list.RemoveAt(1)", expected, actual)
	}
}

func TestListRemoveAtRemovesNode(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(3)
	expected := 2
	// Act
	list.RemoveAt(1)
	actual := list.Count
	// Assert
	if expected != actual {
		simpleAssert(t, "list.Count", expected, actual)
	}
}

func TestRemoveLastReturnsCorrectValue(t *testing.T) {
	// Arrange
	expected := 3
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(expected)
	// Act
	actual := list.RemoveLast()
	// Assert
	if expected != actual {
		simpleAssert(t, "list.RemoveLast()", expected, actual)
	}
}

func TestRemoveLastRemovesNode(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(3)
	expected := 2
	// Act
	list.RemoveLast()
	actual := list.Count
	// Assert
	if expected != actual {
		simpleAssert(t, "list.RemoveLast()", expected, actual)
	}
}

func TestListClearClearsList(t *testing.T) {
	// Arrange
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(3)
	expected := 0
	// Act
	list.Clear()
	actual := list.Count
	// Assert
	if expected != actual {
		simpleAssert(t, "list.Clear()", expected, actual)
	}
}

func TestListSearchReturnsCorrectIndex(t *testing.T) {
	// Arrange
	search := 2
	expected := 1
	list := NewSingleLinkedList(1)
	list.Add(search)
	list.Add(3)
	// Act
	actual := list.Search(search)
	// Assert
	if expected != actual {
		simpleAssert(t, "list.Search(2)", expected, actual)
	}
}

func TestListSearchReturnsMinusOneIfNotInList(t *testing.T) {
	// Arrange
	search := 4
	expected := -1
	list := NewSingleLinkedList(1)
	list.Add(2)
	list.Add(3)
	// Act
	actual := list.Search(search)
	// Assert
	if expected != actual {
		simpleAssert(t, "list.Search(4)", expected, actual)
	}
}
