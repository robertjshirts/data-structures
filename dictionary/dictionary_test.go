package dictionary

import (
	"testing"

	"github.com/robertjshirts/data-structures/kvp"
)

func TestAddAndRetrieveValue(t *testing.T) {
	// Arrange
	key := "key"
	expected := "value"
	dict := NewDict[string, string]()
	// Act
	dict.Add(key, expected)
	actual := *(dict.Get(key))
	// Assert
	if expected != actual {
		t.Fatalf("Expected dict.Get(%s) to return %s, got %s", key, expected, actual)
	}
}

func TestAddAndRetrieveKVP(t *testing.T) {
	// Arrange
	key := "key"
	value := "value"
	expected := kvp.NewKVP(key, value)
	dict := NewDict[string, string]()
	// Act
	dict.AddKVP(expected)
	actual := *(dict.GetKVP(key))
	// Assert
	if expected != actual {
		t.Fatalf("Expected dict.Get(%s) to return %s, got %s", key, expected, actual)
	}
}
