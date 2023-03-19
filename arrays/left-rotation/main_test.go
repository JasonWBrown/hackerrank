package main

import "testing"

func Test_rotLeft(t *testing.T) {
	type tData struct {
		a    []int32
		d    int32
		want []int32
	}

	tests := []tData{
		{
			a:    []int32{1, 2, 3, 4, 5},
			d:    4,
			want: []int32{5, 1, 2, 3, 4},
		},
		{
			a:    []int32{1, 2, 3, 4, 5},
			d:    2,
			want: []int32{3, 4, 5, 1, 2},
		},
		{
			a:    []int32{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
			d:    1,
			want: []int32{2, 3, 4, 5, 1, 2, 3, 4, 5, 1},
		},
	}
	for _, test := range tests {
		got := rotLeft(test.a, test.d)
		if !testEq(got, test.want) {
			t.Errorf("test failed got %v want %v", got, test.want)
		}
	}
}

func testEq(got, want []int32) bool {
	if len(got) != len(want) {
		return false
	}
	for i, _ := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
