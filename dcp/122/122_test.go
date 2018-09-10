package dcp122

import (
	"strconv"
	"testing"
)

func TestMaxCoins(t *testing.T) {
	var tests = []struct {
		grid  [][]int
		coins int
	}{
		{nil, 0},
		{[][]int{{5}}, 5},
		{[][]int{{5, 11}}, 16},
		{[][]int{{5}, {11}}, 16},
		{[][]int{{5, 0}, {0, 11}}, 16},
		{[][]int{{0, 5}, {0, 11}}, 16},
		{[][]int{{0, 5}, {11, 0}}, 11},
		{[][]int{{5, 0}, {11, 0}}, 16},
		{
			[][]int{
				{0, 3, 1, 1},
				{2, 0, 0, 4},
				{1, 5, 3, 1},
			},
			12,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			coins := maxCoins(tt.grid)
			if coins != tt.coins {
				t.Fatalf("got %d, want %d", coins, tt.coins)
			}
		})
	}
}
