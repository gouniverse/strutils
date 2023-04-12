package strutils

import "strings"

// ContainsAnyChar returns true if the string contains
// any of the characters in the provided charset
func ContainsAnyChar(str string, charset string) bool {
	for _, char := range str {
		if strings.Contains(charset, string(char)) {
			return true
		}
	}
	return false
}
