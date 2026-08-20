package helpers

import (
	"slices"
	"strings"
)

func SuffixAfterLast(s string, sep string) string {
	parts := strings.Split(s, sep)
	slices.Reverse(parts)
	return parts[0]
}
