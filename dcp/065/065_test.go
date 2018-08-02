package dcp065

import (
	"reflect"
	"testing"
)

func TestSpiral(t *testing.T) {
	var tests = []struct {
		mat [][]int
		run []int
	}{
		{nil, nil},
		{[][]int{{1}, {2}}, []int{1, 2}},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2, 3, 6, 5, 4}},
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 4, 3}},
		{[][]int{{1, 2}, {3, 4}, {5, 6}}, []int{1, 2, 4, 6, 5, 3}},
		{
			[][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}},
			[]int{1, 2, 3, 4, 5, 10, 15, 20, 19, 18, 17, 16, 11, 6, 7, 8, 9, 14, 13, 12},
		},
	}
	for _, tt := range tests {
		run := spiral(tt.mat)
		if !reflect.DeepEqual(run, tt.run) {
			t.Fatalf("spiral(%+v) returned %+v, want %+v", tt.mat, run, tt.run)
		}
	}
}
