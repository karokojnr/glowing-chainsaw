package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello word",
			expected: []string{"hello", "word"},
		},
		{
			input:    "HELLO word",
			expected: []string{"hello", "word"},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Expected %d words, got %d", len(cs.expected), len(actual))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("Expected %s, got %s", expectedWord, actualWord)
			}
		}

	}
}
