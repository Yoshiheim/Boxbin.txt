package helpers

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Its cut string by maxBytes size
func TruncateByte(s string, maxBytes int) string {
	if len(s) <= maxBytes { // len() returns byte count
		return s
	}
	for !utf8.ValidString(s[:maxBytes]) {
		maxBytes--
	}
	return string(s[:maxBytes])
}

// this function doesn't work like I want.
func SanitizeString(s string) string {
	result := make([]rune, 0, len(s))

	for _, r := range s {
		if unicode.IsPrint(r) || r == '\n' || r == '\t' {
			result = append(result, r)
		}
	}
	return string(result)
}

func DestroySpaces(text string) string {
	text = strings.ReplaceAll(text, "\t", "")
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\b", "")
	text = strings.ReplaceAll(text, "⠀", "")
	text = strings.ReplaceAll(text, " ", "")
	return text
}
