package main

import "testing"

func Test_repeatedString(t *testing.T) {

	type tData struct {
		s    string
		n    int64
		want int64
	}

	tests := []tData{
		{
			s:    "aba",
			n:    10,
			want: 7,
		},
		{
			s:    "a",
			n:    1000000000000,
			want: 1000000000000,
		},
		{
			s:    "ababa",
			n:    3,
			want: 2,
		},
		{
			s:    "b",
			n:    1000000000000,
			want: 0,
		},
		{
			s:    "bbbbb",
			n:    3,
			want: 0,
		},
	}

	for _, test := range tests {
		got := repeatedString(test.s, test.n)
		if got != test.want {
			t.Errorf("test failed got: %d, want: %d", got, test.want)
		}
	}
}
