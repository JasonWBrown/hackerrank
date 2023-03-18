package main

import "testing"

func Test_sockMerchant(t *testing.T) {
	type tData struct {
		n    int32
		ar   []int32
		want int32
	}

	tests := []tData{
		{
			n:    9,
			ar:   []int32{10, 20, 20, 10, 10, 30, 50, 10, 20},
			want: 3,
		},
		{
			n:    10,
			ar:   []int32{10, 20, 20, 10, 10, 30, 50, 10, 20, 30},
			want: 4,
		},
		{
			n:    1,
			ar:   []int32{10},
			want: 0,
		},
	}

	for _, test := range tests {
		got := sockMerchant(test.n, test.ar)
		if got != test.want {
			t.Errorf("test failed got %d, want %d", got, test.want)
		}
	}

}
