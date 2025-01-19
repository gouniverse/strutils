package strutils

import "testing"

func TestRemoveSuffix(t *testing.T) {
	var str = "ABCabc"

	newStr := RemoveSuffix(str, "abc")

	if newStr != "ABC" {
		t.Error("Expected ABC but found:", newStr)

	}
}
