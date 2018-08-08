package dcp089

import (
	"strconv"
	"testing"
)

func TestIsValid(t *testing.T) {
	var nilNode *Node
	var tests = []struct {
		tree *Node
		b    bool // isValid
	}{
		{nil, true},
		{nd(10), true},
		{nd(0, nd(-1)), true},
		{nd(0, nd(+1)), false},
		{nd(0, nilNode, nd(-1)), false},
		{nd(0, nilNode, nd(+1)), true},
		{nd(6, nd(3, nd(1), nd(4)), nd(7, nd(6), nd(8))), true},
		{nd(5, nd(3, nd(1), nd(6)), nd(7, nd(6), nd(8))), false},
		{nd(5, nd(4, nd(3, nd(2, nd(1, nd(0, nilNode, nd(6))))))), false},
		{nd(5, nd(4, nd(3, nd(2, nd(1, nd(0, nilNode, nd(1))))))), true},
		{nd(5, nilNode, nd(6, nilNode, nd(7, nilNode, nd(9, nd(6))))), false},
		{nd(5, nilNode, nd(6, nilNode, nd(7, nilNode, nd(9, nd(8))))), true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if b := isValid(tt.tree); b != tt.b {
				t.Fatalf("isValid returned %v, want %v", b, tt.b)
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
			n.key = Key(v)
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
