package main

import "testing"

// TestAdd tests the Add function.
func TestAdd(t *testing.T) {
	result := Add(3, 5)
	expected := 8

	if result != expected {
		t.Errorf("Add(3, 5) failed, expected %d, got %d", expected, result)
	}

	result = Add(-1, 1)
	expected = 0

	if result != expected {
		t.Errorf("Add(-1, 1) failed, expected %d, got %d", expected, result)
	}
}
