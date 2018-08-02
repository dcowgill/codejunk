package dcp008

import "testing"

func TestUnivals(t *testing.T) {
	var tests = []struct {
		root *node
		n    int
	}{
		{nil, 0},
		{&node{value: 1}, 1},
		{&node{left: &node{value: 2}, right: &node{value: 3}, value: 1}, 2},
		{&node{left: &node{value: 1}, right: &node{value: 1}, value: 1}, 3},
		{
			&node{
				left: &node{value: 1},
				right: &node{
					left: &node{
						left:  &node{value: 1},
						right: &node{value: 1},
						value: 1,
					},
					right: &node{value: 0},
					value: 0,
				},
			},
			5,
		},
	}
	for i, tt := range tests {
		n, _ := univals(tt.root)
		if n != tt.n {
			t.Fatalf("test %d: got count of %d, want %d", i, n, tt.n)
		}
	}
}
