package strutils

import "strings"

func RemoveSuffix(str string, suffix string) string {
	newStr, _ := strings.CutSuffix(str, suffix)
	return newStr
}
