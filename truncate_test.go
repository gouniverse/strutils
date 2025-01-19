package strutils

import (
	"testing"
)

func TestTruncateString(t *testing.T) {
	testCases := []struct {
		name     string
		str      string
		length   int
		ellipsis string
		expected string
	}{
		{
			name:     "Short string, no truncation",
			str:      "Hello",
			length:   10,
			ellipsis: "...",
			expected: "Hello",
		},
		{
			name:     "Basic truncation",
			str:      "Hello, world!",
			length:   10,
			ellipsis: "...",
			expected: "Hello, ...",
		},
		{
			name:     "UTF-8 truncation",
			str:      "你好世界你好世界好世界！",
			length:   10,
			ellipsis: "...",
			expected: "你好世界你好世...",
		},
		{
			name:     "Empty string",
			str:      "",
			length:   5,
			ellipsis: "...",
			expected: "",
		},
		{
			name:     "Zero length",
			str:      "Test",
			length:   0,
			ellipsis: "...",
			expected: "",
		},
		{
			name:     "Negative length",
			str:      "Test",
			length:   -1,
			ellipsis: "...",
			expected: "",
		},
		{
			name:     "Ellipsis longer than length",
			str:      "Test",
			length:   2,
			ellipsis: "......",
			expected: "..",
		},
		{
			name:     "String length equals target length",
			str:      "Test",
			length:   4,
			ellipsis: "...",
			expected: "Test",
		},
		{
			name:     "String length is one more than target length",
			str:      "Tests",
			length:   4,
			ellipsis: "...",
			expected: "T...",
		},
		{
			name:     "Bulgarian truncation",
			str:      "Здравей свят!",
			length:   5,
			ellipsis: "...",
			expected: "Зд...",
		},
		{
			name:     "Bulgarian longer than length",
			str:      "Аз съм програмист",
			length:   10,
			ellipsis: "...",
			expected: "Аз съм ...",
		},
		{
			name:     "Bulgarian short string",
			str:      "Котка",
			length:   10,
			ellipsis: "...",
			expected: "Котка",
		},
		{
			name:     "Emoji truncation",
			str:      "🤔🤔🤔🤔🤔🤔🤔🤔🤔",
			length:   5,
			ellipsis: "...",
			expected: "🤔🤔...",
		},
		{
			name:     "Mixed Emoji and ASCII",
			str:      "Hello🤔world!",
			length:   10,
			ellipsis: "...",
			expected: "Hello🤔w...",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Truncate(tc.str, tc.length, tc.ellipsis)
			if actual != tc.expected {
				t.Errorf("Expected: %q, Actual: %q", tc.expected, actual)
			}
		})
	}
}

func FuzzTruncateString(f *testing.F) {
	testcases := []string{"Hello, world!", "你好世界！", "Short string", ""}
	for _, tc := range testcases {
		f.Add(tc, 5, "...")
	}
	f.Fuzz(func(t *testing.T, str string, length int, ellipsis string) {
		Truncate(str, length, ellipsis)
	})
}
