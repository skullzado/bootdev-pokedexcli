package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		// special symbols are NOT included
		// all words must be LOWERCASED
		// empty input returns nothing
		{
			input:    "Hello, World!",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello   World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Println(actual)
		fmt.Println(c)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(c.expected) != len(actual) {
			t.Errorf("Slice length does not match\nExpected: %d\nActual: %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Words do not match\nExpected: %s\nActual: %s", expectedWord, word)
			}
		}
	}
}
