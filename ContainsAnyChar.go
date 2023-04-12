package strutils

import "strings"

func ContainsAnyChar(str string, chars string) bool {
	for _, char := range str {
		if strings.Contains(chars, string(char)) {
			return true
		}
	}
	return false
}
