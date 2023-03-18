package main

import "testing"

func Test_twoStrings(t *testing.T) {
	type testData struct {
		one  string
		two  string
		want string
	}

	tests := []testData{
		{
			one:  "hello",
			two:  "world",
			want: "YES",
		},
		{
			one:  "hi",
			two:  "world",
			want: "NO",
		},
		{
			one:  "and",
			two:  "art",
			want: "YES",
		},
		{
			one:  "be",
			two:  "cat",
			want: "NO",
		},
	}

	for _, test := range tests {
		got := twoStrings(test.one, test.two)
		if got != test.want {
			t.Errorf("test failed for one: %s two: %s, got: %s want %s", test.one, test.two, got, test.want)
		}
	}

}
