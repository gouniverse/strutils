package strutils

import "testing"

func TestToCamel(t *testing.T) {
	var str = "the_sleeping_fox_catches_no_poultry"

	newStr := ToCamel(str)

	if newStr != "TheSleepingFoxCatchesNoPoultry" {
		t.Error("Expected the_sleeping_fox_catches_no_poultry but found:", newStr)

	}
}
