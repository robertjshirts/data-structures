package kvp

import (
	"testing"
)

func TestKey(t *testing.T) {
	// Arrange
	expected := "key"
	kvp := NewKVP(expected, "")
	// Act
	actual := kvp.Key()
	// Assert
	if expected != actual {
		t.Fatalf("Expected kvp.Key() to be %s, got %s", expected, actual)
	}
}

func TestValue(t *testing.T) {
	// Arrange
	key := "key"
	expected := "value"
	kvp := NewKVP(key, expected)
	// Act
	actual := kvp.Value()
	// Assert
	if expected != actual {
		t.Fatalf("Expected kvp.Value() to be %s, got %s", expected, actual)
	}
}
