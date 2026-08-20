package helpers

// Deduplicate removes duplicate strings from a slice,
// preserving the order of first appearance.
func Deduplicate[T comparable](input []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0, len(input))

	for _, s := range input {
		if _, exists := seen[s]; !exists {
			seen[s] = struct{}{}
			result = append(result, s)
		}
	}
	return result
}
