package strutils

import "testing"

func TestToSnake(t *testing.T) {
	var str = "The Sleeping Fox Catches No Poultry"

	newStr := ToSnake(str)

	if newStr != "the_sleeping_fox_catches_no_poultry" {
		t.Error("Expected the_sleeping_fox_catches_no_poultry but found:", newStr)

	}
}
