package util

import "testing"

func NilAssert[T comparable](t *testing.T, got *T) {
	if got != nil {
		t.Errorf("Expected nil, got %v", *got)
	}
}

func SimpleAssert[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
