package strutils

import "strings"

// ContainsOnly returns true is the string contains only charcters from the specified charset
func ContainsOnly(str string, charset string) bool {
	nonCharacter := func(c rune) bool { return !strings.ContainsAny(charset, string(c)) }
	words := strings.FieldsFunc(str, nonCharacter)
	return str == strings.Join(words, "")
}
