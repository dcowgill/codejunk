package dcp078

import (
	"strconv"
	"strings"
	"testing"
)

func TestMergeLists(t *testing.T) {
	var tests = []struct {
		lists  []*node
		result string
	}{
		{nil, ""},
		{[]*node{nd("a")}, "a,"},
		{[]*node{nd("a"), nd("b")}, "a,b,"},
		{[]*node{nd("b"), nd("a")}, "a,b,"},
		{[]*node{nd("a", nd("b", nd("c"))), nd("d")}, "a,b,c,d,"},
		{[]*node{
			nd("c", nd("f", nd("i"))),
			nd("a", nd("d", nd("g"))),
			nd("b", nd("e", nd("h"))),
		},
			"a,b,c,d,e,f,g,h,i,",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := dump(mergeLists(tt.lists))
			if result != tt.result {
				t.Fatalf("mergeLists returned %q, want %q", result, tt.result)
			}
		})
	}
}

// Shorthand for creating nodes.
func nd(vs ...interface{}) *node {
	var n node
	for _, v := range vs {
		switch v := v.(type) {
		case string:
			n.value = v
		case *node:
			n.next = v
		}
	}
	return &n
}

// Converts the list to a string (to ease comparisons).
func dump(head *node) string {
	var b strings.Builder
	for p := head; p != nil; p = p.next {
		b.WriteString(p.value)
		b.WriteByte(',')
	}
	return b.String()
}
