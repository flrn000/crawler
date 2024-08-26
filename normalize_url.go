package main

import "net/url"

func normalizeURL(targetURL string) (string, error) {
	u, err := url.Parse(targetURL)
	if err != nil {
		return "", err
	}

	normalizedURL := u.Hostname() + u.EscapedPath()

	return normalizedURL, nil
}
