package dcp093

import (
	"strconv"
	"testing"
)

func TestLargestBST(t *testing.T) {
	var nilNode *Node = nil
	var tests = []struct {
		tree *Node
		size int
	}{
		{nil, 0},
		{nd(5), 1},
		{nd(5, nd(4)), 2},
		{nd(5, nilNode, nd(6)), 2},
		{nd(5, nd(6), nd(4)), 1},
		{nd(5, nd(4), nd(6)), 3},
		{nd(5, nd(3, nd(1), nd(4)), nd(7, nd(6), nd(8))), 7},
		{nd(5, nd(3, nd(1), nd(4)), nd(7, nd(6), nd(6))), 3},
		{nd(5, nd(3, nd(4), nd(4)), nd(7, nd(6), nd(8))), 3},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			n := largestBST(tt.tree)
			if n != tt.size {
				t.Fatalf("largestBST returned %d, want %d", n, tt.size)
			}
		})
	}
}

// A convenient shorthand for creating nodes.
func nd(vs ...interface{}) *Node {
	var n Node
	setLeft := false
	for _, v := range vs {
		switch v := v.(type) {
		case int:
			n.value = int64(v)
		case *Node:
			if setLeft {
				n.right = v
			} else {
				n.left = v
				setLeft = true
			}
		}
	}
	return &n
}
