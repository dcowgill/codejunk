package dcp110

import (
	"reflect"
	"strconv"
	"testing"
)

func TestAllPathsToLeaves(t *testing.T) {
	var nilNode *node
	var tests = []struct {
		tree  *node
		paths [][]int
	}{
		{nil, nil},
		{nd(1), [][]int{{1}}},
		{nd(1, nd(2)), [][]int{{1, 2}}},
		{nd(1, nilNode, nd(2)), [][]int{{1, 2}}},
		{nd(1, nd(2, nd(3))), [][]int{{1, 2, 3}}},
		{nd(1, nd(2, nd(3, nd(4), nd(5)))), [][]int{{1, 2, 3, 4}, {1, 2, 3, 5}}},
		{nd(1, nd(2), nd(3, nd(4), nd(5))), [][]int{{1, 2}, {1, 3, 4}, {1, 3, 5}}},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			paths := allPathsToLeaves(tt.tree)
			if !reflect.DeepEqual(paths, tt.paths) {
				t.Fatalf("got %+v, want %+v", paths, tt.paths)
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
