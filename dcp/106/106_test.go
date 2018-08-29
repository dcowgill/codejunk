package dcp106

import (
	"fmt"
	"testing"
)

func TestCanReachEnd(t *testing.T) {
	var tests = []struct {
		hops     []int
		expected bool
	}{
		{[]int{2, 0, 1, 0}, true},
		{[]int{1, 1, 0, 1}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.hops), func(t *testing.T) {
			b := canReachEnd(tt.hops)
			if b != tt.expected {
				t.Fatalf("canReachEnd returned %v, want %v", b, tt.expected)
			}
		})
	}
}
