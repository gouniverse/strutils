package strutils

import "unicode/utf8"

// Truncate truncates a string to a given length, adding an ellipsis if necessary.
func Truncate(str string, length int, ellipsis string) string {
	// Handle trivial cases: empty string, zero/negative length.
	if str == "" || length <= 0 {
		return ""
	}

	runeStr := []rune(str)
	runeEllipsis := []rune(ellipsis)
	ellipsisLen := len(runeEllipsis)

	// No truncation needed.
	if utf8.RuneCountInString(str) <= length {
		return str
	}

	// If the requested length is less than the ellipsis length, return a truncated ellipsis.
	if length <= ellipsisLen {
		return string(runeEllipsis[:length])
	}

	// Calculate the number of runes to keep from the original string.
	truncatedLen := length - ellipsisLen

	// Truncate and add ellipsis.
	return string(runeStr[:truncatedLen]) + ellipsis
}
