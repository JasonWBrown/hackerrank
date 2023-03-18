package main

import "testing"

func Test_jumpingOnClouds(t *testing.T) {
	type testData struct {
		c    []int32
		want int32
	}

	tests := []testData{
		{
			c:    []int32{0, 0, 1, 0, 0, 1, 0},
			want: 4,
		},
		{
			c:    []int32{0, 0, 0, 0, 1, 0},
			want: 3,
		},
		{
			c:    []int32{0, 0, 0, 1, 0, 0},
			want: 3,
		},
	}

	for _, test := range tests {
		got := jumpingOnClouds(test.c)
		if got != test.want {
			t.Errorf("test failed got %d, want %d", got, test.want)
		}
	}
}
