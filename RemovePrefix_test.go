package strutils

import "testing"

func TestRemovePrefix(t *testing.T) {
	var str = "abcABC"

	newStr := RemovePrefix(str, "abc")

	if newStr != "ABC" {
		t.Error("Expected ABC but found:", newStr)

	}
}
