package main

import "testing"

func Test_sherlockAndAnagram(t *testing.T) {
	type tData struct {
		s    string
		want int32
	}
	tests := []tData{
		{
			s:    "kkkk",
			want: 10,
		},
		{
			s:    "ifailuhkqq",
			want: 3,
		},
		{
			s:    "cdcd",
			want: 5,
		},
	}
	for _, test := range tests {
		got := sherlockAndAnagrams(test.s)
		if got != test.want {
			t.Errorf("test failed got %d want %d", got, test.want)
		}
	}
}
