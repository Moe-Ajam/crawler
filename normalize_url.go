package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputUrl string) (string, error) {
	u, err := url.Parse(inputUrl)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	normalizedURL := u.Host + u.Path
	normalizedURL = strings.ToLower(normalizedURL)
	normalizedURL = strings.TrimSuffix(normalizedURL, "/")

	return normalizedURL, nil
}
