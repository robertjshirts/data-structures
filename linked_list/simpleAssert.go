package linked_list

import "testing"

func simpleAssert[T any](t *testing.T, test string, expectedValue T, actualValue T) {
	t.Fatalf("Expected %v to be %v, got %v", test, expectedValue, actualValue)
}
