package main

import "testing"

func Test_checkMagazine(t *testing.T) {
	type tData struct {
		name     string
		magazine []string
		note     []string
		want     bool
	}

	tests := []tData{
		{
			name:     "attack vs Attack ie case sensitive",
			magazine: []string{"attack", "at", "dawn"},
			note:     []string{"Attack", "at", "dawn"},
			want:     false,
		},
		{
			name:     "note can be made",
			magazine: []string{"give", "me", "one", "grand", "today", "night"},
			note:     []string{"give", "one", "grand", "today"},
			want:     true,
		},
		{
			name:     "more of one word in note than in magazine",
			magazine: []string{"two", "times", "three", "is", "not", "four"},
			note:     []string{"two", "times", "two", "is", "four"},
			want:     false,
		},
	}
	for _, test := range tests {
		got := checkMagazine(test.magazine, test.note)
		if got != test.want {
			t.Errorf("test failed got %t want %t", got, test.want)
		}
	}
}
