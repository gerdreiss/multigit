package helpers

import (
	"fmt"
	"net/url"
)

// GetHostName returns the host of the given raw URL
func GetHostName(rawUrl string) (string, error) {
	r, err := url.Parse(rawUrl)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %w", err)
	}
	return r.Host, nil
}
