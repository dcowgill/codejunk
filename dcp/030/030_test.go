package dcp030

import (
	"fmt"
	"testing"
)

func TestRainfall(t *testing.T) {
	var tests = []struct {
		wall []int
		n    int
	}{
		{[]int{}, 0},
		{[]int{5, 0, 0}, 0},
		{[]int{6, 0, 0, 4, 0, 1, 0, 3, 0}, 16},
		{[]int{0, 3, 0, 1, 0, 4, 0, 0, 6}, 16},
		{[]int{4, 0, 1, 0, 3, 0, 6}, 16},
		{[]int{2, 1, 2}, 1},
		{[]int{0, 2, 1, 2, 0}, 1},
		{[]int{3, 0, 1, 3, 0, 5}, 8},
		{[]int{0, 0, 0}, 0},
		{[]int{0, 1, 0}, 0},
		{[]int{1, 0, 10, 0, 10, 0, 1}, 12},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.wall), func(t *testing.T) {
			n := rainfall(tt.wall)
			if n != tt.n {
				t.Fatalf("rainfall returned %d, want %d", n, tt.n)
			}
		})
	}
}
