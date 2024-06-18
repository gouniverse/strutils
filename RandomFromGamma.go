package strutils

import (
	"math/rand"
	"time"
)

// RandomFromGamma generates random string of specified length with the characters specified in the gamma string
func RandomFromGamma(length int, gamma string) string {
	inRune := []rune(gamma)
	out := ""

	for i := 0; i < length; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomIndex := r.Intn(len(inRune))
		pick := inRune[randomIndex]
		out += string(pick)
	}

	return out
}
