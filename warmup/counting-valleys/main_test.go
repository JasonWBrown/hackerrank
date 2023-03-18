package main

import "testing"

func Test_countingValleys(t *testing.T) {
	type testData struct {
		steps int32
		path  string
		want  int32
	}

	tests := []testData{
		{
			steps: 8,
			path:  "UDDDUDUU",
			want:  1,
		},
		{
			steps: 8,
			path:  "UDDDUUUUUDDDDDDDDDUU",
			want:  2,
		},
	}

	for _, test := range tests {
		got := countingValleys(test.steps, test.path)
		if got != test.want {
			t.Errorf("test failed, got %d want %d", got, test.want)
		}
	}
}
