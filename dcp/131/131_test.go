package dcp131

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func TestMakeList(t *testing.T) {
	var tests = []struct {
		values []int
		repr   string
	}{
		{nil, ""},
		{[]int{13}, "13"},
		{[]int{13, 17}, "13, 17"},
		{[]int{1, 2, 4, 8, 16}, "1, 2, 4, 8, 16"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.values), func(t *testing.T) {
			s := reprList(makeList(tt.values...), false)
			if s != tt.repr {
				t.Fatalf("repr(makeList(%v)) returned %q, want %q", tt.values, s, tt.repr)
			}
		})
	}
}

func TestCopyList(t *testing.T) {
	var tests = []struct {
		values []int
	}{
		{nil},
		{[]int{13}},
		{[]int{13, 17}},
		{[]int{1, 2, 4, 8, 16}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.values), func(t *testing.T) {
			ll := makeList(tt.values...)
			s1 := reprList(ll, true)
			s2 := reprList(copyList(ll), true)
			if s1 != s2 {
				t.Fatalf("got %q, want %q", s2, s1)
			}
		})
	}
}

// Creates a list from the given values.
func makeList(values ...int) *node {
	nodes := make([]*node, 0, len(values))
	var head *node
	p := &head
	for _, v := range values {
		*p = &node{value: v}
		nodes = append(nodes, *p)
		p = &(*p).next
	}
	// Set the "random" links.
	for q := head; q != nil; q = q.next {
		q.random = nodes[rand.Intn(len(nodes))]
	}
	return head
}

// Returns a string representation of the list: the nodes' values, separated by
// commas. E.g. the list [1, 2, 3] is represented as "1, 2, 3". Iff "inclRand"
// is true, the values of the "random" links are included in parentheses after
// the nodes' values; e.g. "1 (3), 2 (1), 3 (1)".
func reprList(head *node, inclRand bool) string {
	var b strings.Builder
	sep := ""
	for p := head; p != nil; p = p.next {
		fmt.Fprintf(&b, "%s%d", sep, p.value)
		if inclRand {
			fmt.Fprintf(&b, " (%d)", p.random.value)
		}
		sep = ", "
	}
	return b.String()
}
