package strutils

import (
	"math/rand/v2"
)

// RandomFromGamma generates random string of specified length with the characters specified in the gamma string
func RandomFromGamma(length int, gamma string) string {
	if length <= 0 {
		return ""
	}

	if gamma == "" {
		return ""
	}

	inRune := []rune(gamma)
	out := make([]rune, length) // Pre-allocate for efficiency

	for i := range out {
		randomIndex := rand.IntN(len(inRune))
		out[i] = inRune[randomIndex]
	}

	return string(out)
}
