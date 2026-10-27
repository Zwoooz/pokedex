package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "         hello     world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "THiS sHould be spLit    aNd        tRImMeD aNd LoWeR",
			expected: []string{"this", "should", "be", "split", "and", "trimmed", "and", "lower"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Incorrect length of slice:")
			t.Errorf("Expected: %v", len(c.expected))
			t.Errorf("Actual: %v", len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("incorrect word at index %v:", i)
				t.Errorf("Expected: %v", expectedWord)
				t.Errorf("Actual: %v", word)
			}
		}
	}
}
