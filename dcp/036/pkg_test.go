package dcp036

import (
	"math"
	"strconv"
	"testing"
)

func TestSecondLargest(t *testing.T) {
	var tests = []struct {
		tree   *Node
		answer int64
	}{
		{nil, math.MinInt64},
		{&Node{value: 42}, math.MinInt64},
		{&Node{left: &Node{value: 37}, value: 42}, 37},
		{&Node{right: &Node{value: 99}, value: 42}, 42},
		{&Node{left: &Node{value: 37}, right: &Node{value: 99}, value: 42}, 42},
		{&Node{left: &Node{value: 37}, right: &Node{left: &Node{value: 81}, value: 99}, value: 42}, 81},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			n := secondLargest(tt.tree)
			if n != tt.answer {
				t.Fatalf("secondLargest returned %d, want %d", n, tt.answer)
			}
		})
	}
}
