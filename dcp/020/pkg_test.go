package dcp020

import (
	"strconv"
	"testing"
)

func TestFindIntersectingNode(t *testing.T) {
	var tests = []struct {
		a, b []int
		v    int
	}{
		{[]int{3, 7, 8, 10}, []int{99, 1, 8, 10}, 8},
		{[]int{3, 7, 8, 10}, []int{5, 6, 7, 99, 1, 8, 10}, 8},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 1},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			a := makeList(tt.a)
			b := makeList(tt.b)
			v := findIntersectingNode(a, b)
			if v != tt.v {
				t.Fatalf("findIntersectingNode1(%+v, %+v) returned %d, want %d", tt.a, tt.b, v, tt.v)
			}
		})
	}
}

// Helper: creates a list containing the values in xs.
func makeList(xs []int) *node {
	var head *node
	p := &head
	for _, x := range xs {
		*p = &node{value: x}
		p = &(*p).next
	}
	return head
}
