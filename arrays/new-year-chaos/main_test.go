package main

import "testing"

func Test_minimumBribes(t *testing.T) {
	type tData struct {
		q     []int32
		want  int
		wantE bool
	}

	tests := []tData{
		{
			q:     []int32{2, 1, 5, 3, 4},
			want:  3,
			wantE: false,
		},
		{
			q:     []int32{2, 5, 1, 3, 4},
			want:  -1, //don't test this here
			wantE: true,
		},
		{
			q:     []int32{5, 1, 2, 3, 7, 8, 6, 4},
			want:  -1, //don't test this here
			wantE: true,
		},
		{
			// int32{1, 2, 3, 4, 5, 6, 7, 8},
			// int32{1, 2, 3, 5, 4, 6, 7, 8} step 1
			// int32{1, 2, 5, 3, 4, 6, 7, 8} step 2
			//5-3=2
			// int32{1, 2, 5, 3, 4, 7, 6, 8} step 3
			// int32{1, 2, 5, 3, 7, 4, 6, 8} step 4
			//7-5=2
			// int32{1, 2, 5, 3, 7, 4, 8, 6} step 5
			// int32{1, 2, 5, 3, 7, 8, 4, 6} step 6
			//8-6=2
			// // int32{1, 2, 5, 3, 7, 8, 6, 4} step 7
			q:     []int32{1, 2, 5, 3, 7, 8, 6, 4},
			want:  7, //don't test this here
			wantE: false,
		},
	}

	for _, test := range tests {
		got, gotE := minimumBribesTestable(test.q)
		if (gotE != nil) != test.wantE {
			t.Errorf("test failed wanted error %t but got %s", test.wantE, gotE)
		}
		if test.want > 0 && got != test.want {
			t.Errorf("test failed got %d, want %d", got, test.want)
		}
	}
}
