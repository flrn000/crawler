package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("error fetching website: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("error: Status %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("invalid response content type: %v", contentType)
	}

	result, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(result[:]), nil
}
