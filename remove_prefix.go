package strutils

import "strings"

func RemovePrefix(str string, prefix string) string {
	newStr, _ := strings.CutPrefix(str, prefix)
	return newStr
}
