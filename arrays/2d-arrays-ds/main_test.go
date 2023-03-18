package main

import "testing"

func Test_hourglassSum(t *testing.T) {
	type tData struct {
		arr  [][]int32
		want int32
	}

	tests := []tData{
		{
			arr: [][]int32{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 0, 2, 4, 4, 0},
				{0, 0, 0, 2, 0, 0},
				{0, 0, 1, 2, 4, 0},
			},
			want: 19,
		},
		{
			arr: [][]int32{
				{0, 0, 0, 0, 0, 0},
				{0, -9, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, -8, 0},
				{0, 0, 0, 0, 0, 0},
			},
			want: 0,
		},
		{
			arr: [][]int32{
				{-1, -1, 0, -9, -2, -2},
				{-2, -1, -6, -8, -2, -5},
				{-1, -1, -1, -2, -3, -4},
				{-1, -9, -2, -4, -4, -5},
				{-7, -3, -3, -2, -9, -9},
				{-1, -3, -1, -2, -4, -5},
			},
			want: -6,
		},
	}
	for _, test := range tests {
		got := hourglassSum(test.arr)
		if got != test.want {
			t.Errorf("test failed got: %d, want %d", got, test.want)
		}
	}
}
