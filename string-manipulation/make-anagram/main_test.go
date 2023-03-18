package main

import "testing"

func Test_makeAnagram(t *testing.T) {
	type testData struct {
		one  string
		two  string
		want int32
	}
	tests := []testData{
		{
			one:  "cde",
			two:  "abc",
			want: 4,
		},
		{
			one:  "abcccc",
			two:  "abc",
			want: 3,
		},
		{
			one:  "abccccddddddd",
			two:  "abcd",
			want: 9,
		},
		{
			one:  "abcd",
			two:  "abccccddddddd",
			want: 9,
		},
	}
	for _, test := range tests {
		got := makeAnagram(test.one, test.two)
		if got != test.want {
			t.Errorf("test failed for one: %s two: %s want: %d got %d", test.one, test.two, test.want, got)
		}
	}
}
