package dcp019

import (
	"strconv"
	"testing"
)

func TestMincost(t *testing.T) {
	var tests = []struct {
		costs  [][]int
		answer int
	}{
		{
			[][]int{
				{20, 10, 30}, // house 0
				{60, 40, 10}, // house 1
				{90, 90, 10}, // house 2
			},
			70, // colors: 0(20), 1(40), 2(10)
		},
		// TODO: more tests
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if c := mincost(tt.costs); c != tt.answer {
				t.Fatalf("mincost returned %d, want %d", c, tt.answer)
			}
		})
	}
}
