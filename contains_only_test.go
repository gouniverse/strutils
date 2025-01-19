package strutils

import (
	"testing"
)

func TestContainsOnly(t *testing.T) {
	chars := "ABCDEF"

	if !ContainsOnly("ABC", chars) {
		t.Error("ContainsOnly MUST return true for ABC contains only chars: " + chars)
	}

	if ContainsOnly("XYZ", chars) {
		t.Error("ContainsOnly MUST return false for XYZ contains only chars: " + chars)
	}

	if ContainsOnly("XAZ", chars) {
		t.Error("ContainsOnly MUST return false for XAZ contains only chars: " + chars)
	}
}
