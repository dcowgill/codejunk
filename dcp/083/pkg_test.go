package dcp083

import (
	"fmt"
	"testing"
)

func TestInvertTree(t *testing.T) {
	var nilNode *node = nil
	var tests = []struct {
		tree     *node
		inverted string
	}{
		{nil, "_"},
		{nd("a"), "(a)"},
		{nd("a", nd("b")), "(a _ (b))"},
		{nd("a", nilNode, nd("b")), "(a (b) _)"},
		{nd("a", nd("b", nd("d"), nd("e")), nd("c", nd("f"))), "(a (c _ (f)) (b (e) (d)))"},
	}
	for _, tt := range tests {
		t.Run(repr(tt.tree), func(t *testing.T) {
			s := repr(invertTree(tt.tree))
			if s != tt.inverted {
				t.Fatalf("got %q, want %q", s, tt.inverted)
			}
		})
	}
}

// Returns a string representation of a preorder traversal of n.
func repr(n *node) string {
	if n == nil {
		return "_"
	}
	if n.left != nil || n.right != nil {
		return fmt.Sprintf("(%s %s %s)", n.value, repr(n.left), repr(n.right))
	}
	return fmt.Sprintf("(%s)", n.value)
}

// A convenient shorthand for creating nodes.
func nd(vs ...interface{}) *node {
	var n node
	setLeft := false
	for _, v := range vs {
		switch v := v.(type) {
		case string:
			n.value = v
		case *node:
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
