package strutils

import "strings"

// LeftFrom returns the substring on the left side of the needle
func LeftFrom(str, needle string) string {
	left, _, isFound := strings.Cut(str, needle)

	if !isFound {
		return ""
	}

	return left
}
