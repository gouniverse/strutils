package strutils

import "strings"

// RightFrom returns the substring on the left side of the needle
func RightFrom(str, needle string) string {
	_, right, isFound := strings.Cut(str, needle)

	if !isFound {
		return ""
	}

	return right
}
