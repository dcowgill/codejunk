package dcp012

import (
	"fmt"
	"math/big"
	"testing"
)

func TestStaircase(t *testing.T) {
	var tests = []struct {
		X []int    // counts of steps that may be climbed at once
		n int      // number of steps in the staircase
		r *big.Int // number of ways to climb the staircase
	}{
		{[]int{1, 2}, 4, big.NewInt(5)},    // 1111 112 121 211 22
		{[]int{1, 3, 5}, 5, big.NewInt(5)}, // 11111 113 131 311 5
		{[]int{1, 3, 5}, 6, big.NewInt(8)}, // 111111 1113 1131 1311 3111 33 15 51
		{[]int{1, 3, 5, 7}, 200, fromString("37769780303971247390600001555026771892981")},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d@%v", tt.n, tt.X), func(t *testing.T) {
			r := staircase(tt.n, tt.X)
			if r.Cmp(tt.r) != 0 {
				t.Fatalf("got %s, want %s", r, tt.r)
			}
		})
	}
}

func fromString(s string) *big.Int {
	n := new(big.Int)
	_, b := n.SetString(s, 10)
	if !b {
		panic(fmt.Sprintf("big.Int.SetString(%q, 10) failed", s))
	}
	return n
}
