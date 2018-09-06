package dcp118

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	var tests = []struct {
		input, squares []int
	}{
		{nil, []int{}},
		{[]int{}, []int{}},
		{[]int{3}, []int{9}},
		{[]int{2, 3}, []int{4, 9}},
		{[]int{-4, 3}, []int{9, 16}},
		{[]int{-3, 4}, []int{9, 16}},
		{[]int{1, 2, 3, 4}, []int{1, 4, 9, 16}},
		{[]int{-4, -3, -2, -1}, []int{1, 4, 9, 16}},
		{[]int{-9, -2, 0, 2, 3}, []int{0, 4, 4, 9, 81}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.input), func(t *testing.T) {
			squares := sortedSquares(tt.input)
			if !reflect.DeepEqual(squares, tt.squares) {
				t.Fatalf("got %+v, want %+v", squares, tt.squares)
			}
		})
	}
}
