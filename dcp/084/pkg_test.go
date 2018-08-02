package dcp084

import (
	"strconv"
	"testing"
)

func TestCountIslands(t *testing.T) {
	var tests = []struct {
		mat [][]int
		n   int // number of islands
	}{
		{nil, 0},
		{[][]int{{0}}, 0},
		{[][]int{{1}}, 1},
		{[][]int{{0, 0}, {0, 0}}, 0},
		{[][]int{{1, 0}, {0, 1}}, 2},
		{[][]int{{0, 1}, {1, 0}}, 2},
		{[][]int{{1, 1}, {1, 1}}, 1},
		{[][]int{{1, 1}, {1, 0}}, 1},
		{[][]int{{1, 1}, {0, 1}}, 1},
		{[][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}, {0, 1, 0}, {1, 0, 1}}, 8},
		{[][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 1}, {1, 0, 1}, {1, 0, 1}}, 2},
		{
			[][]int{
				{1, 0, 0, 0, 0},
				{0, 0, 1, 1, 0},
				{0, 1, 1, 0, 0},
				{0, 0, 0, 0, 0},
				{1, 1, 0, 0, 1},
				{1, 1, 0, 0, 1},
			},
			4,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			n := countIslands(tt.mat)
			if n != tt.n {
				t.Fatalf("found %d islands, want %d", n, tt.n)
			}
		})
	}
}
