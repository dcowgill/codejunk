package dcp009

import (
	"fmt"
	"testing"
)

func TestMaxSumNonAdj(t *testing.T) {
	var tests = []struct {
		a []int
		n int
	}{
		{[]int{5, 1, 1, 5}, 10},
		{[]int{2, 4, 6, 8}, 12},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.a), func(t *testing.T) {
			if n := maxSumNonAdj(tt.a); n != tt.n {
				t.Fatalf("got %d, want %d", n, tt.n)
			}
		})
	}
}
