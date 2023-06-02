package strutils

import "unicode"

// UcFirst convert first letter into upper.
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}

	return ""
}
