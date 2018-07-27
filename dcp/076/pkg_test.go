package dcp076

import (
	"strconv"
	"testing"
)

func TestMinColsToRemove(t *testing.T) {
	var tests = []struct {
		mat [][]rune
		n   int
	}{
		{[][]rune{{'c', 'b', 'a'}, {'d', 'a', 'f'}, {'g', 'h', 'i'}}, 1},
		{[][]rune{{'a', 'b', 'c', 'd', 'e', 'f'}}, 0},
		{[][]rune{{'z', 'y', 'x'}, {'w', 'v', 'u'}, {'t', 's', 'r'}}, 3},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if n := minColsToRemove(tt.mat); n != tt.n {
				t.Fatalf("got %d, want %d", n, tt.n)
			}
		})
	}
}
