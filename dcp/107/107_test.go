package dcp107

import (
	"bytes"
	"strconv"
	"testing"
)

func TestPrintBinaryTreeBreadthFirst(t *testing.T) {
	var nilNode *node
	var tests = []struct {
		tree     *node
		expected string
	}{
		{nil, ""},
		{nd(1), "1"},
		{nd(1, nd(2)), "1, 2"},
		{nd(1, nilNode, nd(2)), "1, 2"},
		{nd(1, nd(2), nd(3, nd(4), nd(5))), "1, 2, 3, 4, 5"},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var buf bytes.Buffer
			printBinaryTreeBreadthFirst(&buf, tt.tree)
			if s := buf.String(); s != tt.expected {
				t.Fatalf("printBinaryTreeBreadthFirst printed %q, want %q", s, tt.expected)
			}
		})
	}
}

// A convenient shorthand for creating nodes.
func nd(vs ...interface{}) *node {
	var n node
	setLeft := false
	for _, v := range vs {
		switch v := v.(type) {
		case int:
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
