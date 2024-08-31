package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []reportPage
	}{
		{
			name: "order count descending",
			input: map[string]int{
				"url1": 5,
				"url2": 1,
				"url3": 3,
				"url4": 10,
				"url5": 7,
			},
			expected: []reportPage{
				{url: "url4", count: 10},
				{url: "url5", count: 7},
				{url: "url1", count: 5},
				{url: "url3", count: 3},
				{url: "url2", count: 1},
			},
		},
		{
			name: "alphabetize",
			input: map[string]int{
				"d": 1,
				"a": 1,
				"e": 1,
				"b": 1,
				"c": 1,
			},
			expected: []reportPage{
				{url: "a", count: 1},
				{url: "b", count: 1},
				{url: "c", count: 1},
				{url: "d", count: 1},
				{url: "e", count: 1},
			},
		},

		{
			name:     "empty map",
			input:    map[string]int{},
			expected: []reportPage{},
		},
		{
			name:     "nil map",
			input:    nil,
			expected: []reportPage{},
		},
		{
			name: "one key",
			input: map[string]int{
				"url1": 1,
			},
			expected: []reportPage{
				{url: "url1", count: 1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if actual := sortPages(tc.input); !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("sortPages() = %v, expected %v", actual, tc.expected)
			}
		})
	}
}
