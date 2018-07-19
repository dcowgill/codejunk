package dcp066

import (
	"strconv"
	"testing"
)

func TestBishops(t *testing.T) {
	var tests = []struct {
		bishops []square
		n       int
	}{
		{[]square{{0, 0}, {1, 2}, {2, 2}, {4, 0}}, 2},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			n := numAttackingBishopPairs(tt.bishops)
			if n != tt.n {
				t.Fatalf("numAttackingBishopPairs returned %d, want %d", n, tt.n)
			}
		})
	}
}
