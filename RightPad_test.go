package strutils

import "testing"

func TestRightPad(t *testing.T) {
	var str = "1"

	newStr := RightPad(str, "0", 6)

	if newStr != "100000" {
		t.Error("Expected 100000 but found:", newStr)

	}
}
