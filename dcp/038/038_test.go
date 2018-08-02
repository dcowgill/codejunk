package dcp038

import "testing"

func TestQueens(t *testing.T) {
	var tests = []struct {
		n int // board size
		r int // result
	}{
		{5, 10},
		{6, 4},
		{7, 40},
		{8, 92},
		{9, 352},
		{10, 724},
		{11, 2680},
		{12, 14200},
		{13, 73712},
		// {14, 365596},
		// {15, 2279184},
		// {16, 14772512},
		// {17, 95815104},
		// {18, 666090624},
		// {19, 4968057848},
		// {20, 39029188884},
		// {21, 314666222712},
	}
	for _, tt := range tests {
		r := queens(tt.n)
		if r != tt.r {
			t.Errorf("queens(%d) returned %d, want %d", tt.n, r, tt.r)
		}
	}
}
