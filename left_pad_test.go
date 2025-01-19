package strutils

import "testing"

func TestLeftPad(t *testing.T) {
	var str = "1"

	newStr := LeftPad(str, "0", 6)

	if newStr != "000001" {
		t.Error("Expected 000001 but found:", newStr)

	}
}
