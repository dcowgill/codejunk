package dcp023

import (
	"strconv"
	"testing"
)

func TestBFS(t *testing.T) {
	var tests = []struct {
		adj   [][]bool
		begin point
		end   point
		steps int
	}{
		{
			[][]bool{
				{false, false, false, false},
				{true, true, false, true},
				{false, false, false, false},
				{false, false, false, false},
			},
			point{3, 0},
			point{0, 0},
			7,
		},
		{
			[][]bool{{false}, {false}, {false}, {false}, {false}},
			point{1, 0},
			point{4, 0},
			3,
		},
		{
			[][]bool{{false}, {false}, {true}, {false}, {false}},
			point{1, 0},
			point{4, 0},
			-1,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			steps := bfs(tt.adj, tt.begin, tt.end)
			if steps != tt.steps {
				t.Fatalf("bfs(adj, %+v, %+v) returned %d, want %d", tt.begin, tt.end, steps, tt.steps)
			}
		})
	}
}
