package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
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

	if res.Header.Get("Content-Type") != "text/html" {
		return "", errors.New("response content type is not 'text/html'")
	}

	result, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(result[:]), nil
}
