package strutils

import (
	"strings"
)

func ToCamel(in string) string {
	in = strings.ReplaceAll(in, "_", " ")
	in = strings.ReplaceAll(in, "-", " ")
	words := strings.Split(in, " ")

	for i := 0; i < len(words); i++ {
		words[i] = UcFirst(words[i])
	}

	return strings.Join(words, "")
}
