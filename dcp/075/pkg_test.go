package dcp075

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestIncreasingSubseq(t *testing.T) {
	var tests = []struct {
		a []int
		s []int
	}{
		{nil, nil},
		{[]int{}, nil},
		{[]int{2, 4, 6, 8}, []int{2, 4, 6, 8}},
		{[]int{8, 6, 4, 2}, []int{8}},
		{[]int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}, []int{0, 2, 6, 9, 11, 15}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.a), func(t *testing.T) {
			s := longestIncreasingSubseq(tt.a)
			if !reflect.DeepEqual(s, tt.s) {
				t.Fatalf("got %+v, want %+v", s, tt.s)
			}
		})
	}
}
