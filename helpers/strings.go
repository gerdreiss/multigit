package helpers

import (
	"slices"
	"strings"
)

func SuffixAfterLast(s string, sep string) string {
	parts := strings.Split(s, sep)
	n := len(parts)
	if n == 0 {
		return ""
	}
	if n > 1 {
		slices.Reverse(parts)
	}
	return parts[0]
}
