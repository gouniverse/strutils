package strutils

import (
	"testing"
)

func TestContainsAnyChar(t *testing.T) {
	if ContainsAnyChar("ABC", "abc") {
		t.Error("ContainsAnyChar MUST return false for ABC contains any chars: abc")
	}

	if !ContainsAnyChar("ABC", "aBc") {
		t.Error("ContainsAnyChar MUST return true for ABC contains any chars: aBc")
	}

	if !ContainsAnyChar("ABC", "Abc") {
		t.Error("ContainsAnyChar MUST return true for ABC contains any chars: Abc")
	}

	if !ContainsAnyChar("ABC", "abC") {
		t.Error("ContainsAnyChar MUST return true for ABC contains any chars: abC")
	}
}
