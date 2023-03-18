package main

import (
	"fmt"
	"testing"
)

func Test_alternatingCharacters(t *testing.T) {
	type testData struct {
		input    string
		expected int32
	}
	tests := []testData{
		{
			input:    "AAAA",
			expected: 3,
		},
		{
			input:    "BBBBB",
			expected: 4,
		},
		{
			input:    "ABABABAB",
			expected: 0,
		},
		{
			input:    "BABABA",
			expected: 0,
		},
		{
			input:    "AAABBB",
			expected: 4,
		},
	}
	for _, test := range tests {
		r := alternatingCharacters(test.input)
		if r != test.expected {
			t.Errorf("test failed for input %s, expected %d, got %d", test.input, test.expected, r)
		}
	}
	fmt.Println("end")
}
