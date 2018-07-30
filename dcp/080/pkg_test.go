package dcp080

import (
	"strconv"
	"testing"
)

func TestFindDeepestNode(t *testing.T) {
	var nilNode *node = nil // see nd()
	var tests = []struct {
		root  *node
		value string
	}{
		{nil, ""},
		{nd("x"), "x"},
		{nd("x", nd("y")), "y"},
		{nd("x", nilNode, nd("y")), "y"},
		{nd("a", nd("b", nd("d")), nd("c")), "d"},
		{nd("a", nd("b", nd("d", nd("e"))), nd("c")), "e"},
		{nd("a", nd("b", nd("d", nd("e"))), nd("c", nilNode, nd("f", nd("g", nd("h"))))), "h"},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			value := findDeepestNode(tt.root)
			if value != tt.value {
				t.Fatalf("findDeepestNode returned %q, want %q", value, tt.value)
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
