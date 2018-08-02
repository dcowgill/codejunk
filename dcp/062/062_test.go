package dcp062

import "testing"

func TestDFS(t *testing.T) {
	var tests = []struct {
		n, m   int
		answer int
	}{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 1},
		{2, 2, 2},
		{3, 2, 3},
		{2, 3, 3},
		{3, 3, 6},
		{3, 4, 10},
		{4, 3, 10},
		{4, 5, 35},
		{5, 4, 35},
		{5, 5, 70},
	}
	for _, tt := range tests {
		answer := dfs(tt.n, tt.m, 0, 0)
		if answer != tt.answer {
			t.Fatalf("dfs(%d, %d) returned %d, want %d", tt.n, tt.m, answer, tt.answer)
		}
	}
}
