package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLs(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://domain.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one" class="test-classs">
					<span>domain.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>domain.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://domain.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "nested a tags",
			inputURL: "https://cinema.dev",
			inputBody: `
		<html>
			<body>
				<a href="path/two" class="test-classs">
					<span>domain.dev</span>
				</a>
				<a href="https://stackoverflow.co/teams/ai?utm_medium=referral&utm_content=overflowai">
					<span>domain.dev</span>
				</a>
				<div>
					<p>
						<a href="legal/terms-of-service" target="_blank" rel="noopener noreferrer">
							just some text
						</a>
					</p>
				</div
			</body>
		</html>
		`,
			expected: []string{
				"https://cinema.dev/path/two",
				"https://stackoverflow.co/teams/ai?utm_medium=referral&utm_content=overflowai",
				"https://cinema.dev/legal/terms-of-service",
			},
		},
		{
			name:     "fragment URLs",
			inputURL: "https://fragment.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one#test" class="test-classs">
					<span>fragment.dev</span>
				</a>
				<a href="https://other.com/path/one#anotherone" name="test">
					<span>fragment.dev</span>
				</a>
				<a href="#check">Checking empty fragment</a>
			</body>
		</html>
		`,
			expected: []string{"https://fragment.dev/path/one#test", "https://other.com/path/one#anotherone"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - %s FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}

			actual, err := getProductsFromHTML(tc.inputBody)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}
