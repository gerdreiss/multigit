package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v3"
)

func GetValue(yamlstring string, path string) (any, error) {
	var data any
	err := yaml.Unmarshal([]byte(yamlstring), &data)
	if err != nil {
		return "", err
	}

	parts := strings.Split(path, ".")
	current := data

	for _, part := range parts {
		if current == nil {
			return "", fmt.Errorf("path %q not found", path)
		}

		switch v := current.(type) {
		case map[string]any:
			// Map access
			if val, ok := v[part]; ok {
				current = val
			} else {
				return "", fmt.Errorf("key %q not found", part)
			}

		case []any:
			// Slice access with index
			idx, err := strconv.Atoi(part)
			if err != nil {
				return "", fmt.Errorf("cannot use %q as array index", part)
			}
			if idx < 0 || idx >= len(v) {
				return "", fmt.Errorf("index %d out of range", idx)
			}
			current = v[idx]

		case map[any]any:
			// For YAML with any keys
			if val, ok := v[part]; ok {
				current = val
			} else {
				return "", fmt.Errorf("key %q not found", part)
			}

		default:
			return "", fmt.Errorf("cannot navigate into %T at %q", current, part)
		}
	}

	return current, nil
}

func SetValue(yamlstring string, path string, value string) error {
	return fmt.Errorf("NOT IMPLEMENTED")
}
