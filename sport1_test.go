package main

import (
	"testing"
)

// test function
func TestReturnSports(t *testing.T) {
	actualString := ReturnSports()
	expectedString := "sports"
	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedString, actualString)
	}
}
