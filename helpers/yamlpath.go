package helpers

import (
	"fmt"
	"slices"
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

func SetValue(yamlstring string, path string, value string) (string, error) {
	var data map[string][]any
	err := yaml.Unmarshal([]byte(yamlstring), &data)
	if err != nil {
		return "", err
	}

	var newvalue any = value

	segments := strings.Split(path, ".")
	subpath := segments[2:]
	slices.Reverse(subpath)
	for _, segment := range subpath {
		newvalue = map[string]any{
			segment: newvalue,
		}
	}

	idx, err := strconv.Atoi(segments[1])
	if err != nil {
		return "", err
	}

	gits := data["git"]
	ngits := len(gits)
	if idx > ngits {
		return "", fmt.Errorf("index out of bounds")
	}
	if idx == ngits {
		gits = append(gits, newvalue)
	} else {
		return "", fmt.Errorf("overriding values is not yet implemented")
	}

	data["git"] = gits

	result, err := yaml.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
