package main

import (
	"fmt"
	"strings"
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

// from fib_test.go
func Benchmark_alternatingCharacters_10000(b *testing.B) {
	// run the Fib function b.N times
	var a string
	for i := 0; i < 10000; i++ {
		a = strings.Join([]string{a}, "A")
	}
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		alternatingCharacters(a)
	}
}

// from fib_test.go
func Benchmark_alternatingCharacters_1000(b *testing.B) {
	// run the Fib function b.N times
	b.StopTimer()
	var a string
	for i := 0; i < 1000; i++ {
		a = strings.Join([]string{a}, "A")
	}

	b.ResetTimer()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		alternatingCharacters(a)
	}
}

func Benchmark_alternatingCharacters_100(b *testing.B) {
	// run the Fib function b.N times
	var a string
	for i := 0; i < 100; i++ {
		a = strings.Join([]string{a}, "A")
	}
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		alternatingCharacters(a)
	}
}

func Benchmark_alternatingCharacters_10(b *testing.B) {
	// run the Fib function b.N times
	a := "AAAAAAAAAA"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		alternatingCharacters(a)
	}
}
