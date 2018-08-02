package dcp079

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	var tests = []struct {
		a []int
		b bool
	}{
		{nil, true},
		{[]int{1}, true},
		{[]int{5, 7}, true},
		{[]int{7, 5}, true},
		{[]int{5, 1, 3}, true},
		{[]int{1, 5, 3}, true},
		{[]int{3, 1, 5}, true},
		{[]int{10, 5, 7}, true},
		{[]int{10, 5, 1}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.a), func(t *testing.T) {
			b := solve(tt.a)
			if b != tt.b {
				t.Fatalf("got %v, want %v", b, tt.b)
			}
		})
	}
}
