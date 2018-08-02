package dcp040

import (
	"fmt"
	"testing"
)

func TestFindNonDuplicate(t *testing.T) {
	var tests = []struct {
		a []int
		n int
	}{
		{[]int{7}, 7},
		{[]int{6, 1, 3, 3, 3, 6, 6}, 1},
		{[]int{5, 4, 3, 2, 5, 4, 3, 5, 4, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.a), func(t *testing.T) {
			n := findNonDuplicate(tt.a)
			if n != tt.n {
				t.Fatalf("findNonDuplicate returned %d, want %d", n, tt.n)
			}
		})
	}
}
