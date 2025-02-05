package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Sliced len != expected len")
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedword := c.expected[i]

			if word != expectedword {
				t.Errorf("word at index %v did not match", i)
				t.Errorf("Expected: %s", expectedword)
				t.Errorf("Actual: %s", word)
				t.Fail()
			}
		}
	}
}
