package main

import (
	"strings"
	"testing"
)

func TestCleanText(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   hello world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   hello    world",
			expected: []string{"hello", "world"},
		},
	}

	for _, cs := range cases {
		actual := cleanText(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("\nThe Lengths were not equal \nExpected length %v \nGot length %v\n", len(cs.expected), len(actual))
			continue
		}

		if strings.Join(actual, "") != strings.Join(cs.expected, "") {
			t.Errorf("\nExpected text %v \nGot text %v\n", cs.expected, actual)
			continue
		}
	}

}
