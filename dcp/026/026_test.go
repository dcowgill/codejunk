package dcp026

import (
	"fmt"
	"strings"
	"testing"
)

// For comparing lists in tests.
func listToString(ll *node) string {
	var b strings.Builder
	i := 0
	for p := ll; p != nil; p = p.next {
		fmt.Fprintf(&b, "%d=%q,", i, p.value)
		i++
	}
	return b.String()
}

func TestRemoveKth(t *testing.T) {
	var tests = []struct {
		head   *node
		k      int
		result string
	}{
		// Remove sole remaining element of a list.
		{&node{value: "foo"}, 0, ``},

		// Remove the head.
		{&node{value: "foo", next: &node{value: "bar", next: &node{value: "baz"}}}, 0, `0="bar",1="baz",`},

		// Remove the tail.
		{&node{value: "foo", next: &node{value: "bar", next: &node{value: "baz"}}}, 2, `0="foo",1="bar",`},

		// Remove the middle node.
		{&node{value: "foo", next: &node{value: "bar", next: &node{value: "baz"}}}, 1, `0="foo",1="baz",`},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("list=%s k=%d", listToString(tt.head), tt.k), func(t *testing.T) {
			result := listToString(removeKth(tt.head, tt.k))
			if result != tt.result {
				t.Fatalf("removeKth(head, %d) produced %s, want %s", tt.k, result, tt.result)
			}
		})
	}
}
